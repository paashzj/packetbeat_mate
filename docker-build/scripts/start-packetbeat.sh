#!/bin/bash

nohup $PACKETBEAT_HOME/packetbeat -e -c $PACKETBEAT_HOME/packetbeat.yml >>$PACKETBEAT_HOME/packetbeat.stdout.log 2>>$PACKETBEAT_HOME/packetbeat.stderr.log &