package packet

type Packet interface {
	Serialize([]byte)
	Deserialize([]byte)
	Size() int16
}
