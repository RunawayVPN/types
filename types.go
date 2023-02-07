package types

type AgentRequest struct {
	Jwt string
}

type AgentResponse struct {
	Jwt string
}

type AddConfigRequest struct {
	Hostname    string
	Owner       string
	Description string
}

type RegistrationRequest struct {
	PublicKey string `json:"public_key"`
	SecretKey string `json:"secret_key"`
	Agent     Agent  `json:"agent"`
	AuthToken string `json:"auth_token"`
}

type RegistrationResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	PublicKey string `json:"public_key"`
	AuthToken string `json:"auth_token"`
}

type AuthToken struct {
	Endpoint string
	Roles    []string
}
type Agent struct {
	Name      string `json:"name"`
	Identity  string `json:"identity"`
	PublicKey string `json:"public_key"`
	PublicIP  string `json:"public_ip"`
	Country   string `json:"country"`
	ISP       string `json:"isp"`
}
