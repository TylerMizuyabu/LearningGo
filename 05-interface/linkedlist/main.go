package main

import (
	"fmt"
	"sort"
)

type LinkedList struct {
	Head int
	Next *LinkedList
	Size int
}

func (l LinkedList) Len() int {
	return l.Size
}

func (l LinkedList) Less(i, j int) bool {
	if i > l.Size || j > l.Size || i == j {
		return false
	}
	list := &l
	for x := i; x > 0; x-- {
		list = list.Next
	}
	list2 := &l
	for y := j; y > 0; y-- {
		list2 = list2.Next
	}
	if i < j {
		return list.Head > list2.Head
	} else {
		return list2.Head > list.Head
	}
}

func (l LinkedList) Swap(i, j int) {
	if !(i > l.Size || j > l.Size || i == j) {
		list := &l
		for x := i; x > 0; x-- {
			list = list.Next
		}
		list2 := &l
		for y := j; y > 0; y-- {
			list2 = list2.Next
		}
		temp := list2.Head
		list2.Head = list.Head
		list.Head = temp
	}
}

/*
Append func takes a pointer of a LinkedList so I wouldn't have to return a new
head in order to update the LinkedList. Otherwise l was just a copy of the list
I want to edit. different version of Append in comments below
*/
func (l *LinkedList) Append(value int) {
	l.Size++
	if l.Size == 1 {
		l.Head = value
		l.Next = &LinkedList{}
	} else {
		l.Next.Append(value)
	}
}

/*
Second Version of Append
Would append values to linked list by calling like this:
list = *(list.Append(10))
Is a little more annoying to use. Could perhaps make it better if append took
a pointer to a LinkedList and then returned a LinkedList. This way you could call
it by typing, list = list.Append(&list,10). Will probably make third version

func (l LinkedList) Append(value int) *LinkedList {
	l.Size++
	if l.Size == 1 {
		l.Head = value
		l.Next = &LinkedList{}
	} else {
		l.Next = l.Next.Append(value)
	}
	return &l
}
*/

func main() {
	var list LinkedList
	list.Append(10)
	list.Append(50)
	list.Append(30)
	list.Append(45)
	for current := list; current != (LinkedList{}); current = *current.Next {
		fmt.Println(current)
	}
	list.Swap(1, 2)
	list.Swap(2, 3)
	for current := list; current != (LinkedList{}); current = *current.Next {
		fmt.Println(current)
	}

	fmt.Println(sort.IsSorted(list))

}
