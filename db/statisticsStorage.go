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

func (s *Storage) GetStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stats, err := s.UpdateStatistics()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		log.Printf("Error encoding statistics to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// DB QUERIES
func (s *Storage) UpdateStatistics() ([]Statistics, error) {
	url := "http://www.mfc-25.ru/queue/statistics.json"
	res, err := http.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot do request, error: %v\n", err))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot read response body, error: %v\n", err))
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	statistics := make([]Statistics, 0)

	err = json.Unmarshal([]byte(body), &statistics)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot unmarshal response, error: %v\n", err))
	}

	txn, err := s.db.Begin()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot begin database, error: %v\n", err))
	}

	stmt, err := txn.Prepare(pq.CopyIn("statistics", "id", "average_awaiting_time", "active_work_places_count", "completed_tickets_count",
		"pending_tikets_count"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot prepare the transaction, error: %v\n", err))
	}

	for _, sts := range statistics {
		_, err := stmt.Exec(sts.Id, sts.AverageAwaitingTime, sts.ActiveWorkPlacesCount, sts.CompletedTicketsCount, sts.PendingTicketsCount)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot execute statistics statement, error: %v\n", err))
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

	return statistics, nil
}