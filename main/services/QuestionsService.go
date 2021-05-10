package services

import (
	"../../sols"
	"../dto"
	"../utils"
	"encoding/json"
	"strconv"
	//"fmt"
	"io/ioutil"
	"log"
)

type QuestionsService struct {

}

/*
 * TODO: make sure the instance is singleton
 */
func GetQuestionsService() *QuestionsService {
	return &QuestionsService{}
}

func (quesServ *QuestionsService) ArrayQuestions() string {

	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/arrays/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	var pageContent string
	for _, quesBody := range quesAnsContenet.QuesAnsList {
		var completeQuestionContent string
		//completeQuestionContent +=
		completeQuestionContent += utils.AfterQuestionBreak()

		for _, solBody := range quesBody.SolList {

			solFileName := solBody.SolAddress
			if solFileName == "" {
				continue
			}
			solCode, err := ioutil.ReadFile("./sols/"+solFileName)
			if err != nil {
				log.Fatal(err)
			}
			var answerContent string

			answerContent += utils.NoteBox(utils.GetNoteSegment(string(solCode)))

			answerContent += utils.CollapsibleWithoutPadding("Code", utils.CodeFormattingWithBoundary(utils.RemoveNoteSegmentFromCode(string(solCode))), true)

			if quesBody.SampleInp.Array != nil || quesBody.SampleInp.Number != -9999 {
				answerContent += utils.Collapsible("Sample Input: ", utils.CodeFormatting(getSampleInput(quesBody.SampleInp)), true)
				answerContent += utils.Collapsible("Output: ", utils.CodeFormatting(executeProgramUsingName(solFileName, quesBody.SampleInp.Array, quesBody.SampleInp.Number)), true)
				answerContent += utils.NormalBreak()
			}
			completeQuestionContent += utils.Collapsible("Solution " + strconv.Itoa(solBody.Id), answerContent, false)

			completeQuestionContent += utils.DifferentSolTypeBreak()
		}
		pageContent += utils.Collapsible(strconv.Itoa(quesBody.Id) + ". " + quesBody.Question, completeQuestionContent, true)
	}
	return utils.InHTMLTag(pageContent)
}

func getSampleInput(sampleInp dto.SampleInput) string {

	if sampleInp.Array != nil && sampleInp.Number != -9999 {
		content, err := json.Marshal(sampleInp)
		if err != nil {
			log.Fatal(err)
		}
		return string(content)
	} else if sampleInp.Array != nil {
		content, err := json.Marshal(sampleInp.Array)
		if err != nil {
			log.Fatal(err)
		}
		return string(content)
	} else if sampleInp.Number != -9999 {
		content, err := json.Marshal(sampleInp.Number)
		if err != nil {
			log.Fatal(err)
		}
		return string(content)
	} else {
		return ""
	}


}


func executeProgramUsingName(progName string, intArrayInp []int, intNumInp int) interface{} {
	switch progName {
	case "array_pairs_sum_z_01.go":
		return sols.Array_pairs_sum_z_01(intArrayInp, intNumInp)
	case "array_pairs_sum_z_02.go":
		return sols.Array_pairs_sum_z_02(intArrayInp, intNumInp)
	case "array_pairs_sum_z_03.go":
		return sols.Array_pairs_sum_z_03(intArrayInp, intNumInp)
	case "array_pairs_sum_z_04.go":
		return sols.Array_pairs_sum_z_04(intArrayInp, intNumInp)

	case "array_max_diff_two_elements_larger_after_smaller_01.go":
		return sols.Array_max_diff_two_elements_larger_after_smaller_01(intArrayInp)
	case "array_max_diff_two_elements_larger_after_smaller_02.go":
		return sols.Array_max_diff_two_elements_larger_after_smaller_02(intArrayInp)

	case "array_quicksort_01.go":
		return sols.Array_quicksort_01(intArrayInp)

	case "array_min_max_element_01.go":
		return sols.Array_min_max_element_01(intArrayInp)
	case "array_min_max_element_02.go":
		return sols.Array_min_max_element_02(intArrayInp)

	case "array_mergesort_01.go":
		return sols.Array_mergesort_01(intArrayInp)

	case "array_pair_sum_close_zero_01.go":
		return sols.Array_pair_sum_close_zero_01(intArrayInp)

	case "array_odd_occurring_element_01.go":
		return sols.Array_odd_occurring_element_01(intArrayInp)

	case "array_sep_zero_one_01.go":
		return sols.Array_sep_zero_one_01(intArrayInp)

	case "array_sorted_element_count_more_nby2_01.go":
		return sols.Array_sorted_element_count_more_nby2_01(intArrayInp)
	case "array_sorted_element_count_more_nby2_02.go":
		return sols.Array_sorted_element_count_more_nby2_02(intArrayInp)

	default:
		return ""
	}
}

