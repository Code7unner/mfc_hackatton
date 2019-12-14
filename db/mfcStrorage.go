package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func (s *Storage) GetMFCStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	mfcs, err := s.UpdateMFCStats()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(mfcs)
	if err != nil {
		log.Printf("Error encoding MFCs to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Storage) UpdateMFCStats() ([]MFC, error) {
	rows, err := s.db.Query("SELECT DISTINCT a.id, a.organization_address, b.pending_tickets_count, b.completed_tickets_count " +
		"FROM servers AS a " +
		"INNER JOIN statistics AS b " +
		"ON a.id = b.id")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot send query to database, error: %v", err))
	}

	var mfcs []MFC
	for rows.Next() {
		mfc := MFC{}
		err := rows.Scan(&mfc.Id, &mfc.Address, &mfc.PendingTicketsCount, &mfc.CompletedTicketsCount)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Cannot scan row, error: %v", err))
		}
		mfcs = append(mfcs, mfc)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("Got an error: %v", err))
	}

	return mfcs, nil
}