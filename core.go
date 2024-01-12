package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Artifact struct {
	FileName     string `json:"fileName"`
	DisplayPath  string `json:"displayPath"`
	RelativePath string `json:"relativePath"`
}

const DefaultJobExecutionTime = "10"

func allBuilds() ([]Build, error) {
	builds, err := getAllBuilds()
	var result []Build
	if err != nil {
		return result, err
	}
	for _, build := range builds {
		_inqueueStatusHandler(build)

		if build.BuildStatus == "INPROGRESS" {
			// Calculate the expected completion time
			expectedCompletionTime := build.StartTime.Add(time.Duration(build.ExecutionTime) * time.Second)
			// Compare with the current time
			currentTime := time.Now()
			if currentTime.After(expectedCompletionTime) {
				_, err = updateBuildStatus(strconv.FormatInt(build.ID, 10), "SUCCESS")
			}
		}
	}
	builds, err = getAllBuilds()
	return builds, nil
}

func _handlrParamsForBuildJob(executionTime string, forceFailParam string, forceUnstableParam string) (e string, fP bool, fU bool) {
	if executionTime == "" {
		executionTime = DefaultJobExecutionTime
	}

	var forceFail bool
	if forceFailParam == "True" {
		forceFail = true
	}

	var forceUnstable bool
	if forceUnstableParam == "True" {
		forceUnstable = true
	}
	return executionTime, forceFail, forceUnstable
}

func buildJob(JobName string, executionTime string, forceFailParam string, forceUnstableParam string) (string, error) {

	executionTime, forceFail, forceUnstable := _handlrParamsForBuildJob(executionTime, forceFailParam, forceUnstableParam)

	i, err := strconv.Atoi(executionTime)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Conversion error:", err)
		return "", err
	}
	queueNumber, err := insertBuild(JobName, i, forceFail, forceUnstable)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Insertion error:", err)
		return "", err
	}
	q := strconv.FormatInt(queueNumber, 10)
	return q, nil
}

func _inqueueStatusHandler(build Build) {
	if build.BuildStatus == "INQUEUE" {
		if build.ForceFail {
			_, _ = updateBuildStatus(strconv.FormatInt(build.ID, 10), "FAILURE")
			return
		}

		if build.ForceUnstable {
			_, _ = updateBuildStatus(strconv.FormatInt(build.ID, 10), "UNSTABLE")
			return
		}
		_, _ = updateBuildStatus(strconv.FormatInt(build.ID, 10), "INPROGRESS")
	}
}

func buildInfo(buildNumber string) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	build, err := getBuildByBuildNumber(buildNumber)
	if err != nil {
		return response, err
	}

	_inqueueStatusHandler(build)

	if build.BuildStatus == "INPROGRESS" {
		// Calculate the expected completion time
		expectedCompletionTime := build.StartTime.Add(time.Duration(build.ExecutionTime) * time.Second)
		// Compare with the current time
		currentTime := time.Now()
		if currentTime.After(expectedCompletionTime) {
			_, err = updateBuildStatus(buildNumber, "SUCCESS")
		}
	}

	artifact := Artifact{
		DisplayPath:  build.JobName + "_artifacts.json",
		FileName:     build.JobName + "_artifacts.json",
		RelativePath: build.JobName + "_artifacts.json",
	}

	response["artifacts"] = []Artifact{artifact}
	response["queuId"] = build.ID
	response["timestamp"] = time.Now().UnixMilli() - (build.StartTime.UnixMilli() / int64(time.Millisecond))
	response["result"] = build.BuildStatus
	// FIXME auto insertion
	response["url"] = "http://10.199.30.215:8080/job/folder/job/" + build.JobName + "/" + buildNumber

	return response, nil
}

func queueItem(queueNumber string) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	build, err := getBuildByBuildNumber(queueNumber)
	if err != nil {
		return response, err
	}
	executable := make(map[string]interface{})
	executable["number"] = build.ID
	response["executable"] = executable
	return response, nil
}

func jobInfo(jobName string) (map[string]interface{}, error) {
	response := make(map[string]interface{})
	inQueue, err := getAllInQueueBuilds()
	if err != nil {
		return response, err
	}
	response["buildable"] = true
	response["inQueue"] = inQueue
	return response, nil
}

func stopBuild(buildNumber string) {
	build, err := getBuildByBuildNumber(buildNumber)
	if err != nil {
		fmt.Println(err)
	}
	if build.BuildStatus == "INPROGRESS" {
		_, _ = updateBuildStatus(buildNumber, "ABORTED")
	}
}

type WavesConfig struct {
	Specs  []string `json:"specs"`
	Length float64  `json:"length"`
}

func checkJobType(jobName string) string {
	wavesSubstring := "waves"

	lowercaseJobName := strings.ToLower(jobName)

	if strings.Contains(lowercaseJobName, wavesSubstring) {
		return "GENERATOR"
	} else {
		return "DEFAULT"
	}
}

func getDefaultJobArtifact() map[string]interface{} {
	response := make(map[string]interface{})
	response["link"] = "http://google.com"
	return response
}

func getWavesConfigArtifact() []WavesConfig {
	wavesConfigs := []WavesConfig{
		{
			Specs:  []string{"autowaves_Yw2Ib", "autowaves_ImT7sa28O", "autowaves_4ylh2JSQg5vL"},
			Length: 2407.673870402982,
		},
		{
			Specs:  []string{"autowaves_cxmxM37Px87whMw", "autowaves_ExDZoTy8Vj76c6P"},
			Length: 2724.4167384707857,
		},
		{
			Specs:  []string{"autowaves_pVMpliAan", "autowaves_HaN9WkAbMJUW", "autowaves_LxW0zt5xpnRZ8Q"},
			Length: 3564.0500866046277,
		},
	}
	return wavesConfigs
}
