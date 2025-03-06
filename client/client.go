package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"grpc/chatpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second, // Envia ping a cada 10s de inatividade
			Timeout:             5 * time.Second,  // Tempo para esperar resposta do ping
			PermitWithoutStream: true,             // Permite ping sem RPC ativa
		}),
	)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	client := chatpb.NewChatServiceClient(conn)

	stream, err := client.ServerPush(context.Background(), &chatpb.Empty{})
	if err != nil {
		log.Fatalf("Erro ao abrir stream: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println("Stream encerrado pelo servidor.")
			break
		}
		fmt.Printf("\n%s: %s\n", msg.User, msg.Message)
	}
}
