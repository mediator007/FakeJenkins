package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
}

func jobInfo(c *gin.Context) {
	jobName := c.Param("jobName")
	// TODO 
	response := "Data for Job " + jobName 
    c.String(http.StatusOK, response)
}

func buildInfoHandler(c *gin.Context) {
	// curl localhost:8080/job/ANY_JOB_NAME/<BUILD-NUMBER>/api/json
	// jobName := c.Param("jobName")
	buildNumber := c.Param("buildNumber")
	response, err := buildInfo(buildNumber)
	if err != nil {
		badResponse :=  "Cant get build info with buildNumber: " + buildNumber
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func queueItem(c *gin.Context) {
	queueNumber := c.Param("queueNumber")
	//TODO
	response := "Queue number " + queueNumber
	c.String(http.StatusOK, response)
}

func buildJobHandler(c *gin.Context) {
	// curl -X POST -i localhost:8080/job/ANY_JOB_NAME/buildWithParameters?executionTime=55
	executionTime := c.Query("executionTime")
	response, err := buildJob(executionTime)
	if err != nil {
		response =  "Cant build Job with execTime " + executionTime
		c.JSON(http.StatusBadRequest, response)
	} else {
		header := "some/strange/url/" + response + "/"
		c.Header("Location", header)
		c.JSON(http.StatusAccepted, response)
	}
}

func main() {
	dbInitialization()
	r := gin.Default()
	r.GET("/ping", pong)
	r.GET("/job/:jobName/api/json", jobInfo)
	r.GET("job/:jobName/:buildNumber/api/json", buildInfoHandler)
	r.GET("queue/item/:queueNumber/api/json", queueItem)
	r.POST("job/:jobName/buildWithParameters", buildJobHandler)
	r.Run()
}