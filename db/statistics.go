package db

type Statistics struct {
	Id string
	AverageAwaitingTime int
	ActiveWorkPlacesCount int
	CompletedTicketsCount int
	PendingTicketsCount int
	Errors []*StatisticsError
	CompletedTickets []*CompletedTicket
}

type StatisticsError struct {
	Id string
	Key string
	Value string
}

type CompletedTicket struct {
	Hour int
	Load int
}
