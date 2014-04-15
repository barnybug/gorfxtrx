package gorfxtrx

import (
	"fmt"
)

func ExampleRead() {
	replay := [][]byte{
		[]byte{0xa},
		[]byte{0x1, 0x0, 0x1},
		[]byte{0x2, 0x53, 0x3e, 0x0},
		[]byte{0xc, 0x2f, 0x1, 0x1},
		[]byte{0x0, 0x0},
	}
	ser := NewMockSerialPort(replay)
	dev := Device{ser: ser, debug: false}
	packet, err := dev.Read()
	fmt.Printf("%+v %v\n", packet, err)
	packet, err = dev.Read()
	fmt.Printf("%+v %v\n", packet, err)
	// Output:
	// &{data:[10 1 0 1 2 83 62 0 12 47 1] TransceiverType:83 FirmwareVersion:62} <nil>
	// <nil> <nil>
}
