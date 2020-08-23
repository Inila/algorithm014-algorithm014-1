package main

import (
	"sort"
	"strings"
)

/*
	### 242. 有效的字母异位词
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, b := range s {
		if _, ok := m[b]; !ok {
			m[b] = 0
		}
		m[b]++
	}

	for _, b := range t {
		if _, ok := m[b]; !ok {
			return false
		}
		m[b]--
		if m[b] == 0 {
			delete(m, b)
		}
	}
	return len(m) == 0
}

/*
	### 49. 字母异位词分组
*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	ret := make([][]string, 0)

	for _, str := range strs {
		// 排序处理
		fields := strings.Split(str, "")
		sort.Strings(fields)
		ss := strings.Join(fields, "")

		if _, ok := m[ss]; !ok {
			m[ss] = make([]string, 0)
		}
		m[ss] = append(m[ss], str)
	}

	for _, values := range m {
		ret = append(ret, values)
	}
	return ret
}
