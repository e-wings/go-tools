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

	// RangeInRange测试
	res, err = RangeInRange([]string{"footbal", "basketball"}, []string{"footbal", "basketball", "tennis"})
	if res != true || err != nil {
		t.Error("Test_RangeInRange: 字符串子集测试未通过")
	}
	res, err = RangeInRange([]string{"footbal", "basketball", "wrong"}, []string{"footbal", "basketball", "tennis"})
	if res != false || err == nil {
		t.Error("Test_RangeInRange: 字符串超集测试未通过")
	}
	res, err = RangeInRange([]int{1, 2}, []int{1, 2, 3})
	if res != true || err != nil {
		t.Error("Test_RangeInRange: 整型子集测试未通过")
	}
	res, err = RangeInRange([]int{1, 2, 4}, []int{1, 2, 3})
	if res != false || err == nil {
		t.Error("Test_RangeInRange: 整型超集测试未通过")
	}
	res, err = RangeInRange([]float64{1.1, 2.2}, []float64{1.1, 2.2, 3.3})
	if res != true || err != nil {
		t.Error("Test_RangeInRange: Float子集测试未通过")
	}
	res, err = RangeInRange([]float64{1.1, 2.2, 4.4}, []float64{1.1, 2.2, 3.3})
	if res != false || err == nil {
		t.Error("Test_RangeInRange: Float超集测试未通过")
	}
}
