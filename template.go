package swu

type Template struct {
	Id       string     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Tags     []string   `json:"tags,omitempty"`
	Created  int64      `json:"created,omitempty"`
	Versions []*Version `json:"versions,omitempty"`
}

func (api *Api) Templates() ([]Template, error) {
	return api.Emails()
}

func (api *Api) GetTemplate(templateId string) (*Template, error) {
	resp := new(Template)
	err := api.get("/templates/"+templateId, nil, nil, resp)
	return resp, err
}

func (api *Api) CreateTemplate(version *Version) (*Template, error) {
	resp := new(Template)
	err := api.post("/templates", version, nil, resp)
	return resp, err
}

func (api *Api) CreateTemplateVersion(templateId string, version *Version) (*Template, error) {
	resp := new(Template)
	err := api.post("/templates/"+templateId+"/versions", version, nil, resp)
	return resp, err
}
