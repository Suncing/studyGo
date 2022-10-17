package array

//1.数组解法
func destCity(paths [][]string) string {
	//ar arr = [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	if len(paths) == 0 {
		return ""
	}
	if len(paths) == 1 {
		return paths[0][1]
	}
	var arr []string
	for i := 0; i < len(paths); i++ {
		arr = append(arr, paths[i][0])
	}
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == paths[i][1] {
				break
			}
			if j == len(paths)-1 {
				return paths[i][1]
			}
		}
	}
	return ""
}

//2.hash解法
func destCity2(paths [][]string) string {
	if len(paths) == 0 {
		return ""
	}
	if len(paths) == 1 {
		return paths[0][1]
	}
	seen := map[string]struct{}{}
	for i := 0; i < len(paths); i++ {
		seen[paths[i][0]] = struct{}{}
	}

	for i := 0; i < len(paths); i++ {
		_, ok := seen[paths[i][1]]
		if !ok {
			return paths[i][1]
		}
	}
	return ""
}
