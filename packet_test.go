package gorfxtrx

import (
	"fmt"
)

func ExampleStatus() {
	x, _ := Parse([]byte{0x0d, 0x01, 0x00, 0x01, 0x02, 0x53, 0x3e, 0x00, 0x0c, 0x2f, 0x01, 0x01, 0x00, 0x00})
	//Output:
	status := *x.(*Status)
	fmt.Printf("%+v\n", status)
	// {TranceiverType:83 FirmwareVersion:62}
	fmt.Println(status.TypeString())
	// 433.92MHz
	fmt.Println(status.Devices())
	// [ac arc hideki homeeasy lacrosse oregon x10]
}

func ExampleShortBytes() {
	_, err := Parse([]byte{0x0d, 0x01, 0x00, 0x01, 0x02, 0x53, 0x3e, 0x00, 0x0c, 0x2f, 0x01, 0x01, 0x00})
	fmt.Println(err)
	//Output:
	// Packet too short
}

func ExampleShortData() {
	_, err := Parse([]byte{0x01, 0x01})
	fmt.Println(err)
	//Output:
	// Packet too short
}

func ExampleStatusSend() {
	p, err := NewStatus()
	fmt.Println(err)
	fmt.Println(p.Send())
	//Output:
	// <nil>
	// [13 0 0 1 2 0 0 0 0 0 0 0 0 0]
}

func ExampleResetSend() {
	p, err := NewReset()
	fmt.Println(err)
	fmt.Println(p.Send())
	//Output:
	// <nil>
	// [13 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func ExampleLightingX10() {
	x, _ := Parse([]byte{0x07, 0x10, 0x00, 0x2a, 0x45, 0x05, 0x01, 0x70})
	lighting := *x.(*LightingX10)
	fmt.Printf("%+v\n", lighting)
	fmt.Println(lighting.Type())
	fmt.Println(lighting.Id())
	fmt.Println(lighting.Command())
	//Output:
	// {TypeId:0 SequenceNumber:42 HouseCode:69 UnitCode:5 command:1}
	// X10 lighting
	// e05
	// on
}

func ExampleLightingX10Send() {
	p, _ := NewLightingX10(0x01, "e05", "on")
	fmt.Println(p.Send())
	//Output:
	// [7 16 1 0 69 5 1 0]
}

func ExampleLightingHE() {
	x, _ := Parse([]byte{0x0b, 0x11, 0x00, 0x2a, 0x01, 0x23, 0x45, 0x67, 0x05, 0x02, 0x08, 0x70})
	lighting := *x.(*LightingHE)
	fmt.Printf("%+v\n", lighting)
	fmt.Println(lighting.Type())
	fmt.Println(lighting.Id())
	fmt.Println(lighting.Command())
	//Output:
	// {TypeId:0 SequenceNumber:42 HouseCode:19088743 UnitCode:5 command:2 Level:8}
	// AC
	// 12345675
	// set level
}

func ExampleLightingHENewBad() {
	_, err := NewLightingHE(0x00, "bad", "on")
	fmt.Println(err)
	//Output:
	// id should be 8 characters (eg. 1234567b)
}

func ExampleLightingHESend() {
	p, err := NewLightingHE(0x00, "002A41E6", "on")
	fmt.Println(err)
	fmt.Println(p.Send())
	//Output:
	// <nil>
	// [11 17 0 0 0 2 164 30 6 1 0 0]
}

func ExampleTemp() {
	x, _ := Parse([]byte{0x08, 0x50, 0x02, 0x2a, 0x96, 0x03, 0x81, 0x41, 0x79})
	temp := *x.(*Temp)
	fmt.Printf("%+v\n", temp)
	fmt.Printf("%+v\n", temp.Id())
	fmt.Printf("%+v\n", temp.Type())
	//Output:
	// {TypeId:2 SequenceNumber:42 id:38403 Temp:-32.1 Battery:90 Rssi:7}
	// 96:03
	// THC238/268,THN132,THWR288,THRN122,THN122,AW129/131
}

func ExampleTempHumid() {
	x, _ := Parse([]byte{0x0a, 0x52, 0x01, 0x2a, 0x96, 0x03, 0x81, 0x41, 0x60, 0x03, 0x79})
	temp := *x.(*TempHumid)
	fmt.Printf("%+v\n", temp)
	fmt.Printf("%+v\n", temp.Id())
	fmt.Printf("%+v\n", temp.Type())
	//Output:
	// {TypeId:1 SequenceNumber:42 id:38403 Temp:-32.1 Humidity:96 HumidityStatus:3 Battery:90 Rssi:7}
	// 96:03
	// THGN122/123, THGN132, THGR122/228/238/268
}

func ExampleWind() {
	x, _ := Parse([]byte{0x10, 0x56, 0x01, 0x03, 0x2F, 0x00, 0x00, 0xF7, 0x00, 0x20, 0x00, 0x24, 0x01, 0x60, 0x00, 0x00, 0x59})
	wind := *x.(*Wind)
	fmt.Printf("%+v\n", wind)
	fmt.Println(wind.Id())
	fmt.Println(wind.Type())
	//Output:
	// {data:[16 86 1 3 47 0 0 247 0 32 0 36 1 96 0 0 89] TypeId:1 SequenceNumber:3 id:12032 Direction:247 AverageSpeed:3.2 Gust:3.6 Battery:90 Rssi:5}
	// 2f:00
	// WTGR800
}
