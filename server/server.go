package main

import (
	"log"
	"net"
	"time"

	"grpc/chatpb"

	"google.golang.org/grpc"
)

type chatServer struct {
	chatpb.UnimplementedChatServiceServer
}

func (s *chatServer) ServerPush(_ *chatpb.Empty, stream chatpb.ChatService_ServerPushServer) error {
	messages := []string{
		"Bem-vindo ao chat!",
		"Essa é uma mensagem automática do servidor.",
		"Fique à vontade para interagir.",
		"Mais uma mensagem enviada sem solicitação do cliente.",
	}

	for _, msg := range messages {
		res := &chatpb.ChatMessage{
			User:    "Servidor",
			Message: msg,
		}

		if err := stream.Send(res); err != nil {
			log.Printf("Erro ao enviar mensagem: %v", err)
			return err
		}

		time.Sleep(3 * time.Second)
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	chatpb.RegisterChatServiceServer(grpcServer, &chatServer{})

	log.Println("Servidor gRPC rodando na porta 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao rodar servidor: %v", err)
	}
}
