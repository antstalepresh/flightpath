# Golang API for tracking user's flight path
 
## Story
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

## Goal
To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.


## Required JSON structure
```
[["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
[["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
```

## Commands
`make test` will run the unit tests

`make run` will run the service

`make build` will build the binary file

`make lint` will check linter issues and fix

## Format of API endpoint
### Request
`POST /calculate <LIST OF FLIGHTS>`

```
curl -i -X POST -H "Content-Type:application/json" \
http://127.0.0.1:8080/calculate \
-d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
```
### Response
```
HTTP/1.1 200 OK
Date: Wed, 08 Mar 2023 12:11:08 GMT
Content-Type: application/json
Content-Length: 13

["SFO","EWR"]
```
