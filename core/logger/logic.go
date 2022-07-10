package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/oklookat/toanother/core/datadir"
)

const (
	COLOR_RED    = "\u001b[31m"
	COLOR_GREEN  = "\u001b[32m"
	COLOR_YELLOW = "\u001b[33m"
	COLOR_CYAN   = "\u001b[36m"
	COLOR_WHITE  = "\u001b[37m"
	COLOR_RESET  = "\u001b[0m"
	LOG_NAME     = "./log.txt"
)

const (
	LEVEL_PRINT = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

type Logger struct {
	writeToConsole bool
	level          int
	file           *os.File
}

func New(level int, writeToConsole bool) (*Logger, error) {
	var isOversize = false

	// check is log big.
	exists, err := datadir.IsFileExists(LOG_NAME)
	if err != nil {
		return nil, err
	}
	if exists {
		file, err := datadir.OpenFile(LOG_NAME, os.O_RDONLY)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		stat, err := file.Stat()
		if err != nil {
			return nil, err
		}
		isOversize = stat != nil && stat.Size() > 1000000
	}

	// open file.
	var file *os.File
	var openFlag = 0
	if isOversize {
		openFlag = os.O_RDWR
	} else {
		openFlag = os.O_APPEND
	}
	file, err = datadir.OpenFile(LOG_NAME, openFlag)
	if err != nil {
		return nil, err
	}

	if isOversize {
		// clean file.
		if err = file.Truncate(0); err != nil {
			return nil, err
		}
		if _, err = file.Seek(0, 0); err != nil {
			return nil, err
		}
	}

	return &Logger{
		writeToConsole: writeToConsole,
		level:          level,
		file:           file,
	}, nil
}

func (l *Logger) Print(message string) {
	go l.print(LEVEL_PRINT, message)
}

func (l *Logger) Debug(message string) {
	go l.print(LEVEL_DEBUG, message)
}

func (l *Logger) Info(message string) {
	go l.print(LEVEL_INFO, message)
}

func (l *Logger) Warn(message string) {
	go l.print(LEVEL_WARN, message)
}

func (l *Logger) Error(message string) {
	go l.print(LEVEL_ERROR, message)
}

func (l *Logger) Fatal(message string) {
	l.print(LEVEL_FATAL, message)
}

func (l *Logger) print(level int, message string) {
	if l.level > level {
		return
	}
	var color = ""
	var levelPrefix = ""
	switch level {
	case LEVEL_PRINT:
		levelPrefix = "PRINT"
	case LEVEL_DEBUG:
		levelPrefix = "DEBUG"
	case LEVEL_INFO:
		levelPrefix = "INFO"
	case LEVEL_WARN:
		levelPrefix = "WARN"
		color = COLOR_YELLOW
	case LEVEL_ERROR:
		levelPrefix = "ERROR"
		color = COLOR_RED
	case LEVEL_FATAL:
		levelPrefix = "FATAL"
		color = COLOR_RED
	}
	var currentTime = time.Now().Format("15:04:05")
	var final = fmt.Sprintf("[%v] > [%v] > %v", currentTime, levelPrefix, message)
	if l.writeToConsole {
		println(color + final + COLOR_RESET)
	}
	var _, _ = l.file.WriteString(final + "\n")
	if level == LEVEL_FATAL {
		os.Exit(1)
	}
}
