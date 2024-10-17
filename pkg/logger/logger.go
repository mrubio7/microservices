package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

var Logger *LoggerStruct

type LoggerStruct struct {
	Log        *log.Logger
	ConsoleLog *log.Logger
}

func Initialize() {
	logger := LoggerStruct{}

	// Crear el directorio de logs si no existe
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Crear un archivo de log con la fecha y hora actual
	path := fmt.Sprintf("logs/%d-%d-%d_%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute())

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Crear un multi-writer para escribir tanto en el archivo como en la consola
	fileWriter := io.MultiWriter(f)
	consoleWriter := io.MultiWriter(os.Stdout)

	l := log.New(fileWriter, "", log.Ldate|log.Ltime)
	cl := log.New(consoleWriter, "", log.Ldate|log.Ltime)

	Logger = &logger
	Logger.Log = l
	Logger.ConsoleLog = cl
}

func Trace(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("TRC | "+format, argv...)
	Logger.Log.Println(formattedMessage)

	faint := color.New(color.Faint).SprintfFunc()
	Logger.ConsoleLog.Println(faint(formattedMessage))
}

func Debug(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("DBG | "+format, argv...)
	Logger.Log.Println(formattedMessage)

	green := color.New(color.FgGreen).SprintfFunc()
	Logger.ConsoleLog.Println(green(formattedMessage))
}

func Info(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("INF | "+format, argv...)
	Logger.Log.Println(formattedMessage)

	cyan := color.New(color.FgCyan).SprintfFunc()
	Logger.ConsoleLog.Println(cyan(formattedMessage))
}

func Warning(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("WRN | "+format, argv...)
	Logger.Log.Println(formattedMessage)

	yellow := color.New(color.FgYellow).SprintfFunc()
	Logger.ConsoleLog.Println(yellow(formattedMessage))
}

func Error(format string, argv ...any) {
	formattedMessage := fmt.Sprintf("ERR | "+format, argv...)
	Logger.Log.Println(formattedMessage)

	red := color.New(color.FgRed).SprintfFunc()
	Logger.ConsoleLog.Println(red(formattedMessage))
}
