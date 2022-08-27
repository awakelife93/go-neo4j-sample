# Golang Neo4j Sample

## [Note]

```
1. study & learn golang, neo4j
```

### Description

1. neo4j admin(management browser) - {server domain or ip}:7474/browser
  - default username = neo4j / password = test
  - Because passwordChangeRequired=true, it is basically changed once because it is an unconditional password change when logging in for the first time.
  - If you do not build a new build in the changed state, an error occurs.
2. [neo4j docs](https://neo4j.com/docs/)

## Author

```
2021.08.10 ->
Author: Hyunwoo Park
```

## Getting Started

```
1. go run main.go or go build main.go
2. If you run it with docker-compose Please check each container environment.
  2-1. docker-compose.app.yml = Create and run only golang(app) container.
  2-2. docker-compose.neo4j.yml = Create and run only neo4j container.
  2-3. docker-compose.full.system.yml = Create and run each container for neo4j, and golang(app)
```
