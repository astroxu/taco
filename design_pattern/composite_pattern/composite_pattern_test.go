/**
 * @Author: Sean
 * @Date: 2022/5/30 14:50
 */

package composite_pattern

import (
	"fmt"
	"testing"
)

func TestNewMenuItem(t *testing.T) {
	menuItem := NewMenuItem("Order1", "desc1", 12)
	// 可以直接调用匿名组合的方法
	fmt.Println(menuItem.Name())
}

func TestMenuGroup_Add(t *testing.T) {
	menu := NewMenu("Order", "ordernew")
	menuItem1 := NewMenuItem("Order1", "desc1", 12)
	menuItem2 := NewMenuItem("Order2", "desc2", 15)
	menu.Add(menuItem1)
	menu.Add(menuItem2)
	if len(menu.child) != 2 {
		t.Errorf("number error,should be 2")
	}
}

func TestMenuGroup_Price(t *testing.T) {
	menu := NewMenu("Order", "ordernew")
	menuItem1 := NewMenuItem("Order1", "desc1", 12)
	menuItem2 := NewMenuItem("Order2", "desc2", 15)
	menu.Add(menuItem1)
	menu.Add(menuItem2)
	if menu.Price() != 27 {
		t.Errorf("price error,should be 17")
	}
}

func TestMenuGroup_Print(t *testing.T) {
	menu := NewMenu("Order", "ordernew")
	menuItem1 := NewMenuItem("Order1", "desc1", 12)
	menuItem2 := NewMenuItem("Order2", "desc2", 15)
	menu.Add(menuItem1)
	menu.Add(menuItem2)
	menu.Print()
}
