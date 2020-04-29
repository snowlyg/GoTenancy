package sysinit

import (
	"github.com/snowlyg/go-tenancy/config"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

const UserIDKey = "UserID"

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
)

func init() {

	Redis := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      config.Config.Redis.Addr,
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Password:  "",
		Database:  "",
		Prefix:    "",
		Delim:     "-",
		Driver:    redis.Redigo(), // redis.Radix() can be used instead.
	})

	iris.RegisterOnInterrupt(func() {
		_ = Redis.Close()
	})

	//defer Redis.Close() // close the database connection if application errored.

	Sess.UseDatabase(Redis)
}
