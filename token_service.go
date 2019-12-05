package main

// Authable -
type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

// TokenService -
type TokenService struct {
	repo Repository
}

// Decode -
func (srv *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

// Encode -
func (srv *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
