go build server.go &
go build relay.go &
go build volume.go &
go build heartBeat.go &
go build monitor.go &
go build svarmrMdnsWatcher.go &
go build mDNSprocessor.go &
go build svarmrMessage.go &
go build moduleStarter.go &
go build pitchWrapper.go &
go build pitchProcessor.go &
go build noteProcessor.go &
gcc monitor.c -Os -flto -omonitorc &
cd pitchDetect && gcc -l PortAudio pitchDetect.c -o pitchDetect &
wait
