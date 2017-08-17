package lengthconv
func ITOM(i Incun) Meter{
	return Meter(i*0.0254)
}
func MTOI(m Meter) Incun{
	return Incun(m*39.3700787)
}