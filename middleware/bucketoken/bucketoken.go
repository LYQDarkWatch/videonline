package bucketoken

import (
	"fmt"
)

type Limiter struct {
	cont   int
	bucket chan int
}

func NewConnLimiter(cc int) *Limiter {
	return &Limiter{
		cont:   cc,
		bucket: make(chan int, cc), // buffer channel
	}
}
func (cl *Limiter) GetToken(id int) bool {
	if len(cl.bucket) >= cl.cont {
		fmt.Println("超过限制")
		return false
	}
	cl.bucket <- id
	return true
}
func (cl *Limiter) ReleaseToken() {
	c := <-cl.bucket
	fmt.Println(c)
}
