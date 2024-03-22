package main

import (
	"fmt"
	"time"
)

type acount struct {
	name    string
	balance int64
}

func transfer(source, dest *acount, amount int64) {
	if source.balance < amount {
		fmt.Printf("❌: %s\n", fmt.Sprintf("%s %s", source, dest))
		return
	}
	time.Sleep(time.Second)

	dest.balance += amount
	source.balance -= amount

	fmt.Printf("✅: %s\n", fmt.Sprintf("%s %s", source, dest))
}

func main() {

	andresAcount := acount{name: "Andres", balance: 500}
	yefriAcount := acount{name: "yefri", balance: 900}

	for _, value := range []int{300, 100} {
		transfer(&andresAcount, &yefriAcount, int64(value))
	}

}

func (a acount) String() string {
	return fmt.Sprintf("{%s %d}", a.name, a.balance)
}
