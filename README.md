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
