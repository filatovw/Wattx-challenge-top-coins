package libs

import (
	"log"
	"os"
)

func GetStdLogger(prefix string) *log.Logger {
	return log.New(os.Stdout, prefix+" ### ", log.Llongfile|log.LstdFlags|log.Lmicroseconds)
}
