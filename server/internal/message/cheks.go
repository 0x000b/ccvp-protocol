package message

import "strconv"

func CheckTOR(tor string) int {
	value, err := strconv.Atoi(tor)

	if err != nil {
		return -1
	}

	if value == 1 {
		return 1
	} else if value == 0 {
		return 0
	} else {
		return -1
	}
}

func CheckTOD(tod string) int {
	value, _ := strconv.Atoi(tod)

	if value == 0 {
		return 0
	} else if value == 1 {
		return 1
	} else {
		return -1
	}

}

func CheckLength(len string, tod int) bool {
	length, _ := strconv.Atoi(len)

	if length == 11 && tod == 0 {
		return true
	} else if length == 14 && tod == 1 {
		return true
	} else {
		return false
	}
}
