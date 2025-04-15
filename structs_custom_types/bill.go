package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/google/uuid"
)

// the bill structure
type Bill struct {
	BillID   uuid.UUID // a UUID type from the google/uuid package
	BillDate string //time.Time // a time type from the time package
	BillName string
	BillItems map[string]float64
	BillTotal float64
	BillPaid  bool
	BillPaidDate *string
}

// make bill func

func newBill(billName string) Bill{
	b := Bill{
		BillName: billName,
		BillID: uuid.New(),
		BillItems: make(map[string]float64),
		BillTotal: 0.0,
		BillPaid: false,
		BillDate: time.Now().Format("2006-01-02 15:04:05"),
		BillPaidDate: nil,
	}
	return b

}

//reciever functions (methods) for the bill struct

// diff is the struct/tpe it will be attached to

func (b *Bill) addItem(){
	reader := bufio.NewReader(os.Stdin)

	itemName, _ := getInput(reader, "Enter the name of the item: ") // updated to use getInput
	price, _ := getInput(reader, "Price:") // updated to use getInput

	itemPrice, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Invalid price. Please enter a valid number.")
		return
	}

	b.BillItems[itemName] = itemPrice
	b.BillTotal += itemPrice
}


func (b *Bill) format() string {
	// format the bill as a string
	paid := "No"
	if b.BillPaid {
		paid = "Yes"
	}

	billString := "\n\nBill Name: " + b.BillName + "\n"
	billString += "Bill ID: " + b.BillID.String() + "\n"
	billString += "Bill Date: " + b.BillDate + "\n"
	billString += "Bill Total: " + fmt.Sprintf("%.2f",b.BillTotal) + "\n"
	billString += "Bill Paid: " + paid + "\n"

	if b.BillPaidDate != nil {
		billString += "Bill Paid Date: " + fmt.Sprintf("%v", *b.BillPaidDate) + "\n"
	} else {
		billString += "Bill Paid Date: Not Paid\n"
	}

	billString += "Bill Items: \n"

	for item, price := range b.BillItems {
		billString += fmt.Sprintf("%-25v",item) + fmt.Sprintf("%.2f",price) + "\n"
	}

	return billString + "\n"
}

func (b *Bill) payBill() {
	b.BillPaid = true

	b.BillPaidDate = new(string)
	now := time.Now().Format("2006-01-02 15:04:05")
	b.BillPaidDate = &now
}

func (b *Bill) saveBill(){
	err := os.MkdirAll("bills", 0755)
	if err != nil {
		fmt.Println("Error creating bills directory:", err)
		return
	}

	data := []byte(b.format())




	fileName := fmt.Sprintf("bills/%s.txt", b.BillID.String())
	err = os.WriteFile(fileName, data, 0644)
}
