pkill -f 4816

./svarmr/server 0.0.0.0 4816 &
sleep 2
#./heartBeat localhost 4816 `hostname` &
#./monitor localhost 4816 &
#./net/svarmrMdnsWatcher.go localhost 4816 &
./misc/torrentListener localhost 4816 &
./net/svarmrMdnsAdvertiser localhost 4816 &
