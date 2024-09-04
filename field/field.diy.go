package field

import (
	"database/sql/driver"
)

// Add ...
func (field Field) Add(value driver.Valuer) Field {
	return Field{field.add(value)}
}

// Sub ...
func (field Field) Sub(value driver.Valuer) Field {
	return Field{field.sub(value)}
}

// Mul ...
func (field Field) Mul(value driver.Valuer) Field {
	return Field{field.mul(value)}
}

// Div ...
func (field Field) Div(value driver.Valuer) Field {
	return Field{field.div(value)}
}

// FloorDiv ...
func (field Field) FloorDiv(value driver.Valuer) Int {
	return Int{field.floorDiv(value)}
}

// Floor ...
func (field Field) Floor() Int {
	return Int{field.floor()}
}
