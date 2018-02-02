package bdcrypto

import (
	"fmt"
	"testing"
)

func TestRSAEncryptOfWapBaidu(t *testing.T) {
	b, _ := RSAEncryptOfWapBaidu(DefaultRSAPublicKeyModulus, []byte("123"))
	fmt.Println(b)
	fmt.Println("-------------------")
}

func BenchmarkRSAEncryptOfWapBaidu(b *testing.B) {
	var by = []byte("Pythonphp123sdif8e83")
	for i := 0; i < b.N; i++ {
		RSAEncryptOfWapBaidu(DefaultRSAPublicKeyModulus, by)
	}
}

func TestBase64(t *testing.T) {
	var b = []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	en := Base64Encode(b)
	fmt.Println(string(en))
	de := Base64Decode(en)
	fmt.Println(string(de))
}

func TestReverse(t *testing.T) {
	var b = []byte("1234567890")
	fmt.Println(string(BytesReverse(b)))
}

func BenchmarkReverse(b *testing.B) {
	var by = []byte("Pythonphp123sdif8e83")
	for i := 0; i < b.N; i++ {
		BytesReverse(by)
	}
}
