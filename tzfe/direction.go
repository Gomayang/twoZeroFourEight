package tzfe

type direction [2]int

func leftDir() direction {
	return direction{-1, 0}
}

func rightDir() direction {
	return direction{1, 0}
}

func upDir() direction {
	return direction{0, 1}
}

func downDir() direction {
	return direction{0, -1}
}
