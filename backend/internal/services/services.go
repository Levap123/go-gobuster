package services

type Services struct {
	routes []string
}

func New(routes []string) *Services {
	return &Services{
		routes: routes,
	}
}
