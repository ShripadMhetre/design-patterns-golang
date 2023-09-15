package singleton

func Main() {
	s1 := GetInstance()
	s2 := GetInstance()
	println("s1.GetName()", s1.GetName())
	println("s2.GetName()", s2.GetName())
	println("s1 == s2", s1 == s2)
}
