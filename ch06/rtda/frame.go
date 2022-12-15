package rtda

type Frame struct {
	Lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (this *Frame) LocalVars() LocalVars {
	return this.localVars
}

func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}

func (this *Frame) Thread() *Thread {
	return this.thread
}

func (this *Frame) NextPC() int {
	return this.nextPC
}

func (this *Frame) SetNextPC(nextPC int) {
	this.SetNextPC(nextPC)
}
