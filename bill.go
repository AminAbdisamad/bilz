package main

import (
	"fmt"
	"os"
)


type Bill struct {
	name string
	items map[string]float64
	tip float64


}

func newBill(name string)Bill {
	nb := Bill{
		name: name,
		items:map[string]float64{},
		tip: 0,
	}


	return nb
}

// Receiver function : are functions related to structs or interfaces
func (b *Bill) format()string{
	var total float64 = 0
	var str = "Your bills are\n"
	for key,value := range b.items{
		str += fmt.Sprintf("%-25v ...$%v\n", key+":",value)
		total += value
	}

	str += fmt.Sprintf("%-25v ...$%0.2f\n","Tip:",b.tip)
	str += fmt.Sprintf("%-25v ...$%0.2f\n","Total:",total + b.tip)
	return str

}

func (b *Bill) updateTip(tip float64){
	b.tip = tip
}

func (b *Bill) addItem(name string, price float64){
	b.items[name] = price

}

// Save data
func (b *Bill) saveBill(){
// change data into slice of byte
data := []byte(b.format())
// save data os.WriteFile("fileName.txt",data,permission)
err:= os.WriteFile("data/"+b.name+".txt", data, 0644)
if err != nil{
	panic(err)
}
fmt.Println("Data saved successfully")
}