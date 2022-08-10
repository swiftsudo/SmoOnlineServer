package packet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	packets "github.com/swiftsudo/smoonlineserver/internal/Packet/Packets"
	"github.com/swiftsudo/smoonlineserver/internal/common"
)

// Test Cap Packet
func TestCap(t *testing.T) {
	testPacket := &packets.CapPacket{
		Position: common.Vector3{1, 2, 3},
		Rotation: common.Quaternion{1, 2, 3, 4},
		CapOut:   true,
		CapAnim:  "test",
	}

	testPacketHeader := &PacketHeader{
		Id:         1,
		Type:       CapPacket,
		PacketSize: testPacket.Size(),
	}

	serializedData := SerializeHeaded(testPacketHeader, testPacket)
	fmt.Printf("serialized: %x\n", serializedData)

	outputPacketHeader := DeserializeHeader(serializedData)
	fmt.Printf("output header: %+v\n", outputPacketHeader)
	assert.Equal(t, testPacketHeader, outputPacketHeader)

	outputPacket := Deserialize(serializedData)
	fmt.Printf("output packet: %+v\n", outputPacket)
	assert.Equal(t, testPacket, outputPacket)
}
