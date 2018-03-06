package bdcrypto

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	bkey      = []byte("asfasfawefwfgagasfasfawefwfgag")
	plaintext = "1111"
)

func TestAesStream(t *testing.T) {
	var key [32]byte
	copy(key[:], bkey)

	rd, err := Aes256CFBEncrypt(key, strings.NewReader(plaintext))
	if err != nil {
		t.Error(err)
	}

	by, _ := ioutil.ReadAll(rd)

	t.Log("encd:", by)

	rd, err = Aes256CFBDecrypt(key, bytes.NewReader(by))
	by, _ = ioutil.ReadAll(rd)
	t.Log("decd:", by)

}

func TestAesBlock(t *testing.T) {
	var key [16]byte
	copy(key[:], bkey)

	ciphertext, err := Aes128ECBEncrypt(key, []byte(plaintext))
	if err != nil {
		t.Log(err)
	}

	t.Log("encd:", ciphertext)

	plain, err := Aes128ECBDecrypt(key, ciphertext)
	if err != nil {
		t.Log(err)
	}

	t.Log("decd:", plain)
}
