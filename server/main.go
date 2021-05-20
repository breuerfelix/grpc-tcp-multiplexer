package main

import (
	"log"
	"net"
	"os"

	bridge "github.com/breuerfelix/grpc-tcp-multiplexer/client"
	"google.golang.org/grpc"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "4444"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Panicln("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	broker := bridge.Server{}
	grpcServer := grpc.NewServer()
	bridge.RegisterBridgeServer(grpcServer, &broker)

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalln("Failed to serve grpc server:", err.Error())
	}

}
