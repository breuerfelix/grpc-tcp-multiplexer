package mqtt

import (
	"log"
)

type Broker struct {
	UnimplementedMqttServiceServer
}

func (s *Broker) NewClient(stream MqttService_NewClientServer) error {
	for {
		data, err := stream.Recv()
		if err != nil {
			//log.Println("got error")
			break
		}

		log.Println("Broker got data:", string(data.Data))
		stream.Send(&DataPacket{Data: []byte("Echo: " + string(data.Data))})
	}

	return nil
}
