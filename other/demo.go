package main

// 并发安全 使用channal

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Cash int
}
type Transfer struct {
	Sender    *User
	Recipient *User
	Amount    int
}

func sendCashHandler(transferchan chan Transfer) {
	var val Transfer
	for {
		val = <-transferchan
		val.Sender.sendCash(val.Recipient, val.Amount)
	}
}

func (u *User) sendCash(to *User, amount int) bool {
	if u.Cash < amount {
		return false
	}
	/* 设置延迟Sleep，当多个goroutines并行执行时,便于进行数据安全分析 */
	time.Sleep(500 * time.Millisecond)
	u.Cash = u.Cash - amount
	to.Cash = to.Cash + amount
	return true
}

func main() {
	me := User{Cash: 500}
	you := User{Cash: 500}
	transferchan := make(chan Transfer)
	go sendCashHandler(transferchan)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transfer := Transfer{Sender: &me, Recipient: &you, Amount: 50}
		transferchan <- transfer
		fmt.Fprintf(w, "I have $%d\n", me.Cash)
		fmt.Fprintf(w, "You have $%d\n", you.Cash)
		fmt.Fprintf(w, "Total transferred: $%d\n", (you.Cash - 50))
	})
	http.ListenAndServe(":8089", nil)
}