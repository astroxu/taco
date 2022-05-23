/**
 * @Author: Sean
 * @Date: 2022/5/23 14:13
 */

package cryptor

import (
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

var (
	data = "hello world"
	key  = "abcdefghijklmnop"
)

func TestAesEcbEncrypt(t *testing.T) {
	aesEcbEncrypt := AesEcbEncrypt([]byte(data), []byte(key))
	aesEcbDecrypt := AesEcbDecrypt(aesEcbEncrypt, []byte(key))

	asset := internal.NewAssert(t, "TestAesEcbEncrypt")
	asset.Equal(data, string(aesEcbDecrypt))
}

func TestAesCbcEncrypt(t *testing.T) {
	aesCbcEncrypt := AesCbcEncrypt([]byte(data), []byte(key))
	aesCbcDecrypt := AesCbcDecrypt(aesCbcEncrypt, []byte(key))

	asset := internal.NewAssert(t, "TestAesCbcEncrypt")
	asset.Equal(data, string(aesCbcDecrypt))
}

func TestAesCtrCrypt(t *testing.T) {
	aesCtrCrypt := AesCtrCrypt([]byte(data), []byte(key))
	aesCtrDeCrypt := AesCtrCrypt(aesCtrCrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCtrCrypt")
	assert.Equal(data, string(aesCtrDeCrypt))
}

func TestAesCfbEncrypt(t *testing.T) {
	aesCfbEncrypt := AesCfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesCfbDecrypt(aesCfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesCfbEncrypt")
	assert.Equal(data, string(aesCfbDecrypt))
}

func TestAesOfbEncrypt(t *testing.T) {
	aesOfbEncrypt := AesOfbEncrypt([]byte(data), []byte(key))
	aesCfbDecrypt := AesOfbDecrypt(aesOfbEncrypt, []byte(key))

	assert := internal.NewAssert(t, "TestAesOfbEncrypt")
	assert.Equal(data, string(aesCfbDecrypt))
}
