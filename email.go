package swu

type Email struct {
	Id          string                 `json:"email_id,omitempty"`
	Recipient   *Recipient             `json:"recipient,omitempty"`
	CC          []*Recipient           `json:"cc,omitempty"`
	BCC         []*Recipient           `json:"bcc,omitempty"`
	Sender      *Sender                `json:"sender,omitempty"`
	EmailData   map[string]interface{} `json:"email_data,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Inline      *Attachment            `json:"inline,omitempty"`
	Files       []*Attachment          `json:"files,omitempty"`
	ESPAccount  string                 `json:"esp_account,omitempty"`
	VersionName string                 `json:"version_name,omitempty"`
}

type Attachment struct {
	Id   string `json:"id,omitempty"`
	Data string `json:"data,omitempty"`
}

func (api *Api) Emails() (templates []Template, err error) {
	err = api.get("/templates", nil, nil, &templates)
	return
}

func (api *Api) Send(email *Email) (*LogSend, error) {
	resp := new(LogSend)
	err := api.post("/send", email, nil, resp)
	return resp, err
}
