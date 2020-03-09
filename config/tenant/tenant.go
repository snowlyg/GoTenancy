package tenant

import (
	"github.com/snowlyg/qortenant"
	"github.com/snowlyg/qortenant/backend/database"

	"go-tenancy/config/db"
)

var Tenant *qortenant.QorTenant

func init() {
	Tenant = qortenant.New(database.New(db.DB))
}
