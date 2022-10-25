package main

import (
	"aset-go/controllers/asetcontroller"
	"aset-go/controllers/authcontroller"
	"aset-go/middlewares"
	"aset-go/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	api.HandleFunc("/asets", asetcontroller.Index).Methods("GET")
	api.HandleFunc("/aset/{id}", asetcontroller.Show).Methods("GET")
	api.HandleFunc("/aset", asetcontroller.Create).Methods("POST")
	api.HandleFunc("/aset/{id}", asetcontroller.Update).Methods("PUT")
	api.HandleFunc("/aset", asetcontroller.Delete).Methods("DELETE")
	api.HandleFunc("/file", asetcontroller.UploadFile).Methods("POST")
	api.Use(middlewares.JWTMiddleware)
	log.Fatal(http.ListenAndServe(":8080", r))
}
