package swu

type Sender struct {
	ReplyTo   string     `json:"reply_to,omitempty"`
	Recipient *Recipient `json:"recipient,omitempty"`
}
