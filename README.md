# CHALLENGE

### <a name='About'></a>About

🚀 A simple Rest API using go-chi, that can check or access for github users information with one endpoint. This solves the problem of the github api since they only cater one username per information using there users endpoint. This api accepts a list of github users and returns the user information including the following.

- name
- login
- company
- number of followers
- number of public repos


### <a name='Local'></a>Local Setup

1. Clone the repository:
```
git clone https://github.com/Jeswyrne/chlnge.git
```

2. Start local development server
```
go run .
```

Run all tests

```
go test -v ./...
```

### <a name='Dockersupport'></a>Docker support

You can build and run production with docker

1. Build docker image

```
docker build . -t production-coding-challenge
```

2. Run it with your prefer port

```
docker run -d -p 3001:3000 production-coding-challenge
```

Launch http://localhost:3001