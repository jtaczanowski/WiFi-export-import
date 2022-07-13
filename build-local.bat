go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
SET GOARCH=386
go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
go generate
go build -ldflags="-H windowsgui" -o WiFi-export-import-x86.exe .

SET GOARCH=amd64
go build -ldflags="-H windowsgui" -o WiFi-export-import-x64.exe .
