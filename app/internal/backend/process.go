package backend

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"AssessmentTask/internal/constant"
	"AssessmentTask/internal/util"
	"github.com/shirou/gopsutil/process"
)

type SystemProcess struct {
	ProcessName string `json:"Process Name"`
	ProcessID   int32  `json:"Process ID"`
	IsPrime     bool   `json:"Is Prime"`
}

func AnalysisSystem() {
	processRegexpFilter, isExistRegExp := util.GetEnv("PROCESS_REGEXP_FILTER")
	processes, err := process.Processes()
	if err != nil {
		return
	}

	var systemProcessList []SystemProcess
	for _, p := range processes {
		processName, _ := p.Name()
		pid := p.Pid

		if isExistRegExp {
			match, _ := regexp.MatchString(processRegexpFilter, processName)
			if !match {
				continue
			}
		}

		systemProcessList = append(
			systemProcessList, SystemProcess{
				ProcessName: processName,
				ProcessID:   pid,
				IsPrime:     big.NewInt(int64(p.Pid)).ProbablyPrime(0),
			},
		)
	}

	if len(systemProcessList) == 0 {
		fmt.Println("Информация о процессах не найдена, текущий фильтр: " + processRegexpFilter)

		return
	}

	jsonData, err := json.MarshalIndent(systemProcessList, "", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(constant.Assets+"/processes.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Информация о процессах успешно сохранена в файл processes.json")
}

func AnalysisFile() {
	pathFile, isExistPath := util.GetEnv("PATH_FILE_ANALYSIS")
	if pathFile == "" && !isExistPath {
		fmt.Println("Неизвестный путь до файла для анализ: " + pathFile)

		return
	}

	data, err := os.ReadFile(constant.Src + "/" + strings.TrimLeft(pathFile, "/"))
	if err != nil {
		fmt.Println("Ошибка чтения файла: " + err.Error())

		return
	}

	var systemProcessList []SystemProcess
	err = json.Unmarshal(data, &systemProcessList)

	errorCount := 0

	for _, p := range systemProcessList {
		actualIsPrime := big.NewInt(int64(p.ProcessID)).ProbablyPrime(0)
		if actualIsPrime != p.IsPrime {
			errorCount++
			fmt.Printf("Для процесса %s найдена ошибка.\n", p.ProcessName)
			fmt.Printf(
				"Число (PID) %d на самом деле %s, а в файле указано %s.\n",
				p.ProcessID, defineTypeNumber(actualIsPrime), defineTypeNumber(p.IsPrime),
			)
		}
	}

	fmt.Printf("\nВ файле найдены ошибки: %d шт.\n", errorCount)
}

func defineTypeNumber(isPrime bool) string {
	if isPrime {
		return "простое"
	}

	return "составное"
}
