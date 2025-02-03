# Ticketing Service gRPC API

## Overview
This repository provides a gRPC-based ticketing service that allows users to purchase tickets, retrieve receipts, manage user information, and handle seating arrangements. The service is designed for managing tickets between two locations.

## Requirements
- Go 1.14 or later
- gRPC library for Go
- Protocol Buffer compiler (`protoc`)
- Protobuf definitions for the Ticketing service (located in `ticketing/ticketing.proto`)


# Installation

Clone the repository:

```bash
git clone <repository-url>
cd <repository-directory>

** git url - https://github.com/rajakushwah/MyLocalProjects/tree/master/CloudBeesAPI_APP/ticketing
```

Installation of Required Packages

Use the following command to install the gRPC and Protobuf libraries:

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

Compile Protocol Buffers:

Make sure to have the protoc compiler installed. Then run:

```bash
protoc --go_out=. --go-grpc_out=. ticketing/ticketing.proto
```
# Running the Service

```bash
go run server.go
```
After starting the service, it will listen on localhost:50051.

# gRPC API Endpoints

## 1. Purchase Ticket

- **RPC:** `PurchaseTicket`
- **Request:** `PurchaseRequest`
  - **Email** (string): User's email address
  - **FirstName** (string): User's first name
  - **LastName** (string): User's last name
  - **Section** (string): Section for the ticket
- **Response:** `TicketReceipt`
  - **From** (string): Departure location
  - **To** (string): Arrival location
  - **User** (string): User's full name
  - **PricePaid** (float): Amount paid for the ticket
  - **Section** (string): Seating section
- **Description:** Allows a user to purchase a ticket and returns the receipt.

## 2. Get Receipt

- **RPC:** `GetReceipt`
- **Request:** `GetReceiptRequest`
  - **Email** (string): User's email address
- **Response:** `TicketReceipt`
- **Description:** Retrieves the ticket receipt for a given email address.

## 3. Get Users in Section

- **RPC:** `GetUsersInSection`
- **Request:** `GetUsersInSectionRequest`
  - **Section** (string): The section for which to retrieve users
- **Response:** `UsersInSectionResponse`
  - **Users** (repeated string): List of users in the specified section
- **Description:** Returns a list of users who have purchased tickets in a specified section.

## 4. Remove User

- **RPC:** `RemoveUser`
- **Request:** `RemoveUserRequest`
  - **Email** (string): User's email address
- **Response:** `RemoveUserResponse`
  - **Success** (bool): Indicates if the removal was successful
  - **Message** (string): Message describing the result
- **Description:** Removes a user and their associated ticket information from the system.

## 5. Modify User Seat

- **RPC:** `ModifyUserSeat`
- **Request:** `ModifyUserSeatRequest`
  - **Email** (string): User's email address
  - **NewSection** (string): New section for the user
- **Response:** `TicketReceipt`
- **Description:** Modifies the seating section for an existing user.




