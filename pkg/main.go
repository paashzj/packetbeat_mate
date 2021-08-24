package main

import (
	"os"
	"os/signal"
	"packetbeat_mate/pkg/pb"
	"packetbeat_mate/pkg/util"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	interrupt := make(chan os.Signal, 1)
	pb.Start()
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
