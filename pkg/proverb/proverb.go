package proverb

import (
	"math/rand"
	"net"
	"sync"
	"time"
)

var (
	proverbs = []string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
		"interface{} says nothing.",
		"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		"A little copying is better than a little dependency.",
		"Syscall must always be guarded with build tags.",
		"Cgo must always be guarded with build tags.",
		"Cgo is not Go.",
		"With the unsafe package there are no guarantees.",
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
		"Don't just check errors, handle them gracefully.",
		"Design the architecture, name the components, document the details.",
		"Documentation is for users.",
		"Don't panic.",
	}
	length = len(proverbs)
	rwmu   = sync.RWMutex{}
)

// RandProverb отправляет каждые 3 секунды в conn случайную Go пословицу
func RandProverb(done chan struct{}, conn net.Conn) {
	for {
		select {
		case <-done:
			return
		default:
			conn.Write([]byte(proverb() + "\n"))
			time.Sleep(time.Second * 3)
		}
	}
}

// proverb возвращает случайный элемент из proverbs[]
func proverb() string {
	n := rand.Intn(length)
	rwmu.RLock()
	msg := proverbs[n]
	rwmu.RUnlock()
	return msg
}
