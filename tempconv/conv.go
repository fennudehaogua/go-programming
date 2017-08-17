package tempconv
func CTOF(c Celsius) Fahrenheit{
	return Fahrenheit(c*9/5+32)
}
func FTOC(f Fahrenheit) Celsius{
	return Celsius((f-32)*5/9)
}