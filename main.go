package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("Creating new bill for: ", b.name)
	return b
}

func promptOptions(b Bill){
	var reader = bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose options ( a - add item, s - save the bill, t - add a tip: ", reader)

	// create a swtich
	switch opt {
	case "a":
		name, _ := getInput("Item Name: ",reader)
		price, _:= getInput("Item Price ($s): ",reader)
		p, err := strconv.ParseFloat(price,64)
		if err != nil{
			fmt.Println("Only numbers allowed")
			promptOptions(b)
		}
		// Add Name * Price
		b.addItem(name,p)
		fmt.Println("Item ",name," with a price of $",price, " Added")
		promptOptions(b)
	case "s":
		b.saveBill()
	case "t":
		tip, _ := getInput("Enter a tip ($): ",reader)
		t, err := strconv.ParseFloat(tip,64)
		if err != nil{
			fmt.Println("Only numbers allowed")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip: ($)",tip," Added")
		promptOptions(b)
	default:
		fmt.Println("Invalid Option")
		promptOptions(b)

	}
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
	format:= bill.format()
	fmt.Println(format)
}