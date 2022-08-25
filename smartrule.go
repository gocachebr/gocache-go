package gocache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (api *API) CreateRules(domain string, rules []SmartRule, ruleType string) (*API_Response, error) {
	resp := API_Response{}

	reqBody, err := json.Marshal(rules)
	if err != nil {
		return nil, err
	}

	resp.Response = string(reqBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rules/bulk/%s/%s", api.HostURL, ruleType, domain), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new(SmartRuleResult)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		resp.Response = string(body)
		return &resp, err
	}

	resp.Response = *resp.Response.(*SmartRuleResult)

	return &resp, nil
}

func (api *API) UpdateRules(domain string, rules []SmartRule, id int, ruleType string) (*API_Response, error) {
	resp := API_Response{}

	reqBody, err := json.Marshal(rules)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/rules/bulk/%s/%s/%d", api.HostURL, ruleType, domain, id), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new(SmartRuleResult)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		resp.Response = string(body)
		return &resp, err
	}

	resp.Response = *resp.Response.(*SmartRuleResult)

	return &resp, nil
}

func (api *API) ApplyBulk(domain string, id int, ruleType string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rules/bulk/%s/%s/%d", api.HostURL, ruleType, domain, id), nil)
	if err != nil {
		return nil, err
	}

	// resp.Response = string(req.URL.Path)

	// return &resp, fmt.Errorf("aa")

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new([]SmartRule)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &resp, err
	}

	resp.Response = *resp.Response.(*[]SmartRule)

	return &resp, nil
}

func (api *API) GetRules(domain string, ruleType string) (*API_Response, error) {
	resp := API_Response{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rules/%s/%s", api.HostURL, ruleType, domain), nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new(SmartRuleResult)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		if strings.Contains(string(body), "\"rules\":{}") {
			resp.Response = *resp.Response.(*SmartRuleResult)

			return &resp, nil
		} else {
			return &resp, err
		}
	}

	resp.Response = *resp.Response.(*SmartRuleResult)

	return &resp, nil
}

func (api *API) DeleteBulk(domain string, id int, ruleType string) (*API_Response, error) {
	resp := API_Response{}

	var url string
	if id != -1 {
		url = fmt.Sprintf("%s/rules/bulk/%s/%s/%d", api.HostURL, ruleType, domain, id)
	} else {
		url = fmt.Sprintf("%s/rules/bulk/%s/%s", api.HostURL, ruleType, domain)
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	body, status, err := api.doRequest(req)
	resp.HTTPStatusCode = status

	if err != nil {
		return &resp, err
	}

	resp.Response = new(map[string]interface{})

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return &resp, err
	}

	resp.Response = *resp.Response.(*map[string]interface{})

	return &resp, nil
}
