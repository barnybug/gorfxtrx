package gorfxtrx

import (
	"encoding/binary"
	"fmt"
)

// Struct for the Wind packets
type Wind struct {
	data           []byte
	TypeId         byte
	SequenceNumber byte
	id             uint16
	Direction      uint16
	AverageSpeed   float64
	Gust           float64
	Battery        byte
	Rssi           byte
}

var WindTypes = map[byte]string{
	0x01: "WTGR800",
	0x02: "WGR800",
	0x03: "STR918, WGR918",
	0x04: "TFA",
}

func (self *Wind) Receive(data []byte) {
	self.data = data
	self.TypeId = data[2]
	self.SequenceNumber = data[3]
	self.id = binary.BigEndian.Uint16(data[4:6])
	self.Direction = binary.BigEndian.Uint16(data[6:8])
	self.AverageSpeed = float64(binary.BigEndian.Uint16(data[8:10])) / 10
	self.Gust = float64(binary.BigEndian.Uint16(data[10:12])) / 10
	if self.TypeId == 0x03 {
		self.Battery = (data[16] + 1) * 10
	} else {
		self.Battery = (data[16] & 0x0f) * 10
		self.Rssi = data[16] >> 4
	}
}

func (self *Wind) Id() string {
	return fmt.Sprintf("%02x:%02x", self.id>>8, self.id&0xff)
}

func (self *Wind) Type() string {
	return WindTypes[self.TypeId]
}
