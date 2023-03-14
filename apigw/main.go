package apigw

import (
	"net/http"

	"github.com/JuanVF/personal_bot/common"
	"github.com/gorilla/mux"
)

var router *mux.Router = nil
var logger *common.Logger = common.GetLogger()

func Start() {
	router = mux.NewRouter()

	routers := []RouterHandler{
		AuthRoute{}, PaymentRouter{}, UserRouter{},
	}

	for _, r := range routers {
		r.Handle()
	}

	http.Handle("/", router)

	logger.Log("APIGW", "Server started on :3000")
	http.ListenAndServe(":3000", nil)
}
