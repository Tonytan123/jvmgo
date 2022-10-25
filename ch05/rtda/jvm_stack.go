package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) Push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.Lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}

	top := self._top
	self._top = top.Lower
	top.Lower = nil
	self.size--

	return top
}

func (self *Stack) Top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
