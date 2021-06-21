package main
//
//import "golang.org/x/tour/pic"
//
//func Pic(dx, dy int) [][]uint8 {
//	sY := make([][]uint8, dy)
//	for y := 0; y < dy; y++{
//		sZ := make([]uint8, dx)
//		sY[y] = sZ
//		for x := 0; x < dx; x++{
//			sZ[x] = uint8(x * y)
//		}
//	}
//	return sY
//}
//
//func main() {
//	pic.Show(Pic)
//}