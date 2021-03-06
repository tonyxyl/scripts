package main

import (
  "container/list"
  "fmt"
  "regexp"
)

// 双向链表
func testList() {
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack(2)
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// unsafe.Sizeof(1) 平台相关的, 32位系统返回4, 64位返回8

// 正则包
func testReg() {
  searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" //要查找的字符串
  pat := "[0-9]+.[0-9]+" //模式
  
  ok, _ := regexp.Match(pat, []byte(searchIn))
  ok, _ := regexp.MatchString(pat, searchIn)
  
  re, _ := regexp.Compile(pat)
  str := re.ReplaceAllString(searchIn, "@@@") //将匹配到的部分都替换
  str2 := re.ReplaceAllStringFunc(searchIn, f) //匹配到的部分用函数f替换
}

// 锁机制, sync包
type Info struct {
  mu sync.Mutex
  name string
}

func Update(info *Info) {
  info.mu.Lock()
  info.name = "tony"
  info.mu.Unlock()
}

// 精密计算, math/big包
im = big.NewInt(math.MaxInt64) //声明整数
io = big.NewInt(1)
ip = big.NewInt(1990)
io.Mul(im, im).add(io, im).Div(io, ip)
rm = big.NewRate(math.MaxInt64, 1990) //声明有理数, 参数1为分子, 参数2为分母

// 自定义包和可见性
import "包的本地路径或URL地址"
import "./demo/myTest"  //从当前目录demo/myTest导入
import "project/demo"  //从$GOPATH/src/project/demo导入
import "github.com/xxx/yyy"  //先从远程下载到本地$GOPATH再导入


// 结构体, 可看做类的简化形式
type identifier struct {
  field1 type1
  field2 type2
  field3 map[type]*field3
	Id int `json:"id"`  //字段名映射,
	Name string `json:"name"`
	Admin bool `json:"-"`  //"-"表示跳过该字段
  //...
}
var t *T = new(T) //struct as a pointer
t.field1 = xx
(*t).field2 = yy

var t T  // struct as a value type
t.field1 = 5

t := &T{10, "Chris"} //struct-iteral

// 二叉树结构
type Btree struct {
  pri *Node
  data float64
  su *Node
}

// 非结构体类型
type IntVector []int

type File struct {
  fd      int     // 文件描述符
  name    string  // 文件名
}

//工厂方法以new/New开头
func NewFile(fd int, name string) *File {
  if fd < 0 {
    return nil
  }
  return &File{fd, name}
}

f := NewFile(10, "./test.txt")  //实例化

//对于结构体 type A struct { a, b int} 可以使用y := new(A) 不能使用 y := make(A)
//对于结构体 type B map[string]string 可以使用y := make(B) 不能使用 y := new(B)

// 带标签的结构体
type TagType struct { // tags
  field1 bool   "An important answer"
  field2 string "The name of the thing"
  field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
  ttType := reflect.TypeOf(tt)
  ixField := ttType.Field(ix)
  fmt.Printf("%v\n", ixField.Tag)
}

// 匿名字段
type innerS struct {
  in1 int
  in2 int
}

type outerS struct {
  b int
  c float32
  int  //匿名字段
  innerS  //内嵌结构体
}

//命名冲突问题
type A struct {a int}
type B struct {a, b int}
type C struct {A; B}
var c C;
c.A.a
c.B.a
type D struct {B; b float32}
var d D  
d.b  //来自结构体D的字段b
d.B.b //继承自B的字段b

// 定义方法, recv类似OOP中的this或self; 如果方法中不需要使用recv的值, 可用_替换它
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

// 方法和函数的区别: 函数将变量作为参数, function(recv); 方法在变量上被调用, recv.method()
type TwoInts struct {a, b int}
two1 := new(TwoInts)
two1.a = 10
two1.b = 20
two1.AddItem(30)

two2 := TwoInts{3, 4}
two2.AddItem(5)

func (tn *TwoInts) AddItem(param int) int {
  return tn.a + tn.b + param
}

// 类型和作用在它上面定义的方法必须在同一个包里定义, 但有个间接方法可行
type myTime struct {
  time.time
}
fun (t myTime) first3Chars() string {
  return t.Time.String()[0:3]
}

// 指针方法和值方法都可以在指针或非指针上被调用

// 内嵌类型的方法和继承
type Engine interface {
	Start()
	Stop()
}
type car struct {  // 结构体car的实例将继承接口Engine的方法
	Engine
}

// 类型中嵌入功能, A: 聚合, 包含一个所需功能类型的具名字段; B: 内嵌所需功能类型
type Log struct {
	msg string
}
// 包含所需功能类型具名字段
type Customer struct {
	name string
	log *Log
}
// 内嵌所需功能类型
type Shop struct {
	name string
	Log
}

// 多重继承, 通过接口实现多态
type Camera struct {}
type Phone struct {}
type CameraPhone struct {
	Camera
	Phone
}

// 类型的String()方法和格式化描述符
func (t T) String() string {
	// ...
}
// %v 默认格式; %T 类型的完整规格; %#v 实例的完整输出; %+v 打印结构体会添加字段名; %% 打印百分号; %t 布尔占位符;
// %b 二进制; %c 相应Unicode码点表示的字符; %d 十进制; %o 八进制; %q 带引号; %x 十六进制小写; %X 十六进制大写;
// %U Unicode格式; %e 科学计数法; %f 浮点数

// 垃圾回收和SetFinalizer
runtime.GC()  //显示触发, 默认是自动触发的

var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("%d Kb\n", m.Alloc / 1024) //当前已分配内存Kb

// 在对象obj被GC回收之前, 执行函数func, 比如写日志
runtime.SetFinalizer(obj, func(obj *typeObj))


// 自定义标签处理器
const tagName = "validate"
type User struct {
	Id int `validate:"-"`
	Name string `validate:"presence,min=2,max=32"`
	Email string `validate:"email,required"`
}

user := User{
	Id, 1,
	Name: "John Doe",
	Email: "john@xx",
}
t := reflect.TypeOf(user)
t.Name() //User
t.Kind() //struct
t.NumField() //3
field = t.Field(0)
field.Name //Id
field.Type.Name() //int
tag := field.Tag.Get(tagName)  //-
