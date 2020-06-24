
# go_rule_executor
一个基于go语言实现的规则执行引擎

```go
var ifStmt = `
  aa = "mike" + " wen"
   d += 1.3 * 2.534 +1 
   d = true || false
   if (d) {
     f = 12+ 123.1
   }else {
    f = 1 + 2
	}

   ss = fmt.sprintf("hahah")
`

var methodcall = `
  ss = io.sprintf("hello:%s","mikewen")
 tt = time.now()
 i = 12
 i = "hello world"
 for (i = 2.1; i < 20; i=i+1) {
	tt = time.now()
    io.print("i:",i)
 }
 io.print(tt)
 io.print(ss)
 io.print(i)
`
```
