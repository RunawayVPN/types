package types

type Agent struct {
	Name      string `json:"name"`
	Identity  string `json:"identity"`
	PublicKey string `json:"public_key"`
	PublicIP  string `json:"public_ip"`
	Country   string `json:"country"`
	ISP       string `json:"isp"`
}

type RegistrationRequest struct {
	PublicKey string `json:"public_key"`
	SecretKey string `json:"secret_key"`
	JwtToken  string `json:"jwt"`
}

type RegistrationResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	PublicKey string `json:"public_key"`
	JwtToken  string `json:"jwt"`
}

type HubInfo struct {
	PublicKey string `json:"public_key"`
	AgentJwt  string `json:"agent_jwt"`
}
