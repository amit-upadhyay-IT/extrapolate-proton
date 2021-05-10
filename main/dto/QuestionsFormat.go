package dto

type QuestionAnswer struct {
	QuesAnsList []QuestionAnswerBody `json:"questions_list"`
}

type Solution struct {
	Id         int    `json:"id"`
	SolAddress string `json:"sol_address"`
	Notes      string `json:"notes"`
}

type SampleInput struct {
	Array  []int `json:"array"`
	Number int   `json:"number"`
}

type QuestionAnswerBody struct {
	Id        int         `json:"id"`
	Question  string      `json:"question"`
	SolList   []Solution  `json:"solutions_list"`
	SampleInp SampleInput `json:"sample_input"`
}
