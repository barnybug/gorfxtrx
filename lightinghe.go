package gorfxtrx

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
)

// Struct for the Homeeasy lighting packets
type LightingHE struct {
	TypeId         byte
	SequenceNumber byte
	HouseCode      uint32
	UnitCode       byte
	command        byte
	Level          byte
}

var LightingHETypes = map[byte]string{
	0x00: "AC",
	0x01: "HomeEasy EU",
	0x02: "ANSLUT",
}

var LightingHECommands = map[byte]string{
	0x00: "off",
	0x01: "on",
	0x02: "set level",
	0x03: "group off",
	0x04: "group on",
	0x05: "set group level",
}

var LightingHECommandBytes = reverseByteStringMap(LightingHECommands)

func NewLightingHE(typeId byte, id string, command string) (*LightingHE, error) {
	if len(id) != 8 {
		return nil, errors.New("id should be 8 characters (eg. 1234567b)")
	}
	houseCode, err := strconv.ParseInt(id[:7], 16, 32)
	if err != nil {
		return nil, err
	}
	unitCode, err := strconv.ParseInt(id[7:], 16, 8)
	if err != nil {
		return nil, err
	}
	return &LightingHE{
		TypeId:    typeId,
		HouseCode: uint32(houseCode),
		UnitCode:  byte(unitCode),
		command:   LightingHECommandBytes[command],
	}, nil
}

func (self *LightingHE) Receive(data []byte) {
	self.TypeId = data[2]
	self.SequenceNumber = data[3]
	self.HouseCode = binary.BigEndian.Uint32(data[4:8])
	self.UnitCode = data[8]
	self.command = data[9]
	self.Level = data[10]
}

func (self *LightingHE) Type() string {
	return LightingHETypes[self.TypeId]
}

func (self *LightingHE) Id() string {
	return fmt.Sprintf("%07x%1x", self.HouseCode, self.UnitCode)
}

func (self *LightingHE) Command() string {
	return LightingHECommands[self.command]
}

func (self *LightingHE) Send() []byte {
	b := []byte{0x0b, 0x11, self.TypeId, self.SequenceNumber,
		0, 0, 0, 0, self.UnitCode, self.command, self.Level, 0}
	binary.BigEndian.PutUint32(b[4:8], self.HouseCode)
	return b
}
