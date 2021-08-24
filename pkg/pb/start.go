package pb

import (
	"go.uber.org/zap"
	"packetbeat_mate/pkg/path"
	"packetbeat_mate/pkg/util"
)

func Start() {
	startPb()
}

func startPb() {
	err := generatePbFile()
	if err != nil {
		util.Logger().Error("generate prom config file failed ", zap.Error(err))
		return
	}
	err = util.CallScript(path.PbStartScript)
	util.Logger().Error("run start prom scripts failed ", zap.Error(err))
}

func generatePbFile() (err error) {
	// todo
	return nil
}
