package logger

type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
}

type LogField struct {
	Key string
	Val any
}

func Field(key string, val any) LogField {
	return LogField{
		Key: key,
		Val: val,
	}
}
