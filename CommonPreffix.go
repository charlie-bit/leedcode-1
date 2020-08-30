package Leedcode_1

func LongestCommonPrefix(strs []string) string {
	var preffix string
	for len(strs) > 0 {
		var resp string
		var flag = true
		if len(strs[0]) <= 1 {
			resp = strs[0]
		}else {
			resp = strs[0][:1]
		}
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == 1 {
				if resp != strs[i]{
					flag = false
				}
			}else if len(strs[i]) > 1{
				if resp != strs[i][:1]{
					flag = false
				}
			}else {
				return preffix
			}
		}
		if !flag {
			return preffix
		}
		preffix += resp //定值
		for i, i2 := range strs {
			if len(strs[i]) > 1 {
				strs[i] = i2[1:]
			}else {
				strs[i] = ""
			}
		}
	}
	return preffix
}
