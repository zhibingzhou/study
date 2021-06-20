package common

import (
	"fmt"
	"reflect"
)

type School struct {
	Name    string
	Address string
}

type Student struct {
	Name   string
	Age    int
	Gender string
}

func (sc School) Show() {
	fmt.Println("增加学校名称：", sc.Name, "地址：", sc.Address)
}

func (st Student) Show() {
	fmt.Println("增加学生姓名：", st.Name, "性别", st.Gender, "年龄：", st.Age)
}

func (sc School) FShow(name, address string) {
	fmt.Println("我是有参：增加学校名称：", name, "地址：", address)
}

func (st Student) FShow(name, gender string, age int) {
	fmt.Println("我是有参：增加学生姓名：", name, "性别", gender, "年龄：", age)
}

//通过函数名称来调用
func rfRun(rt interface{}, gv []reflect.Value) {
	//先获得对象
	getValue := reflect.ValueOf(rt)
	//无参数
	if len(gv) <= 0 {
		methodvalue := getValue.MethodByName("Show")
		methodvalue.Call(gv)
	} else {
		//有参数
		methodvalue := getValue.MethodByName("FShow")
		methodvalue.Call(gv)
	}

}

func Test() {
	MapAndSlice()
	//反射
	var Gsh School
	var Gst Student
	//有值
	gshool := []reflect.Value{reflect.ValueOf("zhonghua"), reflect.ValueOf("china")}
	rfRun(Gsh, gshool)
	gstudent := []reflect.Value{reflect.ValueOf("xiaoming"), reflect.ValueOf("Girl"), reflect.ValueOf(20)}
	rfRun(Gst, gstudent)

	Gsh = School{Name: "Nzhonghua", Address: "Nchina"}
	Gst = Student{Name: "Nxiaoming", Age: 20, Gender: "Girl"}
	//无值
	rfRun(Gsh, make([]reflect.Value, 0))
	rfRun(Gst, make([]reflect.Value, 0))

}












/*

 购物车
 订单
     接口（）新增
 新增

 同一个业务，如果有多种方法，如果有不同的实现方法可以用接口定义 ，好比支付业务有多种渠道，但实现的功能是一样的，最终返回一个接口的类，让类去执行，这个是返回接口的实例

 不同的业务逻辑，可以定义同一个方法，让反射去执行，事件，不同的事件有不同的方法 ， 单聊消息事件 ， 群聊消息事件 返回为一个接口，接口反射再执行 ，这个就是ingerface{},返回不知道什么类型的情况
*/

func TestPay() {
	var a AliYun
	GoPay(a)
}

func GoPay(p PayServer) {
	p.Pay()
	p.Notify()
}

type PayServer interface {
	Pay()
	Notify()
}

type AliYun struct {
	User_id string
}

type PinDuo struct {
	SystemId string
}

func (a AliYun) Pay() {
	fmt.Println("我是阿里云支付")
}

func (a AliYun) Notify() {
	fmt.Println("我是阿里云回调")
}

func (a PinDuo) Pay() {
	fmt.Println("我是拼多多支付")
}

func (a PinDuo) Notify() {
	fmt.Println("我是拼多多回调")
}

func MapAndSlice() {
	stringSlice := make([]string, 0)
	stringMap := make(map[string]string)

	sliceType := reflect.TypeOf(stringSlice)
	mapType := reflect.TypeOf(stringMap)

	rMap := reflect.MakeMap(mapType)
	rSlice := reflect.MakeSlice(sliceType, 0, 0)

	k := "first"
	rMap.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf("test"))
	i := rMap.Interface().(map[string]string)
	fmt.Println(i)

	reflect.Append(rSlice, reflect.ValueOf("test slice"))
	strings := rSlice.Interface().([]string)
	fmt.Println(strings)
}
