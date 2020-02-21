# Auth Themes

Auth Themes is a collection of themes for [Auth](https://GoTenancy/libs/auth), which is a Golang authentication framework.

It aimis to reduce repeated code, make you be able to integrate [Auth](https://GoTenancy/libs/auth) into your application with few lines of code.

## Usage

Each theme might has different usage, please refer their own documents, but most of them should be easy as accept an Auth Config to initialize itself.

Here is an example for how to use theme `clean`

```go
import  "GoTenancy/libs/auth_themes/clean"

func main() {
  Auth = clean.New(&auth.Config{
    DB:         db.DB,
    Render:     config.View,
    Mailer:     config.Mailer,
    UserModel:  models.User{},
    Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
  })
}
```

## How to create themes

Although integrate Auth into your application already much easier than write your own solution, it is boring/time costing to repeat yourself again and again.

To avoid this, you could create your own Auth theme.

Usually when write your theme, you can just accept an Auth Config and extend it with some default settings, and prepend Auth's ViewPaths to customize view templates, for example:

```go
func New(config *auth.Config) *auth.Auth {
  if config == nil {
    config = &auth.Config{}
  }
  config.ViewPaths = append(config.ViewPaths, "GoTenancy/libs/auth_themes/clean/views")

  Auth := auth.New(config)
  return Auth
}
```

Refer Theme [Clean](https://GoTenancy/libs/auth_themes/tree/master/clean) for more details
