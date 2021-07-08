package main

import (
	"fmt"
	"time"
)

type BookParserInterface interface {
	GetNumPages() int
}

type BookParser struct {
	book string
}

func (b BookParser) GetNumPages() int {
	return 12
}
func NewBookParser(book string) *BookParser {
	// do very heavy processing
	time.Sleep(5 * time.Second)
	return &BookParser{
		book: book,
	}
}

type LazyBookParserProxy struct {
	book       string
	bookparser *BookParser
}

func (l *LazyBookParserProxy) GetNumPages() int {
	if l.bookparser == nil {
		// instantiate when needed and keep it in cache
		fmt.Println("instantiating")
		l.bookparser = NewBookParser(l.book)
	}
	return 0
}

func NewLazyBookParserProxy(book string) LazyBookParserProxy {
	// We are not instantiating bookparser as it is heavy
	return LazyBookParserProxy{
		book:       book,
		bookparser: nil,
	}
}

func main() {
	lbp := NewLazyBookParserProxy("very huge string")
	fmt.Println("It will take 5 sec as it is first time")
	fmt.Println(lbp.GetNumPages())

	// now it wont take time
	fmt.Println("Wont take time now")
	fmt.Println(lbp.GetNumPages())
}
