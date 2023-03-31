/*
Copyright 2023 Juan Jose Vargas Fletes

This work is licensed under the Creative Commons Attribution-NonCommercial (CC BY-NC) license.
To view a copy of this license, visit https://creativecommons.org/licenses/by-nc/4.0/

Under the CC BY-NC license, you are free to:

- Share: copy and redistribute the material in any medium or format
- Adapt: remix, transform, and build upon the material

Under the following terms:

  - Attribution: You must give appropriate credit, provide a link to the license, and indicate if changes were made.
    You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.

- Non-Commercial: You may not use the material for commercial purposes.

You are free to use this work for personal or non-commercial purposes.
If you would like to use this work for commercial purposes, please contact Juan Jose Vargas Fletes at juanvfletes@gmail.com.
*/
package apigw

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/gorilla/mux"
)

var router *mux.Router = nil
var logger *common.Logger = common.GetLogger()

type Server interface {
	Serve()
}

type HttpsServer struct {
}

type HttpServer struct {
}

// Serve an HTTPS Server
func (serv HttpsServer) Serve() {
	router = mux.NewRouter()

	HandleAllRoutes(router)

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.CurveP384,
				tls.CurveP521,
			},
		},
		Handler: router,
	}

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	if err != nil {
		log.Fatal(err)
	}

	server.TLSConfig.Certificates = []tls.Certificate{cert}

	log.Fatal(server.ListenAndServeTLS("", ""))
}

// Serve an HTTP Server
func (serv HttpServer) Serve() {
	router = mux.NewRouter()

	HandleAllRoutes(router)

	http.Handle("/", router)

	logger.Log("APIGW", "Server started on :443")
	http.ListenAndServe(":443", nil)
}

// Handle All Routes
func HandleAllRoutes(r *mux.Router) {
	routers := []RouterHandler{
		AuthRoute{}, PaymentRouter{}, UserRouter{}, UserHealthRouter{}, UserFitnessRouter{},
	}

	for _, r := range routers {
		r.Handle()
	}
}

func Start() {
	// Define a map of server types with their corresponding instances
	servers := map[string]Server{
		"development": HttpServer{},
		"container":   HttpsServer{},
	}

	// Retrieve the server instance based on the environment
	server, found := servers[common.GetEnvironment()]

	// If the server instance is found, serve the server
	if found {
		server.Serve()
	}
}
