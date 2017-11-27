# GoRedirecter

## Description

A simple web application written in Go, that accepts: `{ imagePath: string; hostName: string; }` as
a query parameters and returns an image

## Examples

```go
  res, err := http.Get(`/redirect/image?image=${imagePath}&host={hostName}`)

  if err != nil {
    // handle error here
  }

  // handle response here
```

## Authors

[Anatoly Belobrovik](https://github.com/AnatolyBelobrovik "Anatoly Belobrovik")
