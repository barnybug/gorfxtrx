package gorfxtrx

import "sort"

// Struct for the Status packet type
type Status struct {
	data            []byte
	TranceiverType  byte
	FirmwareVersion byte
}

var StatusTypes = map[byte]string{
	0x50: "310MHz",
	0x51: "315MHz",
	0x53: "433.92MHz",
	0x55: "868.00MHz",
	0x56: "868.00MHz FSK",
	0x57: "868.30MHz",
	0x58: "868.30MHz FSK",
	0x59: "868.35MHz",
	0x5A: "868.35MHz FSK",
	0x5B: "868.95MHz",
}

func (self *Status) TypeString() string {
	if StatusTypes[self.TranceiverType] != "" {
		return StatusTypes[self.TranceiverType]
	}
	return "unknown"
}

func (self *Status) Devices() []string {
	devs := []string{}
	devs = extend(devs, decode_flags(self.data[7]/0x80, []string{"undecoded"}))
	devs = extend(devs, decode_flags(self.data[8], []string{"mertik", "lightwarerf", "hideki", "lacrosse", "fs20", "proguard"}))
	devs = extend(devs, decode_flags(self.data[9], []string{"x10", "arc", "ac", "homeeasy", "ikeakoppla", "oregon", "ati", "visonic"}))
	sort.Strings(devs)
	return devs
}

func (self *Status) Receive(data []byte) {
	self.data = data
	self.TranceiverType = data[5]
	self.FirmwareVersion = data[6]
}

func (self *Status) Send() []byte {
	return []byte{0x0d, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

// Status packet constructor
func NewStatus() (*Status, error) {
	return &Status{}, nil
}
