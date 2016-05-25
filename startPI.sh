pkill -f 4816

./server 0.0.0.0 4816 &
sleep 2
#./heartBeat localhost 4816 `hostname` &
#./monitor localhost 4816 &
./torrentListener localhost 4816 &
