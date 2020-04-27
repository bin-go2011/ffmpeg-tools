package av

func NewAVFormatContext() *AVFormatContext {
	return avformatAllocContext()
}
func (fmtctx *AVFormatContext) Release() {
	avformatFreeContext(fmtctx)
}

func (fmtctx *AVFormatContext) OpenInput(file string) error {
	fmtctx.file = file
	return avformatOpenInput(fmtctx, file)
}

func (fmtctx *AVFormatContext) Dump() {
	avDumpFormat(fmtctx, fmtctx.file)
}
