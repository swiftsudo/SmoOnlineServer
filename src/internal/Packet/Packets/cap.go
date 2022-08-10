package packets

import (
	"encoding/binary"
	"strings"

	"github.com/swiftsudo/smoonlineserver/internal/common"
)

const (
	capPacketSize int16 = 0x50
	capNameSize   int   = 0x30
)

type CapPacket struct {
	Position common.Vector3
	Rotation common.Quaternion
	CapOut   bool
	CapAnim  string
}

func (p *CapPacket) Serialize(data []byte) {
	// Pack position bytes
	for i, v := range p.Position {
		binary.LittleEndian.PutUint32(data[i*common.Float32Size:], uint32(v))
	}

	// Pack rotation bytes
	for i, v := range p.Rotation {
		binary.LittleEndian.PutUint32(data[common.Vector3Size+i*common.Float32Size:], uint32(v))
	}

	binary.LittleEndian.PutUint16(data[common.Vector3Size+common.QuaternionSize:], uint16(common.BoolToInt(p.CapOut)))
	copy(data[common.Vector3Size+common.QuaternionSize+common.BoolSize:], p.CapAnim)
}

func (p *CapPacket) Deserialize(data []byte) {
	p.Position = common.GetVector3(data)
	p.Rotation = common.GetQuaternion(data[common.Vector3Size:])
	p.CapOut = binary.LittleEndian.Uint16(data[common.Vector3Size+common.ShortSize:]) == 1
	p.CapAnim = string(data[common.Vector3Size+common.QuaternionSize+common.BoolSize : (common.Vector3Size + common.QuaternionSize + common.BoolSize + capNameSize)])
	p.CapAnim = strings.Trim(p.CapAnim, "\x00")
}

func (p *CapPacket) Size() int16 {
	return capPacketSize
}
