package leetcode

func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	// let dp[i][j] is a longest common subsequence at i, j.
	dp := make([][]int, m+1)
	for r := range dp {
		dp[r] = make([]int, n+1)
	}

	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(s)
	//fmt.Println(t)
	//PrintMatrix(&dp)

	return dp[m][n] == len(s)
}

/*
func PrintMatrix(m *[][]int) {
	r, c := len((*m)), len((*m)[0])
	dp_str := ""
	for i := 0; i < r; i++ {
		s := ""
		for j := 0; j < c; j++ {
			s += fmt.Sprintf("%d", (*m)[i][j])
		}
		dp_str += fmt.Sprintf("%s\n", s)
	}
	fmt.Println(dp_str)

}



func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
*/
