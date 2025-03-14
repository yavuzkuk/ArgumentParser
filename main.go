package main

import (
	"fmt"
	"os"

	Structs "github.com/yavuzkuk/ArgumentParser/Struct"
)

func main() {

	arguments := os.Args[1:]
	enteredFlag := Structs.FlagParse(arguments)
	var definedFlags []Structs.Flag

	definedFlags = append(definedFlags, Structs.Flag{
		ShortName: "-h",
		LongName:  "--help",
		Usage:     "-h",
		Desc:      "Show help message",
		Required:  false,
		Value:     "",
		GetValue:  false,
		Default:   false,
	}, Structs.Flag{
		ShortName: "-v",
		LongName:  "--verbose",
		Usage:     "-v",
		Desc:      "More information",
		Required:  true,
		Value:     true,
		GetValue:  false,
		Default:   true,
	}, Structs.Flag{
		ShortName: "-u",
		LongName:  "--url",
		Usage:     "-u <link>",
		Desc:      "Scan target",
		Required:  true,
		Value:     "",
		GetValue:  true,
		Default:   false,
	}, Structs.Flag{
		ShortName: "-w",
		LongName:  "--wordlist",
		Usage:     "-w <wordlist>",
		Desc:      "Specify wordlist",
		Required:  false,
		Value:     "/usr/share/wordlists/Seclist/Discovery/Web-Content/common.txt",
		GetValue:  true,
		Default:   true,
	}, Structs.Flag{
		ShortName: "-c",
		LongName:  "--config",
		Usage:     "-c <config>",
		Desc:      "Specify config file",
		Required:  false,
		Value:     "",
		GetValue:  true,
		Default:   false,
	}, Structs.Flag{
		ShortName: "-f",
		LongName:  "--fast",
		Usage:     "-f",
		Desc:      "Fast scan",
		Required:  true,
		Value:     false,
		Default:   true,
	})

	flags := Structs.GetMain(arguments, enteredFlag, definedFlags)
	fmt.Println(flags)
}
