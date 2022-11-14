package gorfxtrx

// Struct for the SetMode packet type.
type SetMode struct {
	AeBlyss          bool
	Rubicson         bool
	FineoffsetViking bool
	Lighting4        bool
	Rsl              bool
	ByronSX          bool
	Rfu6             bool
	Edisplay         bool
	Mertik           bool
	Lightwarerf      bool
	Hideki           bool
	Lacrosse         bool
	Fs20             bool
	Proguard         bool
	Blindst0         bool
	Blindst1         bool
	X10              bool
	Arc              bool
	Ac               bool
	Homeeasy         bool
	Ikeakoppla       bool
	Oregon           bool
	Ati              bool
	Visonic          bool
}

func (self *SetMode) Send() []byte {
	// return []byte{0x0d, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	var b2i = map[bool]byte{false: 0, true: 1}
	var flag1 byte = b2i[self.AeBlyss] + b2i[self.Rubicson]*2 + b2i[self.FineoffsetViking]*4 + b2i[self.Lighting4]*8 + b2i[self.Rsl]*16 + b2i[self.ByronSX]*32 + b2i[self.Rfu6]*64 + b2i[self.Edisplay]*128
	var flag2 byte = b2i[self.Mertik] + b2i[self.Lightwarerf]*2 + b2i[self.Hideki]*4 + b2i[self.Lacrosse]*8 + b2i[self.Fs20]*16 + b2i[self.Proguard]*32 + b2i[self.Blindst0]*64 + b2i[self.Blindst1]*128
	var flag3 byte = b2i[self.X10] + b2i[self.Arc]*2 + b2i[self.Ac]*4 + b2i[self.Homeeasy]*8 + b2i[self.Ikeakoppla]*16 + b2i[self.Oregon]*32 + b2i[self.Ati]*64 + b2i[self.Visonic]*128
	return []byte{0x0d, 0x00, 0x00, 0x01, 0x03, 0x53, 0x00, flag1, flag2, flag3, 0x00, 0x00, 0x00, 0x00}
}

// SetMode packet constructor
func NewSetMode() (*SetMode, error) {
	return &SetMode{}, nil
}
