package Library

import (
	"fmt"
	"testing"
)

func Read(mbc *MBClient) {
	//ReadCoil := mbc.ReadCoil(1, 5, 10)
	//fmt.Println(ReadCouil)
	fmt.Println("dd")
}
func fun1() int {
	return 100
}

func TestRead(t *testing.T) {
	//mbc := NewClient("127.0.0.1", 502)
	//mbc.Open()
	//Read(mbc)

	if fun1() != 100 {
		t.Error("fun1 is return 100")
	}
}
