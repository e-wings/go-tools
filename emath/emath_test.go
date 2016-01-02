package emath

import (
	"testing"
)

func Test_Max(t *testing.T) {
	sampleArray := []float64{7.1, 9, 8.0}
	res, err := Max(sampleArray)
	if err != nil || res != 9 {
		t.Error("Max测试未通过", err)
	}
	res, err = Max(7.1, "9", 8.0)
	if err != nil || res != 9 {
		t.Error("Max测试未通过", err)
	}
	t.Log("Max测试通过")
}

func Test_Min(t *testing.T) {
	sampleArray := []float64{7.1, 9, 8.0}
	res, err := Min(sampleArray)
	if err != nil || res != 7.1 {
		t.Error("Min测试未通过", err)
	}
	res, err = Min("7.1", 9, 8.0)
	if err != nil || res != 7.1 {
		t.Error("Min测试未通过", err)
	}
	t.Log("Min测试通过")
}
