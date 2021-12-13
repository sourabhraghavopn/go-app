package main

import (
	"fmt"
	env "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)



func  (app App) loadEnv() string {
	err := env.Load("resources/.env")
	if err != nil {
		app.logger.Fatal(err)
	}
	port:="8080"
	env_port, exist := os.LookupEnv("PORT")
	if exist {
		port=env_port
	}else{
		app.logger.Print("Default  ")
	}
	app.logger.Printf("port : %s ",port)
	return  fmt.Sprintf(":%s", port)
}

func  main() {
	fmt.Println("Running")
	app:=&App{
		logger:log.New(os.Stdout,"",log.Ldate|log.Llongfile),
	}
	serve:=&http.Server{
		Addr:app.loadEnv(),
		Handler:app.route(),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30*time.Second,
	}
	err:=serve.ListenAndServe()
	if err != nil {
		app.logger.Fatal(err)
	}
}
