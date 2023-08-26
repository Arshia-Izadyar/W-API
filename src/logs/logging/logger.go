package logging

type Logger interface {
	Init()
	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})

	Fatal(err error, cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}

func NewLogger() Logger {
	return newZapLogger()
}
