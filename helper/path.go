package helper

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var TmpDir = ""

func AutoMkdirAll(tmpFileDir string) error {
	tmpFileDir, _ = filepath.Abs(tmpFileDir)
	_, err := os.Stat(tmpFileDir)
	if err != nil {
		_ = os.MkdirAll(tmpFileDir, os.ModeDir)
	}
	abs, _ := filepath.Abs(tmpFileDir + "/test")
	err = ioutil.WriteFile(abs, []byte("test"), os.ModeAppend)
	_ = os.Remove(abs)
	return err
}
func AutoSetTmpDir() string {
	if runtime.GOOS == "windows" {
		return `C:\Windows\TEMP`
	}
	tmpDirs := []string{
		"/data/local/tmp",
		"/tmp",
		"/data/tmp",
		"/sdcard/tmp",
	}

	for e := range tmpDirs {
		if AutoMkdirAll(tmpDirs[e]) == nil {
			TmpDir = tmpDirs[e]
			break
		}
	}
	if TmpDir != "" {
		_ = os.Setenv("TEMP", TmpDir)
		_ = os.Setenv("TMP", TmpDir)
		_ = os.Setenv("TMPDIR", TmpDir)
	}
	return TmpDir
}
func SetTmpDir(tmpFileDir string) {
	AutoMkdirAll(tmpFileDir)

	_ = os.Setenv("TEMP", tmpFileDir)
	_ = os.Setenv("TMP", tmpFileDir)
	_ = os.Setenv("TMPDIR", tmpFileDir)
}
