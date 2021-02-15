
### 最大公约数
```golang
func gcd(a, b int) int {
    for a != 0 {
        a, b = b % a, a
    }
    return b
}
```