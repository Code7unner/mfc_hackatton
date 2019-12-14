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
func (s *Storage) GetStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conte" +
		"nt-Type", "application/json")

	url := "http://www.mfc-25.ru/queue/statistics.json"
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

	statistics := make([]Statistics, 0)

	err = json.Unmarshal([]byte(body), &statistics)
	if err != nil {
		log.Printf("Cannot unmarshal response, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	txn, err := s.db.Begin()
	if err != nil {
		log.Printf("Cannot begin database, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	stmt, err := txn.Prepare(pq.CopyIn("statistics", "id", "average_awaiting_time", "active_work_places_count", "completed_tickets_count",
		"pending_tikets_count"))
	if err != nil {
		log.Printf("Cannot prepare the transaction, error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, sts := range statistics {
		_, err := stmt.Exec(sts.Id, sts.AverageAwaitingTime, sts.ActiveWorkPlacesCount, sts.CompletedTicketsCount, sts.PendingTicketsCount)
		if err != nil {
			log.Printf("Cannot execute statistics statement, error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		log.Println(sts.Id)
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

	err = json.NewEncoder(w).Encode(statistics)
	if err != nil {
		log.Printf("Error encoding servers to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}