package gen

type DTO struct {
	Teacher
	Subject string
	Hour    int16
}
type Teacher struct {
	Name string
}

var sub = []string{
	"Математика",
}

var pedagogs = []string{
	"Шлюха1",
}

var group_1m_1 = struct {
	Geo, Bio, Him, Fizik, Inf, Math, IstWorl, Lit, Rus, Obj, Angl, Fiz, MuzLit, WorldMuzLit, Solf, Elrad, Ist DTO
}{
	Math: DTO{
		Teacher: Teacher{
			Name: pedagogs[0],
		},
		Subject: sub[0],
		Hour:    10,
	},
}
