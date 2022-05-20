/**
 * @Author: Sean
 * @Date: 2022/5/20 15:20
 */

package generator_pattern

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	ch := Count(1, 50)

	for i := range ch {
		fmt.Println(i)
	}
}
