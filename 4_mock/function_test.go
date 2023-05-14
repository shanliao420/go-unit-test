package __mock

import (
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	line := ProcessFirstLine()
	expectLine := "Hello Mock!😎"
	assert.Equal(t, expectLine, line)
}

// ReadFirstLine 函数读取文字，此测试依赖于文件

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "Hello Mock!"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	expectLine := "Hello Mock!😎"
	assert.Equal(t, expectLine, line)
}

// 使用mock的方式实现对函数对打桩，从而去除对函数或者文件对依赖性

// king@chengxiaofengdeiMac 4_mock % go test function.go function_test.go -cover
// ok      command-line-arguments  0.010s  coverage: 75.0% of statements
