package motor

import (
	"fmt"

	"github.com/jumadimuhammad/task6/control"
	// "github.com/jumadimuhammad/task6/sepeda"
)

type Motor struct {
	Ban, Gigi, Jarak float32
	Warna            string
	// sepeda.Sepeda
}

func (m *Motor) WaktuTempuh(jarak float32) float32 {
	return jarak * 2.5 * 2
}

func (m *Motor) Cepat() float32 {
	return m.Jarak * 2.5 * 2 * 2
}

func (m *Motor) Lambat() float32 {
	return m.Jarak * 2.5 * 2 / 2
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

func FuncMotor() {
	data := Motor{2, 4, 20, "Pink"}

	fmt.Println("Jumlah ban       :", data.Ban)
	fmt.Println("Jumlah gigi      :", data.Gigi)
	fmt.Println("Jumlah gigi      :", data.Warna)
	fmt.Println("Kecepatan awal   :", data.WaktuTempuh(data.Jarak), "Menit")
	// fmt.Print("Kecepatan cepat  : ")
	// fmt.Println(CheckCanRun(&Motor{Jarak: data.Jarak}))
	fmt.Print("Kecepatan cepat  : ")
	fmt.Println(CheckCepat(&Motor{Jarak:data.Jarak}), "Menit")
	fmt.Print("Kecepatan lambat : ")
	fmt.Println(CheckLambat(&Motor{Jarak:data.Jarak}), "Menit")
}
