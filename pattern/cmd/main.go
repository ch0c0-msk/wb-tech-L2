package main

import (
	"fmt"
	"pattern"
)

func main() {
	// Facade
	fmt.Println("\nINFO: Facade example")
	computer := pattern.NewComputer(2400, 100, 500)
	computer.Start()
	fmt.Println()

	// Builder
	fmt.Println("\nINFO: Builder example")
	greenBuilder := pattern.GreenPhoneBuilder{}
	pipeline := pattern.NewFabricPipeline(&greenBuilder)
	pipeline.Construct()
	greenBuilder.GetResult()
	redBuilder := pattern.RedPhoneBuilder{}
	pipeline.ChangeBuilder(&redBuilder)
	pipeline.Construct()
	redBuilder.GetResult()

	// Visitor
	fmt.Println("\nINFO: Visitor example")
	collection := pattern.NewCity()
	visitor1 := &pattern.Tourist{}
	visitor2 := &pattern.Citizen{}
	fmt.Println("Tourist visitor")
	for _, place := range collection.GetPlaces() {
		fmt.Println(place.Accept(visitor1))
	}
	fmt.Println("Local citizen visitor")
	for _, place := range collection.GetPlaces() {
		fmt.Println(place.Accept(visitor2))
	}

	// Command
	fmt.Println("\nINFO: Command example")
	light := &pattern.Light{}
	commandOn := pattern.NewLightOn(light)
	commandOff := pattern.NewLightOff(light)
	fmt.Printf("isLight: %v\n", light.IsLightOn)
	invoker := &pattern.Invoker{}
	invoker.AddCommand(commandOn)
	invoker.ExecuteCommand()
	fmt.Printf("CommandOn execute. isLight: %v\n", light.IsLightOn)
	invoker.AddCommand(commandOff)
	invoker.ExecuteCommand()
	fmt.Printf("CommandOff execute. isLight: %v\n", light.IsLightOn)

	// Chain of responsibility
	fmt.Println("\nINFO: Chain of responsibility example")
	firstHandler := &pattern.ImageHandler{Next: &pattern.TextHandler{}}
	fmt.Println("Send image file...")
	fmt.Println(firstHandler.ProcessRequest("image"))
	fmt.Println("Send text file...")
	fmt.Println(firstHandler.ProcessRequest("text"))
	fmt.Println("Send excel file...")
	fmt.Println(firstHandler.ProcessRequest("excel"))

	// Factory method
	fmt.Println("\nINFO: Factory method example")
	creators := []pattern.Factory{&pattern.KnifeFactory{}, &pattern.SpoonFactory{}}
	products := make([]pattern.Product, len(creators))
	fmt.Println("Product creation...")
	for i, c := range creators {
		products[i] = c.NewProduct()
	}
	fmt.Println("Product using...")
	for _, p := range products {
		p.Use()
	}

	// Strategy
	fmt.Println("\nINFO: Strategy example")
	a := []int{3, 0, 10}
	fmt.Printf("Original list: %v\n", a)
	ctx := pattern.SortContext{}
	midPivotAlg := &pattern.MiddlePivot{}
	ctx.SwitchStrategy(midPivotAlg)
	fmt.Printf("Middle pivot element: %d\n", ctx.GetListPivot(0, len(a)-1, a))
	lomutoPivotAlg := &pattern.LomutoPivot{}
	ctx.SwitchStrategy(lomutoPivotAlg)
	fmt.Printf("Lomuto pivot element: %d\n", ctx.GetListPivot(0, len(a)-1, a))

	// State
	fmt.Println("\nINFO: State example")
	normalState := &pattern.NormalState{}
	wrongState := &pattern.WrongState{}
	server := &pattern.Server{}
	server.SetState(normalState)
	fmt.Println("Normal server state processing:")
	server.Request("Some message")
	fmt.Println("Wrong server state processing:")
	server.SetState(wrongState)
	server.Request("some message")
}
