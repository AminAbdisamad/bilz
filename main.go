package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	// Capturing values when user hits enter
	name, err := r.ReadString('\n')
	// Remove Spaces from User Typed String
	name = strings.TrimSpace(name)
	return name, err

}


func createBill()Bill{
	// Using bufio to get user typed string
	reader:= bufio.NewReader(os.Stdin)
	name, err := getInput("Enter Your Name: ", reader)
	// Print if there's error
	if err != nil {
		fmt.Println("Error: ",err)
	}
	// calling create bill func
	b := newBill(name)
	fmt.Println("Creating new bill for - ", b.name)
	return b
}

func promptOptions(b Bill){
	var reader = bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose options ( a - add item, s - save the bill, p - add a tip: ", reader)
	fmt.Println(opt)
}
func main(){
	// fmt.Println("Hello there")
	// bill := newBill("Mohamed")
	// bill.updateTip(10)
	// bill.addItem("Milk",3.44)
	// bill.addItem("Soup",5.48)
	// bill.addItem("Yogurt",3.94)
	// fmt.Println(bill.format())
	bill:= createBill()
	promptOptions(bill)
}