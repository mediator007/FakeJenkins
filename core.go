package main

import (
	"fmt"
	"strconv"
	"time"
)

type Artifact struct {
	FileName     string `json:"fileName"`
	DisplayPath  string `json:"displayPath"`
	RelativePath string `json:"relativePath"`
}

const DefaultJobExecutionTime = "5"

func allBuilds() ([]Build, error) {
	builds, err := getAllBuilds()
	var result []Build
	if err != nil {
		return result, err
	}
	for _, build := range builds {
		if build.BuildStatus == "INQUEUE" {
			_, err = updateBuildStatus(strconv.FormatInt(build.ID, 10), "INPROGRESS")
		}

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

func buildJob(JobName string, executionTime string) (string, error) {

	if executionTime == "" {
		executionTime = DefaultJobExecutionTime
	}

	i, err := strconv.Atoi(executionTime)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Conversion error:", err)
		return "", err
	}
	queueNumber, err := insertBuild(JobName, i)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Insertion error:", err)
		return "", err
	}
	q := strconv.FormatInt(queueNumber, 10)
	return q, nil
}

func buildInfo(buildNumber string) (map[string]interface{}, error) {

	response := make(map[string]interface{})
	build, err := getBuildByBuildNumber(buildNumber)
	if err != nil {
		return response, err
	}

	if build.BuildStatus == "INQUEUE" {
		_, err = updateBuildStatus(buildNumber, "INPROGRESS")
	}

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
