package main

import "net/http"
func getPerson(id string) string{
	url := "HTTPS://myinfo.api.gov.sg/dev/L0/v1/person/" + id + "/"
	resp, err:= http.Get(url)

	if resp.Status == "200"{

	}else {
		return err.Error()
	}



}
