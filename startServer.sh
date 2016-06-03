#!/bin/bash
cd "$(dirname ${BASH_SOURCE[0]})"
echo "Starting svarmr in " `pwd`
export HOST=localhost
export PORT=4816
pkill -f $PORT
#sh build.sh
sleep 2 && svarmr/heartBeat $HOST $PORT `hostname` &
sleep 2 && svarmr/moduleStarter $HOST $PORT &
sleep 2 && svarmr/usernotify $HOST $PORT &
sleep 2 && svarmr/configManager.pl $HOST $PORT &
echo Server started
svarmr/server $HOST $PORT
