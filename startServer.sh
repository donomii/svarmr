pkill -f 4816
sh build.sh
./server localhost 4816 &
sleep 5
./heartbeat localhost 4816 `hostname`&
./monitor localhost 4816 &
./mdnsWatcher localhost 4816 2> /dev/null &
./mDNSprocessor localhost 4816 &
