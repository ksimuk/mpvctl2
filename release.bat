go build -o mpvctl.exe cmd\mpvctl\mpvctl.go
go build -o mpvstart.exe -ldflags -H=windowsgui cmd\mpvstart\mpvstart.go 
copy mpvctl.exe C:\Tools\
copy mpvstart.exe C:\Tools\