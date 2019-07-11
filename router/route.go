package router

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jou66jou/go-p2p-websocket/common"
	"github.com/jou66jou/go-p2p-websocket/handler"
)

func RunHTTP(port string) error {
	mux := makeMuxRouter()
	httpAddr := port
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func makeMuxRouter() http.Handler {
	rName := common.RouteName
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc(rName["getAllPeers"], handler.GetPeers).Methods("GET")
	muxRouter.HandleFunc(rName["newWS"], handler.NewWS)
	return muxRouter
}
