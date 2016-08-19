package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/shinmyung0/loglite"
)

type logwrapper struct {
	router *mux.Router
}

func (l *logwrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Infof("%s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
	l.router.ServeHTTP(w, r)
}

const serverPort = "8080"

func ListenAndServe() {
	r := mux.NewRouter()
	l := &logwrapper{r}
	v1ApiControllerInit(r)

	log.Info("Running server on port " + serverPort)
	http.Handle("/", l)
	http.ListenAndServe(":"+serverPort, nil)
}
