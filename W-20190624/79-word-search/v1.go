package leetcode

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	wi := 0
	br := len(board)
	if len(board) == 0 {
		return false
	}
	bc := len(board[0])

	type p struct {
		i, j       int
		directions int
	}
	path := make([]*p, len(word))

	used := func(i, j int) bool {
		for _, pp := range path {
			if pp != nil {
				if i == pp.i && j == pp.j {
					return true
				}
			} else {
				return false
			}
		}
		return false
	}

	matchUp := func(i, j *int, ch byte) bool {
		if *i == 0 {
			return false
		}
		if !used((*i-1), *j) && board[*i-1][*j] == ch {
			wi++
			*i = *i - 1
			return true
		}
		return false

	}

	matchDown := func(i, j *int, ch byte) bool {
		if *i >= br-1 {
			return false
		}
		if !used((*i+1), *j) && board[*i+1][*j] == ch {
			wi++
			*i = *i + 1
			return true
		}
		return false
	}

	matchLeft := func(i, j *int, ch byte) bool {
		if *j == 0 {
			return false
		}
		if !used(*i, *j-1) && board[*i][*j-1] == ch {
			wi++
			*j = *j - 1
			return true
		}
		return false
	}

	matchRight := func(i, j *int, ch byte) bool {
		if *j >= bc-1 {
			return false
		}
		if !used(*i, *j+1) && board[*i][*j+1] == ch {
			*j = *j + 1
			wi++
			return true
		}
		return false
	}

	matchDirection := func(startDirection *int, i, j *int) bool {
		defer func() {
			*startDirection = *startDirection * 2
		}()
		switch *startDirection {
		case 1:
			return matchLeft(i, j, word[wi+1])
		case 2:
			return matchRight(i, j, word[wi+1])
		case 4:
			return matchUp(i, j, word[wi+1])
		case 8:
			return matchDown(i, j, word[wi+1])
		}
		return false
	}

	for i := 0; i < br; i++ {
		// 列
		for j := 0; j < bc; j++ {
			ti, tj := i, j
			if board[ti][tj] == word[wi] {
				if wi == len(word)-1 {
					return true
				}
				path[wi] = &p{
					i:          ti,
					j:          tj,
					directions: 1,
				}
				for {
					for path[wi].directions <= 8 {
						if matchDirection(&path[wi].directions, &ti, &tj) {
							if wi == len(word)-1 {
								return true
							}
							path[wi] = &p{
								i:          ti,
								j:          tj,
								directions: 1,
							}
						}
					}
					if wi == 0 {
						break
					}
					// match失败，回退一格
					path[wi] = nil
					wi--
					ti, tj = path[wi].i, path[wi].j
				}
			}
		}
	}
	return false
}
