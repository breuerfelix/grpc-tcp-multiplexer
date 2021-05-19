package main

import (
  "log"
  "net"
  "os"

  "google.golang.org/grpc"
  "github.com/breuerfelix/thesis/mqtt"
)

const (
	CONN_HOST   = "localhost"
	CONN_PORT   = "4444"
	CONN_TYPE   = "tcp"
	BUFFER_SIZE = 1024
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Panicln("Error listening:", err.Error())
		os.Exit(1)
	}
  defer l.Close()

  broker := mqtt.Broker{}
  grpcServer := grpc.NewServer()
  mqtt.RegisterMqttServiceServer(grpcServer, &broker)

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

  if err := grpcServer.Serve(l); err != nil {
    log.Fatalln("Failed to serve grpc server:", err.Error())
  }

}
