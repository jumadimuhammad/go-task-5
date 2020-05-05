package sepeda

import (
	"fmt"

	"github.com/jumadimuhammad/task6/control"
)

type Sepeda struct {
	Ban, Gigi, Jarak float32
	Warna            string
}

func (m *Sepeda) WaktuTempuh(jarak float32) float32 {
	return jarak * 2.5
}

func (m *Sepeda) Cepat() float32 {
	return m.Jarak * 2.5 * 2
}

func (m *Sepeda) Lambat() float32 {
	return m.Jarak * 2.5 / 2
}

func CheckCepat(m control.Maju) float32 {
	return m.Cepat()
}
func CheckLambat(m control.Maju) float32 {
	return m.Lambat()
}

// func CheckCanRun(m control.Maju) (float32, float32) {
// 	return m.Cepat(), m.Lambat()
// }

func FuncSepeda() {

	// var data = [...]Sepeda{
	// 	{2, 3, 30, "Putih"},
	// 	{2, 4, 20, "Hitam"},
	// }

	var data = Sepeda{2, 3, 20, "Merah"}
	fmt.Println("Jumlah ban       :", data.Ban)
	fmt.Println("Jumlah gigi      :", data.Gigi)
	fmt.Println("Jumlah gigi      :", data.Warna)
	fmt.Println("Kecepatan awal   :", data.WaktuTempuh(data.Jarak), "Menit")
	fmt.Print("Kecepatan cepat  : ")
	fmt.Println(CheckCepat(&Sepeda{Jarak:data.Jarak}), "Menit")
	fmt.Print("Kecepatan lambat : ")
	fmt.Println(CheckLambat(&Sepeda{Jarak:data.Jarak}), "Menit")

	// for index, value := range data {
	// 	fmt.Println("Sepeda ke        :", index+1)
	// 	fmt.Println("Jumlah ban       :", value.Ban)
	// 	fmt.Println("Jumlah gigi      :", value.Gigi)
	// 	fmt.Println("Kecepatan awal   :", value.WaktuTempuh(value.Jarak))
	// 	fmt.Print("Kecepatan cepat  : ")
	// 	fmt.Println(CheckCanRun(&Sepeda{Jarak: value.Jarak}))

	// 	fmt.Println("---------------------------------------")
	// }
}
