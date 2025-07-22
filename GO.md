



# GO

## Basic

//当前程序中包名 

package main

//导包

import"fmt"

//创建函数

func main(){

​	//

​	fmt.Println("...")

}

## 变量

- 方法一：var a int //var关键字 + a函数名 + int类型
- 方法二：var b int = 100
- 方法三：var c = 100 //通过值自动匹配变量类型
- 方法四（常用，但不能定义全局变量）：e := 100 //省去var关键字，直接自动匹配
- 多变量声明：var kk, ll = 100, "Anthea"

## 常量

const length int = 10 //不可更改

### tips: iota

可以在const( ) 中添加一个关键字iota，每行的iota都会累加1，第一行的iota默认是0

const(

BEIJING = iota

SHANGHAI

SHENZHEN

)

​	例子：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250624202439249.png" alt="image-20250624202439249" style="zoom: 67%;" />

​	注：**iota只能出现在const中**

## nil（null）

nil 不仅仅是一个简单的空值，它的**类型是动态的**，取决于它被用在什么上下文中。nil 可以是以下类型的零值：

1. **指针 (Pointers)**: var p *int = nil
2. **切片 (Slices)**: var s []int = nil
3. **映射 (Maps)**: var m map[string]int = nil
4. **通道 (Channels)**: var c chan int = nil
5. **函数类型 (Functions)**: var f func() = nil
6. **接口 (Interfaces)**: var i interface{} = nil

**注意：** nil 不能用于像 int, string, struct 这样的基本值类型。它们的零值分别是 0, "", 和一个所有字段都为零值的结构体。

## 函数

​	常规函数，返回多个参数：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250624204530405.png" alt="image-20250624204530405" style="zoom:50%;" />

​	有名称的返回值赋值：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250624204510796.png" alt="image-20250624204510796" style="zoom:50%;" />

### init函数与import导包

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250624220458563.png" alt="image-20250624220458563" style="zoom:50%;" />

匿名导包：可用于只需要其 init() 方法，不需要调包的情况

​	import _ "project/package"

别名导包：可用于需要给导入的包起别名的情况

​	import kun "project/package"

点导包：将包中的全部方法全部导入，使用时可不加包名

​	import . "project/package"

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250624230548030.png" alt="image-20250624230548030" style="zoom: 50%;" />

## 指针

同C语言

## defer关键字

用于函数中，函数体结束时执行。**最常见的用途是进行资源清理，比如关闭文件、释放锁、关闭数据库连接**

**注：多个执行，栈顺序**

defer fmt.Println("main end")

## 普通数组

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625141807791.png" alt="image-20250625141807791" style="zoom: 50%;" />

## 切片slice(动态数组)

相当于传入数组的指针，不限制数组的大小，且可更改数组的值

 <img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625141648583.png" alt="image-20250625141648583" style="zoom:50%;" />

### len / capacity

切片的长度和容量不同，长度表示左指针至有指针的距离，容量标书左指针至底层数组末尾的距离

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625162823448.png" alt="image-20250625162823448" style="zoom:50%;" />

### 扩容机制

切片在扩容时，如果扩容后len的长度超过容量，增将capacity增加至二倍

## 切片操作

a1 = a[0:1] //a1取得a的左闭右开区间

a1 的 capacity 取 a 的 capacity，a1 的 len 只取切片的长度

**注：切片取得的只是索引（浅拷贝），并不是拷贝**

​	深拷贝：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625164525289.png" alt="image-20250625164525289" style="zoom: 67%;" />

## map

map的增删改查

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625172009378.png" alt="image-20250625172009378" style="zoom:50%;" />

## 结构体

### 封装

​	结构体创建

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625173833654.png" alt="image-20250625173833654" style="zoom:50%;" />

​	结构体指针传递

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625173911270.png" alt="image-20250625173911270" style="zoom:50%;" />



​	一般结构体独有方法，最好用指针传参

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625175912240.png" alt="image-20250625175912240" style="zoom: 33%;" />

​	**注：结构体本身包括成员变量的首字母大小写，可决定是否可包外访问（private/pubic）**

### 继承

​	子类直接将父类写入其中即可

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625181820883.png" alt="image-20250625181820883" style="zoom: 50%;" />	

​	可以重写父类的方法，也可以增加子类的方法

​	初始化时，可以用父类嵌套赋值，也可以定义类型后单独赋值

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250625181905129.png" alt="image-20250625181905129" style="zoom:50%;" />

### 多态（interface）

基本要素：

- 有一个父类（有接口）
- 有子类（实现了父类接口的全部方法）
- 父类类型指针指向子类具体数据变量

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250626200433397.png" alt="image-20250626200433397" style="zoom:50%;" />

## 断言interface{}

​	例：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250626203014989.png" alt="image-20250626203014989" style="zoom:50%;" />

​	断言机制：args.(string) //只有args为 interface{} 时，才能使用这种用法

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250626203115593.png" alt="image-20250626203115593" style="zoom:50%;" />

