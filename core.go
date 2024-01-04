package main

import (
	"fmt"
	"strconv"
)

func buildJob(executionTime string) (string, error){
	i, err := strconv.Atoi(executionTime)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Conversion error:", err)
		return "", err
	}
	queueNumber, err := insertBuild(i)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Insertion error:", err)
		return "", err
	}
	q := strconv.FormatInt(queueNumber, 10)
	return q, nil
}

func buildInfo(buildNumber string) map[string]interface{} {


	// Create a map with the key "artifacts."
	response := make(map[string]interface{})

	response["artifacts"] = []string{"artifact 1", "artifact 2"}
	response["queuId"] = "BUILD NUMBER - " + buildNumber

	return response
}