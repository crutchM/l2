package Patterns

import "fmt"

//что-то своего рода полиморфизма, когда нам неважно что мы делаем, нам важно чтобы у того кто это делает был метод вроде execute

type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func NewRestaurant() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {

	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {

	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func perform() {
	r := NewRestaurant()

	tasks := []Command{
		r.MakePizza(1),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	for i, task := range tasks {
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
