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
