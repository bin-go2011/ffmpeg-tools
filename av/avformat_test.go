package av

import (
	"path/filepath"
	"testing"
)

var SAMPLE_VIDEO string

func init() {
	SAMPLE_VIDEO, _ = filepath.Abs("../data/big_buck_bunny.mp4")

}
func TestOpenInput(t *testing.T) {
	fmtctx := NewAVFormatContext()
	defer fmtctx.Release()

	fmtctx.OpenInput(SAMPLE_VIDEO)

	fmtctx.Dump()
}
