package match

var servers = []string{
	"NA-EAST",
	"NA-WEST",
	"EU",
}

var serverIndex = 0

func AllocateServer() string {

	server := servers[serverIndex]

	serverIndex++

	if serverIndex >= len(servers) {
		serverIndex = 0
	}

	return server
}