## 反射

​	变量结构：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250626215050926.png" alt="image-20250626215050926" style="zoom:25%;" />

reflect.TypeOf( ) //返回变量的类型

reflect.TypeOf( ) //返回变量的值

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250627110528840.png" alt="image-20250627110528840" style="zoom:50%;" />

​	获取结构的含有的变量和方法：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250627120823291.png" alt="image-20250627120823291" style="zoom: 33%;" />

## 结构体标签

定义结构体时，可以给变量加标签，以便于对变量进行解释，也可以将结构体变量转换为json

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250627124048762.png" alt="image-20250627124048762" style="zoom:50%;" />

结构体 -> JSON / JSON -> 结构体

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250627160033900.png" alt="image-20250627160033900" style="zoom:50%;" />

orm映射关系



## Go并发的基石——GMP模型

| **复用线程**  | M:N调度模型，G在IO阻塞时会让出M                              | Go是用户态调度，成本极低；Java等是OS线程，成本高。           |
| ------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **利用并行**  | 默认创建与CPU核数相等的P，每个P由一个M执行                   | 启动并行任务的心智负担和资源成本远低于传统线程模型。         |
| **抢占**      | 协作式抢占（函数调用检查 + Sysmon后台监控）                  | Go是用户态抢占，轻量；传统语言依赖OS内核级抢占，笨重。       |
| **全局G队列** | P有本地队列(LRQ)，还有一个全局队列(GRQ)用于负载均衡和**工作窃取(Work Stealing)** | Go将本地队列和工作窃取结合，兼顾了缓存效率和负载均衡，设计非常精妙。 |

 

## goroutine

runtime.Goexit() //当前所有go进行退出

协程 / 匿名协程

 <img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250628120833135.png" alt="image-20250628120833135" style="zoom:50%;" />

## channel

​	无缓冲channel的基本定义和使用：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250628125416353.png" alt="image-20250628125416353" style="zoom:50%;" />

​	有缓冲的channel：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250628191650143.png" alt="image-20250628191650143" style="zoom:50%;" />

​	channel的关闭：

​	c := make(chan int)

​	close(c)

**channel 与 range**

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250628215842352.png" alt="image-20250628215842352" style="zoom:50%;" />

select 与 channel

可以多路监听 channel 状态

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250628220114324.png" alt="image-20250628220114324" style="zoom: 50%;" />

## Go Modules

GOPATH工作模式缺陷：

- 无版本控制概念
- 无法同步一致第三方版本号
- 无法指定当前项目引用的第三方版本号

GOPROXY

GOSUMDB

GONOPROXY

GONOSUMDB

GOPRIVATE：直接找私有仓库

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250629113218330.png" alt="image-20250629113218330" style="zoom: 50%;" />

手动down依赖： 

​	go get ...



## 可变参数：

**例子：**

```
num, err := UserGroupMemberCount(ctx, "user_id = ? and group_id in ?", u.Id, ugids)
```

```
func UserGroupMemberCount(ctx *ctx.Context, where string, args ...interface{}) (int64, error)
```

第三个参数 `args ...interface{}` 会**收集所有剩余的参数**，把它们放进一个名为 `args` 的切片里。所以，在函数内部，`args` 就变成了 `[u.Id, ugids]`。



# Go并发：

临界区：如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，这部分被保护起来的程序，就叫做临界区。

## 同步原语

包含：互斥锁 Mutex、读写锁 RWMutex、并发编排 WaitGroup、条件变量 Cond、Channel 等

适用场景：

- 共享资源。并发地读写共享资源，会出现数据竞争（data race）的问题，所以需要 Mutex、RWMutex 这样的并发原语来保护。
- 任务编排。需要 goroutine 按照一定的规律执行，而 goroutine 之间有相互等待或者依赖的顺序关系，我们常常使用 WaitGroup 或者 Channel 来实现。
- 消息传递。信息交流以及不同的 goroutine 之间的线程安全的数据交流，常常使用 Channel 来实现。

互斥锁 Mutex 就提供两个方法 Lock 和 Unlock：进入临界区之前调用 Lock 方法，退出临界区的时候调用 Unlock 方法：

```
func(m *Mutex)Lock() 

func(m *Mutex)Unlock()
```

## race detector：

Go 提供了一个检测并发访问共享资源是否有问题的工具，帮助我们自动发现程序有没有 data race 的问题

例：

go run -race counter.go

## mutex：

用法一：直接使用

```
// 互斥锁保护计数器 
var mu sync.Mutex

mu.Lock() 
count++ 
mu.Unlock()
```

用法二：结构体中嵌套使用

```
type Counter struct {
    mu    sync.Mutex
    Count uint64
}

counter.Lock() 
counter.Count++ 
counter.Unlock()
```

用法三：结构体嵌套 + 方法嵌套

```
counter.Incr() // 受到锁保护的方法

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}
```



