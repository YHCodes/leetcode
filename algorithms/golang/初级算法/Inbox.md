
### 最大公约数
```golang
func gcd(a, b int) int {
    for a != 0 {
        a, b = b % a, a
    }
    return b
}
```

## Golang 语法点

### 1. string 和 int 类型的相互转换
```golang

import "strconv"

// string 转成 int
int, err := strconv.Atoi(string)

// string 转成 int64
int64, err := strconv.ParseInt(string, 10, 64)

// int 转成 string:
string := strconv.Itoa(int)

// int64 转成 string:
string := strconv.FormatInt(int64, 10)
```

### 2. 对字符串的字符进行排序
```golang
package main

import "fmt"
import "sort"

func main() {
	s := "bca"
    // step1: 把字符串转成 byte 数组
	sb := []byte(s)
    // step2: 排序
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})
    // step3: 数组转成string
	fmt.Println(string(sb))
}
```