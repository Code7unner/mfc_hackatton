package db

type Statistics struct {
	Id                    string             `json:"Id"`
	AverageAwaitingTime   float64            `json:"averageAwaitingTime"`
	ActiveWorkPlacesCount int                `json:"activeWorkPlacesCount"`
	CompletedTicketsCount int                `json:"completedTicketsCount"`
	PendingTicketsCount   int                `json:"pendingTicketsCount"`
	CompletedTickets      []*CompletedTicket `json:"completedTickets"`
}

type CompletedTicket struct {
	ParentId string `json:"parentId"`
	Hour     int    `json:"hour"`
	Load     int    `json:"load"`
}
