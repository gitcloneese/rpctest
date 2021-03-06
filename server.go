package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)
/*
	Doc  Addr
	
 	http://docscn.studygolang.com/pkg/net/rpc/	
	
- the method's type is exported.
- the method is exported.
- the method has two arguments, both exported (or builtin) types.
- the method's second argument is a pointer.
- the method has return type error.

   func (t *T) MethodName(argType T1, replyType *T2) error

*/

type Panda int

func (this *Panda) Getinfo(argType int, replyType *int) error{
	
	fmt.Println("打印对方发过来的内容 ", argType)
	
	
	*replyType = argType + 10000
	
	return nil
	
}

func pandatest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world hello panda")
}

func main() {
	http.HandleFunc("/panda", pandatest)

	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("网络错误")
		panic("网络错误")
	}


	//类实例化
	pd := new(Panda) // Panda是一个类， pd  new一个对象
	
	rpc.Register(pd)
	rpc.HandleHTTP()




	http.Serve(ln, nil)
}
