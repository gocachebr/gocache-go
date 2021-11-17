package gocache

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *API) GetDNSSEC(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dns/dnssec/%s", api.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}

		if status == 200 {
			resp.Response = resp.Response.(map[string]interface{})
		}
	}

	return &resp, nil
}

func (api *API) CreateDNSSEC(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dns/dnssec/%s", api.HostURL, domain), nil)
	if err != nil {
		return &resp, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {


		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = resp.Response.(map[string]interface{})

	}

	return &resp, nil
}

func (api *API) DeleteDNSSEC(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/dns/dnssec/%s", api.HostURL, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return &resp, err
	}

	return &resp, nil
}
