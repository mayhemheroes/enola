package fuzz

func Fuzz(usr []byte) int{
	findAndShowResult(string(usr),"")
	return 0
}