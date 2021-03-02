package main

import (
	"fmt"
)

const (
	sp = 8203
)

const (
	_ = string(sp-8101*(((iota+2)%3)>>1)) +
		string(sp-8098*(((iota+2)%3)>>1)) +
		string(sp-8081*(((iota+2)%3)>>1)) +
		string(sp-8081*(((iota+2)%3)>>1)) +
		string(sp-8105*(((iota+4)%5)>>2)) +
		string(sp-8086*(((iota+4)%5)>>2)) +
		string(sp-8081*(((iota+4)%5)>>2)) +
		string(sp-8081*(((iota+4)%5)>>2)) +
		string(sp-(8155-(iota/10%10))*(1-((((iota+2)%3)>>1)|(((iota+4)%5)>>2)|-((iota-10)>>65535)))) +
		string(sp-(8155-(iota%10))*(1-((((iota+2)%3)>>1)|(((iota+4)%5)>>2))))
	_1
	_2
	_3
	_4
	_5
	_6
	_7
	_8
	_9
	_10
	_11
	_12
	_13
	_14
	_15
)

func main() {
	fmt.Println(_1)
	fmt.Println(_2)
	fmt.Println(_3)
	fmt.Println(_4)
	fmt.Println(_5)
	fmt.Println(_6)
	fmt.Println(_7)
	fmt.Println(_8)
	fmt.Println(_9)
	fmt.Println(_10)
	fmt.Println(_11)
	fmt.Println(_12)
	fmt.Println(_13)
	fmt.Println(_14)
	fmt.Println(_15)
}
