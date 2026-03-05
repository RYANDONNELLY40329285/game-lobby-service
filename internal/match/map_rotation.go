package match

var maps = []string{
	"Nuketown",
	"Rust",
	"Shipment",
	"Terminal",
}

var mapIndex = 0

func NextMap() string {

	mapName := maps[mapIndex]

	mapIndex++

	if mapIndex >= len(maps) {
		mapIndex = 0
	}

	return mapName
}
