package db

type Server struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	IsConnected          bool   `json:"is_connected"`
	WrongProtocol        bool   `json:"wrong_protocol"`
	OrganizationName     string `json:"organization_name"`
	OrganizationFullName string `json:"organization_full_name,omitempty"`
	OrganizationAddress  string `json:"organization_address,omitempty"`
	OrganizationPhone    string `json:"organization_phone,omitempty"`
	OrganizationFax      string `json:"organization_fax,omitempty"`
	OrganizationEmail    string `json:"organization_email,omitempty"`
}