# 即时通信系统

​	架构图：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141233799.png" alt="image-20250630141233799" style="zoom: 33%;" />

​	版本迭代：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141318224.png" alt="image-20250630141318224" style="zoom:50%;" />

# 生态拓展

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141337452.png" alt="image-20250630141337452" style="zoom: 67%;" />

web：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141441612.png" alt="image-20250630141441612" style="zoom:25%;" />

微服务：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141523043.png" alt="image-20250630141523043" style="zoom:25%;" />

容器编排：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141612254.png" alt="image-20250630141612254" style="zoom:25%;" />

服务发现：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141633728.png" alt="image-20250630141633728" style="zoom:25%;" />

存储引擎：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141646666.png" alt="image-20250630141646666" style="zoom:25%;" />

静态建站：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141734403.png" alt="image-20250630141734403" style="zoom:25%;" />

中间件：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141813887.png" alt="image-20250630141813887" style="zoom:25%;" />

爬虫：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250630141857882.png" alt="image-20250630141857882" style="zoom:25%;" />





# Gin + Gorm 框架

## POST vs PUT的区别

**幂等**（一个操作，你执行**一次**和执行**一百次**，对资源产生的**最终结果是完全相同**的）：PUT，GET，DELETE

**非幂等**（每执行一次，都会对资源产生**新的、额外的**影响）：POST

​	注：**PATCH** (更新部分内容): PATCH 方法专门用来对一个资源进行局部修改。

## Delete

### 多个ID

db.Delete(&Todo{}, []int{1, 2, 3})

### where

db.Where("status = ?", "completed").Delete(&Todo{})

### 结构体

db.Delete(&Todo{Title: "学习 Go 语言", Status: "completed"})

## Gin 的路由匹配机制：

**基数树/压缩前缀树 (Radix Tree) ：**合并公共前缀+树状查找



## gin.Context

通过将一次请求所需的所有数据和功能都封装在一个“公文包”里，并确保每个请求都有自己独立的“公文包”，从而让开发者能够以一种清晰、安全、高效的方式来编写和组织中间件与业务逻辑。

```
// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	writermem responseWriter
	Request   *http.Request
	Writer    ResponseWriter

	Params   Params
	handlers HandlersChain
	index    int8
	fullPath string

	engine       *Engine
	params       *Params
	skippedNodes *[]skippedNode

	// This mutex protects Keys map.
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]any

	// Errors is a list of errors attached to all the handlers/middlewares who used this context.
	Errors errorMsgs

	// Accepted defines a list of manually accepted formats for content negotiation.
	Accepted []string

	// queryCache caches the query result from c.Request.URL.Query().
	queryCache url.Values

	// formCache caches c.Request.PostForm, which contains the parsed form data from POST, PATCH,
	// or PUT body parameters.
	formCache url.Values

	// SameSite allows a server to define a cookie attribute making it impossible for
	// the browser to send this cookie along with cross-site requests.
	sameSite http.SameSite
}
```



- **`Request`**: 包含了客户端发来的所有原始请求信息（URL、Header、Body 等）。

- **`Writer`**: 这是一个响应写入器，您最终需要通过它来把响应数据（比如 JSON 或 HTML）写回给客户端。
- **`handlers`**: 记录了这个请求需要依次经过哪些“部门”（即中间件和最终处理函数）的列表。
- **`index`**: 记录了当前公文包已经流转到了哪个部门。`c.Next()` 方法就是通过递增这个 `index` 来将公文包传递给下一个部门的。
- **`Params`**: 如果您的路由是 `/user/:id` 这种形式，`Params` 就用来存放解析出来的路径参数（比如 `id` 的值）。

- **`Keys`**: 这是一个 `map`，它就是公文包里一个**可供所有经手部门读写的“共享便签条”或“夹层”**。
- **`mu`**: 这是一个读写锁，用来保证对 `Keys` 这个 map 的并发读写是安全的。



// 尝试用 c.Get(key) 从 Keys map 中取"user"，如果不存在，就直接 panic

**user := c.MustGet("user").(*models.User)**

## ginx

**`ginx.Dangerous(err)` (处理“危险但不致命”的错误)**

- **作用**: 这个函数通常用来处理那些发生了、需要被记录下来，但**不应该立即中断**整个请求流程的错误。

**`ginx.Bomb(http.StatusForbidden, "forbidden")` (抛出“致命”错误)**

- **作用**: 这个函数用来处理那些**必须立即中断**请求流程的严重错误。

`ginx.NotFound("user not found")` (返回 404 Not Found)

`ginx.BadRequest("invalid parameter: id")` (返回 400 Bad Request)

`ginx.InternalServerError("database connection failed")` (返回 500 Internal Server Error)



## Gorm

“构建器模式” (Builder Pattern) 和 “链式调用” (Method Chaining)

### 设定或指定查询的基础范围

