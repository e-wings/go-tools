/* Package form 包含表单验证等辅助函数
 */
package form

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

func IsFormat(str string, fmtName string) (res bool, err error) {
	switch fmtName {
	case "cn":
		res, err = regexp.MatchString("^\\p{Han}+$", str)
		break
	case "en":
		res, err = regexp.MatchString("^[a-zA-Z ]+$", str)
		break
	case "email":
		res, err = regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, str)
		break
	case "mobile":
		res, err = regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, str)
		break
	case "idCard":
		fmt.Println(str)
		if len(str) == 15 {
			res, err = regexp.MatchString(`^(\d{15})$`, str)
			break
		} else if len(str) == 18 {
			res, err = regexp.MatchString(`^(\d{17})([0-9]|X)$`, str)
			break
		} else {
			res = false
			err = errors.New("身份证号位数不对")
		}
		res = false
		err = nil
		//res, err = regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, str)
		break
	}

	return res, err
}

/* InRange判断元素的值是否包含在一个数组中
   适用于元素类型为整型、Float32、Float64和字符串时
*/
func InRange(i interface{}, r interface{}) (res bool, err error) {
	fmt.Println()
	iv := reflect.ValueOf(i)
	ran := reflect.ValueOf(r)
	//元素为整型时
	if iv.Kind() == reflect.Int {
		val := iv.Int()
		for i := 0; i < ran.Len(); i++ {
			e := ran.Index(i)
			if e.Kind() == reflect.Int {
				ev := e.Int()
				if ev == val {
					res = true
				}
			}
		}
		//元素为Float32或Float64时
	} else if iv.Kind() == reflect.Float32 || iv.Kind() == reflect.Float64 {
		val := iv.Float()
		for i := 0; i < ran.Len(); i++ {
			e := ran.Index(i)
			if e.Kind() == reflect.Float32 || e.Kind() == reflect.Float64 {
				ev := e.Float()
				if ev == val {
					res = true
				}
			}
		}
		//元素为字符串时
	} else if iv.Kind() == reflect.String {
		val := iv.String()
		for i := 0; i < ran.Len(); i++ {
			e := ran.Index(i)
			if e.Kind() == reflect.String {
				ev := e.String()
				if ev == val {
					res = true
				}
			}
		}
	}
	//BUG(adam): 未补全会出错的情况
	return res, nil
}

//判断一个Slice是否为另一个Slice的子集
//支持String、Float、Int
func RangeInRange(slice1 interface{}, slice2 interface{}) (res bool, err error) {
	s1 := reflect.ValueOf(slice1)
	s2 := reflect.ValueOf(slice2)
	l1 := s1.Len()
	l2 := s2.Len()
	l3 := 0
	sType := s1.Index(0).Kind()
	s2Type := s2.Index(0).Kind()
	if sType != s2Type {
		return false, errors.New("类型不一致")
	}

	switch sType {
	case reflect.String:
		for i := 0; i < l1; i++ {
			for j := 0; j < l2; j++ {
				if s1.Index(i).String() == s2.Index(j).String() {
					l3++
				}
			}
		}
	case reflect.Int:
		for i := 0; i < l1; i++ {
			for j := 0; j < l2; j++ {
				if s1.Index(i).Int() == s2.Index(j).Int() {
					l3++
				}
			}
		}
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		for i := 0; i < l1; i++ {
			for j := 0; j < l2; j++ {
				if s1.Index(i).Float() == s2.Index(j).Float() {
					l3++
				}
			}
		}
	}

	if l1 != l3 {
		return false, errors.New("元素超出范围")
	}
	return true, nil
}
