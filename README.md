# About
本项目来源：[ugo](https://github.com/wa-lang/ugo) 
- 代码结构简单，可用于初步了解编译器的实现。实现原理见：[ugo-compiler-book](https://github.com/wa-lang/ugo-compiler-book) 
- 本项目对源代码进行了初步注释，因实现简洁，未作详细注释。
- 如读者有 JavaScript 基础，可参考：[acorn](https://github.com/acornjs/acorn) 的实现

两者思路差异：
- ugo 一次读取所有 token，然后进行语法解析，形成 AST
- acorn 边读取 token ，边分析语法，再形成 AST
- 两者形成的 AST 信息量也有不同，acorn 形成的 AST 信息更完备
  - acorn AST 包括语句的起始结束位置，包括 token 的起始结束位置

# µGo 编程语言

µGo 是 Go 语言的真子集(不含标准库部分), 可以直接作为Go代码编译执行.

- 安装 ugo: `go get github.com/wa-lang/ugo`.
- 实现原理: https://github.com/wa-lang/ugo-compiler-book

## 例子

例子 ([_example/hello.ugo](_example/hello.ugo)):

```go
package main

func main() {
	for n := 2; n <= 30; n = n + 1 {
		var isPrime int = 1
		for i := 2; i*i <= n; i = i + 1 {
			if x := n % i; x == 0 {
				isPrime = 0
			}
		}
		if isPrime != 0 {
			println(n)
		}
	}
}
```

运行:

```
$ ugo run _examples/hello.ugo 
2
3
5
7
11
13
17
19
23
29
```

## 版权

个人学习目的可自由使用.
