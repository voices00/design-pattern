package main

import "fmt"

//	装饰模式是一种结构型设计模式，
//	允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。

func main() {

	pizza := &VeggeMania{}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
