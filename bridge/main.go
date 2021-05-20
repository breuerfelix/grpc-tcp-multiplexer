package main

import (
	bridge "github.com/breuerfelix/grpc-tcp-multiplexer/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

const (
	CONN_HOST   = "localhost"
	CONN_PORT   = "3333"
	CONN_TYPE   = "tcp"
	BUFFER_SIZE = 256
)

func main() {
	// grpc client
	conn, err := grpc.Dial("localhost:4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	log.Println("Connected!")

	defer conn.Close()

	c := bridge.NewBridgeClient(conn)

	// tcp server
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		log.Panicln("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()

	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	runServer(&l, &c)

}

func runServer(listener *net.Listener, client *bridge.BridgeClient) {
	for {
		// Listen for an incoming connection.
		conn, err := (*listener).Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			continue
		}

		go handleConnection(conn, client)
	}
}

func handleConnection(conn net.Conn, client *bridge.BridgeClient) {
	defer conn.Close()

	c := *client
	stream, _ := c.NewClient(context.Background())
	defer stream.CloseSend()

	// forwards broker message to client
	go handleReadConnection(&conn, &stream)

	for {
		data, err := readData(conn)
		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}

		stream.Send(&bridge.DataPacket{Data: data})
	}
}

func handleReadConnection(conn *net.Conn, client *bridge.Bridge_NewClientClient) {
	cStream := *client
	cConn := *conn

	for {
		data, err := cStream.Recv()
		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}

		cConn.Write(data.Data)
	}
}

func readData(conn net.Conn) ([]byte, error) {
	data := make([]byte, 0)
	buffer := make([]byte, BUFFER_SIZE)

	for {
		reqLen, err := conn.Read(buffer)
		if err != nil {
			return nil, err
		}

		data = append(data, buffer[:reqLen]...)
		if reqLen < BUFFER_SIZE {
			break
		}
	}

	return data, nil
}
