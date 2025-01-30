package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	pb "ticketing/ticketing" // Import the generated protobuf code

	"google.golang.org/grpc"
)

func TestPurchaseTicket(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketServiceClient(conn)
	req := &pb.PurchaseRequest{
		Email:     "test@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Section:   "A1",
	}
	receipt, err := client.PurchaseTicket(context.Background(), req)
	fmt.Print(receipt)
	if err != nil {
		t.Fatalf("PurchaseTicket failed: %v", err)
	}
	// if receipt.From != "London" {
	// 	t.Errorf("expected From to be 'London', got %s", receipt.From)
	// }
}

func TestGetReceipt(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketServiceClient(conn)
	req := &pb.GetReceiptRequest{
		Email: "test@example.com",
	}
	receipt, err := client.GetReceipt(context.Background(), req)
	fmt.Print(receipt)
	if err != nil {
		t.Fatalf("GetReceipt failed: %v", err)
	}
}

func TestGetUsersInSection(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketServiceClient(conn)
	req := &pb.GetUsersInSectionRequest{
		Section: "A1",
	}
	users, err := client.GetUsersInSection(context.Background(), req)
	fmt.Print(users)
	if err != nil {
		t.Fatalf("GetUsersInSection failed: %v", err)
	}
}

func TestRemoveUser(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketServiceClient(conn)
	req := &pb.RemoveUserRequest{
		Email: "test@example.com",
	}
	status, err := client.RemoveUser(context.Background(), req)
	fmt.Print(status)
	if err != nil {
		t.Fatalf("RemoveUser failed: %v", err)
	}
}

func TestModifyUserSeat(t *testing.T) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketServiceClient(conn)
	req := &pb.ModifyUserSeatRequest{
		Email:      "test@example.com",
		NewSection: "A2",
	}
	updated_record, err := client.ModifyUserSeat(context.Background(), req)
	fmt.Print(updated_record)
	if err != nil {
		t.Fatalf("ModifyUserSeat failed: %v", err)
	}
}
