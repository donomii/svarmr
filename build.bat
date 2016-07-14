cd svarmr
go build server.go
go build heartBeat.go
gcc monitor.c -Os -flto -omonitorc
go build monitor.go
go build svarmrMessage.go
go build usernotify.go
go build moduleStarter.go
cd ..

cd net
go build relay.go
go build mDNSprocessor.go
go build svarmrMdnsWatcher.go
cd ..

cd pitch
go build pitchWrapper.go
go build pitchProcessor.go
go build noteKeyboard.go
cd detect  gcc -l PortAudio pitchDetect.c -o pitchDetect
cd ..

cd image
go build snapShot.go
go build recogniser.go
cd ..

cd clipboard
go build clipboardProcessor.go
cd ..

cd misc
go build volume.go
go build insertKey.go
go build volumeControllerWindows.go
cd ..


wait
