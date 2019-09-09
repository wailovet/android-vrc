go-bindata-assetfs ../static/...
set GOOS=linux
set GOARCH=arm
go build -o ./vrc
adb push ./vrc /data/local/tmp/
adb shell chmod 755 /data/local/tmp/vrc
adb shell /data/local/tmp/vrc