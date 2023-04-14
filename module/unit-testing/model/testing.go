package model

type Name struct {
	Input   string
	SubTest string
	Status  bool
	Message string
}

var Names []Name = []Name{
	{
		Input:   "adli",
		SubTest: "benar",
		Status:  true,
		Message: "input not adli",
	},
	{
		Input:   "zakir",
		SubTest: "salah",
		Status:  false,
		Message: "input is adli",
	},
}
