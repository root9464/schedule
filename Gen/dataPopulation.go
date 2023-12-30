package gen

type DTO struct {
	Teacher
	Subject string
	Hour    int16
}
type Teacher struct{
	Name string
}

var subjects = []string{
	"Математика", "Русский", "Информатика", "Физика", 
	"История", "Литература", "МДК 08.02", "Архитектура АП",
	"Общество", "ОБЖ", "Физра", "Англ. яз.", "УИПД", "География"}
var teachers = []string{
	"Трушина И.Ю", "Пахилько О.Н", "Морозова М.В", "Имашева К.Б",
	"Швиндина Н.А", "Зернова Е.А", "Имаев М.А", "Нургалиева И.Ю",
	"Чебрукова Т.А", "Новиков Р.Е", "Чернышова Е.А", "Силищева О.И",
	"Новикова М.Н"}
var Data = struct{
	Math, Rus, Inf, Fiz, Hist, Lit, Mdk, Arch, Soc, Obg, Fizr, Angl, UIPD, Geo DTO
}{
	Math: DTO{
		Teacher: Teacher{
			Name: teachers[0],
		},
		Subject: subjects[0],
		Hour:    10,
	},
	Rus: DTO{
		Teacher: Teacher{
			Name: teachers[1],
		},
		Subject: subjects[1],
		Hour:    11,
	},
	Inf: DTO{
		Teacher: Teacher{
			Name: teachers[2],
		},
		Subject: subjects[2],
		Hour:    12,
	}, 
	Fiz: DTO{
		Teacher: Teacher{
			Name: teachers[3],
		},
		Subject: subjects[3],
		Hour:    13,
	},
	Hist: DTO{
		Teacher: Teacher{
			Name: teachers[4],
		},
		Subject: subjects[4],
		Hour:    14,
	},
	Lit: DTO{
		Teacher: Teacher{
			Name: teachers[1],
		},
		Subject: subjects[5],
		Hour:    15,
	},
	Mdk: DTO{
		Teacher: Teacher{
			Name: teachers[5],
		},
		Subject: subjects[6],
		Hour:    16,
	},
	Arch: DTO{
		Teacher: Teacher{
			Name: teachers[6],
		},
		Subject: subjects[7],
		Hour:    17,
	},
	Soc: DTO{
		Teacher: Teacher{
			Name: teachers[7],
		},
		Subject: subjects[8],
		Hour:    18,
	},
	Obg: DTO{
		Teacher: Teacher{
			Name: teachers[8],
		},
		Subject: subjects[9],
		Hour:    19,
	},
	Fizr: DTO{
		Teacher: Teacher{
			Name: teachers[9],
		},
		Subject: subjects[10],
		Hour:    20,
	},
	Angl: DTO{
		Teacher: Teacher{
			Name: teachers[10],
		},
		Subject: subjects[11],
		Hour:    21,
	},
	UIPD: DTO{
		Teacher: Teacher{
			Name: teachers[11],
		},
		Subject: subjects[12],
		Hour:    22,
	},
	Geo: DTO{
		Teacher: Teacher{
			Name: teachers[12],
		},
		Subject: subjects[13],
		Hour:    23,
	},

}

