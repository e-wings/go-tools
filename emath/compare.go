//Package emath 包含多种数学运算函数
package emath

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var arrayFloat = reflect.TypeOf([]float64{})
var floatType = reflect.TypeOf(float64(0))

/*Max比较一组数字、字符串式的数字中最大的一个
 */
func Max(arg ...interface{}) (res float64, err error) {
	fmt.Println("")
	//多个元素比较大小
	if len(arg) > 1 {
		res, err = GetFloat(arg[0])
		for _, v := range arg {
			v, _ := GetFloat(v)
			if v > res {
				res = v
			}
		}
	} else { //对单个数组中元素进行比较
		t := reflect.TypeOf(arg[0])
		v := reflect.ValueOf(arg[0])
		if t != arrayFloat {
			return math.NaN(), errors.New("Max: 传入一个元素时，必须是[]flat64")
		} else {
			res = math.NaN()
			for i := 1; i < v.Len(); i++ {
				if math.IsNaN(res) {
					res = v.Index(0).Float()
				}
				if v.Index(i).Float() > res {
					res = v.Index(i).Float()
				}
			}
		}
	}
	return res, err
}

/* Min比较一组数字、字符串式的数字中最小的一个
 */
func Min(arg ...interface{}) (res float64, err error) {
	fmt.Println("")
	//多个元素比较大小
	if len(arg) > 1 {
		res, err = GetFloat(arg[0])
		for _, v := range arg {
			v, _ := GetFloat(v)
			if v < res {
				res = v
			}
		}
	} else { //对单个数组中元素进行比较
		t := reflect.TypeOf(arg[0])
		v := reflect.ValueOf(arg[0])
		if t != arrayFloat {
			return math.NaN(), errors.New("Max: 传入一个元素时，必须是[]flat64")
		} else {
			res = math.NaN()
			for i := 1; i < v.Len(); i++ {
				if math.IsNaN(res) {
					res = v.Index(0).Float()
				}
				if v.Index(i).Float() < res {
					res = v.Index(i).Float()
				}
			}
		}
	}
	return res, err
}

//GetFloat 将int, float32, 字符串式的数字转换为float64类型
func GetFloat(unk interface{}) (float64, error) {
	v := reflect.ValueOf(unk)
	switch v.Type().String() {
	case "string":
		res, err := strconv.ParseFloat(v.String(), 64)
		return res, err
	default:
		//得到指针所指向的对象（去除指针）
		v = reflect.Indirect(v)
		if !v.Type().ConvertibleTo(floatType) {
			return math.NaN(), errors.New("cannot convert to float64")
		}
		fv := v.Convert(floatType)
		return fv.Float(), nil
	}

}
