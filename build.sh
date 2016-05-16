go build server.go
go build relay.go
go build volume.go
gcc monitor.c -Os -flto 
