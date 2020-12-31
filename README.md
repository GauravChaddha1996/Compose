
# Compose
A blogging backend written with the aim to learn Golang basics, how to best structure things in it, basic CRUD API building and learning mysql (above the usual queries).

### Learnings
* Golang is easy to learn and use. 
* It’s better to keep data layer separate since packages will need to call each other. 
Due to golang structure, they want each package to be self sufficient,  making interfaces for each package and maintaining it becomes quite cumbersome. 
* Hate that it's missing optional parameters in functions, repetitive error handling, necessary parenthesis even with single lines, overloading functions. 

### Code-structure
The code is divided into following packages and files:
1. **Commons**
	* To hold stuff common for other packages - logging, validators, input sanitizers, utils, error handling etc. 
	* Depends on no other package
2. **Data layer**
	* Holds the DB Data models, DAO’s to access them and useful API entities 
	* Depends on commons package
3. **Endpoints**
	* Sub-packages for user, comment & replies, like and articles endpoints
	* Each sub-package holds a handler.go to register different endpoints and their respective handlers in the sub router along with any custom config for security, timeout etc.
	* Depends on commons & data layer packages
4. **Middlewares**
	* Holds all the middlewares - common model (access token, user id) , security, request and response logging, timeout
	* Depends on commons and data layer packages
5. **main.go**
	*  The entry point to this monolith
6. **config.json**
	* Holds the configuration for database and server

## CRUD api standards
Each API in the crud is divided into 3 files
1. models.go 
	1. Holds RequestModel
		1. The model  has necessary fields for the operation (e.g. article id, common model etc.)
		2. **Input validation** struct tags using [go-validator](https://github.com/asaskevich/govalidator) for ids, email, string length, password etc.
		3. **Sanitising input** using [Bluemonday](https://github.com/microcosm-cc/bluemonday) for prevention against XSS attacks
	2. ResponseModel
		1. The model that will be sent back as json result for the API
		2. Json tags for marshalling
2. handler.go
	1. Handles the incoming request and transforms it into the request model
	2. Checks for **Authorization** and security checks (permissions, id existing etc.) and provides model to action.go
	3. Makes a sub-logger from the default [logger library](https://github.com/rs/zerolog)  for **structured logging** (i.e. adds certain relevant field to each API)
3. action.go
	1. Handles the actual actions executing business logic 
	2. Uses one or more DAO’s for the db tasks
	3. Handles db tasks using transaction where necessary
	4. Uses the sub-logger from its handler.go to log steps that were completed or failed with relevant message.
