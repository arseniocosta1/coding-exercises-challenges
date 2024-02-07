// Package distconv provides tools for converting between different distance
// unit
package distconv

import (
	"fmt"
	"slices"
)

type UnitFactoryFunc func(float64) Unit

// Registry map
var unitFactories = map[string]UnitFactoryFunc{
	"ft": func(value float64) Unit {
		return Feet(value)
	},

	"m": func(value float64) Unit {
		return Meter(value)
	},

	"mi": func(value float64) Unit {
		return Mile(value)
	},

	"pc": func(value float64) Unit {
		return Parsec(value)
	},
}

// RegisterUnitFactory register a new distance Unit
func RegisterUnitFactory(unit string, factory UnitFactoryFunc) {
	unitFactories[unit] = factory
}

type Unit interface {
	ToMeter() Meter

	Ratio() float64

	FromMeter(meter Meter) Unit
}

type Feet float64

func (f Feet) Ratio() float64 {
	return 0.3048
}

func (f Feet) FromMeter(meter Meter) Unit {
	return Feet(meter / 0.3048)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (f Feet) ToMeter() Meter {
	return Meter(f * 0.3048)
}

type Meter float64

func (m Meter) Ratio() float64 {
	return 1
}

func (m Meter) FromMeter(meter Meter) Unit {
	return meter
}

func (m Meter) ToMeter() Meter {
	return m
}

func (m Meter) String() string {
	return fmt.Sprintf("%.5fm", m)
}

type Mile float64

func (m Mile) Ratio() float64 {
	return 1609.344
}

func (m Mile) FromMeter(meter Meter) Unit {
	return Mile(meter / 1609.344)
}

func (m Mile) ToMeter() Meter {
	return Meter(m * 1609.344)
}

func (m Mile) String() string {
	return fmt.Sprintf("%gmi", m)
}

type Parsec float64

func (p Parsec) Ratio() float64 {
	return 3.08567758149137e16
}

func (p Parsec) FromMeter(meter Meter) Unit {
	return Parsec(meter / Meter(p.Ratio()))
}

func (p Parsec) String() string {
	return fmt.Sprintf("%gpc", p)
}

func (p Parsec) ToMeter() Meter {
	return Meter(p * 3.08567758149137e16)
}

func NewUnit(value float64, unit string) (Unit, error) {
	if factory, ok := unitFactories[unit]; ok {
		return factory(value), nil
	}
	return nil, fmt.Errorf("unsupported Unit: %s", unit)
}

func Convert(u Unit, to string) (Unit, error) {
	if factory, ok := unitFactories[to]; ok {
		unit := factory(float64(u.ToMeter()))
		return factory(float64(u.ToMeter()) / unit.Ratio()), nil
	}
	return nil, fmt.Errorf("unsupported Unit: %s", to)
}

func AllUnitConversions(u Unit) []Unit {
	meter := u.ToMeter()

	var units []Unit

	keys := make([]string, 0, len(unitFactories))
	for k := range unitFactories {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, k := range keys {
		f := unitFactories[k](0)
		units = append(units, f.FromMeter(meter))
	}

	return units
}
