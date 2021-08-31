package gocache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (api *API) ActivateCertificate(domain string, certType string, enable bool) (*API_Response, error) {
	resp := API_Response{}

	jsonBody := map[string]CertificateInput{
		"certificate": {
			Type:   certType,
			Enable: enable,
		},
	}

	reqBody, err := json.Marshal(jsonBody)
	if err != nil {
		return &resp, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/ssl/certificate/%s", api.HostURL, domain), strings.NewReader(string(reqBody)))
	if err != nil {
		return &resp, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {
		resp.Response = new(CertificateResult)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = *resp.Response.(*CertificateResult)
	}

	return &resp, nil
}

func (api *API) ListCertificates(domain string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ssl/certificate/%s", api.HostURL, domain), nil)
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

		resp.Response = new(CertificateResult)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = *resp.Response.(*CertificateResult)

	}

	return &resp, nil
}

func (api *API) CreateCertificate(domain string, certificate string, privatekey string, enable bool) (*API_Response, error) {

	resp := API_Response{}

	jsonBody := CertificateInput{
		Certificate: certificate,
		Privatekey:  privatekey,
		Enable:      enable,
	}

	reqBody, err := json.Marshal(jsonBody)

	if err != nil {
		return &resp, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ssl/certificate/%s", api.HostURL, domain), strings.NewReader(string(reqBody)))
	if err != nil {
		return &resp, err
	}

	req.Header.Add("Content-Type", "application/json")

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status
	if err != nil {

		jsonErr := json.Unmarshal(body, &resp)

		if jsonErr != nil {
			return &resp, err
		}

	} else {

		resp.Response = new(CertificateResult)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = *resp.Response.(*CertificateResult)

	}

	return &resp, nil
}

func (api *API) DeleteCertificate(domain string) (*API_Response, error) {

	resp := API_Response{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/ssl/certificate/%s", api.HostURL, domain), nil)
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

		resp.Response = new(CertificateResult)

		err = json.Unmarshal(body, &resp)
		if err != nil {
			return &resp, err
		}

		resp.Response = *resp.Response.(*CertificateResult)

	}

	return &resp, nil
}
