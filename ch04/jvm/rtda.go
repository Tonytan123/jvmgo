package jvm

import "jvmgo/ch04/rtda"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PushFrame(frame *rtda.Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *rtda.Frame {
	return self.stack.Top()
}
