package main

import (
	"google.golang.org/grpc"
	"fmt"
	"meetup/grpc/meetup/gen"
	"golang.org/x/net/context"
)

func main() {
	//Create the connection
	conn, err := grpc.Dial("127.0.0.1:9999",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	//Create the Client
	client := gen.NewHelloClient(conn)

	//Create the object request
	var messages []*gen.WelcomMessage
	messages = append(messages,
		&gen.WelcomMessage{
		Person:&gen.Person{
			Age:       34,
			FirstName: "Cecile",
			Gender:    gen.Gender_FEMALE,
			LastName:  "Thiebaut",
		},
		Message:"I love you Gopher",
	})
	req := &gen.WelcomMessages{Messages:messages}

	//Call the Server
	resp,err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.MessageResponse)
}
