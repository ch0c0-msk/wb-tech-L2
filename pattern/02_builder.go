package pattern

import "fmt"

// Director
type FabricPipeline struct {
	builder Builder
}

func NewFabricPipeline(builder Builder) *FabricPipeline {
	return &FabricPipeline{builder: builder}
}

func (f *FabricPipeline) ChangeBuilder(builder Builder) {
	f.builder = builder
}

func (f *FabricPipeline) Construct() {
	f.builder.createPhone()
	f.builder.makeCase()
	f.builder.paint()
}

// Abstract Builder
type Builder interface {
	createPhone()
	makeCase()
	paint()
}

// Concrete Builder
type GreenPhoneBuilder struct {
	phone *Phone
}

func (b *GreenPhoneBuilder) createPhone() {
	b.phone = &Phone{}
}

func (b *GreenPhoneBuilder) makeCase() {
	b.phone.brand = "Apple"
}

func (b *GreenPhoneBuilder) paint() {
	b.phone.color = "green"
}

func (b *GreenPhoneBuilder) GetResult() {
	fmt.Printf("%s color phone. Brand is %s\n", b.phone.color, b.phone.brand)
}

// Another concrete Builder
type RedPhoneBuilder struct {
	phone *Phone
}

func (b *RedPhoneBuilder) createPhone() {
	b.phone = &Phone{}
}

func (b *RedPhoneBuilder) makeCase() {
	b.phone.brand = "Xiaomi"
}

func (b *RedPhoneBuilder) paint() {
	b.phone.color = "Red"
}

func (b *RedPhoneBuilder) GetResult() {
	fmt.Printf("%s color phone. Brand is %s\n", b.phone.color, b.phone.brand)
}

// Product
type Phone struct {
	color, brand string
}
