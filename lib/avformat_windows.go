package av

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var avformatDLL *windows.DLL

var (
	avformatVersionProc *windows.Proc
)

func init() {
	avformatDLL = windows.MustLoadDLL("../win64/avformat-58.dll")
	avformatVersionProc = avformatDLL.MustFindProc("avformat_version")
}

func AvformatVersion() string {
	r1, _, _ := avformatVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}
