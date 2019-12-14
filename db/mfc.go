package db

type MFC struct {
	Id string `'json:"id"`
	Address string `json:"address"`
	PendingTicketsCount int `json:"pendingTicketsCount"`
	CompletedTicketsCount int `json:"completedTicketsCount"`
}