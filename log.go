package swu

type Log struct {
	LogEvent
	Id               string `json:"id,omitempty"`
	RecipientName    string `json:"recipient_name,omitempty"`
	RecipientAddress string `json:"recipient_address,omitempty"`
	Status           string `json:"status,omitempty"`
	EmailID          string `json:"email_id,omitempty"`
	EmailName        string `json:"email_name,omitempty"`
	EmailVersion     string `json:"email_version,omitempty"`
	EventsURL        string `json:"events_url,omitempty"`
}

type LogEvent struct {
	Object  string `json:"object,omitempty"`
	Created int64  `json:"created,omitempty"`
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}

type LogQuery struct {
	Count      int   `json:"count,omitempty"`
	Offset     int   `json:"offset,omitempty"`
	CreatedGT  int64 `json:"created_gt,omitempty"`
	CreatedGTE int64 `json:"created_gte,omitempty"`
	CreatedLT  int64 `json:"created_lt,omitempty"`
	CreatedLTE int64 `json:"created_lte,omitempty"`
}

type LogSend struct {
	Id        string    `json:"log_id,omitempty"`
	Success   bool      `json:"success,omitempty"`
	Status    string    `json:"status,omitempty"`
	ReceiptId string    `json:"receipt_id,omitempty"`
	Email     *LogEmail `json:"email,omitempty"`
}

type LogEmail struct {
	Name        string `json:"name,omitempty"`
	Locale      string `json:"locale,omitempty"`
	VersionName string `json:"version_name,omitempty"`
}

type LogCustomer struct {
	Response
	Logs []Log `json:"logs,omitempty"`
}

func (api *Api) GetLogs(query *LogQuery) (logs []Log, err error) {
	err = api.get("/logs", query, nil, &logs)
	return
}

func (api *Api) GetLog(logId string) (*Log, error) {
	resp := new(Log)
	err := api.get("/logs/"+logId, nil, nil, resp)
	return resp, err
}

func (api *Api) GetLogEvents(logId string) (*LogEvent, error) {
	resp := new(LogEvent)
	err := api.get("/logs/"+logId+"/events", nil, nil, resp)
	return resp, err
}

func (api *Api) GetCustomerLogs(email string, query *LogQuery) (*LogCustomer, error) {
	resp := new(LogCustomer)
	err := api.get("/customers/"+email+"/logs", query, nil, resp)
	return resp, err
}

func (api *Api) ResendLog(logId string) (*LogSend, error) {
	resp := new(LogSend)
	err := api.post("/resend", map[string]string{"log_id": logId}, nil, resp)
	return resp, err
}
