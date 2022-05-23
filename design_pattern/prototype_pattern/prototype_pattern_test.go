/**
 * @Author: Sean
 * @Date: 2022/5/23 14:52
 */

package prototype_pattern

import (
	"fmt"
	"testing"
)

func TestExample_Clone(t *testing.T) {
	origin := New("origin object")
	current := origin.Clone()
	fmt.Println(current.Description)
	// 定制clone对象的属性
	current.Description = "current object"
	fmt.Println(current.Description)
}
