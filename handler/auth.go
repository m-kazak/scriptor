package handler

import (
	_"fmt"
    "net/http"
    "time"

    "scriptor/config"
    "scriptor/crypt"
    "scriptor/model"
    "scriptor/handler/authenticate"
)


//Auth is used to authorize user in application
var Auth = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

    traceHTTPRequest(request)

    auth_err := config.Config.Errors.Auth

    user, err := authenticate.Authenticate(request)
    if err != nil {
        responseMessage(GeneralResponse{auth_err, err.Error()}, response, request)
        return
    }
    
    sessionToken, err := appAuth(user, request.UserAgent())
    if err != nil {
        responseMessage(GeneralResponse{auth_err, err.Error()}, response, request)
        return
    }
    
    exp_time := time.Now().Unix() + 86400
    jwtToken := crypt.JWTSign(crypt.JWTClaim{OID:user.OID, Exp: exp_time, Token:sessionToken})

    responseMessage(LoginResponse{config.Config.Errors.None, "", user.Name, user.Email, jwtToken}, response, request)
})

func appAuth(objUser model.User, strUserAgent string) (string, error) {

    sessionToken := crypt.GetSessionToken(objUser.OID)

    //TODO: SET USER IN DATABASE

    return sessionToken, nil
}