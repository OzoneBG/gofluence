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
Right now it doesn't work properly but:
to build the docker image

`docker build -t gofluence .`

To run it:
`docker run -p 8080:8081 -it gofluence`

# TODO
- [ ] Rename environment variables to be more clear and readable
- [ ] Setup `docker-compose` to work with both the app and `PostgreSQL`
- [ ] Update api to return articles with the username and id
