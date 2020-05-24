package leetcode

// tip: 爬楼梯问题升级版
// 4种情况:
// 1. s[i] == '0' && (s[i-1] == '1' || s[i-1] == '2') ==> dp[i] = dp[i-2] ==> 否则无法编码 直接return
//		==> 只能是 s[i-1] + s[i], 不增加编码情况
//
// 		例子: 1220 ==> (1) ==> (1, 2), (12) ==> (1, 2, 2), (12, 2), (1, 22) ==> (1, 2, 20), (12, 20)
//
// 2. s[i] != '0' && s[i-1] == '1' ==> dp[i] = dp[i-1] + dp[i-2]
// 		==> dp[i-1]为s[i]和s[2]分开编码, 相当于把s[i]添加到 dp[i]的所有情况中
//		==> dp[i-2]为s[i]+s[i-2]组合成一个字符合并到dp[i-2]的所有情况上，并没有增加情况，直接使用dp[i+-12]
//
// 		例子: 1218 ==> (1) ==> (1, 2), (12) ==> (1, 2, 1), (12, 1), (1, 21) ==> (1, 2, 1, 8), (12, 1, 8), (1, 21, 1, 8), (1, 2, 18), (12, 18)
//
// 3. s[i] != '0' && s[i-1] == '2' && s[i] <= '6' ==> dp[i] = dp[i-1] + dp[i-2]
//		==> 同 2
//
// 4. 其他 ==> dp[i] = dp[i-1] 分开编码
// 		例子: 1228 ==> (1) ==> (1, 2), (12) ==> (1, 2, 2), (12, 2), (1, 22) ==> (1, 2, 2, 8), (12, 2, 8), (1, 22, 8)

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	result, dp1, dp2 := 1, 1, 1 // dp1对应dp[i-1] dp2对应dp[i-2]
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0':
			if s[i-1] != '1' && s[i-1] != '2' { // 无法编码的情况
				return 0
			}
			result = dp2
		case s[i] != '0' && s[i-1] == '1':
			result = dp1 + dp2
		case s[i] != '0' && s[i-1] == '2' && s[i] <= '6':
			result = dp1 + dp2
		default:
			result = dp1
		}

		dp2, dp1 = dp1, result
	}

	return result
}