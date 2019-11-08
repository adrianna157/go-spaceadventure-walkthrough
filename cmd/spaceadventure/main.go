package main
import "fmt"
import "github.com/adrianna157/go-spaceadventure-walkthrough/Internal/spaceadventure"

func main(){
	var ps = spaceadventure.PlanetarySystem{"Solar System"}
	spaceadventure.Start(ps)
}
