package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	//"encoding/hex"
	"encoding/base64"
)
const (
    zero  = byte('0')
    one   = byte('1')
    lsb   = byte('[') // left square brackets
    rsb   = byte(']') // right square brackets
    space = byte(' ')
)


func test(val []byte)[]int8{
	var res []int8
	for _, a := range val{
		res = append(res, int8(a))
	}
	return res
}

func BytesToBinaryString(bs []byte) string {
    l := len(bs)
    bl := l*8 + l + 1
    buf := make([]byte, 0, bl)
    buf = append(buf, lsb)
    for _, b := range bs {
        buf = appendBinaryString(buf, b)
        buf = append(buf, space)
    }
    buf[bl-1] = rsb
    return string(buf)
}

// append bytes of string in binary format.
func appendBinaryString(bs []byte, b byte) []byte {
    var a byte
    for i := 0; i < 8; i++ {
        a = b
        b <<= 1
        b >>= 1
        switch a {
        case b:
            bs = append(bs, zero)
        default:
            bs = append(bs, one)
        }
        b <<= 1
    }
    return bs
}

func main() {
	tmp := "e82245d99cc36ced6dce432ed9211987"
	APP_KEY_DeBase64, _ := base64.StdEncoding.DecodeString(tmp)
	str := "access_token=L0Z0FwhhZy9tqBVS7KusByDr7cFGtkvMf3cX5WwWBMwsLDf29mY09w==client_id=3667fomat=jsonopen_id=5e2vJlH3FYmjGmO4j4+ty9sW69VZXDFrgOsqjvwfUKc6EJWhk2WAAA==state=1504254261639378000"
	mac := hmac.New(sha1.New, APP_KEY_DeBase64)
	mac.Write([]byte(str))
	//c := hex.EncodeToString(mac.Sum([]byte("")))
	//fmt.Println(c)
	byte_tmp := mac.Sum(nil)
	fmt.Println(byte_tmp)
	b := BytesToBinaryString(byte_tmp)
	fmt.Println(b)

	a := test(mac.Sum(nil))
	fmt.Println(a)

}

var signum int
var mag []int

func BigInteger(val []int8) string{
	if val[0] < 0{
		mag = makePositive(val);
            	signum = -1;
	}
	return ""
}

func makePositive(a []int8)[]int{
	var keep, k int
	byteLength := len(a)
	for keep=0; keep < byteLength && a[keep] == -1; keep++{}
	for k=keep; k < byteLength && a[k] == 0; k++{}
	var extraByte int
	if k == byteLength{
		extraByte = 1
	}else{
		extraByte = 0
	}
	intLength := ((byteLength - keep + extraByte) + 3) >> 2
	var result []int
	b := byteLength - 1
	for i:=intLength-1; i>=0; i--{
		b = b -1
		result[i] = a[b] & 0xff
		var numBytesToTransfer int
		if b-keep+1 > 3{
			numBytesToTransfer = 3
		}else{
			numBytesToTransfer = b-keep+1
		}
		for j:=8; j <= 8*numBytesToTransfer; j+=8{
			b = b - 1
			result[i] = result[i] | ((a[b] & 0xff) << j)
		}
		mask := -1 >> (8 * (3 - numBytesToTransfer))
		result[i] = ^result[i] & mask

	}
	LONG_MASK := 0xffffffff
	for i:=len(result)-1; i>=0; i--{
		result[i] = int((result[i] & LONG_MASK) + 1)
		if result[i] != 0{
			break
		}
	}
	return result
}
