package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ayushi-khandal09/carZone/driver"
	carHandler "github.com/ayushi-khandal09/carZone/handler/car"
	engineHandler "github.com/ayushi-khandal09/carZone/handler/engine"
	carService "github.com/ayushi-khandal09/carZone/service/car"
	engineService "github.com/ayushi-khandal09/carZone/service/engine"
	carStore "github.com/ayushi-khandal09/carZone/store/car"
	engineStore "github.com/ayushi-khandal09/carZone/store/engine"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default environment variables")
	}

	// Initialize database connection
	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	if db == nil {
		log.Fatal("Database connection failed")
	}

	// Initialize services & handlers
	carStorage := carStore.New(db)
	carService := carService.NewCarService(carStorage)
	engineStorage := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStorage)
	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	// Execute schema
	schemaFile := "store/schema.sql"
	if err := executeSchemaFile(db, schemaFile); err != nil {
		log.Fatalf("Error while executing the schema file: %v", err)
	}

	// Set up routes
	router := mux.NewRouter()
	// Debugging: Print registered routes
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		log.Printf("Registered route: %s %v", path, methods)
		return nil
	})

	router.HandleFunc("/cars/{id}", carHandler.GetCarById).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST") 
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandler.GetEngineById).Methods("GET")
	router.HandleFunc("/engine", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

// Execute schema file for database migrations
func executeSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
