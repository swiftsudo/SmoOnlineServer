package packet

import packets "github.com/swiftsudo/smoonlineserver/internal/Packet/Packets"

func SerializeHeaded(header *PacketHeader, p Packet) []byte {
	data := make([]byte, HeaderSize+p.Size())
	header.Serialize(data)
	p.Serialize(data[HeaderSize:])

	return data
}

func Deserialize(data []byte) Packet {
	header := DeserializeHeader(data)

	var p Packet
	switch header.Type {
	case CapPacket:
		p = &packets.CapPacket{}
	}

	p.Deserialize(data[HeaderSize:])

	return p
}

func SizeOf(p Packet) int16 {
	return HeaderSize + p.Size()
}
