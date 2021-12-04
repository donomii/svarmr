cd svarmr
go build server.go &
go build example.go &
go build heartBeat.go &
gcc monitor.c -Os -flto -lm -omonitorc &
go build monitor.go &
go build svarmrMessage.go &
go build usernotify.go &
go build moduleStarter.go &
cd ..

cd net
go build relay.go &
go build mDNSprocessor.go &
go build svarmrMdnsWatcher.go &
go build svarmrMdnsAdvertiser.go &
cd ..

cd pitch
go build pitchWrapper.go &
go build pitchProcessor.go &
go build noteKeyboard.go &
cd detect && gcc -lm -lportaudio pitchDetect.c -o pitchDetect &
cd ..

cd image
go build snapShot.go &
go build recogniser.go &
go build broadcastImage.go &
cd ..

cd clipboard
go build clipboardProcessor.go &
cd ..

cd misc
go build volume.go &
go build insertKey.go &
go build volumeControllerWindows.go &
go build torrentListener.go &
cd ..

cd systray
go build tray.go
cd ..

cd gui
go build gui.go
cd ..


wait
move svarmr/server .
