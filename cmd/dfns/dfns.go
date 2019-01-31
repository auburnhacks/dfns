package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/auburnhacks/dfns/canattend"
	"github.com/auburnhacks/dfns/events"
	"github.com/auburnhacks/dfns/updatestatus"
	"github.com/gorilla/mux"
)

var (
	addr *string
)

func init() {
	addr = flag.String("addr", "localhost:9000", "hostport")
	flag.Parse()
}

func main() {
	r := mux.NewRouter()
	r.Use(middleware)

	r.HandleFunc("/_readyz", readyz)
	r.HandleFunc("/_healthz", healthz)
	r.Handle("/update_status", &updatestatus.Handler{})
	r.Handle("/can_attend", &canattend.Handler{})
	r.Handle("/events", &events.Handler{})

	s := &http.Server{
		Handler: r,
		Addr:    *addr,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("server started")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("exiting server....")
}

func readyz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addHeader(&w)
		next.ServeHTTP(w, r)
	})
}

func addHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, "+
			"Accept-Encoding, X-CSRF-Token, Authorization")
}
