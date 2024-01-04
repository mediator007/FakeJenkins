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

func buildInfo(c *gin.Context) {
	jobName := c.Param("jobName")
	buildNumber := c.Param("buildNumber")
	// TODO
	response := "Job " + jobName + "--Data for buildNumber " + buildNumber
    c.String(http.StatusOK, response)
}

func queueItem(c *gin.Context) {
	queueNumber := c.Param("queueNumber")
	//TODO
	response := "Queue number " + queueNumber
	c.String(http.StatusOK, response)
}

func buildJob(c *gin.Context) {
	jobName := c.Param("jobName")
	// shortcut for c.Request.URL.Query().Get("lastname")
	// url?commit=123
	commit := c.Query("commit")
	response := "Job name " + jobName + "-- Commit " + commit
	//TODO
	c.String(http.StatusAccepted, response)
}

func main() {
  r := gin.Default()
  r.GET("/ping", pong)
  r.GET("/job/:jobName/api/json", jobInfo)
  r.GET("job/:jobName/:buildNumber/api/json", buildInfo)
  r.GET("queue/item/:queueNumber/api/json", queueItem)
  r.POST("job/:jobName/buildWithParameters", buildJob)
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}