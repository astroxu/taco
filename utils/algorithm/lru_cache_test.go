/**
 * @Author: Sean
 * @Date: 2022/5/16 18:16
 */

package algorithm

import (
	"github.com/taco/utils/internal"
	"testing"
)

func TestLRUCache(t *testing.T) {
	assert := internal.NewAssert(t, "TestLRUCache")
	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	_, ok := cache.Get(0)
	assert.Equal(false, ok)

	v, ok := cache.Get(1)
	assert.Equal(true, ok)
	assert.Equal(1, v)

	v, ok = cache.Get(2)
	assert.Equal(true, ok)
	assert.Equal(2, v)

	cache.Put(3, 3)
	v, ok = cache.Get(1)
	assert.Equal(false, ok)
	assert.NotEqual(1, v)

	v, ok = cache.Get(3)
	assert.Equal(true, ok)
	assert.Equal(3, v)
}
