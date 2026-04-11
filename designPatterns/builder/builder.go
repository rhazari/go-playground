package builder

// House represents the complex object we're building.
type House struct {
	Floors  int
	Windows int
	Garage  bool
}

// HouseBuilder is the builder for constructing House objects.
type HouseBuilder struct {
	house House
}

// NewHouseBuilder initializes a new HouseBuilder with default values.
func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{
		house: House{
			Floors:  1, // Default: single floor
			Windows: 4, // Default: 4 windows
			Garage:  false,
		},
	}
}

// SetFloors sets the number of floors.
func (b *HouseBuilder) SetFloors(floors int) *HouseBuilder {
	b.house.Floors = floors
	return b
}

// SetWindows sets the number of windows.
func (b *HouseBuilder) SetWindows(windows int) *HouseBuilder {
	b.house.Windows = windows
	return b
}

// SetGarage sets whether the house has a garage.
func (b *HouseBuilder) SetGarage(hasGarage bool) *HouseBuilder {
	b.house.Garage = hasGarage
	return b
}

// Build returns the constructed House object.
func (b *HouseBuilder) Build() House {
	return b.house
}