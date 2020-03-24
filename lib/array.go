package lib

//InArrayS 如果 s 在 items 中,返回 true；否则，返回 false。
func InArrayS(items []string, s string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}
