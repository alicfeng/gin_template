package config

import (
	"os"
)

type jwt struct {
	Secret string
}

var Jwt jwt

func init() {
	Jwt.Secret = os.Getenv("SG_JWT_SECRET")
}
