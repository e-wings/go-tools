/* Package form 包含表单验证等辅助函数
 */
package form

import (
	"fmt"
	"reflect"
)

var intType = reflect.TypeOf(1)

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
