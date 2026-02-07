package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const aurRPCURL = "https://aur.archlinux.org/rpc"

type AURResponse struct {
	ResultCount int          `json:"resultcount"`
	Results     []AURPackage `json:"results"`
	Type        string       `json:"type"`
	Version     int          `json:"version"`
}

type AURPackage struct {
	ID             int      `json:"ID"`
	Name           string   `json:"Name"`
	PackageBaseID  int      `json:"PackageBaseID"`
	PackageBase    string   `json:"PackageBase"`
	Version        string   `json:"Version"`
	Description    string   `json:"Description"`
	URL            string   `json:"URL"`
	NumVotes       int      `json:"NumVotes"`
	Popularity     float64  `json:"Popularity"`
	OutOfDate      *int     `json:"OutOfDate"`
	Maintainer     string   `json:"Maintainer"`
	FirstSubmitted int      `json:"FirstSubmitted"`
	LastModified   int      `json:"LastModified"`
	URLPath        string   `json:"URLPath"`
}

// getAURPackageInfo queries the AUR RPC API for package information
func getAURPackageInfo(pkgName string) (*AURPackage, error) {
	params := url.Values{}
	params.Add("v", "5")
	params.Add("type", "info")
	params.Add("arg[]", pkgName)
	
	resp, err := http.Get(aurRPCURL + "?" + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var aurResp AURResponse
	if err := json.NewDecoder(resp.Body).Decode(&aurResp); err != nil {
		return nil, err
	}
	
	if len(aurResp.Results) == 0 {
		return nil, nil
	}
	
	return &aurResp.Results[0], nil
}
