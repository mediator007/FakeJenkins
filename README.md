# FakeJenkins
Tool for mocking Jenkins  
![Screenshot](screenshot.png)  

FE ```localhost:8080```  

## Available API
#### Create Build of DEFAULT Job:  
``` POST /job/folder/job/ANY_JOB_NAME/buildWithParameters?executionTime=10```  
Example:  
```curl -X POST -i localhost:8080/job/folder/job/ANY_JOB_NAME/buildWithParameters?executionTime=10``` ->  
```7``` & ```headers[Location]: some/strange/url/7/```  

#### Get build info:  
``` GET /job/folder/job/ANY_JOB_NAME/<build_number>/api/json```  
Example:  
```curl localhost:8080/job/folder/job/default/1/api/json``` ->   
```{"artifacts":[{"fileName":"ANY_JOB_NAME_artifacts.json","displayPath":"ANY_JOB_NAME_artifacts.json","relativePath":"ANY_JOB_NAME_artifacts.json"}],"queuId":1,"result":"SUCCESS","timestamp":1704975685909,"url":"http://10.199.30.215:8080/job/folder/job/ANY_JOB_NAME/1"}```   
First request after build creation return INQUEUE status. Second - launch build &  
change status to INPROGRESS. Build became SUCCESSFUL if startTime + execTime >= currentTime  

#### Get job info:
```GET /job/folder/job/DEFAULT/api/json```  
Example:  
```curl localhost:8080/job/DEFAULT/api/json``` ->   
```{"buildable":true,"inQueue":[{"ID":9,"ExecutionTime":10,"StartTime":"2024-01-10T13:55:25Z","BuildStatus":"INQUEUE","JobName":"DEFAULT"}]}```  

#### Get queue item:
```GET /queue/item/<queueNumber>/api/json```  
Example:  
```curl localhost:8080/queue/item/6/api/json``` ->   
```{"executable": {"number": 6}}```  

#### Delete all builds:  
```DELETE /deleteAllBuilds```  
Example:   
```curl -X DELETE -i localhost:8080/deleteAllBuilds``` ->  
```"Successful delete all"```  

#### GetArtifact:
```GET /job/folder/job/<jobName>/<buildNumber>/artifact/<artifactName>```

### Launching
Insert instance IP adress to ```core.go/buildInfo/96:response["url"]```  

#### Dev Launch
```go run .```  

#### Prod Launch
For docker compose v2  
```docker compose build```  
```docker compose up -d```  
  
