package log

import (
	"fmt"
	"time"
)

var level int

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

const (
	traceColor  = "\x1b[34;3m"
	debugColor  = "\x1b[36;3m"
	infoColor   = "\x1b[32;3m"
	warnColor   = "\x1b[33;3m"
	errorColor  = "\x1b[31;3m"
	fatalColor  = "\x1b[35;3m"
	cancelColor = "\x1b[0m"
)

const timeFormat = "2006-01-02 15:04:05.999"

func printColor() {
	for color := 30; color < 38; color++ {
		fmt.Printf("%d: ", color)
		for bold := 0; bold < 10; bold++ {
			fmt.Printf("\x1b[%d;%dmHello! \x1b[0m", color, bold)
		}
		fmt.Println()
	}
}

// func print(printLevel int, format string, a ...interface{}) {
// 	// printColor()
// 	if printLevel > level {
// 		// fmt.Println("\x1b[30;1mHello, World!\x1b[0m")
// 		// fmt.Println("\x1b[31;1mHello, Red!\x1b[0m")
// 		// fmt.Println("\x1b[32;0mHello, Green!\x1b[0m")
// 		// fmt.Println("\x1b[32;1mHello, Green!\x1b[0m")
// 		// fmt.Println("\x1b[32;2mHello, Green!\x1b[0m")
// 		// fmt.Println("\x1b[32;3mHello, Green!\x1b[0m")
// 		// fmt.Println("\x1b[33;1mHello, World!\x1b[0m")
// 		// fmt.Println("\x1b[34;1mHello, World!\x1b[0m")
// 		// fmt.Println("\x1b[35;1mHello, World!\x1b[0m")
// 		// fmt.Println("\x1b[36;1mHello, World!\x1b[0m")
// 		// fmt.Println("\x1b[37;1mHello, World!\x1b[0m")
// 		// now := time.Now()
// 		// fmt.Print(now.Format("2006-01-02 15:04:05.999 "))
// 		// fmt.Printf(format, a)
// 	}
// }

func Trace(message string) {
	if level <= TRACE {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[TRACE]", message)
		fmt.Printf("%s%s%s\n", traceColor, message, cancelColor)
	}
}

func Debug(message string) {
	if level <= DEBUG {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[DEBUG]", message)
		fmt.Printf("%s%s%s\n", debugColor, message, cancelColor)

	}
}
func Info(message string) {
	if level <= INFO {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[INFO]", message)
		fmt.Printf("%s%s%s\n", infoColor, message, cancelColor)
	}

}
func Warn(message string) {
	if level <= WARN {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[WARN]", message)
		fmt.Printf("%s%s%s\n", warnColor, message, cancelColor)
	}

}
func Error(message string) {
	if level <= ERROR {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[ERROR]", message)
		fmt.Printf("%s%s%s\n", errorColor, message, cancelColor)

	}
}
func Fatal(message string) {
	if level <= FATAL {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[FATAL]", message)
		fmt.Printf("%s%s%s\n", fatalColor, message, cancelColor)
	}
}
