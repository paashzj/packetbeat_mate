#!/bin/bash

nohup $PACKETBEAT_HOME/heartbeat -e -c $PACKETBEAT_HOME/heartbeat.yml >$PACKETBEAT_HOME/packetbeat.stdout.log 2>$PACKETBEAT_HOME/packetbeat.stderr.log &