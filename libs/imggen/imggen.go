package imggen

func Push(list []int, n int) []int {
	s := list[1:60]
	r := append(s, n)
	return r
}

func Gen(list []int, n float64, max float64) []int {
	n = (n / max) * 42
	res := int(n)
	data := Push(list, res)

	//dc := gg.NewContext(200, 50)
	//for i, v := range data {
	//dc.DrawRectangle(float64(10*i), 50-float64(v), 10, float64(v))
	//dc.SetRGBA(216, 144, 21, 0.6)
	//dc.Fill()
	//}
	//dc.SavePNG(name)

	return data
}
