/**
 * @Author: Sean
 * @Date: 2022/5/5 16:18
 */

package abstract_factory_pattern

import "fmt"

/*
设计思想
	1. 抽象工厂接口
	2. 抽象产品接口
	3. 具体的工厂和产品struct
	4. 使用具体的工厂来创建产品，并返回接口类型值
*/

// 抽象工厂模式（Abstract Factory Pattern）是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

// 在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。
type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

// 具体的产品
type ConcreteProduct struct {
	Name string
}

func (conProduct *ConcreteProduct) Describe() {
	fmt.Println(conProduct.Name)
}

// 具体的工厂
type ConcreteFactory struct{}

func (conFactory *ConcreteFactory) CreateProduct() Product {
	return &ConcreteProduct{Name: "KG"}
}
