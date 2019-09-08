go-bindata-assetfs ../static/...
GOOS=linux GOARCH=arm go build -o ./vrc
$HOME/Desktop/android-platform-tools/adb push ./vrc /data/local/tmp/
$HOME/Desktop/android-platform-tools/adb shell chmod 755 /data/local/tmp/vrc
$HOME/Desktop/android-platform-tools/adb shell /data/local/tmp/vrc