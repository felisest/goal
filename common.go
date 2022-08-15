package goal

func grow_container(cur_cap int) int {

	mnum := 512
	i := 0
	for ; (cur_cap > mnum<<i) && (i < 4); i++ {
	}

	return cur_cap + cur_cap/(1<<i)
}