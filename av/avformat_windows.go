package av

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type AVFormatContext struct {
	handle uintptr
	file   string
}

var avformatDLL *windows.DLL

var (
	avformatVersionProc,
	avformatAllocContextProc,
	avformatFreeContextProc,
	avformatOpenInputProc,
	avDumpFormatProc *windows.Proc
)

func init() {
	avformatDLL = windows.MustLoadDLL("../lib/avformat-58.dll")

	avformatVersionProc = avformatDLL.MustFindProc("avformat_version")
	avformatAllocContextProc = avformatDLL.MustFindProc("avformat_alloc_context")
	avformatFreeContextProc = avformatDLL.MustFindProc("avformat_free_context")
	avformatOpenInputProc = avformatDLL.MustFindProc("avformat_open_input")
	avDumpFormatProc = avformatDLL.MustFindProc("av_dump_format")
}

func avformatVersion() string {
	r1, _, _ := avformatVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func avformatAllocContext() *AVFormatContext {
	r1, _, _ := avformatAllocContextProc.Call()
	return &AVFormatContext{
		handle: r1,
	}
}

func avformatFreeContext(fmtctx *AVFormatContext) {
	avformatFreeContextProc.Call(fmtctx.handle)
}

func avformatOpenInput(fmtctx *AVFormatContext, file string) error {
	f, _ := syscall.BytePtrFromString(file)

	r1, _, err := avformatOpenInputProc.Call(uintptr(unsafe.Pointer(&(fmtctx.handle))), uintptr(unsafe.Pointer(f)),
		uintptr(0), uintptr(0))

	if r1 == 0 {
		return err
	}

	return nil

}

func avDumpFormat(fmtctx *AVFormatContext, file string) {
	f, _ := syscall.BytePtrFromString(file)

	avDumpFormatProc.Call(fmtctx.handle, uintptr(0), uintptr(unsafe.Pointer(f)), uintptr(0))
}
