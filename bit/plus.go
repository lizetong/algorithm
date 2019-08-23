package bit

/*
描述
给出两个整数 a 和 b , 求他们的和。

样例 1:
输入:  a = 1, b = 2
输出: 3
样例解释: 返回a+b的结果.

样例 2:
输入:  a = -1, b = 1
输出: 0
样例解释: 返回a+b的结果.
*/
/*
// 主要利用异或运算来完成
// 异或运算有一个别名叫做：不进位加法
// 那么a ^ b就是a和b相加之后，该进位的地方不进位的结果
// 然后下面考虑哪些地方要进位，自然是a和b里都是1的地方
// a & b就是a和b里都是1的那些位置，a & b << 1 就是进位
// 之后的结果。所以：a + b = (a ^ b) + (a & b << 1)
// 令a' = a ^ b, b' = (a & b) << 1
// 可以知道，这个过程是在模拟加法的运算过程，进位不可能
// 一直持续，所以b最终会变为0。因此重复做上述操作就可以
// 求得a + b的值。
*/
func Plus(a, b int) int {

	return 0
}
