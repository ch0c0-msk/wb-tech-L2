package pattern

import "fmt"

// Product interface

type Product interface {
	Use()
}

// Concrete products
type Knife struct{}

func (k *Knife) Use() {
	fmt.Println("Knife is using")
}

type Spoon struct{}

func (s *Spoon) Use() {
	fmt.Println("Spoon is using")
}

// Factory interface
type Factory interface {
	NewProduct() Product
}

// Factory implementations
type KnifeFactory struct{}

func (f *KnifeFactory) NewProduct() Product {
	return &Knife{}
}

type SpoonFactory struct{}

func (f *SpoonFactory) NewProduct() Product {
	return &Spoon{}
}
