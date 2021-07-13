# marvelapi

## About
marvel rest api built in golang. This service is runned to port 8080 (by default)
serving :
1. Get All Character IDs (host:8080/characters)
2. Get Character Info (Id,Name,Description) By id (host:8080/characters/<id>)

## Dependencies
- Golang 1.16+
- Redis Server (:6379)

## To run
1. clone this repository
2. make sure you have all dependencies ready and setup
3. Configure the `.env` files with your credentials needed from marvel: https://developer.marvel.com/
4. execute `make run` on the main project folder

## Development Notes

### To test
`make test`

### To generate mockery for unit tests mock interface
```bash
mockery --all --recursive=true --inpackage
```