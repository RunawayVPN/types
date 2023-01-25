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
