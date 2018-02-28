package models

import (
	"encoding/json"
	"fmt"
	"os"
)

type ECManifest struct {
    ReqId    int    	`json:"reqId"`
	Title    string 	`json:"title"`
	Command  []string	`json:"command"`
	Baseline string 	`json:"baseline"`
}

type ECManifestResult struct {
	ECManifest
    Output   string `json:"output"`
    DateExe  string `json:"dateExe"`
}

type ECResult struct {
	ECManifest
	HostExec		string `json:"host"`
	StdOutput		[]string `json:"stdOutput"`
	StdErrOutput	[]string `json:"stdErrOutput"`
	DateExe			string `json:"dateExe"`
}

func (p ECManifest) ToString() string {
	return ToJson(p)
}

func ToJson(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}