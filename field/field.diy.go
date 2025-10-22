package field

// Between ...
func (field genericsField[T]) Between(left T, right T) Expr {
	return field.between([]any{left, right})
}

// NotBetween ...
func (field genericsField[T]) NotBetween(left T, right T) Expr {
	return Not(field.Between(left, right))
}

// Add ...
func (field genericsField[T]) Add(value T) genericsField[T] {
	return newGenerics[T](field.add(value))
}

// Sub ...
func (field genericsField[T]) Sub(value T) genericsField[T] {
	return newGenerics[T](field.sub(value))
}

// Mul ...
func (field genericsField[T]) Mul(value T) genericsField[T] {
	return newGenerics[T](field.mul(value))
}

// Div ...
func (field genericsField[T]) Div(value T) genericsField[T] {
	return newGenerics[T](field.div(value))
}

// Mod ...
func (field genericsField[T]) Mod(value T) genericsField[T] {
	return newGenerics[T](field.mod(value))
}

// FloorDiv ...
func (field genericsField[T]) FloorDiv(value T) genericsField[T] {
	return newGenerics[T](field.floorDiv(value))
}

// Floor ...
func (field genericsField[T]) Floor() genericsField[T] {
	return newGenerics[T](field.floor())
}
