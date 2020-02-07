package router

var defaultDriver = &hashDriver{}


// Goto redirects the user to the given path using the default route driver.
func Goto(path string) {
	defaultDriver.gotoPath(path)
}