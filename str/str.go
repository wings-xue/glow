package str

import "sort"

type strSuffix []string

func (suffix strSuffix) Len() int           { return len(suffix) }
func (suffix strSuffix) Swap(i, j int)      { suffix[i], suffix[j] = suffix[j], suffix[i] }
func (suffix strSuffix) Less(i, j int) bool { return len([]rune(suffix[i])) < len([]rune(suffix[j])) }

func Suffix(s string) []string {
	out := make([]string, 0)
	if s == "" {
		return out
	}
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if string(rs[:i]) != "" {
			out = append(out, string(rs[:i]))
		}

	}
	out = append(out, s)
	sort.Sort(strSuffix(out))
	return out
}

type strPrefix []string

func (suffix strPrefix) Len() int           { return len(suffix) }
func (suffix strPrefix) Swap(i, j int)      { suffix[i], suffix[j] = suffix[j], suffix[i] }
func (suffix strPrefix) Less(i, j int) bool { return len([]rune(suffix[i])) > len([]rune(suffix[j])) }

func Prefix(s string) []string {
	out := make([]string, 0)
	if s == "" {
		return out
	}
	rs := []rune(s)

	for i := 1; i < len(rs); i++ {
		if string(rs[i:]) != "" {
			out = append(out, string(rs[i:]))
		}

	}
	sort.Sort(strPrefix(out))
	return out
}
