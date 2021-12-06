package swu

type Version struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	HTML      string `json:"html,omitempty"`
	Text      string `json:"text,omitempty"`
	Subject   string `json:"subject,omitempty"`
	Created   int64  `json:"created,omitempty"`
	Published bool   `json:"published,omitempty"`
}

func (api *Api) GetTemplateVersion(templateId, versionId string) (*Version, error) {
	resp := new(Version)
	err := api.get("/templates/"+templateId+"/versions/"+versionId, nil, nil, resp)
	return resp, err
}

func (api *Api) UpdateTemplateVersion(templateId, versionId string, version *Version) (*Version, error) {
	resp := new(Version)
	err := api.put("/templates/"+templateId+"/versions/"+versionId, version, nil, resp)
	return resp, err
}