- **`.Model(&User{})`**: 通过传入一个**结构体实例**，让 GORM 根据约定（比如将结构体名转为蛇形复数）**自动推断**出表名（`User` -> `users`）。

- **`.Table("users")`**: **直接通过字符串**来指定表名

```
// 假设你的用户表名叫 my_users，而不是 users
db.Table("my_users").Where("id = ?", 1).First(&user)
```

- `.Raw()` 也是一种设定查询范围的方法，但它更加底层。它直接让你用**原生的 SQL 语句**来代替 GORM 的查询构建器。
-  隐式模型指定 - “看不见的 .Model()”，很多**终结方法**（Finisher Methods）可以**隐式地**设定模型，让你在简单查询中可以省略 `.Model()`。

```
var user User
// GORM 从 &user 这个参数的类型推断出要查询的是 users 表
db.First(&user, 1)

var users []User
// GORM 从 &users 这个切片的元素类型推断出要查询的是 users 表
db.Find(&users)
```



### 套壳子（构建器模式）方法：

1. 筛选与条件 (Filtering and Conditions)

这是最常用的一类，相当于 SQL 中的 `WHERE` 子句。

- `Where(query, args...)`: **最核心的筛选方法**。
  - 示例: `db.Where("name = ? AND age > ?", "jinzhu", 20)`
- `Not(query, args...)`: 与 `Where` 相反，查询**不满足**条件的结果。
  - 示例: `db.Not("name = ?", "jinzhu")`
- `Or(query, args...)`: 添加 `OR` 条件。
  - 示例: `db.Where("name = ?", "jinzhu").Or("name = ?", "lisa")`
- `Find(dest, conds...)`: `Find` 虽然是终结方法，但它也可以直接接收 `Where` 条件作为参数，算是一种便捷写法。
  - 示例: `db.Find(&users, "name <> ?", "jinzhu")`

2. 排序 (Ordering)

相当于 SQL 中的 `ORDER BY` 子句。

- `Order(value)`: 指定排序字段和顺序。
  - 示例: `db.Order("age desc, name asc")` (按年龄降序，再按名字升序)

3. 限制与分页 (Limiting & Pagination)

用于控制返回结果的数量，是实现分页功能的关键。

- `Limit(limit)`: 限制返回的最大记录数，相当于 `LIMIT`。
  - 示例: `db.Limit(10)` (最多返回10条)
- `Offset(offset)`: 跳过指定数量的记录，相当于 `OFFSET`。
  - 示例: `db.Offset(5)` (跳过前5条)
- **组合使用实现分页**:
  - `db.Limit(10).Offset(20)` (获取第3页的数据，每页10条)

4. 选择特定字段 (Selecting Specific Fields)

相当于 SQL 中的 `SELECT column1, column2`。默认情况下，GORM 会查询所有字段 (`SELECT *`)。

- `Select(query, args...)`: 指定要查询的字段，可以提升性能。
  - 示例: `db.Select("name", "age").Find(&users)` (只查询 name 和 age 字段)

5. 关联查询与预加载 (Joins and Preloading)

用于处理表之间的关联关系。

- `Joins(query, args...)`: 执行 SQL `JOIN` 操作，连接多个表。
  - 示例: `db.Joins("JOIN users ON users.id = emails.user_id")`
- `Preload(query, args...)`: **GORM 的特色功能**。用于高效地加载关联数据（“预加载”），以避免 N+1 查询问题。
  - 示例: `db.Preload("Orders").Find(&users)` (查询所有用户，并一次性查出这些用户关联的所有订单)

6. 分组 (Grouping)

相当于 SQL 中的 `GROUP BY` 和 `HAVING`。

- `Group(name)`: 按指定字段分组。
  - 示例: `db.Group("role").Count(&count)` (按角色分组统计人数)
- `Having(query, args...)`: 对分组后的结果进行筛选。
  - 示例: `db.Group("role").Having("COUNT(*) > ?", 10).Find(&results)` (找出成员超过10人的角色)



### “终结方法” (Finisher Methods)：

#### **1. 查询记录 (Retrieving Records)**

- `Find(&users)`: **最常用**。查询所有满足条件的结果，并将它们填充到一个**结构体切片**中。
- `First(&user)`: 获取按主键升序排序的**第一条**记录。如果找不到，会返回 `gorm.ErrRecordNotFound` 错误。
- `Last(&user)`: 获取按主键降序排序的**第一条**记录。
- `Take(&user)`: **不保证顺序**地获取一条记录。
- `Pluck("column", &values)`: 只查询**单列**数据，并将结果填充到一个基本类型的切片中（比如 `[]int64`, `[]string`）。

#### **2. 创建记录 (Creating Records)**

- `Create(&user)`: 将一个或多个结构体对象作为新记录**插入**到数据库中。

#### **3. 更新记录 (Updating Records)**

