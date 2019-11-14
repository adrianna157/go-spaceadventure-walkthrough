package main
import "fmt"
import "github.com/adrianna157/go-spaceadventure-walkthrough/Internal/spaceadventure"

func main(){
	
	 spaceadventure.Start(
		spaceadventure.PlanetarySystem{
			Name:"Solar System", Planets: []spaceadventure.Planet{
				spaceadventure.Planet{"Tatooine", "Desert planet"},
			},
  		},
 	)	
}
