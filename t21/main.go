package main

import (
	"fmt"
	"strings"
)

/*
	Задание 21

	Реализовать паттерн «адаптер» на любом примере.
*/

//StringList is Adaptee
type StringList struct {
	rows []string
}

func (sl StringList) getString() string {
	return strings.Join(sl.rows, "\n")
}

func (sl *StringList) add(value string) {
	sl.rows = append(sl.rows, value)
}

//TextAdapter is Adapter
type TextAdapter struct {
	StringList
}

func (ta TextAdapter) getText() string {
	return ta.getString()
}

func getTextAdapter() TextAdapter {
	adapter := TextAdapter{}
	adapter.add("line 1")
	adapter.add("line 2")
	return adapter
}

//Client
func main() {
	adapter := getTextAdapter()
	text := adapter.getText()
	//text: line 1
	// line 2

	fmt.Println(text)
}
