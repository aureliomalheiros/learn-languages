package math

type Number interface {
	int | float32
}

func Sum[T Number](x, y T) T{
	return x + y 
}

func Multiple[T Number](x, y T) T{
	return x * y 
}

func Div[T Number](x, y T) T{
	return x / y
}

func Sub[T Number](x, y T) T{
	return x - y
}