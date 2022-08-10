package common

import "encoding/binary"

type Vector3 [3]float32    // list of 3 floats
type Quaternion [4]float32 // list of 4 floats (x, y, z, w) or (vector3, scalarPart)

// GetVector3 returns Vector3 from data
func GetVector3(data []byte) Vector3 {
	return Vector3{
		float32(binary.LittleEndian.Uint32(data[:Float32Size])),
		float32(binary.LittleEndian.Uint32(data[Float32Size : Float32Size*2])),
		float32(binary.LittleEndian.Uint32(data[Float32Size*2 : Float32Size*3])),
	}
}

// GetQuaternion returns Quaternion from data
func GetQuaternion(data []byte) Quaternion {
	return Quaternion{
		float32(binary.LittleEndian.Uint32(data[:Float32Size])),
		float32(binary.LittleEndian.Uint32(data[Float32Size : Float32Size*2])),
		float32(binary.LittleEndian.Uint32(data[Float32Size*2 : Float32Size*3])),
		float32(binary.LittleEndian.Uint32(data[Float32Size*3 : Float32Size*4])),
	}
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
