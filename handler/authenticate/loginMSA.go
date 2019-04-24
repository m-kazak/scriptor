package authenticate

import (
    "net/http"
    "net/url"
    "errors"
    "io/ioutil"
    "strings"
    "time"

    "scriptor/config"
    "scriptor/crypt"
    "scriptor/model"

    "github.com/bitly/go-simplejson"
)

func loginMSA(strCode, strNonce string) (model.User, error){
    
    var user model.User
    
    params := url.Values{}
    params.Set("client_id", config.Config.App.MS.ClientId)
    params.Set("client_secret", config.Config.App.MS.ClientSecret)
    params.Set("redirect_uri", config.Config.App.MS.RedirectUri)
    params.Set("scope", config.Config.App.MS.Scope)
    params.Set("grant_type", config.Config.App.MS.GrantType)
    params.Set("code", strCode)
    
    resp, err := http.PostForm(config.Config.App.MS.AuthURL, params)
    
    if err != nil {
        return user, errors.New("Can't send request for getting access token.")
    }

    body, _ := ioutil.ReadAll(resp.Body)
    json, _ := simplejson.NewJson(body)

    idToken, err := json.Get("id_token").String()
    if err != nil {
        errType, _ := json.Get("error").String()
        errDesc, _ := json.Get("error_description").String()
        return user, errors.New(errType + " : " + errDesc)
    }

    oid, username, name, email, aud, exp, nonce := parseJWT(idToken)

    if strNonce != nonce {
        return user, errors.New("Not valid nonce parameter.")
    }

    if exp < time.Now().Unix() {
        return user, errors.New("MSA token expired.")
    }

    user = model.User {0, "ms", username, name, oid, email, aud, 1, "", ""}

    return user, nil
}

func parseJWT(strJWT string) (string, string, string, string, string, int64, string) {

    parts := strings.Split(strJWT, ".")
    
    btJwt, _ := crypt.DecodeSegment(parts[1])

    json, _ := simplejson.NewJson(btJwt)

    oid, _ := json.Get("oid").String()
    username, _ := json.Get("preferred_username").String()
    name, _ := json.Get("name").String()
    email, _ := json.Get("email").String() 
    aud, _ := json.Get("aud").String()
    exp, _ := json.Get("exp").Int64()
    nonce, _ := json.Get("nonce").String()

    config.Loggy.Debugf("<MSA Auth> OID:[%s]; UserName:[%s]; Name:[%s]; EMail:[%s]; AUD:[%s]; Exp:[%v]; Nonce:[%s]", oid, username, name, email, aud, exp, nonce)

    return oid, username, name, email, aud, exp, nonce
}