package service

import (
	"meetup/grpc/meetup/gen"
	"golang.org/x/net/context"
	"fmt"
)

type WelcomService struct {
}

//Server method
func (s *WelcomService) SayHello(ctx context.Context, in *gen.WelcomMessages) (*gen.ThankYouResponse, error) {
	var returnMessage string
	fmt.Println(fmt.Sprintf("Receive %d message",len(in.Messages)))
	for _, p := range in.Messages {
		fmt.Println(fmt.Sprintf("A message from %s : %s",p.Person.FirstName,p.Message))
		returnMessage += fmt.Sprintf(" Thanks  %s ",p.Person.FirstName)
	}
	response := gen.ThankYouResponse{
		MessageResponse: returnMessage,
	}
	return &response, nil
}
