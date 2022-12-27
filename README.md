# Client for DaData.ru

Forked from github.com/ekomobile/dadata/v2.

DaData API v2

Implemented [Clean](https://dadata.ru/api/clean/) and [Suggest](https://dadata.ru/api/suggest/) methods.

## Installation

`go get github.com/drgi/dadata-client`

## Configuration

### Credentials

`DADATA_API_KEY` and `DADATA_SECRET_KEY` environment variables are used by default to authenticate client.

Custom credential provider may be used by implementing `client.CredentialProvider` interface.

Also, there is a "simple" credential provider `client.Credentials` you may utilize.

```go
creds := client.Credentials{
    ApiKeyValue:    "<YOUR_API_KEY>",
    SecretKeyValue: "<YOUR_SECRET_KEY>",
}

api := NewSuggestApi(client.WithCredentialProvider(&creds))
```

### HTTP client

HTTP client may be overridden with custom one:

```go
httpClient := &http.Client{}

api := NewSuggestApi(WithHttpClient(httpClient))
```

## Licence

MIT see [LICENSE](LICENSE)
