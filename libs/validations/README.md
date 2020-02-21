# Validations

Validations provides a means to [*validate*](https://en.wikipedia.org/wiki/Data_validation) [GORM](https://github.com/jinzhu/gorm) models when creating and updating them.

### Register GORM Callbacks

Validations uses [GORM](https://github.com/jinzhu/gorm) callbacks to handle *validations*, so you will need to register callbacks first:

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/validations"
)

func main() {
  db, err := gorm.Open("sqlite3", "demo_db")

  validations.RegisterCallbacks(db)
}
```

### Usage

After callbacks have been registered, attempting to create or update any record will trigger the `Validate` method that you have implemented for your model. If your implementation adds or returns an error, the attempt will be aborted.

```go
type User struct {
  gorm.Model
  Age uint
}

func (user User) Validate(db *gorm.DB) {
  if user.Age <= 18 {
    db.AddError(errors.New("age need to be 18+"))
  }
}

db.Create(&User{Age: 10})         // won't insert the record into database, as the `Validate` method will return error

var user User{Age: 20}
db.Create(&user)                  // user with age 20 will be inserted into database
db.Model(&user).Update("age", 10) // user's age won't be updated, will return error `age need to be 18+`

// If you have added more than one error, could get all of them with `db.GetErrors()`
func (user User) Validate(db *gorm.DB) {
  if user.Age <= 18 {
    db.AddError(errors.New("age need to be 18+"))
  }
  if user.Name == "" {
    db.AddError(errors.New("name can't be blank"))
  }
}

db.Create(&User{}).GetErrors() // => []error{"age need to be 18+", "name can't be blank"}
```

## [Govalidator](https://github.com/asaskevich/govalidator) integration

Qor [Validations](https://GoTenancy/libs/validations) supports [govalidator](https://github.com/asaskevich/govalidator), so you could add a tag into your struct for some common *validations*, such as *check required*, *numeric*, *length*, etc.

```
type User struct {
  gorm.Model
  Name           string `valid:"required"`
  Password       string `valid:"length(6|20)"`
  SecurePassword string `valid:"numeric"`
  Email          string `valid:"email"`
}
```

## Customize errors on form field

If you want to display errors for each form field in [QOR Admin](http://GoTenancy/libs/admin), you could register your error like this:

```go
func (user User) Validate(db *gorm.DB) {
  if user.Age <= 18 {
    db.AddError(validations.NewError(user, "Age", "age need to be 18+"))
  }
}
```

## Try it out for yourself

Checkout the [http://demo.getqor.com/admin/products/1](http://demo.getqor.com/admin/products/1) demo, change `Name` to be a blank string and save to see what happens.

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
