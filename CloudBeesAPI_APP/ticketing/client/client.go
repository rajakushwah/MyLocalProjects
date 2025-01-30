package main

import (
	"context"
	"fmt"
	"log"

	pb "ticketing/ticketing"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTicketServiceClient(conn)

	// Purchase a ticket
	purchaseResp, err := client.PurchaseTicket(context.Background(), &pb.PurchaseRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Section:   "A",
	})
	if err != nil {
		log.Fatalf("could not purchase ticket: %v", err)
	}
	fmt.Printf("Purchased Ticket: %+v\n", purchaseResp)

	// Get Receipt
	receiptResp, err := client.GetReceipt(context.Background(), &pb.GetReceiptRequest{
		Email: "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("could not get receipt: %v", err)
	}
	fmt.Printf("Receipt: %+v\n", receiptResp)

	// Get Users in Section
	usersResp, err := client.GetUsersInSection(context.Background(), &pb.GetUsersInSectionRequest{
		Section: "A",
	})
	if err != nil {
		log.Fatalf("could not get users: %v", err)
	}
	fmt.Printf("Users in Section A: %+v\n", usersResp)

	// Modify User Seat
	modifyResp, err := client.ModifyUserSeat(context.Background(), &pb.ModifyUserSeatRequest{
		Email:      "john.doe@example.com",
		NewSection: "B",
	})
	if err != nil {
		log.Fatalf("could not modify seat: %v", err)
	}
	fmt.Printf("Modified Seat: %+v\n", modifyResp)

	// Remove User
	response, err := client.RemoveUser(context.Background(), &pb.RemoveUserRequest{
		Email: "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("could not remove user: %v", err)
	}

	fmt.Printf("Response: %+v\n", response) // Print response details

}
