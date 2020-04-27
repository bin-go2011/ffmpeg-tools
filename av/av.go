package av

import (
	"fmt"
)

func Version() {
	fmt.Printf("libavutil\t%s\n", avutilVersion())
	fmt.Printf("libavcodec\t%s\n", avcodecVersion())
	fmt.Printf("libavformat\t%s\n", avformatVersion())
}
