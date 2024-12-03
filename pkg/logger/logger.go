package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

func init() {
	Initialize()
}

var Logger *LoggerStruct

type LoggerStruct struct {
	ConsoleLog *log.Logger
}

func Initialize() {
	logger := LoggerStruct{}

	consoleWriter := io.MultiWriter(os.Stdout)

	cl := log.New(consoleWriter, "", log.Ldate|log.Ltime)

	Logger = &logger
	Logger.ConsoleLog = cl
}

func Trace(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("TRC | "+format, argv...)

	faint := color.New(color.Faint).SprintfFunc()
	Logger.ConsoleLog.Println(faint(formattedMessage))
}

func Debug(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("DBG | "+format, argv...)

	green := color.New(color.FgGreen).SprintfFunc()
	Logger.ConsoleLog.Println(green(formattedMessage))
}

func Info(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("INF | "+format, argv...)

	cyan := color.New(color.FgCyan).SprintfFunc()
	Logger.ConsoleLog.Println(cyan(formattedMessage))
}

func Warning(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("WRN | "+format, argv...)

	yellow := color.New(color.FgYellow).SprintfFunc()
	Logger.ConsoleLog.Println(yellow(formattedMessage))
}

func Error(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("ERR | "+format, argv...)

	red := color.New(color.FgRed).SprintfFunc()
	Logger.ConsoleLog.Println(red(formattedMessage))
}
