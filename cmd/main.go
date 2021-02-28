package main

import (
	"my_grpc_do_zero/application/http"
	"my_grpc_do_zero/application/grpc"
	"fmt" 
)

func main() { 

	fmt.Println("Programa Inicializado")
	
	// Startar o webserver
	webserver := http.NewWebServer()
	go webserver.Serve()

	// Startar o grpc Server
	grpc.StartGrpcServer()

}

