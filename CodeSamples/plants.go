package main

import "fmt"
import "errors"
import "strings"

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	status    PlantStatus
	capacity  float64
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) GeneratePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%.0f\n", "Type:", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
		fmt.Printf("%-20s%.0f\n", "Status:", p.status)

		//fmt.Printf("%-20s%.1f\n", "Utilization:", gridLoad/capacity*100)
		println()
	}
}

func GeneratePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func GeneratePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", gridLoad)
	fmt.Printf("%-20s%.1f\n", "Utilization:", gridLoad/capacity*100)

}

func (pg *PowerGrid) GenerateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Load:", pg.load)
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.1f\n", "Utilization:", pg.load/capacity*100)
}

func RequestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")
	fmt.Scanln(&option)

	if option != "1" && option != "2" {
		err = errors.new("Invalid option selected")
	}

	return
}

func main() {
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

	if option, err := RequestOption(); err == nil {
		fmt.Println("")

		switch option {
		case "1":
			grid.GeneratePlantReport()
			//GeneratePlantCapacityReport(plantCapacities)
		case "2":
			// TODO
			//GeneratePowerGridReport(activePlants, plantCapacities, gridLoad)
			grid.GenerateGridReport()
		}
	} else {
		fmt.Println(err.Error())
	}
}
