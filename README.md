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
- Load test 
- Performance test 

### DB design 

### API swagger doc 
- Go to: http://localhost:3000/docs/index.html#/
- https://blog.logrocket.com/documenting-go-web-apis-with-swag/ 


### Architecture 


### Apply clean code 
- Code quality: golint, 

### High availability, scalable and fault tolerant API service 
- Add retry mechanism
- 


### How to run 
- make setup && make build
- Base URL: http://localhost:3000/api/v1




### TODO next 
- Add more testing for good testing coverage 
- Add input validators 
- Cover more error handling, crash recovery
- Security 
- Add metrics, logging, alert 
