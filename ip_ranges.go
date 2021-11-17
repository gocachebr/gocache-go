package gocache

import (
	"fmt"
	"net/http"
	"strings"
)

func (api *API) ListIpRanges() ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.gocache.com.br/ips"), nil)
	if err != nil {
		return nil, err
	}

	body, _, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}

	resp := make([]string,0)

	for _,line := range strings.Split(strings.TrimRight(string(body), "\n"), "\n"){
		resp = append(resp,line)
	}

	return resp, nil
}
