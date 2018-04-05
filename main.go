package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/an-jun/go-wechat/wx"
)

const (
	logLevel = "dev"
	port     = 1234
	token    = "d027us7DzEzod7So0ddTE27KZyEUZoFo"
)

func get(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}

	if len(client.Query.Echostr) > 0 {
		w.Write([]byte(client.Query.Echostr))
		return
	}

	w.WriteHeader(403)
	return
}

func post(w http.ResponseWriter, r *http.Request) {

	client, err := wx.NewClient(r, w, token)

	if err != nil {
		log.Println(err)
		w.WriteHeader(403)
		return
	}

	client.Run()
	return
}

func main() {
	handler := &HttpHandler{}
	server := http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 0,
	}

	log.Println(fmt.Sprintf("Listen: %d", port))
	log.Fatal(server.ListenAndServe())
}
