package main

import "fmt"


type Bill struct {
	name string
	items map[string]float64
	tip float64


}

func newBill(name string)Bill {
	nb := Bill{
		name:name,
		items:map[string]float64{
			"coffee":5.99,
			"tea":3.89,
		},
		tip:0.0,
	}

	return nb
}

// Receiver function : are functions related to structs or interfaces
func (b Bill) format()string{
	var total float64 = 0
	var str = "Your bills are\n"
	for key,value := range b.items{
		str += fmt.Sprintf("%-25v ...$%v\n", key+":",value)
		total += value
	}

	str += fmt.Sprintf("%-25v ...$%v\n","Tip:",b.tip)
	str += fmt.Sprintf("%-25v ...$%0.2f\n","Total:",total + b.tip)
	return str

}

func (b *Bill) updateTip(tip float64){
	b.tip = tip
}

func (b Bill) addItem(name string, price float64){
	b.items[name] = price

}