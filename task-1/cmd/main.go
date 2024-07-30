package main

import (
	"github.com/mohaali482/a2sv-backend-learning-path/task-1/pkg"
)

func main() {
	pkg.ClearScreen()

	pkg.PrintWelcomeMessage()

	pkg.ClearScreen()

	name := pkg.GetNameFromUser()

	pkg.ClearScreen()

	subjects := pkg.GetSubjects()

	pkg.ClearScreen()

	pkg.PrintClosingMessage(name, &subjects)
}
