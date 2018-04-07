package bdcrypto

import (
	"github.com/iikira/BaiduPCS-Go/pcsutil"
)

// BytesReverse 反转字节数组
func BytesReverse(b []byte) []byte {
	length := len(b)
	for i := 0; i < length/2; i++ {
		b[i], b[length-i-1] = b[length-i-1], b[i]
	}
	return b
}

// StringReverse 反转字符串
func StringReverse(s string) string {
	return pcsutil.ToString(BytesReverse(pcsutil.ToBytes(s)))
}
