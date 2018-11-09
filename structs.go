package helpers

import "fmt"

func ToStr(instance interface{}) string  {
	return fmt.Sprintf("%+v\n", instance)
}
