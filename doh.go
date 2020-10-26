package madoh

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"

	madns "github.com/multiformats/go-multiaddr-dns"
)

var DefaultDOHResolver = &madns.Resolver{Backend: &DOHResolver{
	Host:   "https://dns.google/resolve",
	Client: http.DefaultClient,
}}

type DOHResolver struct {
	Host   string
	Client *http.Client
}

type jsonResponse struct {
	Answer []struct {
		Type int
		Data string
	}
}

// Perform the DNS over HTTPS lookup using JSON interface
func (r *DOHResolver) lookup(name string, type_ int) ([]string, error) {
	u, err := url.Parse(r.Host)
	if err != nil {
		panic(err)
	}
	u.RawQuery = fmt.Sprintf(
		"name=%s&type=%d",
		url.QueryEscape(name),
		type_,
	)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/dns-json")
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Got HTTP error from resolver: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response jsonResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	answers := make([]string, 0, len(response.Answer))
	for _, a := range response.Answer {
		if a.Type == type_ {
			answers = append(answers, a.Data)
		}
	}
	return answers, nil
}

func (r *DOHResolver) LookupIPAddr(_ context.Context, host string) ([]net.IPAddr, error) {
	records, err := r.lookup(host, 1)
	if err != nil {
		return nil, err
	}
	result := make([]net.IPAddr, 0, len(records))
	for _, r := range records {
		result = append(result, net.IPAddr{
			IP:   net.ParseIP(r),
			Zone: "",
		})
	}
	return result, nil
}

func (r *DOHResolver) LookupTXT(_ context.Context, host string) ([]string, error) {
	records, err := r.lookup(host, 16)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0, len(records))
	for _, r := range records {
		results = append(results, strings.Trim(r, "\""))
	}
	return results, nil
}
