package environment

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

func (e Env) IsDevelopment() bool {
	return e == Development
}

func (e Env) IsProduction() bool {
	return e == Production
}

func (e Env) String() string {
	return string(e)
}
