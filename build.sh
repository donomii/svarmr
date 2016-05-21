go build server.go &
go build relay.go &
go build volume.go &
go build heartBeat.go &
go build monitor.go &
go build mdnsWatcher.go &
go build mDNSprocessor.go &
go build svarmrMessage.go &
go build moduleStarter.go &
gcc monitor.c -Os -flto -omonitorc &
wait