- `Update("column", "value")`: 更新单个字段。
- `Updates(User{Name: "new_name", Age: 25})`: 使用结构体或 `map` 更新多个字段。
- `Save(&user)`: **智能更新或创建**。如果提供的对象包含了主键，就更新该记录；如果没有主键，就创建一条新记录。

#### **4. 删除记录 (Deleting Records)**

- `Delete(&User{}, 10)`: 删除指定条件的记录。

#### **5. 聚合函数 (Aggregates)**

- `Count(&count)`: 统计满足条件的记录**总数**。

#### **6. 执行原生 SQL (Raw SQL)**

- `Exec("UPDATE users SET name = ? WHERE id = ?", "new_name", 1)`: 执行不返回结果行的原生 SQL（如 UPDATE, DELETE, INSERT）。
- `Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&user)`: 执行一个原生的 `SELECT` 查询，并将结果扫描到结构体中。



# HTTP

## URL:

这个结构体是 Go 标准库用来表示一个经过解析的 URL 的。一个 URL（比如 `https://user:pass@example.com/path?query=1#ref`）会被拆解成以下各个部分存放在这个结构体里：



- `Scheme`: **协议**，比如 `"https"` 或 `"http"`。
- `Opaque`: **非透明数据**。用于一些非标准的 URL 格式，比如 `mailto:user@example.com`，此时 `Opaque` 就是 `"user@example.com"`。
- `User`: **用户信息**。包含用户名和密码，比如 `user:pass`。
- `Host`: **主机(和端口)**。比如 `"example.com"` 或 `"example.com:8080"`。
- `Path`: **路径**。URL 中主机名之后、查询参数之前的部分，比如 `"/path"`。它存储的是解码后的形式（比如 `%2f` 会被存为 `/`）。
- `RawPath`: **原始路径**。存储未解码的、编码形式的路径。只有当它和 `Path` 的编码不同时才会被设置。
- `RawQuery`: **查询字符串**。URL 中 `?` 之后、`#` 之前的部分，但不包含 `?` 本身。比如 `"query=1"`。
- `Fragment`: **片段标识符**。URL 中 `#` 之后的部分，但不包含 `#` 本身，比如 `"ref"`。通常用于定位到页面内的某个部分。
- `RawFragment`: **原始片段**。存储编码形式的片段字符串。

简单来说，`url.URL` 结构体就是把一个复杂的 URL 字符串，拆解成了一份清晰、易于程序读取和修改的“清单”。



## HTTP 请求：

- `Method`: 请求方法，比如 `"GET"`, `"POST"`, `"PUT"` 等。
- `URL`: 一个指向 `url.URL` 结构体的指针，包含了上面解释过的所有 URL 信息。
- `Proto`: HTTP 协议版本，比如 `"HTTP/1.1"`。
- `Header`: 一个 `map`，包含了所有的 HTTP 请求头，比如 `Content-Type`, `User-Agent` 等。
- `Body`: 请求体。对于 `GET` 请求，它通常是空的；对于 `POST` 或 `PUT` 请求，它包含了要发送的数据（比如 JSON 数据）。
- `ContentLength`: 请求体的长度。
- `Host`: 请求的目标主机名。对于服务器收到的请求，它来自于 HTTP 的 `Host` 头。
- `RemoteAddr`: 发起请求的客户端的网络地址（IP:Port）。

**总结：`Request` 结构体是客户端请求的完整封装，`director` 函数通过修改这个结构体的字段，来改变即将被代理转发出去的请求的最终形态。**



## HTTP 头 (HTTP Headers)

### **常见的标准 HTTP 头：**

- `Host`: 我想访问哪个网站。
- `User-Agent`: 我是什么浏览器/客户端。
- `Content-Type`: 我发送的正文内容是什么格式的（比如 `application/json`）。
- `Accept`: 我希望你返回什么格式的内容给我。

### 自定义的 HTTP 头

**API 密钥认证**: `X-API-Key: a-very-long-and-secret-key-string`

- 有些服务不用用户名密码，而是用一个固定的 API Key 来认证。

**多租户识别**: `X-Tenant-ID: company-A`

- 如果后端是一个多租户的集群，可以用这个头来告诉后端，这次请求是属于哪个租户（公司）的，以便进行数据隔离。

**追踪和调试**: `X-Request-Source: nightingale-alerting-engine`

- 可以用一个头来标记这个请求的来源，方便在排查问题时，从海量的日志中快速定位是哪个模块发出的请求。

**传递用户信息**: `X-User-Name: zhangsan`

- 在一些信任代理的场景下，可以直接通过头信息告诉后端当前操作用户的身份。



## Http报错：

http.StatusBadRequest

http.StatusInternalServerError

http.StatusBadGateway

http.StatusForbidden

## http.RoundTripper

`http.RoundTripper` 是 Go 语言 `net/http` 包中一个非常核心的**接口 (interface)**。

