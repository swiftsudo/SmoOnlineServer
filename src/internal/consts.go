package internal

import packet "github.com/swiftsudo/smoonlineserver/internal/Packet"

const (
	CostumeNameSize int = 0x20
	HeaderSize          = packet.HeaderSize
)

// PacketMap is a map of PacketType to Packet

// PacketIdMap

var (
	MapNames = map[string]string{
		"cap":     "CapWorldHomeStage",
		"cascade": "WaterfallWorldHomeStage",
		"sand":    "SandWorldHomeStage",
		"lake":    "LakeWorldHomeStage",
		"wooded":  "ForestWorldHomeStage",
		"cloud":   "CloudWorldHomeStage",
		"lost":    "ClashWorldHomeStage",
		"metro":   "CityWorldHomeStage",
		"sea":     "SeaWorldHomeStage",
		"snow":    "SnowWorldHomeStage",
		"lunch":   "LavaWorldHomeStage",
		"ruined":  "BossRaidWorldHomeStage",
		"bowser":  "SkyWorldHomeStage",
		"moon":    "MoonWorldHomeStage",
		"mush":    "PeachWorldHomeStage",
		"dark":    "Special1WorldHomeStage",
		"darker":  "Special2WorldHomeStage",
	}
)
