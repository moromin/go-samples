package sample

func GetName(id int64) string {
	switch id {
	case 1:
		return "one"
	case 2:
		return "second"
	default:
		return "No name"
	}
}
