package main

import "fmt"

//
// --- Product Family 1: Burger ---
//

type Burger interface {
	Prepare()
}

type BasicBurger struct{}
func (b BasicBurger) Prepare() {
	fmt.Println("Preparing Basic Burger with bun, patty, and ketchup!")
}

type StandardBurger struct{}
func (s StandardBurger) Prepare() {
	fmt.Println("Preparing Standard Burger with bun, patty, cheese, and lettuce!")
}

type PremiumBurger struct{}
func (p PremiumBurger) Prepare() {
	fmt.Println("Preparing Premium Burger with gourmet bun, premium patty, cheese, lettuce, and secret sauce!")
}

type BasicWheatBurger struct{}
func (b BasicWheatBurger) Prepare() {
	fmt.Println("Preparing Basic Wheat Burger with bun, patty, and ketchup!")
}

type StandardWheatBurger struct{}
func (s StandardWheatBurger) Prepare() {
	fmt.Println("Preparing Standard Wheat Burger with bun, patty, cheese, and lettuce!")
}

type PremiumWheatBurger struct{}
func (p PremiumWheatBurger) Prepare() {
	fmt.Println("Preparing Premium Wheat Burger with gourmet bun, premium patty, cheese, lettuce, and secret sauce!")
}

//
// --- Product Family 2: Garlic Bread ---
//

type GarlicBread interface {
	Prepare()
}

type BasicGarlicBread struct{}
func (b BasicGarlicBread) Prepare() {
	fmt.Println("Preparing Basic Garlic Bread with butter and garlic!")
}

type CheeseGarlicBread struct{}
func (c CheeseGarlicBread) Prepare() {
	fmt.Println("Preparing Cheese Garlic Bread with extra cheese and butter!")
}

type BasicWheatGarlicBread struct{}
func (b BasicWheatGarlicBread) Prepare() {
	fmt.Println("Preparing Basic Wheat Garlic Bread with butter and garlic!")
}

type CheeseWheatGarlicBread struct{}
func (c CheeseWheatGarlicBread) Prepare() {
	fmt.Println("Preparing Cheese Wheat Garlic Bread with extra cheese and butter!")
}

//
// --- Abstract Factory Interface ---
//

type MealFactory interface {
	CreateBurger(burgerType string) Burger
	CreateGarlicBread(breadType string) GarlicBread
}

//
// --- Concrete Factory: SinghBurger ---
//

type SinghBurger struct{}

func (s SinghBurger) CreateBurger(burgerType string) Burger {
	switch burgerType {
	case "basic":
		return BasicBurger{}
	case "standard":
		return StandardBurger{}
	case "premium":
		return PremiumBurger{}
	default:
		fmt.Println("Invalid burger type!")
		return nil
	}
}

func (s SinghBurger) CreateGarlicBread(breadType string) GarlicBread {
	switch breadType {
	case "basic":
		return BasicGarlicBread{}
	case "cheese":
		return CheeseGarlicBread{}
	default:
		fmt.Println("Invalid garlic bread type!")
		return nil
	}
}

//
// --- Concrete Factory: KingBurger ---
//

type KingBurger struct{}

func (k KingBurger) CreateBurger(burgerType string) Burger {
	switch burgerType {
	case "basic":
		return BasicWheatBurger{}
	case "standard":
		return StandardWheatBurger{}
	case "premium":
		return PremiumWheatBurger{}
	default:
		fmt.Println("Invalid burger type!")
		return nil
	}
}

func (k KingBurger) CreateGarlicBread(breadType string) GarlicBread {
	switch breadType {
	case "basic":
		return BasicWheatGarlicBread{}
	case "cheese":
		return CheeseWheatGarlicBread{}
	default:
		fmt.Println("Invalid garlic bread type!")
		return nil
	}
}

//
// --- Main ---
//

func main() {
	burgerType := "basic"
	garlicBreadType := "cheese"

	var mealFactory MealFactory
	mealFactory = KingBurger{} // choose SinghBurger{} or KingBurger{}

	burger := mealFactory.CreateBurger(burgerType)
	garlicBread := mealFactory.CreateGarlicBread(garlicBreadType)

	if burger != nil {
		burger.Prepare()
	}
	if garlicBread != nil {
		garlicBread.Prepare()
	}
}