- **数据类型**: 它是一个接口，不是一个具体的结构体。任何类型只要实现了 `RoundTrip(*Request) (*Response, error)` 这个方法，就满足了 `http.RoundTripper` 接口的要求。
- **用途**: 它的核心用途是**执行一个单一的 HTTP 事务**。它接收一个 `*http.Request`（请求），负责将其发送到服务器，并返回一个 `*http.Response`（响应）和可能的错误。`http.Transport` 就是 `http.RoundTripper` 接口最常见的一个实现，它负责管理 TCP 连接、连接池、TLS握手等所有底层网络细节。通过自定义 `RoundTripper`，可以实现客户端请求的各种高级功能，如重试、日志记录、模拟测试（mocking）等。

# GMP模型---刘丹冰

## GMP可视化的GMP编程

基本的trace编程：

- 创建trace文件 f, err := os.Create("trace.out")
- 启动trace trace.Start(f)
- 停止trace trace.Stop()

解析 trace.out 文件：go tool trace trace.out

​	G、M、P的信息：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704162010470.png" alt="image-20250704162010470" style="zoom:25%;" />

第二种方式：Windows通过配置环境变量Debug trace查看GMP信息

$env:GODEBUG = "schedtrace=1000"

## 场景模拟

​	场景一：G1在执行过程中如果想创建新的G3，优先将G3放入与G1相同的队列中

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704163955465.png" alt="image-20250704163955465" style="zoom:50%;" />

​	场景二：G1执行完后，通过G0调度G2进入P

 									<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704163926580.png" alt="image-20250704163926580" style="zoom: 50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704164340370.png" alt="image-20250704164340370" style="zoom:50%;" />

场景四：为什么要将本地队列前一半的G与新的G打包放在全局队列中？
	A. 实现负载均衡

​	当一个 P 的工作多到本地队列都放不下时，意味着它“过于繁忙”。而此时系统中可能还有其他 P 是空闲的。把一部分工作（一半的 LRQ + 新的 G）扔到“公共池子”（GRQ）里，就是为了让那些空闲的 P 能找到工作做，从而让整个系统的 CPU 资源得到更充分的利用。

​	B. 保证公平性

​	**为什么需要“随机打散”？**
​	这主要是为了避免**饥饿** 。如果不打散，直接按顺序放进 GRQ，可能会导致某些 Goroutine 长时间得不到执行。随机化处理打破了这种潜在的顺序性，让每个被放入 GRQ 的 Goroutine 都有更公平的被执行机会。



**本地队列的两种访问模式:**

1. **LIFO (后进先出) for 本地 P**: 当 P1 自己执行时，它倾向于从 LRQ 的**尾部**取 G 来执行。因为刚被创建的 G（在队尾）和正在执行的 G（创建者 G1）很可能在数据上具有**缓存亲和性 (Cache Affinity)**。比如 G1 创建 G2 去处理刚生成的数据，这些数据很可能还在 P1 对应的 CPU 核心的 L1/L2 缓存里。立即执行 G2 会非常快。这遵循了“后进先出”的原则。
2. **FIFO (先进先出) for 窃取者**: 当一个空闲的 P2 来 P1 的 LRQ “窃取”任务时，它会从 LRQ 的**头部**偷。为什么？因为头部的 G 是**等待时间最长**的 G。把等待最久的 G 偷走去执行，是一种更公平的策略，可以防止任务饿死。

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704164409645.png" alt="image-20250704164409645" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704164436178.png" alt="image-20250704164436178" style="zoom:50%;" />

**场景6：自旋线程**

- 唤醒：当一个 Goroutine (G) 需要运行时，却没有空闲的 P 可以用时，调度器就会唤醒或创建一个 M，让它进入“自旋”状态去寻找工作。

- 绑定：自旋线程一旦发现有空闲的 P 出现，就会立即尝试与之绑定
- 自旋：自旋期间
  - 1. 检查全局队列 (GRQ)
  - 2. 检查网络轮询器
  - 3. 从别的 P 那里偷

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704170915880.png" alt="image-20250704170915880" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704193411880.png" alt="image-20250704193411880" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250704194153681.png" alt="image-20250704194153681" style="zoom:50%;" />

# GC垃圾回收机制---刘丹冰

## GoV1.3 标记和清除（mark and sweep）

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705095329377.png" alt="image-20250705095329377" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705095448656.png" alt="image-20250705095448656" style="zoom:50%;" />

## GoV1.5 三色标记法

第一步，就是只要是新创建的对象,默认的颜色都是标记为“白色

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705100259975.png" alt="image-20250705100259975" style="zoom:50%;" />

第二步,每次GC回收开始,然后从根节点开始遍历所有对象，把遍历到的对象从白色集合放入“灰色”集合

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705100332521.png" alt="image-20250705100332521" style="zoom:50%;" />

第三步,遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色对象放入黑色集合

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705100409526.png" alt="image-20250705100409526" style="zoom:50%;" />

第四步,重复第三步,直到灰色中无任何对象

![image-20250705100449602](C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705100449602.png)

