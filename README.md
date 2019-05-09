# gofluence
A simple wiki page like api

# Docker status
Currently the docker container builds and runs but must specify environment vari
ables that is:

* PORT - listening port
* DB_USER - database username
* DB_PWD - database password
* DB_NAME - database name
* TOKEN_PWD - token password

# Building and running
## To run the application on docker:
`docker-compose up`

## To build the application and run it locally
`make`


# TODO
- [ ] Rename environment variables to be more clear and readable
- [X] Setup `docker-compose` to work with both the app and `PostgreSQL`
- [ ] Update api to return articles with the username and id
