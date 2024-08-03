package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	mybill := newBill("mario's bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.updateTip(10)
	mybill.format()

	fmt.Println(mybill)
	fmt.Println(mybill.format())

	// myBill := createBill()
	// promptOptions(myBill)

	// fmt.Println(myBill.format())

	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}

	// mysteryBox := interface{}(10)
	mysteryBox := interface{}("ssss")
	describeValue(mysteryBox)

	retrievedInt, ok := mysteryBox.(int)
	if ok {
		fmt.Println("Retreived int: ", retrievedInt)
	} else {
		fmt.Println("Value is not an integer")
	}

	cVal, err := performCalculation(4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cVal)
	}

}

type Shape interface {
	Area() float64
}
type Measurable interface {
	Perimeter() float64
}
type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	length float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.length
}

type CalculationError struct {
	msg string
}

func (ce CalculationError) Error() string {
	return ce.msg
}

func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid input",
		}
	}
	return math.Sqrt(val), nil
}

func describeGeometry(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}
func describeShape(s Shape) float64 {
	return s.Area()
}

func describeValue(t interface{}) {
	fmt.Printf("Type: %T, Value: %v", t, t)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added: ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added: ", tip)
		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill: ", b)
		b.save()
		fmt.Println("Saved")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}
