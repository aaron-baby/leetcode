package hash_table

import (
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		cnt := make(map[uint8]int)
		for i := 0; i < len(str); i++ {
			cnt[str[i]] += 1
		}
		var key string
		for i := 0; i < 26; i++ {
			if n, ok := cnt['a'+uint8(i)]; ok {
				key += string('a'+uint8(i)) + strconv.Itoa(n)
			}
		}
		groups[key] = append(groups[key], str)
	}
	var res [][]string
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}
