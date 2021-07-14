package main

import (
	"fmt"
	"regexp"
)

func main() {
	//const text = "My email is onyourwill@gmail.com 1"
	//const reg = "onyourwill@gmail.com"
	//compile, err := regexp.Compile(reg)
	//if err != nil {
	//	panic(err)
	//}
	const text = `
	My email is onyourwill@gmail.com
	another is hhh@ggg.org
	and ddd@dfdf.edu
`
	const reg = `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z.]+)`
	compile := regexp.MustCompile(reg)

	match := compile.FindString(text)
	fmt.Println(match)

	allString := compile.FindAllString(text, -1)
	fmt.Println(allString)

	submatch := compile.FindAllStringSubmatch(text, -1)
	for _, m := range submatch {
		fmt.Println(m)
	}
}
