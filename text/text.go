package text

// ErrorHelpInfo is information to be appended to all error messages
var ErrorHelpInfo string

func init() {
	ErrorHelpInfo = `
	If this may be an issue with Kombustion, or happens repeatedly
	please file a bug report https://github.com/KablamoOSS/kombustion/issues/new?template=bug_report.md
`
}
