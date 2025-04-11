package main//118
import (
	"fmt"//引入打印包
	"unsafe"//引入看数据类型占多少存储空间的包
	"strings"//分割字符的包
	"strconv"//转换类型的包
	"sort"//升序排序包
	"errors"//异常的包
	"time"//time包
	"encoding/json"
	//T"dsadad"//起了个别名T
	"sync"//协程的包，互斥锁的包
)
func f()(int,int){
	return 15,1
}//匿名变量
func h1() func()int{
	i := 10
	return func()int{
		return i+1
	}
}//闭包
func h3 (){
	fmt.Println("开始")
	defer func (){
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("结束")
}//defer，先赋值后执行而不是直接跳过
func h4 ()int{
	var h5 int
	defer func(){
		h5 ++
	}()
	return h5
}//匿名返回值，改不到值，先赋值，在defer最后return，因为已经赋值完了
func h6()(h7 int){
	defer func(){
		h7 ++
	}()
	return 5//返回的值，返回到h7
}//命名返回值，可以改到值
func h8 (h9 int, i1 int)int{
	defer func (){
		err := recover()
		if err != nil{
			fmt.Println("err:",err)
		}
	}()
	return h9/i1
}//recover的使用
func i2 (i3 string) error{
	if i3 == "main.go"{
		return nil
	}else{
		return errors.New("读取文件失败")//引入errors包
	}
}
func i4(){
	defer func(){
		i6 := recover()
		if i6 != nil{
			fmt.Println("有毛病")
		}
	}()
	i5 := i2("sad")
	if i5 != nil{
		panic (i5)
	}
}
type person struct{//格式  type 名字 struct
    name string//名字 类型
	age  int
    sex string
}//使用的时候必须实例化，分配内存
type person1 struct{
	Name string  `json:"姓名"` //打个tag标签
	Age int
	Sex string
}//名称首字母要大写，如果要转换成josn字符串的话
func (l4 person) l5(){
	fmt.Printf("%v,%v\n",l4.name,l4.age)
}//结构体参与定义函数
type Usber interface{
	start()
	stop()
}//接口里面有函数，通过结构体或自定义类型实现
type Phone struct{
	Name1 string
}//如果要实现接口(简单来说就是赋值操作一类的)就要实现里面的所有函数
func (m3 Phone)start(){
	fmt.Println(m3.Name1,"启动")
}
func(m3 Phone)stop(){
	fmt.Println(m3.Name1,"关机")
}
func m8(m9 interface{}){
	switch m9.(type){
	case int :
		fmt.Println("int")
	case string :
		fmt.Println("string")
	default :
       fmt.Println("失败")
	}   
}//这种类型断言方式只能搭配switch使用，里面的type是固定格式
var n2 sync.WaitGroup//先声明才能用Done和Wait
func test1(){
	for n1 := 0; n1<10; n1++{
		fmt.Println("test")
	    time.Sleep(time.Millisecond*100)
	}
	n2.Done()//协程计数器减一
}
var n9 sync.Mutex//互斥锁，先声明变量
var o1 sync.RWMutex//读写互斥锁
func main()  {
	fmt.Print("hello world\n")//\n是换行的意思
	fmt.Println("1.可以换行输出 2.输出多结果中间有空格")
	var a int = 10//1.变量声明方式 var 变量名 变量类型 赋值（变量名字不能是数字开头）
	var b = 3//2.变量类型可以省略
	 c  := 5//3.局部变量形式（变量定义完必须使用）
	fmt.Printf("a=%v b=%v c=%v a的类型是%T\n",a,b,c,a)//格式化输出
	//变量定义和初始化
	//第一种
	var username string
	username = "张三"
	fmt.Println(username)
	//第二种
	var username1 = "李四"
	fmt.Println(username1)
	//第三种（注意变量 名称 不能重复声明，但可以重复使用）
	username2 := 99
	fmt.Println(username2)
	//一次声明多个变量
	//第一种
	var d,e string//d,e都是string类型的
	d = "ddd"
	e = "eee"
	fmt.Println(d,e)
    //第二种
	var(
		username3 string
		username4 int
		username5 string
	)
	username3 = "王五"
	username4 = 55
	username5 = "一米五五"
	fmt.Println(" 姓名是",username3,"\n","体重是",username4,"\n","身高是",username5)
	//第三种
	username6,username7,username8 := "吕六" ,66 ,"一米八六"
	fmt.Println(username6,username7,username8)
	//匿名变量
	// func f()(int,int){
	// 	return 15,1
	// }
	username9,_ := f()//因为变量定义完必须使用，所以只想用一个值可以用匿名变量，那个不用的值用_表示
	fmt.Println(username9)
	_,username10 := f()
	fmt.Println(username10)//_可以重复声明
    //常量
	const g = 666//常量定义的时候就得赋值，常量的值不能被改变
	fmt.Println(g)
	//定义多个常量的方法和变量相同
	//第一种
	const h,i = 88,66
	fmt.Println(h,i)
	//第二种
	const (
		j = 1314
		k
	)//常量声明后面不赋值默认和第一个值一样
	fmt.Println(j,k)
	//常量计数器iota,搭配常量使用
	const l = iota//默认值为0
	fmt.Println(l)
	//定义多个变量的时候相当于累加1
	const (
		m = iota
		_//可以用_跳过这个值
		n	
	)
	fmt.Println(m,n)//输出的是0和2
	//iota可以中间插队
	const(
		o = iota
		p = 85
		q = iota
		r = 69
		s = iota
	)
    fmt.Println(o,p,q,r,s)
    //整型    int和uint
	//注意事项
	//类型不同的不能相加减
	var t int8 = 12
	var u int32 = 44
	fmt.Println(t,u)
    //类型转换
	//高位向低位转换 注意截尾问题
	fmt.Println(int32(t)+u)
	//类型推导定义变量int跟计算机的位数有关
	fmt.Print("占了",unsafe.Sizeof(c),"个字节\n")//没法看string类型
	v := "哥哥"
	fmt.Println(v)
	//%v是原值输出，%T是类型输出，%d是10进制，%02d宽度2不够补0，%b是二进制，%o是8进制，%x是16进制，%f是浮点输出（保留6位小数点），%.2f保留两位，%t是bool类型，%c是原样输出，%p是地址后面赋值是&符号意思是取地址
    w := 3.1415926
    fmt.Printf("原值是：%v,保留4位小数是:%.4f,类型是:%T\n",w,w,w)
    //科学计数法表示规则3.14e2表示3.14乘10的2次方
	//3.14e-2表示3.14除以10的2次方
	x := 3.14e-2
	fmt.Println(x)
	//float有精度丢失的问题
	//bool=真假。默认值是假，不能参与运算和转换
	//string默认值是空
	//转义字符串：//输出一个/，/"输出一个双引号，
	//输出多行字符串``,还可用于包含一个字符串，这样就不用引号前面加/了
	y := `吕宏磊是个1
大傻蛋`
	fmt.Println(y)
    //字符串常用操作
	fmt.Println(len(y))//len()求字符串长度，求的是字节数，一个汉字占3个字节,一个字母/一个数字占一个字节
	fmt.Println(v+y)//+拼接字符串
	fmt.Println(fmt.Sprintf("%v %v",v,y))//fmt.Sprintf也可以拼接和+同理，原理是返回打印的字符串
	//strings.Split 分割字符串
	var z = "123-456-789"
	a1 := strings.Split(z,"-")//分割完是切片类型
	fmt.Println(a1)
	fmt.Println(strings.Join(a1,"*"))//strings.Join是链接操作
	//strings.Contains（A,B）判断A中有没有B，返回值是ture和false，strings.HasPrefix判断字符串前面有没有，strings.HasSuffix判断后面有没有
	//strings.Index(A,B)B字符串从前往后第一次出现在A字符串的第几位，注意从零开始数
	// strings.LastIndex（A,B）B字符串从后往前第一次出现在A字符串的第几位，查找不到返回-1
	var a2 = 'a'
	fmt.Printf("值是%v,类型是%T,原样输出是%c\n",a2,a2,a2)//字符返回的是对应的码值，是int类型，汉字对应utf8
	//字符有byte和rune类型，后者遍历用range
    a3 := "你好C语言"
	runea3 := []rune(a3)//修改字符串需要先转换
	runea3[1] = '大'//数组从0开始计数
	fmt.Println(runea3)//同上返回的是码值而不是字符串
	fmt.Println(string(runea3))//在转换为字符串
	//转换为string类型
	//第一种
	a4 := 20
	a5 := fmt.Sprintf("%d",a4)//返回string类型 20 的值
	fmt.Printf("%v,%T\n",a5,a5)
	//第二种，用strconv包
	a6 := 33
	a7 := strconv.FormatInt(int64(a6),10)//参数1是固定格式，参数2是进制
	fmt.Printf("%v,%T\n",a7,a7)//int转化为string类型
	a8 := 3.1415926
	a9 := strconv.FormatFloat(float64(a8),'f',4,64)//参数1是固定格式，参数2是格式化类型，参数3是保留小数位-1代表不格式化，格式化类型32/64
	fmt.Printf("%v,%T\n",a9,a9)//float转化为string类型
	//string转化为int类型
	b1 := "123456"
	b2,_ := strconv.ParseInt(b1,10,64)//结果返回两个值，0表示转化失败   第一个参数确定哪个值，第二个参数几进制，第三个参数返回位数8/16/32/64
	fmt.Printf("%v,%T\n",b2,b2)
	//string转化为float类型
	b3 := "4.656968"
	b4,_ :=strconv.ParseFloat(b3,64)
	fmt.Printf("%v,%T\n",b4,b4)
	//算数运算符
	//除法，整数除完不保留小数，浮点型保留小数（被除数是浮点型，除数可以不是）
	//取余计算公式=被除数-（被除数/除数）*除数
	//自增++和自减--不是运算符，不能和赋值运算一块使用，--/++只能出现在后面如A++
	//关系运算符==/！=/>/<..
	//逻辑运算符&&/||
	//逻辑与和逻辑或的短路问题
	//赋值运算符+=/-=/*=先加减在赋值
	//位运算符是针对二进制数进行运算的
	//左移几位是乘以2的几次方，右移是除以2的及几次方保留整数
	b5 := 8
	fmt.Println(b5<<2)
    //if语句，{}不能省略并且必须紧挨着条件
	//第一种
	b6 := 33
	if b6 > 20{
		fmt.Println("666老铁")
	}//区别当前区域全局变量
	//第二种
	if b7 := 44;b7 > 30{
		fmt.Println("666老铁")
	}//局部变量
	//第一种   for循环格式：for 1初始语句；2条件表达式；3结束语句{}  执行顺序是1 2 {} 3 2 {} 3......
	
	for b8 := 1;b8 <= 10;b8++{
		fmt.Println(b8)
	}
    //第二种   初始化语句可以去掉但是 ； 必须有
	//第三种   初始化和结束语句都可以去掉  注意死循环 不用加；
	//第四种   都可以去掉但是要注意什么时候跳出循环
    //for嵌套，99乘法表
	for b9 := 1; b9 <= 9; b9++{
		for c1 := 1; c1 <= b9; c1++{
			fmt.Printf("%vX%v=%v    ",b9,c1,b9*c1)
		}
		fmt.Println("")
	}
	//for range（遍历字符串，切片,通道）
	c4 := "为什么不回消息whss"
	for c2, c3 := range c4{
		fmt.Printf("位置是：%v , 字符是：%c\n",c2,c3)
	}
	//switch case
	c5 := ".html"
	switch c5 {
    case ".html", "123"://一个case可以有多个分支
	    fmt.Println("html")
		break//可以不写
	case "cs"://case后面可以跟表达式例如判断语句，但是switch后面不能跟变量如c5
		fmt.Println("cs")
		break
	default	:
        fmt.Println("找不到")
	}
    //穿透 fallthrought 语句，只能穿透一层，位置跟break一样
	//break在for循环中是退出当前循环
	//lable是跳出整个循环,名字可以随便起
c6 :
    for c7 := 0 ;c7 <= 2 ;c7++{
		for c8 := 0 ; c8 < 10 ;c8++{
			if c8 == 4{
				break c6
			}
			fmt.Printf("%v,%v\n",c7,c8)
		}
	}
	//continue语句跳过当前循环，只能在for循环中使用	
	//continue+lable是跳出这次循环定位到label继续执行
	//goto+lable（别名）是定位到lable（别名）继续执行
	//数组（长度不可以被改变）
    //数组的长度也是类型的一部分，可以用%T查看
	//第一种声明数组的方法
	var c9 [3] int 
	c9[0] = 1
	c9[1] = 3
	c9[2] = 5
	fmt.Printf("%T,%v\n",c9,c9)
	//第二种
	var d1 = [2] int{2,3}
	fmt.Println(d1)
    //第三种  类型推导方式  同上
    //第四种
	d2 := [...]int{0:25,23,25,215,265,2545,10:15}//冒号前面是下标
	fmt.Println(d2)
	fmt.Println(len(d2))//len()可以打印数组长度
	//数组是值类型，新建一个内存存储数据
	//切片是引用类型，共享内存
	//二维数组以前面的中括号为主导可以用[...]表示，后面的不行
	var d3 = [...][2]string{
		{"qqqq","ddd"},
		{"dasd","www"},
	}//注意逗号问题
	fmt.Println(d3)
	fmt.Println(len(d3))//以行为主
	//切片，切片默认值是nil
	//[]内没有数字就是切片
	//基于数组定义切片
	d4 := [5]int{11,22,33,44,5}
	d5 := d4[:]
	fmt.Printf("%v,%T\n",d5,d5)
    d6 := d4[1:4]
	fmt.Printf("%v,%T\n",d6,d6)//包含1不包含4，这是下标，截取数组的一段
	d7 := d4[2:]
	fmt.Printf("%v,%T\n",d7,d7)
    //同理右边也可以这么写
	//基于切片的切片，基于数组的切片同理
    d8 := []int{11,22,33,44,55,66}
	d9 := d8[2:]
	e1 := d8[:3]
	fmt.Printf("%v,长度：%v,容量: %v\n",d9,len(d9),cap(d9))
    fmt.Printf("%v,长度：%v,容量: %v\n",e1,len(e1),cap(e1))//容量是当前第一个元素到底层数组最后一个元素的个数
	//make函数构造切片格式make([]元素类型，元素数量(长度)，切片容量)
    //给切片扩容需要用到append（）方法
	var e2 []int
	e2 = append(e2,15)//扩容
	fmt.Println(e2) 
    //append还可以合并切片
	e3 := []string{"dsada","dsadas"}
	e4 := []string{"qqqq","eeeee"}
	e3 = append(e3,e4...)//注意形式
	fmt.Println(e3)
	//复制切片，这样就不会影响原来的值，用copy（想要的，被复制的）方法
    e5 := []int{1,2,3,4,5}
	e6 := make([]int,5,5)//make 定义切片
	copy(e6,e5)
	fmt.Println(e6)
	//字符串不能修改 所以要化成byte和rune形式当成切片去修改，直接输出的话是码值因为被拆分了
	//选择排序，冒泡排序，sort包
    sort.Ints(e5)//升序
	fmt.Println(e5)
    sort.Sort(sort.Reverse(sort.IntSlice(e5)))//降序
	fmt.Println(e5)
	//map数据类型（也是引用数据类型）引用数据类型定义的时候就得去分配内存空间（make（））而不是像数据类型可以先声明后再赋值相当于分配内存空间
	//第一种
    e7 := make(map[string]int)
	e7["年龄"] = 23
	e7["身高"] = 173
	e7["体重"] = 60
	fmt.Println(e7)
	//第二种
	e8 := map[string]int{
		"身高":155,
		"体重":55,
		"年龄":24,
	}
	fmt.Println(e8)
	fmt.Printf("%T\n",e8)
	//for range 循环遍历数据
	for f3,f4 := range e8{
		fmt.Println(f3,f4)
	}
	//查找
	e9,f1 := e8["年龄"]
	fmt.Println(e9,f1)
    //删除  delete（对象，key）
	delete (e8,"身高")
	fmt.Println(e8)
    //定义切片在切片内放置用户信息map类型的切片
	f2 := make([]map[string]int,3,3)
	if f2[0] ==nil{
		f2[0] = make (map[string]int)
		f2[0]["身高"] = 173
		f2[0]["体重"] = 60
	}
	if f2[1] == nil{
		f2[1] = make(map[string]int)
		f2[1]["身高"] = 155
		f2[1]["体重"] = 55
	}
	fmt.Println(f2)//初始值是nil
	for f5,f6 := range f2{
		for f7,f8 := range f6{
			fmt.Printf("%v,%v,%v\n",f5,f7,f8)

		}
	}
    //切片类型的map
    f9 := make(map[string][]string)
	f9["小吕"] = []string{
		"xxx",
        "ooo",
		"ggg",
	}
	f9["小李"] = []string{
		"yyy",
		"kkk",
		"lll",
	}
	fmt.Println(f9)
    //数组是值类型
	//map和切片是引用类型
	//对k进行排序（签名算法）
	g1 := make(map[int]int)
	g1[7] = 12
	g1[2] = 46
	g1[8] = 100
	g1[5] = 33
	var g2 []int
	for g3,_ := range g1{//值，脚标
		g2 = append(g2,g3)
	}
	sort.Ints(g2)
	for _,g4 := range g2{
		fmt.Printf("key:%v,value:%v\n",g4,g1[g4])
	}
	//统计单词出现的次数
	g5 := "sss,aaa,bbb,bbb,ccc,qqq,ddd,sss"
	g6 := strings.Split(g5,",")
	g7 := make(map[string]int)
	for _,g8 :=range g6{
		g7[g8]++//作用于外部的g7
	}
	fmt.Println(g7)
    //函数格式
	//func 函数名(接收参数 数据类型，参数 数据类型 ) （返回的类型，返回的类型）{}
	//调用的时候格式，函数名（传入的参数）
    //函数名(参数 ... 数据类型) 返回的类型{}传入的数据形式是切片
	//两种可以结合使用 但是可变参数不能放在前面
	//用函数定义一个数据类型
	//格式   type 函数名 func （参数类型，参数类型） 返回值类型，类型是func （int，int）int 类型 
	//函数作为另一个函数的参数或者作为返回值可以用匿名函数，匿名函数没有函数名 其他一样
    func (g8,g9 int){
		fmt.Println(g9+g8)
	}(10,20)//匿名函数自执行函数
	//没有封装的函数不能放到另一个函数里面，要用匿名函数
	//闭包，常驻内存，不污染全局17
	//格式，函数嵌套一个函数然后返回一个函数
    h2 := h1()
	fmt.Println(h2())//闭包
	//defer延迟执行倒叙，注册正序因为先赋值后执行
    h3()
    fmt.Println(h4())
	fmt.Println(h6())
    //panic（抛出一个异常）/recover（只有在defer调用的函数中有效）（==nil是没有异常）
	fmt.Println(h8(10,0))
    i4 ()
	//time包 和 日期函数
	i7 := time.Now()
	fmt.Println(i7)
    i8 := i7.Year() 
	i9 := i7.Month() 
	j1 := i7.Day() 
	j2 := i7.Hour() 
	j3 := i7.Minute() 
	j4 := i7.Second()
	fmt.Printf("%02d-%02d-%02d   %02d:%02d:%02d\n",i8,i9,j1,j2,j3,j4) 
	//Format格式化输出时间
	j5 := i7.Format("2006-01-02 03:04:05")//固定形式 03是12小时制，15是24小时制
    fmt.Println(j5)
	//时间戳1970-1-1 8点整 到现在的秒数
	j6 := i7.Unix()//秒
	fmt.Println(j6)
	j7 := i7.UnixNano()//毫秒
	fmt.Println(j7)
	//把时间戳转换为日期
	j8 := 1743038095571573100
	j9 := time.Unix(0,int64(j8))//类型是int64，前面是秒 后面是毫秒
	k1 := j9.Format("2006-01-02 03:04:05")
	fmt.Println(k1)
	//日期转化为时间戳
	//格式
	k2 := "2025-03-27 09:15:32"
	k3 := "2006-01-02 03:04:05"
	k4,_ := time.ParseInLocation(k3,k2,time.Local)//(格式，时间，固定形式)返回值（时间，error）
	fmt.Println(k4.Unix())
	//时间操作函数(Add（+），Sub，Equal...)
	i7 = i7.Add(time.Hour)
	fmt.Println(i7.Format("2006-01-02 03:04:05"))
	//定时器
    k5 := time.NewTicker(time.Second)
	k6 := 3
	for k7 := range k5.C{
		k6--
		fmt.Println(k7)
		if k6 == 0{
			k5.Stop()//终止定时器
			break

		}
	}
	//休眠
	fmt.Println("aaa")
	time.Sleep(time.Second*2)//休眠
	fmt.Println("bbb")
	//指针  类型是*int/*string... 
	k8 := 12
	k9 := &k8
	fmt.Println(k8)
	*k9 = 14//利用指针改了目标的值
	fmt.Println(k8)
	//make和new都是声明内存空间用于引用类数据类型
	//定义一个指针
	l1 := new(int)
	fmt.Printf("%v,%T,%v\n",l1,l1,*l1)
	//类型别名
	//格式
	//type 取的名字 = 自带的数据类型例如int/float...(在main函数外定义)打印出来的还是以前的类型
    //结构体实例化方法  实例化才能分配内存  打印结构体一般用%#v  结构体大写的代表公有的，小写的代表私有的（结构体是值类型）
	//第一种
    var l2 person
	l2.name = "吕宏磊"
	l2.sex = "男"
	l2.age = 23
	fmt.Printf("%#v\n",l2)
	//第二种   结构体也可以是指针
	//p := new(person)
	//p. age = ...底层实际实现的是（*p=...）
	//第三种   
	//p := &person{}
	//第四种
	l3 := person{
		name:"里老虎",
		age:25,
		sex:"女",
	}
	fmt.Printf("%#v\n",l3)
	//第五种  类型是指针 可以改动对应的值因为是指针指向的是源地址的数据
	//  l3 := &person{
	// 	name:"里老虎",
	// 	age:25,
	// 	sex:"女",
	// }
	//第七种 赋值的时候前面的key可以不写，但是数据要跟结构体对应
	l2.l5()//结构参与定义函数
	//结构体匿名字段（没有名字省略了），数据类型不能重复
	//结构体也可以嵌套
	//修改有嵌套的结构体，首先去找本结构体的，然后再去找嵌套的，如果是同一级的结构体则会出错，所以要指定是那个结构体里面的
	//JSON字符串
	//结构体转换成json字符串，结构体里面的数据首字母是大写的，因为要是公有的,私有的只能在本个包使用，公有的在别的包也能使用
	var l6 = person1{
		Name : "林徽因",
		Age : 12,
		Sex : "女",
	}
	fmt.Println(l6)
	l7,_ := json.Marshal(l6)//调用的返回值是byte类型的切片和error
	l8 := string(l7)
	fmt.Printf("%v\n",l8)
	//josn字符串转换成结构体
	l9 := `{"Name":"林徽因","Age":12,"Sex":"女"}`	
	var m1 person1
	m2 := json.Unmarshal([]byte(l9),&m1)//返回的error类型，赋的值是byte类型的切片和地址（因为要改变值所以用的是地址），因为m1定义的不是指针类型，如果是指针类型的创建方式，则不用&
	if m2 != nil{//表示失败
		fmt.Println(m2)
	}
	fmt.Printf("%#v\n",m1)
	//josn转化成嵌套结构体与转化成普通结构体同理
	//一个项目只有一个可执行的文件main包
	//先执行init函数在执行main函数，导用多个包的时候嵌套在最里面的init函数先执行
    //import引用完第三方包 然后再终端输入go mod tidy进行下载
	//接口  是一种类型    是一种规范要满足里面的所有函数   一组函数的集合也可以没有 但是接口中不包含变量
    m4 := Phone{
		Name1:"小米手机",
	}
	m4.start()
    //空接口  可以以当作数据类型使用  当作任意类型
	//类型断言 看看空接口代表的是什么类型
	//格式
	var m5 interface{}
	m5 = "dasdasdsad"
	m6,m7 := m5.(string)//m6是值，m7是判断是不是ture  也可以断言结构体类型是不是string类型
	if m7{
		fmt.Println("字符串",m6)
	}else{
		fmt.Println("失败")
	}
    m8("dsad")//断言的另一种格式
    //接口也可以嵌套 前提也是要有里面嵌套的所有函数
	//接口当作数据类型没法直接访问里面的细节，只能一个整块的去访问，可以用类型断言因为第一个返回的值是原来的数据
	//并发多个竞争一个谁抢到谁执行，并行一个时间段可以有多个线程
	//goroutine协程，调度开销小  并行执行
    //格式是go  紧跟一个函数，有一个问题如果协程的速度慢，main函数执行的快，不管协程有没有结束程序都终止了
    n2.Add(1)//计数器加一
	go test1()
	n2.Wait()//等待协程执行完毕
    //channel管道也是一种特殊的类型 引用数据类型
	n3 := make(chan int,3)//int类型的管道,容量是3
	n3<-10
	n3<-20
	n3<-30
	n4 := <-n3
	fmt.Println(n4)
	<-n3//只取值不赋值
	n5 := <-n3
	fmt.Println(n5)
	//管道的阻塞
	//for range遍历管道没有key只有一个返回值，循环遍历要记住关闭管道
	n6 := make(chan int,10)
	for n7 := 1;n7<=10;n7++{
		n6<-n7
	}
	close(n6)//关闭管道
	for n8 := range n6{
		fmt.Println(n8)
	}
    //单向管道
	//x := make(chan<- int,10 )//只写管道
	//y := make(<-chan int,10)//只读管道
	//select多路复用(类似并发执行)
	o3 := make(chan int, 3)
	o3 <- 1
	o3 <- 2
	o3 <- 3
	o4 := make(chan int, 3 )
	o4 <-4
	o4 <- 5
	o4 <-6
    for {
		select {
		case o2 := <- o3:
		    fmt.Printf("o3,value:%v\n",o2)
		case o2 := <- o4:
			fmt.Printf("o4,value:%v\n",o2)
		default:
			fmt.Println("over")
			return	
		}
	}
	//协程的panic的运用，如果多个协程同时执行有一个错的协程会影响其他协程的执行，所以用到panic
	//互斥锁
}//入口

