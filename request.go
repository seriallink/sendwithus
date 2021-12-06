package swu

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	Endpoint        = "https://api.sendwithus.com/api/v1"
	APIHeaderClient = "golang-0.0.1"
)

// Map data por post in the request
type Params map[string]interface{}

// Map extra request headers
type Headers map[string]string

// Make request and return the response
func (api *Api) execute(method string, path string, params interface{}, headers Headers, model interface{}) error {

	// init vars
	var url = Endpoint + path

	// init an empty payload
	payload := strings.NewReader("")

	// check for params
	if params != nil {

		// marshal params
		b, err := json.Marshal(params)
		if err != nil {
			return err
		}

		// set payload with params
		payload = strings.NewReader(string(b))

	}

	// set request
	request, _ := http.NewRequest(method, url, payload)
	request.SetBasicAuth(api.key, "")
	request.Header.Set("X-SWU-API-CLIENT", APIHeaderClient)
	request.Header.Add("ACCEPT", "application/json")
	request.Header.Add("CONTENT-TYPE", "application/json")

	// add extra headers
	if headers != nil {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// read response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// to test if the response body is an error
	erm := &ErrorMessage{}

	// check for error message
	if err = json.Unmarshal(data, erm); err == nil && erm.Code != 0 {
		return erm
	}

	// verify status code
	if NotIn(response.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent, http.StatusContinue) {

		if len(data) > 0 {
			return errors.New(string(data))
		}

		return errors.New(response.Status)

	}

	// some services have empty response
	if len(data) == 0 {
		return nil
	}

	// parse data
	return json.Unmarshal(data, model)

}

// Execute GET requests
func (api *Api) get(path string, params interface{}, headers Headers, model interface{}) error {
	return api.execute(http.MethodGet, path, params, headers, model)
}

// Execute POST requests
func (api *Api) post(path string, params interface{}, headers Headers, model interface{}) error {
	return api.execute(http.MethodPost, path, params, headers, model)
}

// Execute PUT requests
func (api *Api) put(path string, params interface{}, headers Headers, model interface{}) error {
	return api.execute(http.MethodPut, path, params, headers, model)
}

// Execute PATCH requests
func (api *Api) patch(path string, params interface{}, headers Headers, model interface{}) error {
	return api.execute(http.MethodPatch, path, params, headers, model)
}

// Execute DELETE requests
func (api *Api) delete(path string, params interface{}, headers Headers, model interface{}) error {
	return api.execute(http.MethodDelete, path, params, headers, model)
}
