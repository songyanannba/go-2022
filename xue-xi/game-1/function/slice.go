package function

func CheckInNumberSlice(ele uint64, dataList []uint64) bool {

	for _, v := range dataList {
		if v == ele {
			return false
		}
	}
	return true
}

func DelEleInSlice(ele uint64, dataList []uint64) []uint64 {

	var newList []uint64
	for _, v := range dataList {
		if v != ele {
			newList = append(newList, v)
		}
	}

	return newList
}
