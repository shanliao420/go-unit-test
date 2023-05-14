package __benchmark

import (
	"github.com/bytedance/gopkg/lang/fastrand"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = 100 + i
	}
}

func SelectServer() int {
	return ServerIndex[fastrand.Intn(10)]
}
