package main

import "fmt"

//UrbanGarden is a struct which contains the properties of an Urban Garden
type UrbanGarden struct {
	Name       string //Name of the Garden
	Area       float64 //Size of the Garden (sq. ft)
	Location   string //Location of the Garden
	Plants     []string //List of Plants
	AirQuality int //Air Quality Rating
	Aesthetic  string //Aesthetic Appeal
}

//CreateGarden is a function which creates an Urban Garden with given properties
func CreateGarden(name string, area float64, location string, plants []string, airQ int, aesthetic string) *UrbanGarden {
	return &UrbanGarden{name, area, location, plants, airQ, aesthetic}
}

//IncreaseAirQuality is a function which increases the air quality rating of an Urban Garden
func (garden *UrbanGarden) IncreaseAirQuality() {
	garden.AirQuality += 5
}

//AddPlant is a function which adds a plant to Urban Garden
func (garden *UrbanGarden) AddPlant(plant string) {
	garden.Plants = append(garden.Plants, plant)
}

//EnhanceAesthetic is a function which enhances the aesthetic appeal of an Urban Garden
func (garden *UrbanGarden) EnhanceAesthetic(aesthetic string) {
	garden.Aesthetic = aesthetic
}

//DisplayGarden is a function which displays the properties of an Urban Garden
func (garden *UrbanGarden) DisplayGarden() {
	fmt.Printf("Name: %v\n", garden.Name)
	fmt.Printf("Area: %v sq. ft.\n", garden.Area)
	fmt.Printf("Location: %v\n", garden.Location)
	fmt.Printf("Plants: %v\n", garden.Plants)
	fmt.Printf("Air Quality: %v\n", garden.AirQuality)
	fmt.Printf("Aesthetic Appeal: %v\n", garden.Aesthetic)
}

func main() {
	// Creating an Urban Garden
	myGarden := CreateGarden("MyUrbanGarden", 200, "NewYork", []string{"Rose", "Lilac"}, 70, "Beautiful")

	//Increasing the air Quality
	myGarden.IncreaseAirQuality()

	//Adding Plants
	myGarden.AddPlant("Sunflower")

	//Enhancing the Aesthetic Appeal
	myGarden.EnhanceAesthetic("Stunning")

	//Displaying the properties
	myGarden.DisplayGarden()
}