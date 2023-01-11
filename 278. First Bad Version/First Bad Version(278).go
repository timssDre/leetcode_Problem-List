package main

func main() {
	firstBadVersion(5)
}

func firstBadVersion(n int) int {
	for {
		if isBadVersion(n) == false {
			n++
			break
		}
		n--
	}
	return n
}

func isBadVersion(i int) bool {
	if i > 3 {
		return true
	}
	return false
}
