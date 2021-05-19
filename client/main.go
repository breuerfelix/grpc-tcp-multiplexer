package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const (
	CONN_HOST   = "localhost"
	CONN_PORT   = "3333"
	CONN_TYPE   = "tcp"
	BUFFER_SIZE = 1024
)

func main() {
	log.Println("Starting stresser...")

	for i := 0; i < 20000; i++ {
		go dummy()
    log.Println("Client started:", i)

		min := 1
		max := 10
		interval := rand.Intn(max-min) + min
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	log.Println("Stresser going ...")
  dummy()
}

func dummy() {
	// tcp client
	conn, _ := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	for {
		text := "hallo test"
		fmt.Fprintf(conn, text)

    data, _ := readData(conn)
    if string(data) != text {
      log.Println("wrong!")
    }
		//log.Println(string(message))

		min := 5
		max := 20
		interval := rand.Intn(max-min) + min
		time.Sleep(time.Duration(interval) * time.Second)
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
