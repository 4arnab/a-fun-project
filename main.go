// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// var reader = bufio.NewReader(os.Stdin)

// type User struct { // struct in go land are equivalent to objects in javascript
// 	FirstName string
// 	LastName  string
// 	BirthDate string
// 	Created   time.Time
// }

// func (user User) OutPutUserData() {
// 	fmt.Printf("%v", user.FirstName)
// }

// func main() {

// 	var newUser User // type declaration for user struct

// 	firstName := getUserData("please enter first name")
// 	lastName := getUserData("please enter lastname")
// 	birthDate := getUserData("please enter birthdate")
// 	created := time.Now()

// 	newUser = User{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		BirthDate: birthDate,
// 		Created:   created,
// 	}

// 	fmt.Println(newUser)

// }

// func getUserData(promtText string) string {
// 	fmt.Println(promtText)

// 	userInput, _ := reader.ReadString('\n')
// 	return strings.TrimSpace(userInput)
// }
// func double(number *int) int {
// 	return *number * 2
// }

// func newUser() *User {

// 	return &User{
// 		FirstName: "",
// 		LastName:  "",
// 		BirthDate: "",
// 		Created:   time.Now(),
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Id               string
	Title            string
	ShortDescription string
	Price            float64
}

func (p *Product) printInformation() {
	fmt.Printf("Product Title : %v\nDescription: %v\nPrice: %.2f\n", p.Title, p.ShortDescription, p.Price)
}

func (p *Product) store() {
	file, _ := os.Create(p.Id + ".txt")

	content := fmt.Sprintf("Product Title:%v\nDescription:%v\nPrice:%.2f\n", p.Title, p.ShortDescription, p.Price)
	file.WriteString(content)
	file.Close()
}

func main() {

	var createdProduct = getProductInformation()
	createdProduct.printInformation()
	createdProduct.store()
}

func createProduct(title, shortDescription, Id string, price float64) *Product {
	return &Product{
		Id:               Id,
		Title:            title,
		ShortDescription: shortDescription,
		Price:            price,
	}
}

func getProductInformation() *Product {
	fmt.Println("Please enter the product information\n ------------------")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var productId = getInput("ID: ", reader)
	var productName = getInput("Product name: ", reader)
	var description = getInput("Description: ", reader)

	// formatting the product price
	var priceInput = getInput("Price: ", reader)
	var price, _ = strconv.ParseFloat(priceInput, 64)

	return createProduct(productName, description, productId, price)
}

func getInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)

	enteredText, _ := reader.ReadString('\n')
	return strings.TrimSpace(enteredText)
}
