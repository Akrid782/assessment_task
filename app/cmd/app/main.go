package main

import (
	"AssessmentTask/internal/backend"
	"AssessmentTask/internal/util"
)

func main() {
	mod, isExistMod := util.GetEnv("MOD")

	if mod == "FILE_ANALYSIS" && isExistMod {
		backend.AnalysisFile()

		return
	}

	backend.AnalysisSystem()
}
