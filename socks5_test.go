package go_socks5

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOne(t *testing.T) {
	a := 0x05
	fmt.Printf("%d %T \n",a,a)

	ac := []byte{0x05,0x00,0x04,0x09,0x11,0x01}
	fmt.Println(int(ac[len(ac)-2])<<8)
	fmt.Println(int(ac[len(ac)-1]))
	fmt.Println(ac[len(ac)-2])
	port := strconv.Itoa(int(ac[len(ac)-2])<<8 | int(ac[len(ac)-1]))
	fmt.Println(port)

	fmt.Println(int(4352 | 1))
}