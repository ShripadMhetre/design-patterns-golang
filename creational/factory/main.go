package factory

func Main() {
	windows, _ := getOS("windows")
	linux, _ := getOS("linux")
	mac, _ := getOS("mac")

	windows.specs()
	linux.specs()
	mac.specs()
}
