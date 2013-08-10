package log

import (
	"fmt"
)

// logger.trace('Entering cheese testing');
// logger.debug('Got cheese.');
// logger.info('Cheese is Gouda.');
// logger.warn('Cheese is quite smelly.');
// logger.error('Cheese is too ripe!');
// logger.fatal('Cheese was breeding ground for listeria.');

func Debug(format string, a ...interface{}) {
	fmt.Printf(format, a)
}
