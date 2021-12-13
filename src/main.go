package main

import (
	"fmt"
	env "github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
)



func  (app App) loadEnv(path string) string {
	err := env.Load(path)
	if err != nil {
		app.logger.Fatal(err)
	}
	port:="8080"
	definedPort, exist := os.LookupEnv("PORT")
	if exist {
		port= definedPort
	}else{
		app.logger.Print("Default  ")
	}
	app.logger.Printf("port : %s ",port)
	return  fmt.Sprintf(":%s", port)
}

func  main() {
	fmt.Println("Running")
	logger:=log.New(os.Stdout,"",log.Ldate|log.Llongfile)
	app:=&App{
		logger:logger,
	}
	address:=app.loadEnv("resources/.env-dev")
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}).Handler(app.route(UrlRepoImpl{
		conn: app.loadConnection(),
		logger: logger}))
	serve:=&http.Server{
		Addr:address,
		Handler:handler,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30*time.Second,
	}
	err:=serve.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}
