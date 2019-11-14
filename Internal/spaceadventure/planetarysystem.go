package spaceadventure

type PlanetarySystem struct {
	Name string
	//dynamic array in go is called slice
	Planets []Planet
}

func (ps PlanatarySystem) NumberOfPlanets() int{
	return len(ps.Planets)
}
