package manage

func (stuMgs StudentMgs) isExist(id int) bool {
	_, exists := stuMgs.Students[id]
	return exists
}
