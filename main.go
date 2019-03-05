package main

import (
	"net/http"
	"flag"
	"fmt"
	"time"

	"scriptor/handler"
	"scriptor/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	option = flag.String("o", "", "Choose option: start")
)

func init() {
	config.LoadConfiguration("./config.cfg")
	config.LoadLogger(config.Config.Logger.File, config.Config.Logger.Flag, config.Config.Logger.Level)
	config.Loggy.Debug("qweS")
}

func main() {
	flag.Parse()
	
	switch op := *option; op {
	case "start":
		startServer()
	default:
		fmt.Println("You need to choose option")
	}
}

func startServer() {
	hnd := mux.NewRouter()

	hnd.Handle("/", handler.StatusHandler).Methods("GET")

	srv := &http.Server{
		Addr:           ":8000",
		Handler:        handlers.CORS()(hnd),
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	srv.ListenAndServe()
}