package av

func NewAVFormatContext() *AVFormatContext {
	return avformatAllocContext()
}
func (fmtctx *AVFormatContext) Release() {
	avformatFreeContext(fmtctx)
}

func (fmtctx *AVFormatContext) OpenInput(file string) error {
	return avformatOpenInput(fmtctx, file)
}

func (fmtctx *AVFormatContext) Dump(file string) {
	avDumpFormat(fmtctx, file)
}
