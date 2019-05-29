package mutex

type Mutex chan struct{}

//Lock locks the funcction
func (m Mutex) Lock() {
	m <- struct{}{}

}

//Unlock unlocks the fuction
func (m Mutex) Unlock() {
	<-m
}

func New() Mutex {
	return make(Mutex, 1)
}
