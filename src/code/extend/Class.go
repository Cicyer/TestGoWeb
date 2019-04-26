package extend

//用于解决任意结构体泛型转换的测试
type Class interface {
	newInstance() interface{}
}
