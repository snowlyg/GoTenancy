package beego_session_test

import (
	"encoding/json"
	"testing"

	"GoTenancy/libs/session/beego_session"
	"GoTenancy/libs/session/test"
	"github.com/astaxie/beego/session"
)

func TestAll(t *testing.T) {
	config := `{"cookieName":"gosessionid","enableSetCookie":true,"gclifetime":3600,"ProviderConfig":"{\"cookieName\":\"gosessionid\",\"securityKey\":\"beegocookiehashkey\"}"}`
	conf := new(session.ManagerConfig)
	if err := json.Unmarshal([]byte(config), conf); err != nil {
		t.Fatal("json decode error", err)
	}

	globalSessions, _ := session.NewManager("memory", conf)
	go globalSessions.GC()

	engine := beego_session.New(globalSessions)
	test.TestAll(engine, t)
}
