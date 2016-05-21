#!/bin/bash
export HOST=localhost
export PORT=4816
./svarmrMessage $HOST $PORT start-module usernotify  &
./svarmrMessage $HOST $PORT start-module svarmrMdnsWatcher  &
./svarmrMessage $HOST $PORT start-module mDNSprocessor  &
sleep 1
./svarmrMessage $HOST $PORT user-notify "All modules loaded"
./monitor $HOST $PORT
