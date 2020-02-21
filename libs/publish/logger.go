package publish

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

// LoggerInterface logger interface used to print publish logs
type LoggerInterface interface {
	Print(...interface{})
}

// Logger default logger used to print publish logs
var Logger LoggerInterface

func init() {
	Logger = log.New(os.Stdout, "\r\n", 0)
}

func stringify(object interface{}) string {
	if obj, ok := object.(interface {
		Stringify() string
	}); ok {
		return obj.Stringify()
	}

	scope := gorm.Scope{Value: object}
	for _, column := range []string{"Description", "Name", "Title", "Code"} {
		if field, ok := scope.FieldByName(column); ok {
			return fmt.Sprintf("%v", field.Field.Interface())
		}
	}

	if scope.PrimaryField() != nil {
		if scope.PrimaryKeyZero() {
			return ""
		}
		return fmt.Sprintf("%v#%v", scope.GetModelStruct().ModelType.Name(), scope.PrimaryKeyValue())
	}

	return fmt.Sprint(reflect.Indirect(reflect.ValueOf(object)).Interface())
}

func stringifyPrimaryValues(primaryValues [][][]interface{}, columns ...string) string {
	var values []string
	for _, primaryValue := range primaryValues {
		var primaryKeys []string
		for _, value := range primaryValue {
			if len(columns) == 0 {
				primaryKeys = append(primaryKeys, fmt.Sprint(value[1]))
			} else {
				for _, column := range columns {
					if column == fmt.Sprint(value[0]) {
						primaryKeys = append(primaryKeys, fmt.Sprint(value[1]))
					}
				}
			}
		}
		if len(primaryKeys) > 1 {
			values = append(values, fmt.Sprintf("[%v]", strings.Join(primaryKeys, ", ")))
		} else {
			values = append(values, primaryKeys...)
		}
	}
	return strings.Join(values, "; ")
}
