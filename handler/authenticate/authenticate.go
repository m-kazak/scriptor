package authenticate

import (
	_"fmt"
    "net/http"
    "errors"

    "scriptor/model"
)


//Authenticate is used to authenticate user based on provider
func Authenticate (request *http.Request) (model.User, error){

    var user model.User
    var err error

    request.ParseForm()

    provider := request.FormValue("provider")
    if provider == "" {
        return user, errors.New("No authentication provider was passed.")
    }

    code := request.FormValue("code")
    if code == "" {
        return user, errors.New("No code was passed.")
    }
    
    nonce := request.FormValue("nonce")
    if code == "" {
        return user, errors.New("No nonce was passed.")
    }
    
    switch provider {
	case "ms":
        user, err = loginMSA(code, nonce)
        if err != nil {
            return user, err
        }
    default:
		return user, errors.New("No such authentication provider.")
    }

    return user, nil
}