第五步:回收所有的自色标记表的对象.也就是回收垃圾

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705100536572.png" alt="image-20250705100536572" style="zoom:50%;" />

缺点：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705114602251.png" alt="image-20250705114602251" style="zoom:50%;" />

### 强弱三色不变式：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705115113456.png" alt="image-20250705115113456" style="zoom:33%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250705115126879.png" alt="image-20250705115126879" style="zoom:50%;" />

### 屏障机制

- 插入屏障：对象被引用时，触发的机制（插入屏障仅在堆中使用，不在栈中使用）

  操作：将黑色引用的白色直接变为灰色（满足强三色不变式）

- 删除屏障：对象被删除时，触发的机制（保护灰色对象到白色对象的路径不会断）

  操作：被删除的对象，如果自身为灰色或者白色，那么被标记为灰色（满足弱三色不变式）

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711135320498.png" alt="image-20250711135320498" style="zoom:33%;" />

## Go V1.8的三色标记法+混合写屏障机制

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711135537095.png" alt="image-20250711135537095" style="zoom:33%;" /> 

**场景一**：对象被一个堆对象删除引用，成为栈对象的下游

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140159697.png" alt="image-20250711140159697" style="zoom:50%;" />

![image-20250711140209815](C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140209815.png)

**场景二**：对象被一个栈对象删除引用，成为另一个栈对象的下游

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140521392.png" alt="image-20250711140521392" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140539672.png" alt="image-20250711140539672" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140555125.png" alt="image-20250711140555125" style="zoom:50%;" />

**场景三**：对象被一个堆对象删除引用，成为另一个堆对象的下游

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711140948252.png" alt="image-20250711140948252" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711141009383.png" alt="image-20250711141009383" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711141028356.png" alt="image-20250711141028356" style="zoom:50%;" />

**场景四**：对象从一个栈对象删除引用，成为另一个堆对象的下游

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711141332570.png" alt="image-20250711141332570" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711141355295.png" alt="image-20250711141355295" style="zoom:50%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711141410988.png" alt="image-20250711141410988" style="zoom:50%;" />

## 总结

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250711143033099.png" alt="image-20250711143033099" style="zoom:50%;" />

