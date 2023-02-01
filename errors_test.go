package gerrors

import (
	"fmt"
	"os"
	"testing"
)

var (
	orderNotFind = NewCodeMsg(10000, "order not find")
)

func TestAddStack(t *testing.T) {
	fmt.Println(AddStack(orderNotFind))
}

func TestNew(t *testing.T) {

	err1 := New(1000, "业务错误1")
	fmt.Println(Resp(err1))

	err2 := Wrap(err1, "包装错误")
	fmt.Println("err22:=", err2)

	err3 := WrapCode(err1, 2000, "WrapCode")
	fmt.Println("err3:=", err3)

	//fmt.Printf("%+v\n", err1)
	//err1 = Wrap(err1, "exec0 wrap")
	//err1 = Wrap(err1, "exec1 wrap")/
	//fmt.Printf("%s\n", wrap2().Error())
	//fmt.Printf("%#+v\n", err1)
	err := wrap2()
	//fmt.Println(fmt.Sprintf("%+v \n", err))
	//fmt.Printf("%#+v\n", err)
	fmt.Printf("%s\n", err)
	// 打印结果为
	//#0 exec2 wrap D:/fotoable/open/gerrors/errors_test.go:36
	//#1 exec1 wrap D:/fotoable/open/gerrors/errors_test.go:43
	//#2 exec0 wrap D:/fotoable/open/gerrors/errors_test.go:50
	//#e open 1: The system cannot find the file specified.

	//glog.Error(context.TODO(), err)
	fmt.Printf("----------------------------------------------------------- \n")
	//glog.Errorf(context.TODO(), "%s\n", err)

	fmt.Println(Err(err))
	fmt.Printf("----------------------------------------------------------- \n")
	fmt.Println(Resp(err))
}

func wrap2() error {
	if err := wrap1(); err != nil {
		return WrapCode(err, 1004, "exec2 wrap")
	}
	return nil
}

func wrap1() error {
	if err := wrap0(); err != nil {
		return Wrap(err, "exec1 wrap")
	}
	return nil
}

func wrap0() error {
	if err := openFile(); err != nil {
		return New(1002, "exec0 wrap")
	}
	return nil
}

func openFile() error {
	if _, err := os.Open("1"); err != nil {
		//glog.Error(context.TODO(), err)
		return err
	}
	return nil
}
