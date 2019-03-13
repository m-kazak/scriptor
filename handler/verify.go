package handler

import (
	"net/http"

	"scriptor/config"
	"scriptor/crypt"
)

//VerifySession is used to verify token for give access to Angular's modules 
var Verify = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	traceHTTPRequest(request)
	
	session := request.FormValue("session")
	if session == "" {
		responseMessage(GeneralResponse{config.Config.Errors.ModuleAccess, "No session was passed"}, response, request)
        return
	}
	
	err := crypt.JWTVerify(session)
	if err != nil {
		config.Loggy.Infoln(err.Error())
		responseMessage(GeneralResponse{config.Config.Errors.ModuleAccess, err.Error()}, response, request)
        return
	} 
	responseMessage(GeneralResponse{config.Config.Errors.None, ""}, response, request)
})