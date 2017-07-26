package imggen

import (
    "github.com/fogleman/gg"
)

func Push(list []int, n int) []int {
    s := list[1:20]
    r := append(s, n)
    return r
}

func Gen(list []int, n float64, max float64, name string) []int {
    n = (n/max)*50
    res := int(n)
    data := Push(list, res)

    dc := gg.NewContext(200, 50)
    for i,v := range(data){
        dc.DrawRectangle(float64(10*i), 50-float64(v), 10, float64(v))
        dc.SetRGBA(150, 0, 60, 0.8)
        dc.Fill()
    }
    dc.SavePNG(name)

    return data
}
