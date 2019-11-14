package spaceadventure

// When applied

import "fmt"
// "Shall I randomly choose a planet for you to visit? (Y or N)"
func Start(planetarySystem PlanetarySystem){
   printWelcome(planetarySystem)
   PrintGreeting(getResponse("What is your name?"))
   fmt.Println("Let's go on an adventure!")
   travel(promptForRandomOrSpecificDestination("Shall I randomly choose a planet for you to visit? Y or N")
}

func main() {
	printWelcome()
	printGreeting(getName())
	fmt.Println("Let's go on an adventure!")
	// travel()
}

func printWelcome(planetarySystem PlanetarySystem) {
	fmt.Println("Welcome to the %s!/n", planetarySystem.Name)
	fmt.Printf("There are %d planets to explore./n", planetarySystem.NumberOfPlanets())
}

func getName() string {
	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	return name
}

func printGreeting(name string) {
	fmt.Printf("Nice to meet you, %s. My name is Eliza, I'm an old friend of Siri.\n", name)
}

func travel(randomDestination bool) {
	if (randomDestination) {
		travelToPlanet()
	} else {
		travelToPlanet(getResponse("Name the planet you would like to visit?"))
	}
}

func promptForRandomOrSpecificDestination(prompt string) bool {
	var choice string
	for choice != "Y" && choice != "N" {
		choice = getResponse(prompt)
		if choice == "Y" {
			return true
		} else if choice == "N" {
			return false
		} else {
			fmt.Println("Sorry, I didn't get that.")
		}
	}
	return false
}

func getResponse(prompt string) string {
	var response string
	fmt.Println(prompt)
	fmt.Scan(&response)
	return 
}

func travelToRandomPlanet() {
	fmt.Println("Traveling to Jupiter...")
	fmt.Println("Arrived at Jupiter. The large red spot appears ominous.")
}

func travelToPlanet(planetName string) {
	fmt.Printf("Traveling to %s...\n", planetName)
	fmt.Println("Arrived at Neptune. A very cold planet, furthest from the sun.")
}
