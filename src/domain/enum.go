package domain


//type Alias1 = string
type List struct {
	Program_name int
	OrderByDate_Desc int
	OrderByDate_Asc int
	OrderByLocation int
}

var Enum = &List{
	OrderByDate_Desc: 1,
	OrderByDate_Asc:  2,
	OrderByLocation:  3,
	Program_name: 4,
}

type Member struct {
	Silver string
	Gold string
	Platinum string
}

var EnumMember = &Member{
	Silver:   "1",
	Gold:     "2",
	Platinum: "3",
}
