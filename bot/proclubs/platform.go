package proclubs

import "fmt"


type Platform int

const (
	PlatformNextGeneration	Platform = iota
	PlatformOldGeneration
	PlatformNX
	MaxPlatforms
)

var platformStrings = map[Platform]string {
	PlatformNextGeneration:	"common-gen5",
	PlatformOldGeneration:	"common-gen4",
	PlatformNX:				"nx",
}

var platformNames = map[Platform]string {
	PlatformNextGeneration:	"Next Generation",
	PlatformOldGeneration:	"Old Generation",
	PlatformNX:				"NX",
}

func PlatformString(platform Platform) string {
	return platformStrings[platform]
}

func PlatformName(platform Platform) string {
	return platformNames[platform]
}

func StringToPlatform(name string) (Platform, error) {
	for i := range MaxPlatforms {
		if platformStrings[i] == name {
			return i, nil
		}
	}

	return Platform(0), fmt.Errorf("platform string doesn't exist")
}
