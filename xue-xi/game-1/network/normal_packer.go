package network

import (
	"encoding/binary"
	"io"
)

type NormalPacker struct {
	Order binary.ByteOrder
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{Order: order}
}

func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buff := make([]byte, 8+8+uint64(len(message.Data)))

	p.Order.PutUint64(buff[:8], uint64(len(buff)))
	p.Order.PutUint64(buff[8:16], message.Id)

	copy(buff[16:], message.Data)
	return buff, nil
}

func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {

	/*if err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second)) ; err != nil {
		return nil ,err
	}*/

	buffer := make([]byte, 8+8)
	_, err := io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}

	totalLen := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:])
	dataLen := totalLen - 16

	dataBuff := make([]byte, dataLen)
	_, err = io.ReadFull(reader, dataBuff)
	if err != nil {
		return nil, err
	}

	m := &Message{
		Id:   id,
		Data: dataBuff,
	}

	//fmt.Println("NuPack" ,m.Id ,string(m.Data))
	return m, nil
}
