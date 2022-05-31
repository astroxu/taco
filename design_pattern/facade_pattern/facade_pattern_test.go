/**
 * @Author: Sean
 * @Date: 2022/5/31 15:37
 */

package facade_pattern

import (
	"testing"
)

func TestNewFacade(t *testing.T) {
	music := Music{"love"}
	video := Video{1}
	count := Count{12, 30, 5}
	facade := NewFacade(music, count, video)
	facade.PrintServerInfo()
}
