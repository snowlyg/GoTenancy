package scs_test

import (
	"testing"

	"GoTenancy/libs/session/scs"
	"GoTenancy/libs/session/test"
	scssession "github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
)

func TestAll(t *testing.T) {
	engine := scssession.NewManager(memstore.New(0))
	manager := scs.New(engine)
	test.TestAll(manager, t)
}
