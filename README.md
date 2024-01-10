# FakeJenkins
Tool for mocking Jenkins  
![Screenshot](screenshot.png)  

FE ```localhost:8080```  

### Available API
Create Build of DEFAULT Job:  
``` POST /job/ANY_JOB_NAME/buildWithParameters?executionTime=10```  

Get build info:  
``` GET /job/ANY_JOB_NAME/<build_number>/api/json```  

Delete all builds:  
```DELETE /deleteAllBuilds```  


### Dev Launch

```go run .```  

### Prod Launch
For docker compose v2  
```docker compose build```  
```docker compose up -d```  
  
