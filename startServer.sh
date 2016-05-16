./server localhost 4816 &
sleep 5
./heartbeat localhost 4816 `hostname`&
./monitor localhost 4816 &
