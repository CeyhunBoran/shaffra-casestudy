package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CeyhunBoran/shaffra-casestudy/internal/config"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/handlers"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/repositories"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/services"
	"github.com/CeyhunBoran/shaffra-casestudy/internal/utils"
	"github.com/CeyhunBoran/shaffra-casestudy/pkg/logging"
	"github.com/gorilla/mux"
)

func main() {
	config.InitConfig()

	conf := config.Conf
	db, err := utils.NewDB(
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			conf.DbHost, conf.DbUser, conf.DbPass, conf.DbName, conf.DbPort, conf.DbSsl))
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	userService := services.NewUserService(*repositories.NewUserRepository(db))

	r := mux.NewRouter()
	userRoutes := r.PathPrefix("/api/users").Subrouter()
	userRoutes.Use(logging.LoggingMiddleware)
	handler := handlers.NewUserHandler(*userService)

	userRoutes.HandleFunc("", handler.CreateUser).Methods("POST")
	userRoutes.HandleFunc("/{id}", handler.GetUser).Methods("GET")
	userRoutes.HandleFunc("/{id}", handler.UpdateUser).Methods("PUT")
	userRoutes.HandleFunc("/{id}", handler.DeleteUser).Methods("DELETE")

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
