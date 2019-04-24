package config

import (
	"encoding/json"
	"fmt"
	"os"
	"crypto/rand"
	"io"

	"github.com/m-kazak/loggy"
)

//Configuration setup of application
type Configuration struct {
	Server struct {
		TraceHTTP	    int `json:"traceHTTP"`
		SecretBytes	    int `json:"secret_byte"`
		SecretSalt	    string `json:"secret_salt"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Dbname   string `json:"dbname"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
	App struct {
		MS struct {
			AuthURL      string `json:"auth_url"`
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
			RedirectUri  string `json:"redirect_uri"`
			Scope		 string `json:"scope"`
			GrantType	 string `json:"grant_type"`
		} `json:"ms"`
	} `json:"app"`
	Logger struct {
		File	    string  `json:"file"`
		Flag	    int 	`json:"flag"`
		Level	    int 	`json:"level"`
	} `json:"logger"`
	Errors struct {
		None			int `json:"none"`
		Auth			int `json:"auth"`
		ModuleAccess	int `json:"module_access"`
		Logout			int `json:"logout"`
		Session			int `json:"session"`
		Trace			int `json:"trace"`
		UserInfo		int `json:"user_info"`
		LicenseInfo		int `json:"license_info"`
	} `json:"errors"`
}

//Config is a global variable
var Config Configuration

//LoadConfiguration setup config
func LoadConfiguration(file string) {

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		fmt.Println("Can't load configuration file:" + err.Error())
		os.Exit(1)
	}

	if Config.Server.SecretSalt == "" {
		if Config.Server.SecretBytes == 0 {
			Config.Server.SecretSalt = generateJWTSalt(32)
		} else {
			Config.Server.SecretSalt = generateJWTSalt(Config.Server.SecretBytes)
		}
	}
	
	fmt.Println("Configuration loaded")
}

func generateJWTSalt(bytes int) string {
	key := make([]byte, bytes)
	rand.Read(key)
	
	return fmt.Sprintf("%x", key)
}

//Logger for application
var Loggy *loggy.Logger

//LoadLogger setup logger Loggy
func LoadLogger(file string, flag int, initLevel int) {
	
	var writer io.Writer

	if file != "" {
		logFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Logging in terminal")
			writer = os.Stderr
		} else {
			writer = logFile
			io.WriteString(writer, "\r\nLog file is setup correctly\r\n")
		}
	} else {
		writer = os.Stderr
	}

	lvl := loggy.LogLevel(initLevel)
	Loggy = loggy.New(writer, flag, lvl)
}