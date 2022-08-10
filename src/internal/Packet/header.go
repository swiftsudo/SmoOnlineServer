package packet

import (
	"encoding/binary"

	"github.com/swiftsudo/smoonlineserver/internal/common"
)

const (
	StaticSize = 20
	HeaderSize = StaticSize
)

type Guid int16 // todo: find valid guid type

type PacketHeader struct {
	Id         Guid       // size is 16 bits
	Type       PacketType // size is 16 bits
	PacketSize int16      // size is 16 bits

	// Unused?
	// Size = StaticSize
}

// Serialize Packet Header to data
func (h *PacketHeader) Serialize(data []byte) {
	binary.LittleEndian.PutUint16(data[:common.ShortSize], uint16(h.Id))
	binary.LittleEndian.PutUint16(data[common.ShortSize:], uint16(h.Type))
	binary.LittleEndian.PutUint16(data[common.ShortSize+2:], uint16(h.PacketSize))
}

// Deserialize Packet Header from data
func DeserializeHeader(data []byte) *PacketHeader {
	h := &PacketHeader{}

	h.Id = Guid(binary.LittleEndian.Uint16(data[:common.ShortSize]))
	h.Type = PacketType(binary.LittleEndian.Uint16(data[common.ShortSize:]))
	h.PacketSize = int16(binary.LittleEndian.Uint16(data[common.ShortSize+2:]))

	return h
}
