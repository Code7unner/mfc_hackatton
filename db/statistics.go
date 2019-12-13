package db

type Statistics struct {
	Id                    string             `json:"Id"`
	AverageAwaitingTime   int                `json:"averageAwaitingTime"`
	ActiveWorkPlacesCount int                `json:"activeWorkPlacesCount"`
	CompletedTicketsCount int                `json:"completedTicketsCount"`
	PendingTicketsCount   int                `json:"pendingTicketsCount"`
	Errors                []*StatisticsError `json:"errors"`
	CompletedTickets      []*CompletedTicket `json:"completedTickets"`
}

type StatisticsError struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CompletedTicket struct {
	Hour int `json:"hour"`
	Load int `json:"load"`
}
