package main

import (
	// "fmt"
	"os"
	// "io"
	// "strings"
	"text/template"
	"log"
)

// sample one
// func main() {
// 	name:= "Girls love Feji"

// 	tpl := `
// 	<!DOCTYPE html>
// 	<html>
// 	<head>
//   		<meta charset="utf-8">
//   		<meta name="viewport" content="width=device-width">
//   		<title>Go playground</title>
// 	</head>
// 	<body>
// 		<h1>` + name + `</h1>
// 	</body>
// 	</html>
// 	`
// 	fmt.Println(tpl)
// }


// sample two
// func main() {
// 	name:= "Girls love Feji"

// 	str := `
// 	<!DOCTYPE html>
// 	<html>
// 	<head>
//   		<meta charset="utf-8">
//   		<meta name="viewport" content="width=device-width">
//   		<title>Go playground</title>
// 	</head>
// 	<body>
// 		<h1>` + name + `</h1>
// 	</body>
// 	</html>
// 	`
// 	nf, err := os.Create("index.html")
// 	if err != nil {
// 		log.Fatal("error creating file", err)
// 	}
// 	defer nf.Close()
// 	io.Copy(nf, strings.NewReader(str))
// }

// sample three: import html file
func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}


