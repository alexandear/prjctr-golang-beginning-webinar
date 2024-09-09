package pizza

type option struct {
	dough  string
	sauce  string
	cheese string
}

type Option func(*option)

func WithDough(dough string) Option {
	return func(o *option) {
		o.dough = dough
	}
}

func WithSauce(sauce string) Option {
	return func(o *option) {
		o.sauce = sauce
	}
}

func WithCheese(cheese string) Option {
	return func(o *option) {
		o.cheese = cheese
	}
}

func NewPizza(options ...Option) *Pizza {
	o := &option{}
	for _, opt := range options {
		opt(o)
	}
	return &Pizza{
		dough:  o.dough,
		sauce:  o.sauce,
		cheese: o.cheese,
	}
}
