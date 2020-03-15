package sysinit

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID, AllowReclaim: true})
)

func init() {

	Redis := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      "127.0.0.1:6379",
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
