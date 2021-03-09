package Logger

import (
	"elearn100/Pkg/e"
	"log"
	"os"
)

type Logger interface {
	CreateFile()
	LogInfo(string)
	LogError(string)
	LogFinal(string)
	LogSuccess(string)
}

type Log struct {
}

func (l Log) CreateFile() *os.File {
	dir := e.CreateLogDir()
	logFile, err := os.OpenFile(dir+"/"+"error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Printf("create file %s failed, err %s \n", dir+"/"+"error.log", err)
	}
	defer logFile.Close()
	return logFile
}

func NewLogger() *Log {
	return &Log{}
}
func (l Log) LogError(con string) {
	logFile := l.CreateFile()
	defer logFile.Close()
	fileCon := "[ERROR] " + e.GetCurrentTime() + con
	logFile.Write([]byte(fileCon))
}

func (l Log) LogInfo(con string) {
	logFile := l.CreateFile()
	defer logFile.Close()
	fileCon := "[INFO] " + e.GetCurrentTime() + con
	logFile.Write([]byte(fileCon))
}

func (l Log) LogFinal(con string) {
	logFile := l.CreateFile()
	defer logFile.Close()
	fileCon := "[FINAL] " + e.GetCurrentTime() + con
	logFile.Write([]byte(fileCon))
}

func (l Log) LogSuccess(con string) {
	logFile := l.CreateFile()
	defer logFile.Close()
	fileCon := "[SUCCESS] " + e.GetCurrentTime() + con
	logFile.Write([]byte(fileCon))
}
