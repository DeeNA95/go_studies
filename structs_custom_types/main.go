package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput( reader *bufio.Reader, prompt string) (string, error){
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

//reader uses bufio
func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput(reader, "Enter the name of the customer: ")

	b := newBill(name)
	fmt.Println(b.BillName,"bill created successfully!")
	return b
}

func promptOptions(b *Bill) bool {
	fmt.Println("1. Add item")
	fmt.Println("2. Pay bill")
	fmt.Println("3. Void bill")
	fmt.Println("4. Save Bill")

	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput(reader, "Select an option: ")

	switch option {
	case "1":
		b.addItem()
		return true
	case "2":
		fmt.Println(b.format())
		b.payBill()
		return false
	case "3":
		return false
	case "4":
		fmt.Println("Saving bill...")
		b.saveBill()
		fmt.Println("Bill saved successfully!")
		return false
	default:
		fmt.Println("\nInvalid option. Please try again.\n")
		return true
	}
}


func main() {
	bill1 := createBill()
	a := true

	fmt.Println(bill1.format())

	for a {
		a = promptOptions(&bill1)
	}

	if bill1.BillPaid {
	fmt.Println(bill1.format())
	}

}
