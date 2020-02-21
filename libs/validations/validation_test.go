package validations_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"GoTenancy/libs/qor/test/utils"
	"GoTenancy/libs/validations"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name           string `valid:"required"`
	Password       string `valid:"length(6|20)"`
	SecurePassword string `valid:"numeric"`
	Email          string `valid:"email,uniqEmail~Email already be token"`
	CompanyID      int
	Company        Company
	CreditCard     CreditCard
	Addresses      []Address
	Languages      []Language `gorm:"many2many:user_languages"`
}

func (user *User) Validate(db *gorm.DB) {
	govalidator.CustomTypeTagMap.Set("uniqEmail", govalidator.CustomTypeValidator(func(email interface{}, context interface{}) bool {
		var count int
		if db.Model(&User{}).Where("email = ?", email).Count(&count); count == 0 || email == "" {
			return true
		}
		return false
	}))
	if user.Name == "invalid" {
		db.AddError(validations.NewError(user, "Name", "invalid user name"))
	}
}

type Company struct {
	gorm.Model
	Name string
}

func (company *Company) Validate(db *gorm.DB) {
	if company.Name == "invalid" {
		db.AddError(errors.New("invalid company name"))
	}
}

type CreditCard struct {
	gorm.Model
	UserID int
	Number string
}

func (card *CreditCard) Validate(db *gorm.DB) {
	if !regexp.MustCompile("^(\\d){13,16}$").MatchString(card.Number) {
		db.AddError(validations.NewError(card, "Number", "invalid card number"))
	}
}

type Address struct {
	gorm.Model
	UserID  int
	Address string
}

func (address *Address) Validate(db *gorm.DB) {
	if address.Address == "invalid" {
		db.AddError(validations.NewError(address, "Address", "invalid address"))
	}
}

type Language struct {
	gorm.Model
	Code string
}

func (language *Language) Validate(db *gorm.DB) error {
	if language.Code == "invalid" {
		return validations.NewError(language, "Code", "invalid language")
	}
	return nil
}

func init() {
	db = utils.TestDB()
	validations.RegisterCallbacks(db)
	tables := []interface{}{&User{}, &Company{}, &CreditCard{}, &Address{}, &Language{}}
	for _, table := range tables {
		if err := db.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}
		db.AutoMigrate(table)
	}
}

func TestGoValidation(t *testing.T) {
	user := User{Name: "", Password: "123123", Email: "a@gmail.com"}

	result := db.Save(&user)
	if result.Error == nil {
		t.Errorf("Should get error when save empty user")
	}

	if result.Error.Error() != "Name can't be blank" {
		t.Errorf("Error message should be equal `Name can't be blank`")
	}

	user = User{Name: "", Password: "123", SecurePassword: "AB123", Email: "aagmail.com"}
	result = db.Save(&user)
	messages := []string{"Name can't be blank",
		"Password is the wrong length (should be 6~20 characters)",
		"SecurePassword is not a number",
		"Email is not a valid email address"}
	for i, err := range result.GetErrors() {
		if messages[i] != err.Error() {
			t.Errorf(fmt.Sprintf("Error message should be equal `%v`, but it is `%v`", messages[i], err.Error()))
		}
	}

	user = User{Name: "A", Password: "123123", Email: "a@gmail.com"}
	result = db.Save(&user)
	user = User{Name: "B", Password: "123123", Email: "a@gmail.com"}
	if result := db.Save(&user); result.Error.Error() != "Email already be token" {
		t.Errorf("Should get email alredy be token error")
	}
}

func TestSaveInvalidUser(t *testing.T) {
	user := User{Name: "invalid"}

	if result := db.Save(&user); result.Error == nil {
		t.Errorf("Should get error when save invalid user")
	}
}

func TestSaveInvalidCompany(t *testing.T) {
	user := User{
		Name:    "valid",
		Company: Company{Name: "invalid"},
	}

	if result := db.Save(&user); result.Error == nil {
		t.Errorf("Should get error when save invalid company")
	}
}

func TestSaveInvalidCreditCard(t *testing.T) {
	user := User{
		Name:       "valid",
		Company:    Company{Name: "valid"},
		CreditCard: CreditCard{Number: "invalid"},
	}

	if result := db.Save(&user); result.Error == nil {
		t.Errorf("Should get error when save invalid credit card")
	}
}

func TestSaveInvalidAddresses(t *testing.T) {
	user := User{
		Name:       "valid",
		Company:    Company{Name: "valid"},
		CreditCard: CreditCard{Number: "4111111111111111"},
		Addresses:  []Address{{Address: "invalid"}},
	}

	if result := db.Save(&user); result.Error == nil {
		t.Errorf("Should get error when save invalid addresses")
	}
}

func TestSaveInvalidLanguage(t *testing.T) {
	user := User{
		Name:       "valid",
		Company:    Company{Name: "valid"},
		CreditCard: CreditCard{Number: "4111111111111111"},
		Addresses:  []Address{{Address: "valid"}},
		Languages:  []Language{{Code: "invalid"}},
	}

	if result := db.Save(&user); result.Error == nil {
		t.Errorf("Should get error when save invalid language")
	}
}

func TestSaveAllValidData(t *testing.T) {
	user := User{
		Name:       "valid",
		Company:    Company{Name: "valid"},
		CreditCard: CreditCard{Number: "4111111111111111"},
		Addresses:  []Address{{Address: "valid1"}, {Address: "valid2"}},
		Languages:  []Language{{Code: "valid1"}, {Code: "valid2"}},
	}

	if result := db.Save(&user); result.Error != nil {
		t.Errorf("Should get no error when save valid data, but got: %v", result.Error)
	}
}
