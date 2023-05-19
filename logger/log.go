//
//
// @filename: logger/log.go
// @author: Krisna Pranav
// @license COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package log

import (
	"os"
	"time"

	"github.com/fatih/color"
)

type Log struct {
	Message   string
	Type      Type
	CreatedAt time.Time
}

type Type uint

const (
	TypeDefault Type = iota
	TypeWarning
	TypePanic
	TypeFatal
	TypeDebug
)

func Default(data string) {
	New(data, TypeDefault)
}

func Debug(data string) {
	New(data, TypeDebug)
}

func Warning(data string) {
	New(data, TypeWarning)
}

func Fatal(data string) {
	New(data, TypeFatal)
}

func Panic(data string) {
	New(data, TypePanic)
}

func New(message string, logType Type) {
	log := &Log{
		Message:   message,
		CreatedAt: time.Now(),
	}

	ld := color.New(color.FgGreen)

	switch logType {
	case TypePanic:
		log.Type = TypePanic

		l := color.New(color.FgRed)
		_, _ = l.Println(log.String())
		panic(1)
	case TypeFatal:
		log.Type = TypeFatal

		l := color.New(color.FgRed)
		_, _ = l.Println(log.String())
		os.Exit(1)
	case TypeWarning:
		log.Type = TypeWarning

		c := color.New(color.FgYellow)
		_, _ = c.Println(log.String())
	case TypeDebug:
		log.Type = TypeDebug

		if os.Getenv("DEBUG") == "true" {
			c := color.New(color.FgBlue)
			_, _ = c.Println(log.String())
		}
	default:
		log.Type = TypeDefault
		_, _ = ld.Println(log.String())
	}
}

func (l *Log) String() string {
	output := "[" + l.CreatedAt.Format(time.RFC822) + "] > " + l.Message

	return output
}
