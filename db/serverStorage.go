package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Storage) GetServerStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servers, err := s.UpdateServerStats()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(servers)
	if err != nil {
		log.Printf("Error encoding servers to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// DB QUERIES
func (s *Storage) UpdateServerStats() ([]Server, error) {
	url := "http://www.mfc-25.ru/queue/server.json"
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot do request, error: %v\n", err))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot read response body, error: %v\n", err))
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	servers := make([]Server, 0)

	err = json.Unmarshal([]byte(body), &servers)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot unmarshal response, error: %v\n", err))
	}

	txn, err := s.db.Begin()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot begin database, error: %v\n", err))
	}

	stmt, err := txn.Prepare(pq.CopyIn("servers", "id", "name", "is_connected", "wrong_protocol",
		"organization_name", "organization_fullname", "organization_address", "organization_phone", "organization_fax", "organization_email"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot prepare the transaction, error: %v\n", err))
	}

	for _, srv := range servers {
		_, err := stmt.Exec(srv.Id, srv.Name, srv.IsConnected, srv.WrongProtocol, srv.OrganizationName, srv.OrganizationFullName, srv.OrganizationAddress,
			srv.OrganizationPhone, srv.OrganizationFax, srv.OrganizationEmail)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot execute servers statement, error: %v\n", err))
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot execute the statement, error: %v\n", err))
	}
	err = stmt.Close()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot close the statement, error: %v\n", err))
	}

	err = txn.Commit()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot commit the transaction, error: %v\n", err))
	}

	return servers, nil
}