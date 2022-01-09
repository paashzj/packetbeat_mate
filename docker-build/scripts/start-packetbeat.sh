#!/bin/bash

mkdir $PACKETBEAT_HOME/logs
nohup $PACKETBEAT_HOME/packetbeat -e -c $PACKETBEAT_HOME/packetbeat.yml >>$PACKETBEAT_HOME/logs/packetbeat.stdout.log 2>>$PACKETBEAT_HOME/logs/packetbeat.stderr.log &

