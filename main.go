package main

import "fmt"

func main(){
	fmt.Println("Hello there")
	bill := newBill("Mohamed")
	bill.updateTip(10)
	bill.addItem("Milk",3.44)
	bill.addItem("Soup",5.48)
	bill.addItem("Yogurt",3.94)
	fmt.Println(bill.format())
}