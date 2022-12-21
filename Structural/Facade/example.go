package main

func main() {
	// Use the Facade Cafe API to create coffee drinks
	// instead of directly interacting with the complex Coffee API

	makeAmericano(8.0)

	// Make a 12 ounce Latte
	makeLatte(12.0, true)
}
