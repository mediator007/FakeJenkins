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

func buids(c *gin.Context) {
	builds, err := allBuilds()
	if err != nil {
		badResponse := "Cant get all builds"
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, builds)
	}
}

func deleteAllBuildsHandler(c *gin.Context) {
	err := deleteAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Delete all error")
	} else {
		c.JSON(http.StatusOK, "Successful delete all")
	}
}

func jobInfoHandler(c *gin.Context) {
	jobName := c.Param("jobName")
	response, err := jobInfo(jobName)
	if err != nil {
		badResponse := "Cant get job info with jobNAme: " + jobName
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func buildInfoHandler(c *gin.Context) {
	buildNumber := c.Param("buildNumber")
	response, err := buildInfo(buildNumber)
	if err != nil {
		badResponse := "Cant get build info with buildNumber: " + buildNumber
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func queueItemHandler(c *gin.Context) {
	queueNumber := c.Param("queueNumber")
	response, err := queueItem(queueNumber)
	if err != nil {
		badResponse := "Cant get queue item with queueNumber: " + queueNumber
		c.JSON(http.StatusBadRequest, badResponse)
	} else {
		c.JSON(http.StatusOK, response)
	}
}

func buildJobHandler(c *gin.Context) {
	jobName := c.Param("jobName")
	queryParams := c.Request.URL.Query()
	executionTime := queryParams.Get("executionTime")
	forceFail := queryParams.Get("forceFail")
	forceUnstable := queryParams.Get("forceUnstable")

	response, err := buildJob(jobName, executionTime, forceFail, forceUnstable)
	if err != nil {
		response = "Cant build Job with execTime " + executionTime
		c.JSON(http.StatusBadRequest, response)
	} else {
		header := "some/strange/url/" + response + "/"
		c.Header("Location", header)
		c.JSON(http.StatusAccepted, response)
	}
}

func returnStatic(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getArtifactFileHandler(c *gin.Context) {
	jobName := c.Param("jobName")
	jobType := checkJobType(jobName)
	if jobType == "GENERATOR" {
		c.JSON(http.StatusOK, getWavesConfigArtifact())
		return
	}
	c.JSON(http.StatusOK, getDefaultJobArtifact())
}

func stopBuildHandler(c *gin.Context) {
	buildNumber := c.Param("buildNumber")
	stopBuild(buildNumber)
	c.JSON(http.StatusAccepted, "Build stopped")
}

func main() {
	dbInitialization()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", returnStatic)
	r.GET("/ping", pong)
	r.GET("/builds", buids)
	r.DELETE("/deleteAllBuilds", deleteAllBuildsHandler)
	r.GET("/job/:folder/job/:jobName/api/json", jobInfoHandler)
	r.GET("job/:folder/job/:jobName/:buildNumber/api/json", buildInfoHandler)
	r.GET("job/:folder/job/:jobName/:buildNumber/artifact/:artifactFile", getArtifactFileHandler)
	r.GET("queue/item/:queueNumber/api/json", queueItemHandler)
	r.POST("job/:folder/job/:jobName/buildWithParameters", buildJobHandler)
	r.POST("job/:folder/job/:jobName/:buildNumber/stop", stopBuildHandler)
	r.Run()
}
