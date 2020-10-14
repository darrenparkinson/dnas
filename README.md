# DNAS - A simple Cisco DNA Spaces Go Library

[![GoDoc](https://godoc.org/github.com/darrenparkinson/dnas?status.svg)](https://godoc.org/github.com/darrenparkinson/dnas)
[![PkgGoDev](https://pkg.go.dev/badge/darrenparkinson/dnas)](https://pkg.go.dev/github.com/darrenparkinson/dnas)
[![Go Report Card](https://goreportcard.com/badge/github.com/darrenparkinson/dnas)](https://goreportcard.com/report/github.com/darrenparkinson/dnas)
[![published](https://static.production.devnetcloud.com/codeexchange/assets/images/devnet-published.svg)](https://developer.cisco.com/codeexchange/github/repo/darrenparkinson/dnas)


This repository is intended as a simple to use library for the Go language to interact with the [Cisco DNA Spaces API](https://developer.cisco.com/docs/dna-spaces).

DNA Spaces is "*an industry leading indoor location & IoT-as-a Service platform*" from Cisco.  You can find more information [on their site](https://dnaspaces.cisco.com/) and also [sign up for a free trial](https://dnaspaces.cisco.com/#earlyaccess).

In order to use this library, you must have access to DNA Spaces.  As of now, there is currently no sandbox for DNA Spaces and so you will either need an existing DNA Spaces tenant or you will need to sign up for a trial.  In addition, you will need [Go 1.13 or above](https://golang.org/).  

## Installing

You can install the library in the usual way as follows:

```sh
$ go get github.com/darrenparkinson/dnas
```

## Usage

In your code, import the library:

```go
import "github.com/darrenparkinson/dnas"
```

You can then construct a new DNA Spaces client, and use the various services on the client to access different parts of the DNA Spaces API.  For example:

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

* [Map Service](https://github.com/darrenparkinson/dnas#map-service)
* [Active Clients Service](https://github.com/darrenparkinson/dnas#active-clients-service)
* [Access Points Service](https://github.com/darrenparkinson/dnas#access-points-service)
* [Clients History Service](https://github.com/darrenparkinson/dnas#clients-history-service)
* [Notifications Service](https://github.com/darrenparkinson/dnas#notifications-service)

NOTE: Using the context package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then context.Background() can be used as a starting point.

## Examples

There are some examples of usage in the [examples](examples) folder.  To run these, `git clone` the repository and and run them from the top level folder, e.g.:

```sh
$ go run examples/count/main.go
$ go run examples/floors/main.go
$ go run examples/clients/main.go
$ go run examples/history/main.go
```

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

In addition, the error is wrapped in with the original response from Cisco, e.g:

```
2020/10/13 09:00:00 dnas: internal error: There are 155478 records in request time range, more than the limit 50000. Please reduce the time range.
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

| Method | Endpoint                    | Status      | Function    |
|--------|-----------------------------|-------------|-------------|
| GET    | /history                    | Implemented | GetHistory  |
| GET    | /history/records/count      | Implemented | GetCount    |
| GET    | /history/clients            | Implemented | ListClients |
| GET    | /history/clients/{deviceId} | Implemented | GetClient   |

Note that `GetHistory` uses the `/history` api endpoint which returns CSV data that is converted to a struct.  Please note the restrictions on that API:

> *Retrieve small amount clients history to csv format. If startTime and endTime is not given, the time period is last 24 hours. If records amount is more than 50K, the user receives error response and indicate the time range needs to be reduced.*

This error response will be delivered as an `ErrInternalError` with the supplied message from Cisco.  You can check this with:

```go
records, err := c.HistoryService.GetHistory(ctx, &dnas.HistoryParameters{StartTime: dnas.String("1602576000000"), EndTime: dnas.String("1602662400000")})
if errors.Is(err, dnas.ErrInternalError) {
	...
}
```

Note that you can also use Go "time" to provide the times:

```go
fromTime := time.Now().Add(time.Hour*-2).UnixNano() / int64(time.Millisecond)
toTime := time.Now().UnixNano() / int64(time.Millisecond)
history, err := c.HistoryService.GetHistory(ctx,
    &dnas.HistoryParameters{
		StartTime: dnas.String(strconv.FormatInt(fromTime, 10)),
        EndTime:   dnas.String(strconv.FormatInt(toTime, 10)),
    })
```


An example of using history count:

```go
h, _ := c.HistoryService.GetCount(ctx, &dnas.HistoryCountParameters{FloorID: dnas.String("123467890abcdef")})
log.Printf("%+v\n", h)
```

List client history for the last 24 hours (note: this only provides a list of mac addresses):

```go
h, _ := c.HistoryService.ListClients(ctx, &dnas.HistoryClientsParameters{Ssid: dnas.String("YourSSID")})
log.Printf("%+v\n", h)
```

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