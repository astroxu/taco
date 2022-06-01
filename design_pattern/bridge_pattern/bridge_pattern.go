/**
 * @Author: Sean
 * @Date: 2022/6/1 10:00
 */

package bridge_pattern

import "fmt"

/*
分离抽象部分和实现部分， 桥接设计的核心在于抽象接口和组合抽象接口的结构体
设计思想：
	1. 抽象接口，(实现该接口的具体struct，可扩展多个struct)
	2. 属性为抽象接口的struct Phone（桥接层）
	3. 与Phone组合模式的具体struct（可以是多个struct）
*/

// Software 抽象接口
type Software interface {
	Run()
}

//具体类型CPU和Storage
type Cpu struct{}

func (c *Cpu) Run() {
	fmt.Println("this.is cpu run")
}

type Storage struct{}

func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

// 桥接层结构体
type Phone struct {
	software Software
}

// 赋值具体Software
func (s *Phone) SetSoftware(soft Software) {
	s.software = soft
}

// Apple 结构体
type Apple struct {
	phone Phone
}

func (p *Apple) SetShape(soft Software) {
	p.phone.SetSoftware(soft)
}

func (p *Apple) Print() {
	p.phone.software.Run()
}

// Huawei 结构体
type Huawei struct {
	phone Phone
}

func (p *Huawei) SetShape(soft Software) {
	p.phone.SetSoftware(soft)
}

func (p *Huawei) Print() {
	p.phone.software.Run()
}
