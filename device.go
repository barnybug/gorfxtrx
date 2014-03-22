package gorfxtrx

import (
	"io"
	"log"
	"time"

	"github.com/tarm/goserial"
)

type Device struct {
	ser io.ReadWriteCloser
}

func Open(devname string) (*Device, error) {
	dev := Device{}

	c := &serial.Config{Name: devname, Baud: 38400}
	ser, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	dev.ser = ser

	log.Println("Sending reset")
	reset, _ := NewReset()
	err = dev.Send(reset)
	if err != nil {
		return nil, err
	}

	return &dev, nil
}

func (self *Device) Close() {
	self.ser.Close()
}

func (self *Device) Read() (Packet, error) {
	buf := make([]byte, 257)
	for {
		i, err := self.ser.Read(buf)
		if i == 0 && err == io.EOF {
			// empty read, sleep a bit recheck
			time.Sleep(time.Millisecond * 200)
			continue
		}
		if err != nil {
			return nil, err
		}
		if i == 0 {
			continue
		}
		offset := i
		for remain := int(buf[0]) + 1 - i; remain > 0; remain -= i {
			i, err = self.ser.Read(buf[offset:])
			if i == 0 && err == io.EOF {
				time.Sleep(time.Millisecond * 200)
				continue
			}
			if err != nil {
				return nil, err
			}
			offset += i
		}

		// log.Println("Received:", buf[0:offset])
		return Parse(buf[0:offset])
	}
}

func (self *Device) Send(p OutPacket) error {
	buf := p.Send()
	_, err := self.ser.Write(buf)
	return err
}
