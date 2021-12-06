package swu

type Api struct {
	key string
}

func New(key string) *Api {
	return &Api{
		key: key,
	}
}
