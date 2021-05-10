package resources

import (
	"../services"
	"fmt"
	"net/http"
)


type RunProgramResource struct {
	runProgServ *services.RunProgramService
}

// TODO: make this singleton
func GetRunProgramResource() *RunProgramResource {
	return &RunProgramResource{runProgServ:services.GetRunProgramService()}
}

/*
 * /v1/run
 */
func (runResource *RunProgramResource) Run(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",  runResource.runProgServ.Run())
}
