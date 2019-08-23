package bit

// 获取0-n之间的所有偶数
func Even(n int) []int {
	var array []int
	for i := 0; i < n; i++ {
		if i&1 == 0 {
			array = append(array, i)
		}
	}
	return array
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
func Swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func Shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func Transformation(a int) int {
	return ^a + 1 // Go语言取反方式
}

//获取最大值
func GetMaxNum(a int, b int) int {
	c := a - b
	x := (c>>31)&1 ^ 1
	y := x ^ 1
	return a*x + b*y
}
