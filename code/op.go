package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	random_date := currentTime.AddDate(-24, 0, 0)
	fmt.Println("current day:", currentTime.Format("2006-1-02"), "| random date:", random_date.Format("2006-1-02"))
}
