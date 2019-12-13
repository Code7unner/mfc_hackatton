package parser

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Service struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Duration int    `json:"duration"`
}

func parseExcel(path string) ([]Service, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open exel file, error: %v", err)
	}

	services := make([]Service, 0)

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		for _, subrow := range row {
			if subrow == "" {
				continue
			}

			ignoredWords := []string{"Наименование", "Государственное", "Раздел", "Перечень", "Управление", "Районные", "Департамент", "№"}
			if s := strings.Split(subrow, " ")[0]; stringInSlice(s, ignoredWords) {
				continue
			}

			if _, err := strconv.Atoi(strings.Split(subrow, "")[0]); err == nil {
				continue
			}

			if strings.Split(subrow, "")[0] == " " {
				subrow = strings.Join(strings.Split(subrow, "")[1:], "")
			}

			if strings.Split(subrow, "")[0] == "-" {
				subrow = strings.Join(strings.Split(subrow, "")[1:], "")
			}

			for strings.Split(subrow, "")[0] == " " || strings.Split(subrow, "")[0] == "\t" {
				subrow = strings.Join(strings.Split(subrow, "")[1:], "")
			}

			if strings.Contains(subrow, "   ") {
				continue
			}

			re := regexp.MustCompile(`(?s)\((.*)\)`)
			subrow := re.ReplaceAllString(subrow, "")

			services = append(services, Service{
				Id:       len(services),
				Text:     subrow,
				Duration: getRandomDuration(),
			})
		}
	}

	return services, nil
}

func Parse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s, err := parseExcel("assets/data.xlsx")
	if err != nil {
		log.Printf("Error parsing excel file, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(s)
	if err != nil {
		log.Printf("Error encoding services to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getRandomDuration() int {
	return rand.Intn(10*60-2*60) + 2*60
}
