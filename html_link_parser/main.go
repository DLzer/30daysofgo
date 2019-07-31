package main

 import (
     "fmt"
     "strings"

     link "./pkg/link"
 )

 var exampleHtml = `
     <html>
     <body>
         <h1>Hello!</h1>
         <a href="/link">A link to another page</a>
     </body>
     </html>
     `

 func main() {
     r := strings.NewReader(exampleHtml)
     link, err := link.Parse(r)
     if err != nil {
         panic(err)
     }
     fmt.Printf("%+v\n", links)
 }
