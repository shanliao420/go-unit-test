package __unit

import (
	"fmt"
	"os"
	"testing"
)

var (
	data string
)

func TestHelloTom(t *testing.T) {
	fmt.Println(data)
	output := HelloTom()
	expectOutput := "Tom"
	if expectOutput != output {
		t.Errorf("Expected %s do not match actual %s ", expectOutput, output)
	}
}

func TestMain(m *testing.M) {

	// 测试前：数据装在、初始化配置

	fmt.Println("Hello testing!")

	data = "shanliao is cool!"

	code := m.Run()

	// 测试后：资源释放

	fmt.Println("Goodbye testing!")

	os.Exit(code)

}
