func makePaths(x, y int) [][][2]int {
	paths := make([][][2]int, 1)
	paths[0] = make([][2]int, 1)
	paths[0][0] = [2]int{1, 1}

	for i := 1; i < x+y-1; i++ {
		next_paths := make([][][2]int, 0)
		for _, v := range paths {
			last_coordinate := v[i-1]
			if last_coordinate[0] < x {
				var path_to_right [][2]int
				path_to_right = append(path_to_right, v...)
				path_to_right = append(path_to_right, [2]int{last_coordinate[0] + 1, last_coordinate[1]})
				next_paths = append(next_paths, path_to_right)
			}

			if last_coordinate[1] < y {
				var path_to_bottom [][2]int
				path_to_bottom = append(path_to_bottom, v...)
				path_to_bottom = append(path_to_bottom, [2]int{last_coordinate[0], last_coordinate[1] + 1})
				next_paths = append(next_paths, path_to_bottom)
			}
		}
		paths = next_paths
		// p(next_paths)
	}
	return paths
}
