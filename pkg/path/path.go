package path

import (
	"os"
	"path/filepath"
)

// packetbeat
var (
	PbHome   = os.Getenv("PACKETBEAT_HOME")
	PbConfig = filepath.FromSlash(PbHome + "/packetbeat.yml")
)

// mate
var (
	PbMatePath    = filepath.FromSlash(PbHome + "/mate")
	PbScripts     = filepath.FromSlash(PbMatePath + "/scripts")
	PbStartScript = filepath.FromSlash(PbScripts + "/start-packetbeat.sh")
)