- **理解GOGC和GOMEMLIMIT**：Go提供了GOGC环境变量来调整GC的触发频率。简单来说，GOGC=100（默认值）表示当新分配的内存达到上次GC后存活内存的100%时，触发下一次GC。
  - **对于内存敏感、CPU资源充足的应用**，可以调低GOGC，让GC更频繁，及时回收内存。[[5](https://www.google.com/url?sa=E&q=https%3A%2F%2Fvertexaisearch.cloud.google.com%2Fgrounding-api-redirect%2FAUZIYQH-fXJwxi5TiFUssjKGjbsKataaxv0lZyDT3AhMyXFUpoacwOHv0DlkNkHu7LcqF6S1mJZ3C9q_F8uNZ7At8Zrs8mUC1kzOTZ5NJ0ZxhFhjO0TE-WqIFiZ5AxyVVLXt2YFXBtjUOQXpWo-SSESIjArLzimqdhUVhTIqWb3YqSsG)]
  - **对于延迟敏感、内存资源充足的应用**，可以调高GOGC，减少GC频率，降低GC带来的CPU开销和STW风险。Go 1.19引入的GOMEMLIMIT则提供了一个更直观的内存限制“软顶”，当程序内存使用接近这个值时，GC会变得更加积极，这对于防止在容器等内存受限环境中被OOM Kill非常有帮助。
- **分析性能问题**：当应用出现性能抖动或延迟尖刺时，有经验的工程师会使用go tool trace或pprof来分析。如果在火焰图中看到大量时间消耗在runtime.gcBgMarkWorker等GC相关的函数上，结合对三色标记和混合写屏障的理解，就能推断出可能是堆内存过大或对象分配过于频繁，导致GC标记阶段压力巨大。
- **编写GC友好的代码**：
  - **场景一（堆对象删除引用，成为栈对象下游）**：这种情况通常是安全的。因为栈上的对象会被重新扫描，即使它引用了一个之前被认为是“白色”的对象，这个白色对象也会被重新标记为灰色，从而避免被错误回收。
  - **场景二（堆对象删除引用，成为另一个堆对象下游）**：这就是混合写屏障要解决的核心问题。如果没有屏障机制，一个黑色的堆对象直接引用一个白色的堆对象，就有可能导致这个白色对象在本次GC中被漏掉并回收。混合写屏障通过保护被删除引用的对象（如果它是灰色或白色，则标记为灰色）或新建立引用的对象（将黑色引用的白色对象变为灰色），来保证对象的存活性。了解这一点，可以帮助您理解为什么某些操作（特别是在并发场景下修改共享数据结构）的开销会比预想的要高，因为它们可能触发了写屏障。

**结论**

总而言之，学习Go的GC机制，不是为了去重新实现它，而是为了能够与它“和谐共处”。这就像开赛车，您不需要自己能造发动机，但您必须深刻理解发动机的特性，才能在赛道上发挥出它的极限性能。对于Go工程师来说，GC知识是区分“能用”和“卓越”的重要分水岭，它能让您在面对复杂的性能和内存问题时，做到心中有数、游刃有余。

# Nightingale --- 夜莺

## 跑通：

- go run ./cmd/center/main.go init-db

指令解析：

**./cmd/center/main.go**：

By Gemini

1. main.go 是一个**启动器 (Launcher)**。
2. flag.Parse() 负责**分拣**出命令行的标志和参数。
3. center.Initialize() 是一个**调度器 (Dispatcher)**。
4. 它首先检查有没有像 init-db 这样的**一次性任务**。如果有，就去执行这个任务，然后**立刻退出程序 os.Exit(0)**。
5. 如果**没有**发现任何一次性任务的子命令，它才会继续执行**启动 Web 服务器**的常规流程。

- 使用 username : root pwd : root.2020 登录

界面成功啦~：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250706231651626.png" alt="image-20250706231651626" style="zoom:50%;" />

- 实现自己采集自己的主机数据

下载官方推荐采集器：categraf

我当时的最新版本：

wget https://github.com/flashcatcloud/categraf/releases/download/v0.4.13/categraf-v0.4.13-linux-amd64.tar.gz

tar -zxvf

启动：

sudo ./categraf

- 实现时序数据库**VictoriaMetrics**：

我当时的最新版本：

wget https://github.com/VictoriaMetrics/VictoriaMetrics/releases/download/v1.120.0/victoria-metrics-linux-amd64-v1.120.0.tar.gz

tar -zxvf

启动：

./victoria-metrics-prod

- 改categraf的发送端口

[[writers]] 

url = "http://127.0.0.1:17000/prometheus/v1/write"

改为：

[[writers]] 

url = "http://127.0.0.1:8428/api/v1/write"

- 数据源管理：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250707002758272.png" alt="image-20250707002758272" style="zoom:50%;" />

- 成功！

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250707003009255.png" alt="image-20250707003009255" style="zoom:50%;" />

重新启动：

ps aux | grep  victoria-metrics-prod

ps aux | grep  nightingale

sudo lsof -i :20090

ps aux | grep  categraf

./victoria-metrics-prod

go run ./cmd/center/main.go

sudo ./categraf

查找命令：

grep -r "auth/login"



## 数据库外封装一层功能接口：

例如：**UserGetByUsername是UserGet的封装：**

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250707163917845.png" alt="image-20250707163917845" style="zoom:33%;" />

优点：

- **可读性和语义化**：能够被立即理解，表达的是**意图**，而不是**实现细节**

- **可维护性**：将 users 表中的 password 列重命名为 pwhash，只需修改一个UserGetByUsername，不需要修改每个UserGet
- **解耦与抽象**：不应该关心用的是 Postgres、MySQL，还是像 MongoDB 这样的 NoSQL 数据库
- **可测试性**：轻松地**模拟**一个假的 GetUserByUsername 函数，让它只返回一个硬编码的用户，在完全隔离的环境中测试其他逻辑

# gRPC

单体架构： 

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716100932167.png" alt="image-20250716100932167" style="zoom:50%;" />

微服务架构：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716101044937.png" alt="image-20250716101044937" style="zoom:50%;" />

gRPC：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716101407991.png" alt="image-20250716101407991" style="zoom:50%;" />

Protobuf （Protocol Buffers）：

是由 Google 开发的一种**与语言无关、与平台无关、可扩展的数据序列化机制**。可高效地序列化结构化数据，常用于网络通信和数据存储等场景。

程序无需关心对方使用的是什么编程语言或操作系统，只需通过 Protobuf 定义好的格式进行数据交换。高性能、低延迟。

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716102000089.png" alt="image-20250716102000089" style="zoom:33%;" />

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716110532278.png" alt="image-20250716110532278" style="zoom:50%;" />

SSL/TLS协议：

<img src="C:\Users\X1Kun\AppData\Roaming\Typora\typora-user-images\image-20250716130549691.png" alt="image-20250716130549691" style="zoom:50%;" />

1.保密 (Confidentiality)

​	TLS 会用一套复杂的数学方法把内容变成密文。

2.完整性 (Integrity)

​	TLS 会根据原文内容算出一个独一无二的“数字指纹”，接收者收到信息后，会用同样的方法再算一遍“数字指纹”。如果算出的指纹和信息上附带的指纹对不上，就说明信息在传输过程中被改动

3.认证 (Authentication) 

 发送者会先出示一个由“权威机构”（CA）颁发的“身份证”（SSL/TLS 证书），上面有名字和公钥。接受者可以去CA验证。确认无误后，才开始和他通话。这就保证了连接的网站就是真的银行网站，而不是一个看起来很像的钓鱼网站。
