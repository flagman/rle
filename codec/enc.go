package codec
import (
	"errors"
)

const (
	String uint8 = 1

)

var CannotSerializeNilName = errors.New("New Run cannot serialize nil name")
var CannotSerializeNilData = errors.New("New Run cannot serialize nil data")

/*
	DataType   uint8
	NameLength uint8
	DataLength int32
	Name       []byte
	Data       []byte
 */
type Run []byte

func NewRun(dtype uint8, name []byte, data []byte) (Run, error) {
	size, err := RunSize(name, data)
	if err != nil {
		return nil, err
	}
	run := make([]byte, size)
	run[0] = byte(dtype)
	run[1] = byte(len(name))

	datalen := len(data)
	run[2] = byte((datalen & 0xF000) >> 24)
	run[3] = byte((datalen & 0x0F00) >> 16)
	run[4] = byte((datalen & 0x00F0) >> 8)
	run[5] = byte(datalen & 0x000F)

	i := 0
	for ; i < len(name); i++ {
		run[i+6] = name[i]
	}

	for j := 0; j < len(data); j++ {
		pos := j+i+6
		run[pos] = data[j]
	}

	return run, nil
}

// Calculates the run size based on name and data using the
// 2 bytes for type, name length; and 4 bytes data length.
//
// type => 1
// name length => 1
// data length => 4
// name => N
// data => M
func RunSize(name []byte, data []byte) (int, error) {
	if name == nil {
		return 0, CannotSerializeNilName
	}
	if data == nil {
		return 0, CannotSerializeNilData
	}
	return 6 + len(data) + len(name), nil
}

// Len return the number of bytes that encode the Run.
func (r Run) Len() int {
	return len(r)
}

// DataLength binary decodes the data length from the underlying bytes.
func (r Run) DataLength() int {
	n := int(r[2]) + int(r[3]) + int(r[4]) + int(r[5])
	return n
}

// NameLength binary decodes the name length from the underlying bytes.n
func (r Run) NameLength() int {
	return int(r[1])
}

// DataType binary decodes the data type from the underlying bytes.
func (r Run) DataType() uint8 {
	return uint8(r[0])
}

func (r Run) Name() []byte {
	start := 6
	end := start + int(r[1])
	return r[start:end]
}

func (r Run) Data() []byte {
	start := 6 + r.NameLength()
	end := start + r.DataLength()
	return r[start:end]
}
