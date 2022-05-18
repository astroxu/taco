/**
 * @Author: Sean
 * @Date: 2022/5/18 16:56
 */

package factory_method_pattern

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	k       Kind    = 1
	m       Kind    = 2
	n       Kind    = 3
	balance float32 = 100.00
)

func TestGeneratePayment(t *testing.T) {
	payment, _ := GeneratePayment(k, balance)
	fmt.Println(reflect.TypeOf(payment).Elem().String())
	if reflect.TypeOf(payment).Elem().String() != "factory_method_pattern.CashPay" {
		t.Error("factory method generate error")
	}

	payment, _ = GeneratePayment(n, balance)
	if payment != nil {
		t.Error("factory method params has error")
	}
}

func TestCashpay_Pay(t *testing.T) {
	payment, _ := GeneratePayment(1, balance)
	_ = payment.Pay(20)
	//cash:=reflect.New(reflect.TypeOf(payment).Elem()).Interface().(*CashPay)reflect新的对象
	fmt.Println("reflect interface", reflect.ValueOf(payment).Interface().(*CashPay))
	if cash, ok := payment.(*CashPay); ok {
		fmt.Println(reflect.TypeOf(cash))
		if cash.Balance != float32(80) {
			t.Error("结算错误")
		}
	}
}
