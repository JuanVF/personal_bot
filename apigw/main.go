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
		AuthRoute{}, PaymentRouter{}, UserRouter{},
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
