/**
 * @Author: Sean
 * @Date: 2022/5/29 11:30
 */

package cryptor

import (
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

var (
	key8 = "abcdefgh"
)

func TestDesEcbEncrypt(t *testing.T) {
	desEcbEncrypt := DesEcbEncrypt([]byte(data), []byte(key8))
	desEcbDecrypt := DesEcbDecrypt(desEcbEncrypt, []byte(key8))

	assert := internal.NewAssert(t, "TestDesEcbEncrypt")
	assert.Equal(data, string(desEcbDecrypt))
}

func TestDesCbcEncrypt(t *testing.T) {
	desCbcEncrypt := DesCbcEncrypt([]byte(data), []byte(key8))
	desCbcDecrypt := DesCbcDecrypt(desCbcEncrypt, []byte(key8))

	assert := internal.NewAssert(t, "TestDesCbcEncrypt")
	assert.Equal(data, string(desCbcDecrypt))
}

func TestDesCtrCrypt(t *testing.T) {
	desCtrCrypt := DesCtrCrypt([]byte(data), []byte(key8))
	desCtrDeCrypt := DesCtrCrypt(desCtrCrypt, []byte(key8))

	assert := internal.NewAssert(t, "TestDesCtrCrypt")
	assert.Equal(data, string(desCtrDeCrypt))
}

func TestDesCfbEncrypt(t *testing.T) {
	desCfbEncrypt := DesCfbEncrypt([]byte(data), []byte(key8))
	desCfbDecrypt := DesCfbDecrypt(desCfbEncrypt, []byte(key8))

	assert := internal.NewAssert(t, "TestDesCfbEncrypt")
	assert.Equal(data, string(desCfbDecrypt))
}

func TestDesOfbEncrypt(t *testing.T) {

	desOfbEncrypt := DesOfbEncrypt([]byte(data), []byte(key8))
	desOfbDecrypt := DesOfbDecrypt(desOfbEncrypt, []byte(key8))

	assert := internal.NewAssert(t, "TestDesOfbEncrypt")
	assert.Equal(data, string(desOfbDecrypt))
}
