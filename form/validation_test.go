package form

import (
	"testing"
)

func Test_InRange(t *testing.T) {
	res, err := InRange(1, []int{0, 1, 2, 3})
	if res != true || err != nil {
		t.Error("Test_InRange: 测试整型未通过")
	}
	res, err = InRange(4, []int{0, 1, 2, 3})
	if res != false || err != nil {
		t.Error("Test_InRange: 测试整型未通过")
	}
	var a float64 = 3.14
	res, err = InRange(a, []float64{0, 1, 2.1, 3.14})
	if res != true || err != nil {
		t.Error("Test_InRange: 测试Float64未通过")
	}
	res, err = InRange(a, []float64{0, 1, 2.1, 3.1411})
	if res != false || err != nil {
		t.Error("Test_InRange: 测试Float64未通过")
	}
	var b float32 = 3.14
	res, err = InRange(b, []float32{0, 1, 2.1, 3.14})
	if res != true || err != nil {
		t.Error("Test_InRange: 测试Float32未通过")
	}
	res, err = InRange(b, []float32{0, 1, 2.1, 3.1411})
	if res != false || err != nil {
		t.Error("Test_InRange: 测试Float32未通过")
	}
	res, err = InRange("apple", []string{"apple", "tree", "sss"})
	if res != true || err != nil {
		t.Error("Test_InRange: 测试字符串未通过")
	}
	res, err = InRange("apple2", []string{"apple", "tree", "sss"})
	if res != false || err != nil {
		t.Error("Test_InRange: 测试字符串未通过")
	}
}
