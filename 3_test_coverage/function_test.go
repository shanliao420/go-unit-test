package __test_coverage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudgePassTrue(t *testing.T) {
	isPass := JudgePass(70)
	assert.Equal(t, true, isPass)
}

// king@chengxiaofengdeiMac 3_test_coverage % go test function.go function_test.go -cover
// ok      command-line-arguments  0.012s  coverage: 66.7% of statements

// 测试的代码覆盖率为66.7% 因为总结3行代码，跑了2行代码

func TestJudgePassFail(t *testing.T) {
	isPass := JudgePass(59)
	assert.Equal(t, false, isPass)
}

// king@chengxiaofengdeiMac 3_test_coverage % go test function.go function_test.go -cover
// ok      command-line-arguments  0.010s  coverage: 100.0% of statements

// 添加测试用例，增加代码覆盖率，一般测试代码覆盖率能达到50%-60%
