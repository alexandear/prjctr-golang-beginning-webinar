package pizza

type Pizza struct {
	dough  string
	sauce  string
	cheese string
}

type Builder interface {
	SetDough(dough string) Builder
	SetSauce(sauce string) Builder
	SetCheese(cheese string) Builder
	Build() *Pizza
}

type DominosBuilder struct {
	pizza *Pizza
}

func NewDominosBuilder() *DominosBuilder {
	return &DominosBuilder{pizza: &Pizza{}}
}

func (b *DominosBuilder) SetDough(dough string) Builder {
	b.pizza.dough = "Dominos " + dough
	return b
}

func (b *DominosBuilder) SetSauce(sauce string) Builder {
	b.pizza.sauce = sauce
	return b
}

func (b *DominosBuilder) SetCheese(cheese string) Builder {
	b.pizza.cheese = cheese
	return b
}

func (b *DominosBuilder) Build() *Pizza {
	return b.pizza
}

type IlMolinoBuilder struct {
	pizza *Pizza
}

func NewIlMolinoBuilder() *IlMolinoBuilder {
	return &IlMolinoBuilder{pizza: &Pizza{}}
}

func (b *IlMolinoBuilder) SetDough(dough string) Builder {
	b.pizza.dough = "Il Molino " + dough
	return b
}

func (b *IlMolinoBuilder) SetSauce(sauce string) Builder {
	b.pizza.sauce = sauce
	return b
}

func (b *IlMolinoBuilder) SetCheese(cheese string) Builder {
	b.pizza.cheese = cheese
	return b
}

func (b *IlMolinoBuilder) Build() *Pizza {
	return b.pizza
}
