package builder

import (
	"fmt"
	"testing"
)

func Test_Builder(t *testing.T) {
	// Usage: Build a basic house
	basicHouse := NewHouseBuilder().Build()
	fmt.Printf("Basic House: %+v\n", basicHouse) // Output: Basic House: {Floors:1 Windows:4 Garage:false}

	// Usage: Build a custom house with chaining
	customHouse := NewHouseBuilder().
		SetFloors(2).
		SetWindows(8).
		SetGarage(true).
		Build()
	fmt.Printf("Custom House: %+v\n", customHouse) // Output: Custom House: {Floors:2 Windows:8 Garage:true}
}