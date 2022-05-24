/**
 * @Author: Sean
 * @Date: 2022/5/24 14:34
 */

package proxy_pattern

import "testing"

func TestProxyObject(t *testing.T) {
	obj := &Object{action: "run"}

	proxyObj := new(ProxyObject)
	proxyObj.object = obj
	proxyObj.ObjDo("run")
}
