# api-online-book-store

### Features 
- Register/ Login with email/password 
- Authentication with JWT - get JWT token and use them in authorization header for response.
- DB migration 
- Swagger doc 
- Environment configs
- Docker
- CD/CI build actions 
### Tech stack 
- Go 1.19 - Framework: Gin 
- Database: MySQL, Redis 
- For testing purpose I use sqlite 
- Unit test 

### DB design 
![Screenshot 2023-02-20 at 7.52.35 PM.png](..%2F..%2F..%2F..%2Fvar%2Ffolders%2Fqn%2F22kldy0x4bq_9zf3f47txs9m0000gn%2FT%2FTemporaryItems%2FNSIRD_screencaptureui_Bo3AoI%2FScreenshot%202023-02-20%20at%207.52.35%20PM.png)
### API swagger doc 
- Go to: http://localhost:3000/docs/index.html#/ 


### Architecture 


### Apply clean code 
- Code quality: golint

### High availability, scalable and fault tolerant API service 
- Add retry mechanism


### How to run 
- make up
- Base URL: http://localhost:3000/api/v1
- Test API use Postman
- Run unit test: go test -v
### For testing API 
- Register by email and password 
- For auth token simplicity, you can just login with your email and password, copy paste the token
- Then use the token in Bearer token header to request to other APIs 

### For improvement 
- Add more testing for good testing coverage
- Folder structure: add layer of abstraction, separate controller and service layer
- Add input validators 
- Cover more error handling, crash recovery
- Security 
- Add metrics, logging middleware 
- Add concurrency for better performance 
