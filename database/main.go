package database

// ClientInterface is an interface
type ClientInterface interface {
	Client() RedisClientInterface
}

type client struct {
}

// NewClient implements ClientInterface
func NewClient() ClientInterface {
	return &client{}
}

func (c client) Client() RedisClientInterface {
	return NewRedisClient()
}
