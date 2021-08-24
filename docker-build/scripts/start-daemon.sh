#!/bin/bash


nohup $METRICBEAT_HOME/mate/packetbeat_mate >$METRICBEAT_HOME/packetbeat_mate.stdout.log 2>$METRICBEAT_HOME/packetbeat_mate.stderr.log
