package main
import (
	"fmt"
)

type Fruit struct{
	name string
	color string
	quantity int

}


func main(){
	m := Fruit{
		name: "pineapple",
		color: "yellow",
		quantity: 21,
	}
	fmt.Println("My favourite fruits is a ", m.name, "and its color is " , m.color,  "and i currently have ",m.quantity, "in my fridge")

	//with a dot notation 
	// var h Fruit
	// 	h.name = "pineapple"
	// 	h.color = "yellow"
	// 	h.quantity = 21
	
	// fmt.Println("My favourite fruits is a ", h.name, "and its color is " , h.color,  "and i currently have ",h.quantity, "in my fridge")


}
