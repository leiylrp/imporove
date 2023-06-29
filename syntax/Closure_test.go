package syntax

import "testing"

// 函数计数器
func TestClosure(t *testing.T) {
	counter := newCounter()
	counter()
	counter()
}

// 传递函数实现闭包 中间件
func TestPrintN(t *testing.T) {
	pt := timer(printN)
	pt(1)
	pt(10)
}