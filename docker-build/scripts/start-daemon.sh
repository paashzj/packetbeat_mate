#!/bin/bash

mkdir -p $PACKETBEAT_HOME/logs
nohup $PACKETBEAT_HOME/mate/packetbeat_mate >>$PACKETBEAT_HOME/logs/packetbeat_mate.stdout.log 2>>$PACKETBEAT_HOME/packetbeat_mate.stderr.log

