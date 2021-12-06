package swu

type Response struct {
	Success bool   `json:"success,omitempty"`
	Status  string `json:"status,omitempty"`
}
