package Structs

import (
	"fmt"
	"strings"
)

type Flag struct {
	ShortName string
	LongName  string
	Usage     string
	Desc      string
	Required  bool
	Value     interface{}
	GetValue  bool
	Default   bool
}

func Ternary(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func Help(flags []Flag) {
	space := strings.Repeat(" ", 10)
	for _, flag := range flags {
		fmt.Println(flag.ShortName, "/", flag.LongName, space, flag.Usage, "|", flag.Desc, Ternary(flag.Required, "| Zorunlu", "| Zorunlu değil"))
	}

}

func FlagParse(arguments []string) map[string]string {

	mapArgument := make(map[string]string)
	for index, argument := range arguments {
		argument = strings.TrimSpace(argument)
		if argument == "-h" || argument == "--help" {
			mapArgument[argument] = "help"
		} else {
			if strings.HasPrefix(argument, "-") || strings.HasPrefix(argument, "--") {
				if index < len(arguments)-1 {
					if strings.HasPrefix(arguments[index+1], "--") || strings.HasPrefix(arguments[index+1], "-") {
						// aradaki tekli elemanlar
						mapArgument[argument] = "EMPTY"
					} else {
						// normal eleman
						mapArgument[argument] = arguments[index+1]
					}
				} else if index+1 == len(arguments) {
					// son eleman
					mapArgument[argument] = "EMPTY"
				}
			}
		}
	}

	return mapArgument
}

func RequiredCheck(enteredFlags map[string]string, flags []Flag) map[string]string {
	for _, flag := range flags {
		var counter int = 0
		var valueErr int = 0

		for key, value := range enteredFlags {

			if flag.Required {
				if !flag.GetValue && flag.Default && flag.Required {
					var boolValue string = ""
					if flag.Value.(bool) {
						boolValue = "true"
					} else {
						boolValue = "false"
					}
					enteredFlags[flag.LongName] = boolValue
				} else if key != flag.ShortName && key != flag.LongName {
					counter++
				} else if flag.GetValue && (flag.ShortName == key || flag.LongName == key) && value == "EMPTY" {
					valueErr++
				}
			} else if !flag.Required && flag.GetValue && flag.Default {
				// default deger kontolü yapıyor
				enteredFlags[flag.LongName] = flag.Value.(string)
			}

			if counter == len(enteredFlags) {
				panic("Zorunlu flag girilmedi " + flag.ShortName + "/" + flag.LongName)
			}

			if valueErr > 0 {
				panic("Veri girilmedi " + flag.ShortName + "/" + flag.LongName)
			}

		}
	}
	return enteredFlags

}

func OutOfScopeFlag(enteredFlags map[string]string, flags []Flag) {

	for key, _ := range enteredFlags {
		var outOfScopeErr int = 0
		for _, flag := range flags {

			if key != flag.ShortName && key != flag.LongName {
				outOfScopeErr++
			}
		}

		if outOfScopeErr == len(flags) {
			panic("Geçersiz bir flag girdiniz --> " + key)
		}

	}
}

func GetFlagsValue(enteredFlags map[string]string, flags []Flag) {

}

func GetMain(arguments []string, enteredFlag map[string]string, flags []Flag, banner string) map[string]string {
	if len(arguments) == 0 {
		Help(flags)
	} else if enteredFlag["--help"] == "help" || enteredFlag["-h"] == "help" {
		Help(flags)
	} else {
		OutOfScopeFlag(enteredFlag, flags)
		enteredFlag := RequiredCheck(enteredFlag, flags)
		GetFlagsValue(enteredFlag, flags)

		return enteredFlag
	}
	return make(map[string]string)
}
