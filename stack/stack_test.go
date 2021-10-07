package stack

import (
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	ss := New()
	if ss.IsEmpty() {
	} else {
		t.Errorf("empty stack, exp: true; act:false.")
	}
	ss.data = append(ss.data, 1)
	if !ss.IsEmpty() {
	} else {
		t.Errorf("Non empty stack, exp:false, act:true.")
	}
}

func TestPeek(t *testing.T) {
	ss := New()
	ele, err := ss.Peek()
	if ele == 0 && err.Error() == "Stack is empty" {
	} else {
		t.Errorf("empty stack peek(), exp:0, Stack is empty; act:%v, %v\n", ele, err)
	}
	ss.data = append(ss.data, 1)
	ele1, err := ss.Peek()
	if ele1 == 1 && err == nil {
	} else {
		t.Errorf("[]int{1} stack peek, exp:1, nil; act:%v, %v\n", ele1, err)
	}
	ss.data = append(ss.data, 2, 3)
	ele2, err := ss.Peek()
	if ele2 == 3 && err == nil {
	} else {
		t.Errorf("[]int{1,2,3} stack peek, exp:3, nil, act:%v, %v\n", ele2, err)
	}
}

func TestPush(t *testing.T) {
	ss := New()
	ss.Push(1)
	if reflect.DeepEqual(ss.data, []int{1}) {
	} else {
		t.Errorf("push element 1 to stack, exp:[]int{1}, act:%v\n", ss.data)
	}
}

func TestPop(t *testing.T) {
	ss := New()
	ele, err := ss.Pop()
	if ele == 0 && err.Error() == "Stack is empty" {
	} else {
		t.Errorf("empty stack pop, exp:0, Stack is empty; act:%v, %v\n", ele, err)
	}
	ss.data = append(ss.data, 1)
	ele1, err := ss.Pop()
	if reflect.DeepEqual([]int{}, ss.data) && ele1 == 1 && err == nil {
	} else {
		t.Errorf("[]int{1} stack pop, exp:[]int{}, 1, nil; act: %v, %v, %v\n", ss.data, ele1, err)
	}

	ss.data = append(ss.data, 1, 2)
	ele2, err := ss.Pop()
	if reflect.DeepEqual([]int{1}, ss.data) && ele2 == 2 && err == nil {
	} else {
		t.Errorf("[]int{1} stack pop, exp:[]int{1}, 2, nil; act: %v, %v, %v\n", ss.data, ele2, err)
	}
}
