package client

import (
	"log"
)

type Server struct {
	UnimplementedBridgeServer
}

func (s *Server) NewClient(stream Bridge_NewClientServer) error {
	for {
		data, err := stream.Recv()
		if err != nil {
			log.Println("got error")
			break
		}

		// TODO implement your logic here
		// this this echoes all received data
		stream.Send(&DataPacket{Data: []byte(string(data.Data))})
	}

	return nil
}
