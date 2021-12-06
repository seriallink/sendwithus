package swu

type DripCampaign struct {
	Recipient  *Recipient        `json:"recipient,omitempty"`
	CC         []*Recipient      `json:"cc,omitempty"`
	BCC        []*Recipient      `json:"bcc,omitempty"`
	Sender     *Sender           `json:"sender,omitempty"`
	EmailData  map[string]string `json:"email_data,omitempty"`
	Tags       []string          `json:"tags,omitempty"`
	ESPAccount string            `json:"esp_account,omitempty"`
	Locale     string            `json:"locale,omitempty"`
}

func (api *Api) ActivateDripCampaign(campaignId string, campaign *DripCampaign) error {
	err := api.post("/drip_campaigns/"+campaignId+"/activate", campaign, nil, nil)
	return err
}
