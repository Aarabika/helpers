package helpers

import "fmt"

func ToString(instance interface{}) string  {
	return fmt.Sprintf("%+v\n", instance)
}
