package av

import (
	"fmt"

	"golang.org/x/sys/windows"
)

var avutilDLL *windows.DLL

var (
	avutilVersionProc *windows.Proc
)

func init() {
	avutilDLL = windows.MustLoadDLL("../lib/avutil-56.dll")
	avutilVersionProc = avutilDLL.MustFindProc("avutil_version")
}

func AvutilVersion() string {
	r1, _, _ := avutilVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}
