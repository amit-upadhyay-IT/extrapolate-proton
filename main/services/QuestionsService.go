package services

import (
	"../../sols"
	"../../sols/graph"
	"../../sols/system_design"
	"../../sols/tree"
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

	return formQuestionAnswerBody(quesAnsContenet, "sols", true)

}


func (quesServ *QuestionsService) RecursionQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/recursion/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols", true)
}

func (quesServ *QuestionsService) DynamicProgrammingQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/dynamic_programming/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols", true)
}

func (quesServ *QuestionsService) TreeQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/tree/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols/tree", true)
}

func (quesServ *QuestionsService) SystemDesignQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/system_design/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols/system_design", true)
}

func (quesServ *QuestionsService) GraphQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/graph/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols/graph", true)
}

func (quesServ *QuestionsService) DesignPatternQuestions() string {
	// reading the questions and answers json
	contents, _ := ioutil.ReadFile("./raw/design_patterns/question_answer.json")
	var quesAnsContenet dto.QuestionAnswer
	err := json.Unmarshal(contents, &quesAnsContenet)
	if err != nil {
		log.Fatal(err)
	}

	return formQuestionAnswerBody(quesAnsContenet, "sols_design", false)
}



func formQuestionAnswerBody(quesAnsContenet dto.QuestionAnswer, dirName string, isCodeDisplay bool) string {

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
			solCode, err := ioutil.ReadFile("./"+dirName+"/"+solFileName)
			if err != nil {
				log.Fatal(err)
			}
			var answerContent string

			answerContent += utils.NoteBox(utils.GetNoteSegment(string(solCode)))

			if isCodeDisplay {
				answerContent += utils.CollapsibleWithoutPadding("Code", utils.CodeFormattingWithBoundary(utils.RemoveNoteSegmentFromCode(string(solCode))), true)
			}

			if quesBody.SampleInp.Array != nil || quesBody.SampleInp.Number != -9999 {
				answerContent += utils.Collapsible("Sample Input: ", utils.CodeFormatting(getSampleInput(quesBody.SampleInp)), true)
				answerContent += utils.Collapsible(
					"Output: ", utils.CodeFormatting(
						executeProgramUsingName(
							solFileName,
							quesBody.SampleInp.Array,
							quesBody.SampleInp.Array2,
							quesBody.SampleInp.Number,
							quesBody.SampleInp.Number2)),
							true)
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
	} else if sampleInp.Number != -9999 && sampleInp.Number2 != 9999 {
		num1 := sampleInp.Number
		num2 := sampleInp.Number2
		return "num1: " + strconv.Itoa(num1) + " num2: " + strconv.Itoa(num2)
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


func executeProgramUsingName(progName string, intArrayInp, intArrayInp2 []int, intNumInp int, intNumInp2 int) interface{} {
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

	case "array_find_element_occur_more_nby2_unsorted_01.go":
		return sols.Array_find_element_occur_more_nby2_unsorted_01(intArrayInp)
	case "array_find_element_occur_more_nby2_unsorted_02.go":
		return sols.Array_find_element_occur_more_nby2_unsorted_02(intArrayInp)

	case "array_find_element_occur_more_nbyk_unsorted_01.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_find_element_occur_more_nbyk_unsorted_01(cpy, intNumInp)
	case "array_find_element_occur_more_nbyk_unsorted_02.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_find_element_occur_more_nbyk_unsorted_02(cpy, intNumInp)
	case "array_find_element_occur_more_nbyk_unsorted_03.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_find_element_occur_more_nbyk_unsorted_03(cpy, intNumInp)

	case "array_triplets_for_sum_x_01.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_triplets_for_sum_x_01(cpy, intNumInp)
	case "array_triplets_for_sum_x_02.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_triplets_for_sum_x_02(cpy, intNumInp)
	case "array_triplets_for_sum_x_03.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_triplets_for_sum_x_03(cpy, intNumInp)


	case "array_smallest_positive_missing_number_01.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_smallest_positive_missing_number_01(cpy)
	case "array_smallest_positive_missing_number_02.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_smallest_positive_missing_number_02(cpy)
	case "array_smallest_positive_missing_number_03.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_smallest_positive_missing_number_03(cpy)


	case "array_distribute_minimum_coin_in_student_01.go":
		cpy := make([]int, len(intArrayInp))
		copy(cpy, intArrayInp)
		return sols.Array_distribute_minimum_coin_in_student_01(cpy)



	case "recursion_power_function_01.go":
		return sols.Recursion_power_function_01(intNumInp, intNumInp2)
	case "recursion_power_function_02.go":
		return sols.Recursion_power_function_02(intNumInp, intNumInp2)
	case "recursion_power_function_03.go":
		return sols.Recursion_power_function_03(intNumInp, intNumInp2)

	case "recursion_tower_of_hanoi_01.go":
		return sols.Recursion_tower_of_hanoi_01(intNumInp)

	case "recursion_print_binary_01.go":
		return sols.Recursion_print_binary_01(intNumInp)
	case "recursion_print_decimal_01.go":
		return sols.Recursion_print_decimal_01(intNumInp)
	case "recursion_die_combination_sum_01.go":
		return sols.Recursion_die_combination_sum_01(intNumInp, intNumInp2)



	case "dynamic_knapsack01_01.go":
		return sols.Dynamic_knapsack01_01(intArrayInp, intArrayInp2, intNumInp)
	case "dynamic_knapsack01_02.go":
		return sols.Dynamic_knapsack01_02(intArrayInp, intArrayInp2, intNumInp)
	case "dynamic_knapsack01_03.go":
		return sols.Dynamic_knapsack01_03(intArrayInp, intArrayInp2, intNumInp)


	case "tree_count_of_binary_trees_01.go":
		return tree.Tree_count_of_binary_trees_01(intNumInp)
	case "tree_count_no_of_nodes_in_binary_tree_01.go":
		return tree.Tree_count_no_of_nodes_in_binary_tree_01()
	case "tree_count_no_of_leaf_nodes_in_binary_tree_01.go":
		return tree.Tree_count_no_of_leaf_nodes_in_binary_tree_01()
	case "tree_count_no_of_nonleaf_nodes_in_binary_tree_01.go":
		return tree.Tree_count_no_of_nonleaf_nodes_in_binary_tree_01()
	case "tree_count_no_of_full_nodes_in_binary_tree_01.go":
		return tree.Tree_count_no_of_full_nodes_in_binary_tree_01()
	case "tree_count_height_binary_tree_01.go":
		return tree.Tree_count_height_binary_tree_01()
	case "tree_count_height_binary_tree_02.go":
		return tree.Tree_count_height_binary_tree_02()
	case "tree_binary_search_tree_basics.go":
		return tree.Tree_binary_search_tree_basics()
	case "tree_verify_if_bst_01.go":
		return tree.Tree_verify_if_bst_01()


	case "system_design_consistent_hashing.go":
		return system_design.System_design_consistent_hashing()

	case "graph_basics.go":
		return graph.Graph_basics()
	default:
		return ""
	}
}

