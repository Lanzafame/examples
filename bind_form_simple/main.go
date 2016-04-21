package main

import (
	"fmt"
	"html/template"

	"github.com/kataras/iris"
)

type Visitor struct {
	Username string
	Mail     string
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.ExecuteTemplate(formTemplate, "")
	})

	iris.Post("/form_action", func(ctx *iris.Context) {
		visitor := &Visitor{}
		err := ctx.ReadForm(visitor)
		if err != nil {
			fmt.Println("Error when reading from form: " + err.Error())
		}
		fmt.Printf("\n Visitor: %v", visitor)
	})

	fmt.Println("Server is running at :8080")
	iris.Listen()
}

var formTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
<form action="/form_action" method="post">
<input type="text" name="Username" />
<br/>
<input type="text" name="Mail" />
<hr/>
<input type="submit" value="Send data" />

</form>
</body>
</html>
`))
