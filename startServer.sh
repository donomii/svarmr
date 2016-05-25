#!/bin/bash
export HOST=localhost
export PORT=4816
pkill -f $PORT
sh build.sh
./server $HOST $PORT &
sleep 2
./heartBeat $HOST $PORT `hostname` &
./moduleStarter $HOST $PORT &
./usernotify $HOST $PORT &
echo Server started
