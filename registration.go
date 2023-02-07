package types

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
