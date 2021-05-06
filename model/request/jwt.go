package request

import (
	"fmt"

	"github.com/kataras/iris/v12/middleware/jwt"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID         uuid.UUID
	ID           string
	Username     string
	Nickname     string
	AuthorityId  string
	LoginType    int
	AuthType     int
	CreationDate int64
	ExpiresIn    int
}

func (c *CustomClaims) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("%w: %s", jwt.ErrMissingKey, "user_id")
	}

	return nil
}

func (c *CustomClaims) GetID() string {
	return c.ID
}

func (c *CustomClaims) GetAuthorityId() string {
	return c.ID
}
