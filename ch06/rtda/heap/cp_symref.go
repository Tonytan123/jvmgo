package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolveClass() *Class {
	if self.class == nil {
		self.ResolveClass()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
