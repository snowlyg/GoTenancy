package tenant

import (
	"github.com/snowlyg/qortenant/backend/database"

	"go-tenancy/config/db"
)

func init() {
	database.New(db.DB)
}
