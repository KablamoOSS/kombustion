package printer

import (
	"fmt"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
)

var instantiated *spinner.Spinner
var once sync.Once

// Create a singleton to the Spinner
// This ensures we only output to one line
func getPrinter() *spinner.Spinner {
	once.Do(func() {
		instantiated = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	})

	return instantiated
}

// Progress message with a spinner
func Progress(message string) {
	spinner := getPrinter()
	spinner.Suffix = fmt.Sprintf("  %s", message)
	spinner.Color("yellow")
	spinner.Start()
}

// Step prints a line console and stops the spinner
func Step(message string) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Yellow.Color(chalk.Dim.TextStyle("➜")), chalk.Bold.TextStyle(message))
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// SubStep prints a line console, at a given indent and stops the spinner
func SubStep(message string, indent int, last bool) {
	var indentString string

	for i := 1; i <= indent; i++ {
		indentString = fmt.Sprintf("   %s", indentString)
	}

	icon := ""

	switch indent {
	case 1:
		icon = "└─"
	default:
		icon = "├─"
	}

	if last {
		icon = "└─"
	}

	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s%s %s \n", chalk.Dim.TextStyle(indentString), chalk.Dim.TextStyle(icon), chalk.Dim.TextStyle(message))
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Finish prints message to the console and stops the spinner with success.
// This is best used to indicated the end of a task
func Finish(message string) {
	spinner := getPrinter()
	spinner.FinalMSG = fmt.Sprintf("%s  %s \n", chalk.Green.Color("✔"), chalk.Bold.TextStyle(message))
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Error prints an error to the screen. As it's intended reader is a user of your program,
// it expects both the error message, a way for the reader to resolve the error, and if
// possible a link to futher information.
// If the error doesn't have a link, pass a blank string ""
func Error(err error, resolution string, link string) {
	spinner := getPrinter()

	errMessage := fmt.Sprintf("%s  Error: %s \nHow to fix: %s \n", chalk.Red.Color("✖"), chalk.Red.Color(err.Error()), chalk.Dim.TextStyle(resolution))

	if link != "" {
		errMessage = fmt.Sprintf("%s\n More information: %s", errMessage, chalk.Dim.TextStyle(link))
	}

	spinner.FinalMSG = errMessage
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
}

// Fatal prints an error in the exact same way as Error, except it prefixes with "Fatal",
// and the function ends with a panic
func Fatal(err error, resolution string, link string) {
	spinner := getPrinter()

	errMessage := fmt.Sprintf("%s  Fatal: %s \nHow to fix: %s \n", chalk.Red.Color("✖"), chalk.Red.Color(err.Error()), chalk.Dim.TextStyle(resolution))

	// Add the link if a valid one was supplied
	if link != "" {
		errMessage = fmt.Sprintf("%s\n More information: %s", errMessage, chalk.Dim.TextStyle(link))
	}

	spinner.FinalMSG = errMessage
	spinner.Start()
	time.Sleep(2 * time.Second)
	spinner.Stop()
	panic(1)
}
