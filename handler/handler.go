package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	_"errors"
	"strings"

	"scriptor/config"
)

//StatusHandler is used for check server state
var StatusHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	_, err := traceHTTPRequest(request)
	if err != nil {
		responseMessage(GeneralResponse{config.Config.Errors.Trace, err.Error()}, response, request)
        return
	}
	responseMessage(GeneralResponse{config.Config.Errors.None, ""}, response, request)
	return
})

//responseMessage response to requests
func responseMessage(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

//traceHTTPRequest trace HTTP communication
func traceHTTPRequest(r *http.Request) (int, error) {
	
	if (config.Config.Server.TraceHTTP == 1) {

		fmt.Println(r.Method, r.URL, r.Proto)
		fmt.Println("Host: ", r.Host)
		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				fmt.Println(name, h)
			}
		}
			if r.Method == "POST" {
			r.ParseForm()
			fmt.Println(r.Form.Encode())
		}
	}
	
	return 1, nil
}