# DNAS - A simple Cisco DNA Spaces Go Library

[![GoDoc](https://godoc.org/github.com/darrenparkinson/dnas?status.svg)](https://godoc.org/github.com/darrenparkinson/dnas)
[![PkgGoDev](https://pkg.go.dev/badge/darrenparkinson/dnas)](https://pkg.go.dev/github.com/darrenparkinson/dnas)
[![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/dnas)](https://goreportcard.com/report/github.com/darrenparkinson/dnas)


This repository is intended as a simple to use library for the Go language to interact with the [Cisco DNA Spaces API](https://developer.cisco.com/docs/dna-spaces).

## Installing

```sh
go get github.com/darrenparkinson/dnas
```

## Usage

```go
import "github.com/darrenparkinson/dnas"
```

Construct a new DNA Spaces client, then use the various services on the client to access different parts of the DNA Spaces API.  For example:

```go
d, _ := dnas.NewClient(apikey, region, nil)
floors, err := d.ActiveClients.Floors(context.Background())
```

Some API methods have optional parameters that can be passed, for example:

```go
d, _ := dnas.NewClient(apikey, region, nil)
opt := &dnas.ClientParameters{Associated: dnas.Bool(true), DeviceType: dnas.String("CLIENT")}
count, err := d.ActiveClients.Count(context.Background(), opt)
```

The services of a client divide the API into logical chunks and correspond to the structure of the DNA Spaces API documentation at https://developer.cisco.com/docs/dna-spaces/#!dna-spaces-location-cloud-api

NOTE: Using the context package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then context.Background() can be used as a starting point.

## Authentication

Authentication is provided by an API Key as outlined [in the documentation](https://developer.cisco.com/docs/dna-spaces/#!getting-started/getting-started).  You are able to provide the API Key as part of initialisation using `NewClient`.  

## Region

Cisco supports two regions with DNA Spaces.  You must provide the region you are using to `NewClient` on initialisation.  You can tell which region you are using by the URL you use for DNA Spaces.

| URL          | Region Value |
|--------------|--------------|
| dnaspaces.io | `io`         |
| dnaspaces.eu | `eu`         |


## Helper Functions

Most structs for resources use pointer values.  This allows distinguishing between unset fields and those set to a zero value.  Some helper functions have been provided to easily create these pointers for string, bool and int values as you saw above and here, for example:

```go
opts := &dnas.ClientParameters{
    Associated: dnas.Bool(true),
    DeviceType: dnas.String("CLIENT"),
}
```

## Errors

In the [documentation](https://developer.cisco.com/docs/dna-spaces), Cisco identifies four returned errors.  These are provided as constants so that you may check against them:

| Code | Error                | Constant           |
|------|----------------------|--------------------|
| 400  | Bad Request          | `ErrBadRequest`    |
| 401  | Unauthorized Request | `ErrUnauthorized`  |
| 403  | Forbidden            | `ErrForbidden`     |
| 500  | Internal Error       | `ErrInternalError` |

All other errors are returned as `ErrUnknown`

As an example:

```go
count, err := c.ActiveClients.Count(ctx, &dnas.ClientParameters{})
if errors.Is(err, dnas.ErrUnauthorized) {
	log.Fatal("Sorry, you're not allowed to do that.")
}
```

# Roadmap

Currently this library only implements some of the functionality.  It is intended that this API will support all endpoints as they are required.  Feel free to log issues if there are specific endpoints you'd like, or see Contributing.  

The following oulines the available endpoints and their status in relation to implementation in this library.  As you can see, still a way to go.

## Map

| Method | Endpoint                      | Status          |
|--------|-------------------------------|-----------------|
| POST   | /map                          | Not Implemented |
| GET    | /map/hierarchy                | Not Implemented |
| GET    | /map/elements/{elementId}     | Not Implemented |
| DELETE | /map/elements/{elementId}     | Not Implemented |
| GET    | /map/images/floor/{imageName} | Not Implemented |

## Active Clients

| Method | Endpoint        | Status          |
|--------|-----------------|-----------------|
| GET    | /clients        | Not Implemented |
| GET    | /clients/count  | Implemented     |
| GET    | /clients/floors | Implemented     |

## Access Points

| Method | Endpoint            | Status          |
|--------|---------------------|-----------------|
| GET    | /accessPoints       | Not Implemented |
| GET    | /accessPoints/count | Not Implemented |

## Clients History

| Method | Endpoint                    | Status          |
|--------|-----------------------------|-----------------|
| GET    | /history                    | Not Implemented |
| GET    | /history/records/count      | Not Implemented |
| GET    | /history/clients            | Not Implemented |
| GET    | /history/clients/{deviceId} | Not Implemented |

## Notifications

| Method | Endpoint                                   | Status          |
|--------|--------------------------------------------|-----------------|
| GET    | /notifications                             | Not Implemented |
| POST   | /notifications                             | Not Implemented |
| PUT    | /notifications                             | Not Implemented |
| GET    | /notifications/{subscriptionId}            | Not Implemented |
| DELETE | /notifications/{subscriptionId}            | Not Implemented |
| GET    | /notifications/{subscriptionId}/statistics | Not Implemented |


# Contributing

Since all endpoints would ideally be covered, contributions are always welcome.  Adding new methods should be relatively straightforward.

# Versioning

In general dnas follows [semver](https://semver.org/) for tagging releases of the package.  As yet, it is still in development and has not had a tag added, but will in due course.  Since it is still in development, you may expect some changes.

# License

This library is distributed under the MIT license found in the [LICENSE](LICENSE) file.