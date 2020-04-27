package av

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var avcodecDLL *windows.DLL

var (
	avcodecVersionProc *windows.Proc
)

func init() {
	avcodecDLL = windows.MustLoadDLL("../lib/avcodec-58.dll")
	avcodecVersionProc = avcodecDLL.MustFindProc("avcodec_version")
}

func avcodecVersion() string {
	r1, _, _ := avcodecVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}
