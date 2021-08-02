package hw04lrucache

import (
	"fmt"
	"strings"

)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	// List // Remove me after realization.
	size  int
	front *ListItem
	back  *ListItem
	// Place your code here.
}

func (l list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{v, nil, nil}

	if frontItem := l.front; frontItem != nil {
		frontItem.Prev = item
		item.Next = frontItem
	}
	l.front = item
	if l.back == nil {
		l.back = item
	}
	l.size++
	// l.PrintList()
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{v, nil, nil}

	// slog.Debug("back Item: ", backItem)
	if backItem := l.Back(); backItem != nil {
		backItem.Next = item
		item.Prev = backItem
	}
	l.back = item
	if l.front == nil {
		l.front = item
	}
	l.size++
	// l.PrintList()
	return item
}

func (l *list) Remove(i *ListItem) {
	leftItem := i.Prev
	rightItem := i.Next

	switch {
	case leftItem == nil && rightItem == nil:
		l.front = nil
		l.back = nil
	case leftItem == nil:
		rightItem.Prev = nil
		l.front = rightItem
	case rightItem == nil:
		leftItem.Next = nil
		l.back = leftItem
	default:
		leftItem.Next = rightItem
		rightItem.Prev = leftItem
	}

	l.size--
	// l.PrintList()
}

func (l *list) MoveToFront(i *ListItem) {
	value := i.Value
	l.Remove(i)
	l.PushFront(value)
	// l.PrintList()
}

func (l *list) String() string {
	var out strings.Builder
	out.WriteString(fmt.Sprintf("List: size: %d", l.Len()))

	for item := l.Front(); item != nil; item = item.Next {
		out.WriteString(fmt.Sprintf("\n%p: , %v", item, item))
	}
	return out.String()
}

func NewList() List {
	l := new(list)
	l.back = nil
	l.back = nil
	return l
}
