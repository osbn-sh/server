package entity

type VoteOut struct {
	Id           int  `json:"id"`
	UniversityId *int `json:"university_id"`
	MajorId      *int `json:"major_id"`
	AdminBy      *int `json:"admin_by"`
	//must be hash
	Email string `json:"email"`
	Name  string `json:"name"`
}

type OptionVoteResult struct {
	OptionID    int
	OptionName  string
	Weight      int
	AverageRate float64
	VoteCount   int
}

type MyVote struct {
	Id         int
	OptionId   int
	Rate       int
	OptionName string
}
