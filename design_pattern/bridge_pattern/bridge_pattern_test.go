/**
 * @Author: Sean
 * @Date: 2022/6/1 10:10
 */

package bridge_pattern

import "testing"

func TestCpu_Run(t *testing.T) {
	cpu := &Cpu{}
	apple := new(Apple)
	apple.SetShape(cpu)
	apple.Print()
}

func TestStorage_Run(t *testing.T) {
	storage := &Storage{}
	hw := new(Huawei)
	hw.SetShape(storage)
	hw.Print()
}
