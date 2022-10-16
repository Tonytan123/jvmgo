package jvm

import "jvmgo/ch04/rtda"

type Stack struct {
	maxSize uint
	size    uint
	_top    *rtda.Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *rtda.Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.Lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *rtda.Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.Lower
	top.Lower = nil
	self.size--

	return top
}

func (self *Stack) Top() *rtda.Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
