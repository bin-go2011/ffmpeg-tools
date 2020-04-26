module ffmpeg-tools/ffprobe

go 1.13

replace golang.org/x/sys v0.0.0-20190405154228-4b34438f7a67 => github.com/golang/sys v0.0.0-20200323222414-85ca7c5b95cd

require (
	av v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f
)

replace av => ../lib
