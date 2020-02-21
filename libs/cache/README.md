# Cache Store

Cache Store provide a way to cache web content to accelerate the speed of your site.

## Memcached Cache Store Usage

```go
import "GoTenancy/libs/cache/memcached"

func main() {
  client := memcached.New(&Config{Hosts: []string{"127.0.0.1:11211"}, NameSpace: "qor_demo_v1"})

  // Save value `Hello World` with key `hello_world` into cache store
  err := client.Set("hello_world", "Hello World")

  // Get saved value with key `hello_world`
  result, err := client.Get("hello_world")

  // Save marshalled value of user into cache store
  err := client.Set("user", user)

  // Unmarshal saved value into user2
  err := client.Unmarshal("user", &user2)

  // Fetch saved value with key `hello_world`; if the key does not exist, the
  // returned result of `func` will be saved into cache store under passed key
  result, err := client.Fetch("hello_world", func() interface{} {
    return "..."
  })

  // Delete saved value
  err := client.Delete(key string)
}
```

## Memory Cache Store Usage

```go
import "GoTenancy/libs/cache/memory"

func main() {
  client := memory.New()
  // Same API as memcached cache store
}
```

## Redis Cache Store Usage

```go
import "GoTenancy/libs/cache/redis"

func main() {
  client = New(&redis.Options{Addr: "127.0.0.1:6379",
    Password: "",   // no password set
    DB:       0,    // use default DB
    PoolSize: 100,
  })
  // Same API as memcached cache store
}
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
