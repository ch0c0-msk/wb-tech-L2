package pattern

// Visitor interface
type Visitor interface {
	VisitCentre(c *Centre) string
	VisitRestaurant(c *Restaurant) string
}

// Element to visit
type Place interface {
	Accept(v Visitor) string
}

// Visitor implementation
type Tourist struct{}

func (t *Tourist) VisitCentre(c *Centre) string {
	return c.MakePhoto()
}

func (t *Tourist) VisitRestaurant(r *Restaurant) string {
	return r.BuyNationalFood()
}

// Another Visitor implementation
type Citizen struct{}

func (c *Citizen) VisitCentre(centre *Centre) string {
	return centre.TakeWalk()
}

func (c *Citizen) VisitRestaurant(r *Restaurant) string {
	return r.BuyLunch()
}

// Place implementation Centre
type Centre struct{}

func (c *Centre) Accept(v Visitor) string {
	return v.VisitCentre(c)
}

func (c *Centre) MakePhoto() string {
	return "Some photo"
}

func (c *Centre) TakeWalk() string {
	return "Walking in centre"
}

// Place implementation Restaurant
type Restaurant struct{}

func (r *Restaurant) Accept(v Visitor) string {
	return v.VisitRestaurant(r)
}

func (r *Restaurant) BuyNationalFood() string {
	return "Some national food"
}

func (r *Restaurant) BuyLunch() string {
	return "Some lunch"
}

// Collection of elements to visit
type City struct {
	places []Place
}

func NewCity() *City {
	return &City{places: []Place{&Restaurant{}, &Centre{}}}
}

func (c *City) GetPlaces() []Place {
	return c.places
}
