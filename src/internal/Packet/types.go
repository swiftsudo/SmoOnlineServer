package packet

// A short in C# is a signed 16-bit integer.
type PacketType int16

const (
	UnknownPacket PacketType = iota
	InitPacket
	PlayerPacket
	CapPacket
	GamePacket
	TagPacket
	ConnectPacket
	DisconnectPacket
	CostumePacket
	ShinePacket
	CapturePacket
	ChangeStagePacket
	CommandPacket
)
