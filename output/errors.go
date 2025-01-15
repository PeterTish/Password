package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	//val, ok := value.(int)
	//if ok {
	//	color.Red("Код ошибки: %d", val)
	//	return
	//}
	//strValue, ok := value.(string)
	//if ok {
	//	color.Red(strValue)
	//	return
	//}
	//errorValue, ok := value.(error)
	//if ok {
	//	color.Red(errorValue.Error())
	//	return
	//}
	//color.Red("Неизвестный тип ошибки")

	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}
