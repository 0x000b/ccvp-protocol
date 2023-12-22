package log

import (
	"log"
)

var (
	Blue   = "\033[34m"
	Red    = "\033[31m"
	White  = "\033[0m"
	Yellow = "\033[33m"
	Purple = "\033[35m"
	Gray   = "\033[37m"
	Green  = "\033[32m"
)

func ErrorLog(message string) {
	log.Print(Red + "[ERROR]" + White + ": " + message)
}

func RequestErrorLog(message string, client string) {
	log.Print(Red + "[REQUEST ERROR]" + White + ": " + message + " " + Green + client + White)
}
func InfoLog(message string) {
	log.Print(Blue + "[INFO]" + White + ": " + message)
}

func ConnectionLog(client string) {
	log.Print(Yellow + "[CONNECTION ESTABLISHED]" + White + ": " + client)
}

func ClosedLog(client string) {
	log.Print(Yellow + "[CONNECTION CLOSED]" + White + ": " + client)
}

func RequestLog(message string, client string) {
	log.Print(Purple + "[REQUEST]" + White + ": " + message + " " + Green + client + White)
}

func ValidateResponseLog(stt string, client string) {
	log.Print(Green + "[VALIDATION RESPONSE]" + White + ": " + stt + " " + Green + client + White)
}

func TypesResponseLog(stt string, client string) {
	log.Print(Green + "[TYPES RESPONSE]" + White + ": " + stt + " " + Green + client + White)
}

func Log(message string) {
	log.Print(Gray + "[LOG]" + White + ": " + message)
}
