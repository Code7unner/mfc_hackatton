package scheduler

import (
	"github.com/mfc_hackatton/db"
	"log"
	"time"
)

func Schedule(st *db.Storage) {
	time.Sleep(1 * time.Second)
	for true {
		if _, err := st.UpdateServerStats(); err != nil {
			log.Printf("Cannot get any response from /api/server, error: %v", err)
		}
		if _, err := st.UpdateStatistics(); err != nil {
			log.Printf("Cannot get any response from /api/statistics, error: %v", err)
		}
		log.Println("Database has just been updated.")
		time.Sleep(15 * time.Second)
	}
}
