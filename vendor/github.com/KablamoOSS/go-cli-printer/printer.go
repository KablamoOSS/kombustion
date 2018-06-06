package printer

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
)

var instantiated *spinner.Spinner
var once sync.Once

var verbose bool
var color string
var spinnerStyle int

var previousProgressMessage string
var previousStepMessage string
var previousSubStepMessage string

func init() {
	verbose = false
	color = "yellow"
	spinnerStyle = 14
}

// Init the spinner with a verbose flag, and color.
// This is optional, and allows some customisation over the printer.
// You should do invoke this as early as posisble, before the first printer function
// is called.
func Init(initVerbose bool, initColor string, initSpinner int) {
	verbose = initVerbose
	color = initColor
	spinnerStyle = initSpinner
	getPrinter()
}

// Create a singleton to the Spinner
// This ensures we only output to one line
func getPrinter() *spinner.Spinner {
	if instantiated == nil {
		once.Do(func() {
			instantiated = spinner.New(spinner.CharSets[spinnerStyle], 100*time.Millisecond)
			instantiated.Writer = os.Stderr
		})
	}
	return instantiated
}

// Progress message with a spinner
func Progress(message string) {
	spinner := getPrinter()
	if message != previousProgressMessage {
	previousProgressMessage = message
	spinner.Suffix = fmt.Sprintf("  %s", message)
	spinner.Color(color)
}
spinner.Start()
}

// Step prints a line console and stops the spinner
func Step(message string) {
	if message != previousStepMessage {
		previousStepMessage = message
		spinner := getPrinter()

		spinner.Stop()
		fmt.Println(fmt.Sprintf("%s  %s", chalk.Yellow.Color("➜"), chalk.Bold.TextStyle(message)))
	}
}

// SubStep prints a line console, at a given indent and stops the spinner
func SubStep(message string, indent int, last bool) {
	if message != previousSubStepMessage {
		previousSubStepMessage = message
		// Substeps are only printed if the verbose flag is set at init
		if verbose {
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

			spinner.Stop()
			fmt.Println(fmt.Sprintf("%s%s %s", chalk.Dim.TextStyle(indentString), chalk.Dim.TextStyle(icon), chalk.Dim.TextStyle(message)))
		}
	}
}

// Finish prints message to the console and stops the spinner with success.
// This is best used to indicated the end of a task
func Finish(message string) {
	spinner := getPrinter()
	spinner.Stop()
fmt.Println(fmt.Sprintf("%s  %s", chalk.Green.Color("✔"), chalk.Bold.TextStyle(message)))
}

// Error prints an error to the screen. As it's intended reader is a user of your program,
// it expects both the error message, a way for the reader to resolve the error, and if
// possible a link to futher information.
// If the error doesn't have a link, pass a blank string ""
func Error(err error, resolution string, link string) {
	spinner := getPrinter()

	spinner.Stop()
	errMessage := fmt.Sprintf("%s  Error: %s \nHow to fix: %s \n", chalk.Red.Color("✖"), chalk.Red.Color(err.Error()), chalk.Dim.TextStyle(resolution))

	if link != "" {
		errMessage = fmt.Sprintf("%s\n More information: %s", errMessage, chalk.Dim.TextStyle(link))
	}

	fmt.Println(errMessage)
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

	spinner.Stop()
	fmt.Println(errMessage)
	os.Exit(1)
}
