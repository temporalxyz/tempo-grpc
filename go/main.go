// main.go
package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "tempo/gen" // Adjust this import path if needed.
)

func main() {
	// Connect to the gRPC endpoint.
	endpoint := "check the docs for endpoints" // TODO: Replace with the actual endpoint.
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to %q: %v", endpoint, err)
	}
	defer conn.Close()

	// Create a new TransactionStreamClient.
	client := pb.NewTransactionStreamClient(conn)

	// Build the StartStream request.
	authToken := "dm JC, cavey, ben, or j for auth token" // TODO: Replace with your auth token.
	req := &pb.StartStream{
		AuthToken: authToken,
	}

	// Create context
	ctx := context.Background();

	// Open the transaction stream.
	stream, err := client.OpenTransactionStream(ctx, req)
	if err != nil {
		log.Fatalf("Failed to open transaction stream: %v", err)
	}

	// Process transactions from the stream in a loop.
	for {
		// Recv() blocks until a message is received, or an error occurs.
		msg, err := stream.Recv()
		if err == io.EOF {
			// The stream was closed by the server.
			fmt.Println("TransactionStream closed!")
			return
		}
		if err != nil {
			log.Fatalf("Invalid stream message! Error: %v", err)
		}

		// Destructure the Transaction message.
		slot := msg.GetSlot()     // Slot in which the transaction was found.
		index := msg.GetIndex()   // Transaction's index within the slot.
		payload := msg.GetPayload() // The transaction's bytes.

		// Process the transaction.
		snipe(slot, index, payload)
	}
}

// snipe processes an individual Transaction message.
// Replace this with your actual business logic.
func snipe(slot uint64, index uint64, payload []byte) {
	// fmt.Printf("Snipe: slot=%d, index=%d, payload=%v\n", slot, index, payload)
	fmt.Printf("Snipe: slot=%d, index=%d\n", slot, index)
}
