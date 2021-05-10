package services


type RunProgramService struct {

}

func GetRunProgramService() *RunProgramService {
	return &RunProgramService{}
}

func (runServ *RunProgramService) Run() string {
	return ""
}
