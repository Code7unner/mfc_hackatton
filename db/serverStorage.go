package db

import (
	"bytes"
	"encoding/json"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

// DB QUERIES
func (s *Storage) GetServerStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "http://www.mfc-25.ru/queue/server.json"
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot do request, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot read response body, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	servers := make([]Server, 0)

	err = json.Unmarshal([]byte(body), &servers)
	if err != nil {
		log.Printf("Cannot unmarshal response, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	txn, err := s.db.Begin()
	if err != nil {
		log.Printf("Cannot begin database, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	stmt, err := txn.Prepare(pq.CopyIn("servers", "id", "name", "is_connected", "wrong_protocol",
		"organization_name", "organization_full_name", "organization_address", "organization_phone", "organization_fax", "organization_email"))
	if err != nil {
		log.Printf("Cannot prepare the transaction, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	for i := range servers {
		srv := servers[i]
		_, err := stmt.Exec(srv.Id, srv.Name, srv.IsConnected, srv.WrongProtocol, srv.OrganizationName, srv.OrganizationFullName, srv.OrganizationAddress,
			srv.OrganizationPhone, srv.OrganizationFax, srv.OrganizationEmail)
		if err != nil {
			log.Printf("Cannot execute the statement, error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("Cannot execute the statement, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = stmt.Close()
	if err != nil {
		log.Printf("Cannot close the statement, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = txn.Commit()
	if err != nil {
		log.Printf("Cannot commit the transaction, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(servers)
	if err != nil {
		log.Printf("Error encoding servers to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}