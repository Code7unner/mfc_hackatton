package db

type Server struct {
	Id string
	Name string
	IsConnected bool
	WrongProtocol bool
	OrganizationName string
	OrganizationFullName string
	OrganizationAddress string
	OrganizationPhone string
	OrganizationFax string
	OrganizationEmail string
}
