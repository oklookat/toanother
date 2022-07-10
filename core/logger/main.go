package logger

var Log *Logger

func Init() error {
	var err error
	Log, err = New(LEVEL_DEBUG, false)
	return err
}
