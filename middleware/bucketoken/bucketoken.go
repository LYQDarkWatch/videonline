package bucketoken

import (
	"fmt"
)

//type ConnLimiter struct {
//	concurrentConn int  //connect的个数
//	bucket chan int
//}
//
//func NewConnLimiter(cc int) *ConnLimiter {
//	return &ConnLimiter{
//		concurrentConn: cc,
//		bucket:         make(chan int,cc),
//	}
//}
//
//func (cl *ConnLimiter) GetConn() bool {
//	if len(cl.bucket) >= cl.concurrentConn{
//		log.Printf("Reached the rate limitation")
//		return false
//	}
//	cl.bucket <- 1
//	return true
//}
//
//func (cl *ConnLimiter) ReleaseCoon() {
//	 c :=<- cl.bucket
//	 log.Printf("New connection coming: %d",c)
//}

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
