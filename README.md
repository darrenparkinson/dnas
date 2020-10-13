# DNAS - A simple Cisco DNA Spaces Go Library

[![GoDoc](https://godoc.org/github.com/darrenparkinson/dnas?status.svg)](https://godoc.org/github.com/darrenparkinson/dnas)
[![PkgGoDev](https://pkg.go.dev/badge/darrenparkinson/dnas)](https://pkg.go.dev/github.com/darrenparkinson/dnas)
[![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/dnas)](https://goreportcard.com/report/github.com/darrenparkinson/dnas)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/darrenparkinson/dnas)


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
floors, err := d.ActiveClientsService.ListFloors(context.Background())
```

Some API methods have optional parameters that can be passed, for example:

```go
d, _ := dnas.NewClient(apikey, region, nil)
opt := &dnas.ClientParameters{Associated: dnas.Bool(true), DeviceType: dnas.String("CLIENT")}
count, err := d.ActiveClientsService.GetCount(context.Background(), opt)
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

## Pagination

Where pagination is provided, Cisco provides the Page and Limit query parameters as part of the request parameters for a given endpoint. Cisco specifies these as strings, so the helper function `dnas.String` will be necessary:

```go
clients, err := c.ActiveClientsService.ListClients(ctx, &dnas.ClientParameters{Limit: dnas.String("1"), Page: dnas.String("1")})
```

By way of an example, you might use the following to work through multiple pages:

```go
count := 1
for {
    ac, err := c.ActiveClientsService.ListClients(ctx, &dnas.ClientParameters{Associated: dnas.Bool(true), DeviceType: dnas.String("CLIENT"), Limit: dnas.String("1"), Page: dnas.String(fmt.Sprint(count))})
    if err != nil {
        log.Fatal(err)
    }
    log.Println(len(ac.Results), ac.Results[0].MacAddress, ac.Results[0].IPAddress)
    count++
    if ac.MorePage == false {
        break
    }
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
count, err := c.ActiveClientsService.GetCount(ctx, &dnas.ClientParameters{})
if errors.Is(err, dnas.ErrUnauthorized) {
	log.Fatal("Sorry, you're not allowed to do that.")
}
```

# Roadmap

Currently this library only implements some of the functionality.  It is intended that this API will support all endpoints as they are required.  Feel free to log issues if there are specific endpoints you'd like, or see Contributing.  

The following oulines the available endpoints and their status in relation to implementation in this library.  As you can see, still a way to go.

## Map Service

| Method | Endpoint                      | Status          | Function      |
|--------|-------------------------------|-----------------|---------------|
| POST   | /map                          | Not Implemented |               |
| GET    | /map/hierarchy                | Implemented     | GetHierarchy  |
| GET    | /map/elements/{elementId}     | Implemented     | GetMapElement |
| DELETE | /map/elements/{elementId}     | Not Implemented |               |
| GET    | /map/images/floor/{imageName} | Not Implemented |               |

For Map Hierarchy, the `InclusionExclusionRegion` can have a variable number of vertices and additonal items, so this has been specified as a `[]map[string]interface{}`.  Clearly this isn't ideal, so if anyone has any better way to do this, I'd be glad to hear it.  In the meantime though, you'd have to access them something like this:

```go
fmt.Println(h.Map[0].RelationshipData.Children[0].RelationshipData.Children[0].Details.InclusionExclusionRegion[0]["type"])
fmt.Println(h.Map[0].RelationshipData.Children[0].RelationshipData.Children[0].Details.InclusionExclusionRegion[0]["vertices"])
```

## Active Clients Service

| Method | Endpoint        | Status      | Function    |
|--------|-----------------|-------------|-------------|
| GET    | /clients        | Implemented | ListClients |
| GET    | /clients/count  | Implemented | GetCount    |
| GET    | /clients/floors | Implemented | ListFloors  |

## Access Points Service

| Method | Endpoint            | Status      | Function         |
|--------|---------------------|-------------|------------------|
| GET    | /accessPoints       | Implemented | ListAccessPoints |
| GET    | /accessPoints/count | Implemented | GetCount         |

Note that `GetCount` accepts a status in order to return the count of access points for that given status.  For this purpose, the `dnas.AccessPointStatus` constant can be used and is one of: All, Active, Inactive, Missing, e.g:

```go
ac, err := c.AccessPointsService.GetCount(ctx, dnas.Inactive)
if err != nil {
    log.Fatal(err)
}
log.Printf("Inactive Access Points: %d\n", ac.Count)
```

Also note that `ListAccessPoints` only supports listing "missing" access points at this time as per the [Cisco documentation](https://developer.cisco.com/docs/dna-spaces/#!dna-spaces-location-cloud-api)

## Clients History Service

| Method | Endpoint                    | Status          |
|--------|-----------------------------|-----------------|
| GET    | /history                    | Not Implemented |
| GET    | /history/records/count      | Not Implemented |
| GET    | /history/clients            | Not Implemented |
| GET    | /history/clients/{deviceId} | Not Implemented |

## Notifications Service

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