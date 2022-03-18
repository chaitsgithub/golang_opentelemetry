package main

import (
	"log"
	"net/http"

	"github.com/chaitsgithub/golang_opentelemetry/go-middleware-metrics/middleware"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := mux.NewRouter()

	metricsMiddleware := middleware.NewMetricsMiddleware()

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/lemon", lemonHandler).Methods(http.MethodGet)
	r.HandleFunc("/potato", potatoHandler).Methods(http.MethodPost)

	r.Use(metricsMiddleware.Metrics)

	http.ListenAndServe(":8080", r)
}

func lemonHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In lemonHandler")
	w.WriteHeader(200)
	w.Write([]byte("Lemon"))
}

func potatoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("In potatoHandler")
	w.WriteHeader(200)
	w.Write([]byte("Potato"))
}
