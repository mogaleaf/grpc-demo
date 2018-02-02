package main

import (
	"google.golang.org/grpc"
	"meetup/grpc/meetup/gen"
	"meetup/grpc/meetup/service"
	"net"
	"fmt"
)

func main(){
	println("Start GRPC SERVER")
	//Listener on the address
	lis,e := net.Listen("tcp", "127.0.0.1:9999")
	if e != nil {
		fmt.Println(e)
	}
	//Create Grpc Server
	grpcServer := grpc.NewServer()
	//Create the Service
	helloService := &service.WelcomService{}
	//Register the Service
	gen.RegisterHelloServer(grpcServer, helloService)
	//Start Listening
	grpcServer.Serve(lis)
}