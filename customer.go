package swu

type CustomerResponse struct {
	Response
	Customer *Customer `json:"customer,omitempty"`
}

type Customer struct {
	Object  string `json:"object,omitempty"`
	Email   string `json:"email,omitempty"`
	Created int64  `json:"created,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

func (api *Api) GetCustomer(email string) (*CustomerResponse, error) {
	resp := new(CustomerResponse)
	err := api.get("/customers/"+email, nil, nil, resp)
	return resp, err
}

func (api *Api) SaveCustomer(customer *Customer) (*Response, error) {
	resp := new(Response)
	err := api.post("/customers", customer, nil, resp)
	return resp, err
}

func (api *Api) DeleteCustomer(email string) (*Response, error) {
	resp := new(Response)
	err := api.delete("/customers/"+email, nil, nil, resp)
	return resp, err
}
