package db

type Server struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	IsConnected          bool   `json:"isConnected"`
	WrongProtocol        bool   `json:"wrongProtocol"`
	OrganizationName     string `json:"organizationName,omitempty"`
	OrganizationFullName string `json:"organizationFullName,omitempty"`
	OrganizationAddress  string `json:"organizationAddress,omitempty"`
	OrganizationPhone    string `json:"organizationPhone,omitempty"`
	OrganizationFax      string `json:"organizationFax,omitempty"`
	OrganizationEmail    string `json:"organizationEmail,omitempty"`
}
