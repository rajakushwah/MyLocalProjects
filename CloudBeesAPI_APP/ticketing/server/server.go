package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "ticketing/ticketing"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTicketServiceServer
	mu      sync.Mutex
	tickets map[string]*pb.TicketReceipt
	seats   map[string]string
}

// Purchase Ticket for user with name , email
func (s *server) PurchaseTicket(ctx context.Context, req *pb.PurchaseRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	email := req.Email
	receipt := &pb.TicketReceipt{
		From:      "London",
		To:        "France",
		User:      fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		PricePaid: 20.0,
		Section:   req.Section,
	}

	s.tickets[email] = receipt
	s.seats[email] = req.Section

	return receipt, nil
}

// Get the Receipt for a user using email
func (s *server) GetReceipt(ctx context.Context, req *pb.GetReceiptRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.tickets[req.Email]
	if !exists {
		return nil, fmt.Errorf("no receipt found for email: %s", req.Email)
	}

	return receipt, nil
}

// Get the list of users in a section
func (s *server) GetUsersInSection(ctx context.Context, req *pb.GetUsersInSectionRequest) (*pb.UsersInSectionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []string
	for email, section := range s.seats {
		if section == req.Section {
			users = append(users, email)
		}
	}

	return &pb.UsersInSectionResponse{Users: users}, nil
}

// RemoveUser implementation
func (s *server) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tickets, req.Email)
	delete(s.seats, req.Email)
	log.Printf("Attempting to remove user: %s", req.Email)
	return &pb.RemoveUserResponse{
		Success: true,
		Message: "User removed successfully.",
	}, nil
}

func (s *server) ModifyUserSeat(ctx context.Context, req *pb.ModifyUserSeatRequest) (*pb.TicketReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	receipt, exists := s.tickets[req.Email]
	if !exists {
		return nil, fmt.Errorf("no receipt found for email: %s", req.Email)
	}
	receipt.Section = req.NewSection
	s.seats[req.Email] = req.NewSection
	return receipt, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterTicketServiceServer(s, &server{
		tickets: make(map[string]*pb.TicketReceipt),
		seats:   make(map[string]string),
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server is running on port :50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
