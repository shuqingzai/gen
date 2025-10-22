package field

import (
	"database/sql/driver"
)

// Between ...
func (field Field) Between(left driver.Valuer, right driver.Valuer) Expr {
	return field.between([]interface{}{left, right})
}

// NotBetween ...
func (field Field) NotBetween(left driver.Valuer, right driver.Valuer) Expr {
	return Not(field.Between(left, right))
}

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

// Mod ...
func (field Field) Mod(value driver.Valuer) Field {
	return Field{field.mod(value)}
}

// FloorDiv ...
func (field Field) FloorDiv(value driver.Valuer) Field {
	return Field{field.floorDiv(value)}
}

// Floor ...
func (field Field) Floor() Field {
	return Field{field.floor()}
}
