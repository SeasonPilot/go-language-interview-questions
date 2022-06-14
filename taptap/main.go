package main

// 给定一个正整数 x, 可以对x做如下变换:
// x 为奇数时，x => x-1 or x => x+1
// x 为偶数时，x => x/2
// 则至少对x做多少次变换可以变换为1

// GetSteps BFS 解法
func GetSteps(x int) int {
	if x <= 1 {
		return 0
	}

	q := []int{x}
	ret := 0

	for len(q) > 0 {
		var tmp []int
		for i := 0; i < len(q); i++ {
			if q[i] == 1 {
				return ret
			}
			if q[i]&1 == 1 { //奇数
				n1 := q[i] - 1
				tmp = append(tmp, n1)

				n2 := q[i] + 1
				tmp = append(tmp, n2)
			} else {
				n3 := q[i] >> 1
				tmp = append(tmp, n3)
			}
		}
		q = tmp
		ret++
	}
	return ret
}

// DP  没做出来
func GetSteps1(x int) int {
	if x <= 1 {
		return 0
	}

	// 状态  操作次数、奇偶性
	dp := [][2]int{}
	// 选择  x-1  x+1  x/2
	// 状态转移
	var i int
	for x != 1 {
		if x&1 == 1 { // 奇数
			dp[i][1] = dp[i-1][1]

		} else {
			dp[i][0] = (dp[1] - 1, dp[1] + 1)
		}
		i++
	}

	// base case
	dp[0][1] = 0

}
