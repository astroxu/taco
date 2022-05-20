/**
 * @Author: Sean
 * @Date: 2022/5/20 15:36
 */

package abstract_factory_pattern

import (
	"testing"
)

func TestConcreteFactory_CreateProduct(t *testing.T) {
	conFactory := &ConcreteFactory{}

	product := conFactory.CreateProduct()

	conProduct := product.(*ConcreteProduct)
	if conProduct.Name != "KG" {
		t.Errorf("abstract factory can not create the concrete product")
	}
}
