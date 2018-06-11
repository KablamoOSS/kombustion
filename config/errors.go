package config

// ErrorHelpInfo is the default fallback error information
var ErrorHelpInfo string

func init() {
	ErrorHelpInfo = `--
If this may be an issue with Kombustion, or happens repeatedly please file a bug report:
https://github.com/KablamoOSS/kombustion/issues/new?template=bug_report.md`
}
