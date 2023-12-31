package gen

type DTO struct {
	Teacher
	Subject string
	Hour    int16
	Quality float64
}
type Teacher struct{
	Name string
}

var Subjects = []string{
	"Математика", "Русский", "Информатика", "Физика", 
	"История", "Литература", "МДК 08.02", "Архитектура АП",
	"Общество", "ОБЖ", "Физра", "Англ. яз.", "УИПД", "География"}
var Teachers = []string{
	"Трушина И.Ю", "Пахилько О.Н", "Морозова М.В", "Имашева К.Б",
	"Швиндина Н.А", "Зернова Е.А", "Имаев М.А", "Нургалиева И.Ю",
	"Чебрукова Т.А", "Новиков Р.Е", "Чернышова Е.А", "Силищева О.И",
	"Новикова М.Н"}
var Data = struct{
	Math, Rus, Inf, Fiz, Hist, Lit, Mdk, Arch, Soc, Obg, Fizr, Angl, UIPD, Geo DTO
}{
	Math: DTO{
		Teacher: Teacher{
			Name: Teachers[0],
		},
		Subject: Subjects[0],
		Hour:    10,
	},
	Rus: DTO{
		Teacher: Teacher{
			Name: Teachers[1],
		},
		Subject: Subjects[1],
		Hour:    11,
	},
	Inf: DTO{
		Teacher: Teacher{
			Name: Teachers[2],
		},
		Subject: Subjects[2],
		Hour:    12,
	}, 
	Fiz: DTO{
		Teacher: Teacher{
			Name: Teachers[3],
		},
		Subject: Subjects[3],
		Hour:    13,
	},
	Hist: DTO{
		Teacher: Teacher{
			Name: Teachers[4],
		},
		Subject: Subjects[4],
		Hour:    14,
	},
	Lit: DTO{
		Teacher: Teacher{
			Name: Teachers[1],
		},
		Subject: Subjects[5],
		Hour:    15,
	},
	Mdk: DTO{
		Teacher: Teacher{
			Name: Teachers[5],
		},
		Subject: Subjects[6],
		Hour:    16,
	},
	Arch: DTO{
		Teacher: Teacher{
			Name: Teachers[6],
		},
		Subject: Subjects[7],
		Hour:    17,
	},
	Soc: DTO{
		Teacher: Teacher{
			Name: Teachers[7],
		},
		Subject: Subjects[8],
		Hour:    18,
	},
	Obg: DTO{
		Teacher: Teacher{
			Name: Teachers[8],
		},
		Subject: Subjects[9],
		Hour:    19,
	},
	Fizr: DTO{
		Teacher: Teacher{
			Name: Teachers[9],
		},
		Subject: Subjects[10],
		Hour:    20,
	},
	Angl: DTO{
		Teacher: Teacher{
			Name: Teachers[10],
		},
		Subject: Subjects[11],
		Hour:    21,
	},
	UIPD: DTO{
		Teacher: Teacher{
			Name: Teachers[11],
		},
		Subject: Subjects[12],
		Hour:    22,
	},
	Geo: DTO{
		Teacher: Teacher{
			Name: Teachers[12],
		},
		Subject: Subjects[13],
		Hour:    23,
	},

}

