/**
 * @Author: Sean
 * @Date: 2022/5/24 14:14
 */

package decorator_pattern

import (
	"fmt"
	"testing"
)

func TestCreateAppleDecorator(t *testing.T) {
	var comp Component = &Fruit{Count: 8, Description: "水果统称"}
	result := CreateAppleDecorator(comp, "apple", 20)
	fmt.Println(result.Describe())
	re := result.GetCount()
	if re != 28 {
		t.Errorf("装饰错误，期待结果为%d", re)
	}
}
