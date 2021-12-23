package pb

import (
	"github.com/paashzj/gutil"
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
	stdout, stderr, err := gutil.CallScript(path.PbStartScript)
	util.Logger().Info("output is ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	util.Logger().Error("run start coredns scripts failed ", zap.Error(err))
}

func generatePbFile() (err error) {
	// todo
	return nil
}
