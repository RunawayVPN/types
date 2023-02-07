package types

type Agent struct {
	Name      string `json:"name"`
	Identity  string `json:"identity"`
	PublicKey string `json:"public_key"`
	PublicIP  string `json:"public_ip"`
	Country   string `json:"country"`
	ISP       string `json:"isp"`
}

type HubInfo struct {
	PublicKey string `json:"public_key"`
	AgentJwt  string `json:"agent_jwt"`
}
