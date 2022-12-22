package dadata

import (
	"net/url"

	"github.com/drgi/dadata-client/api/clean"
	"github.com/drgi/dadata-client/api/profile"
	"github.com/drgi/dadata-client/api/stat"
	"github.com/drgi/dadata-client/api/suggest"
	"github.com/drgi/dadata-client/client"
)

const (
	// EndpointURL is a base API endpoint.
	EndpointURL = "https://dadata.ru/api/v2/"
	// EndpointURLSuggest is a suggestion API endpoint.
	EndpointURLSuggest = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/"
)

var (
	endpointURL        *url.URL
	endpointURLSuggest *url.URL
)

func init() {
	var err error

	endpointURL, err = url.Parse(EndpointURL)
	if err != nil {
		panic(err)
	}

	endpointURLSuggest, err = url.Parse(EndpointURLSuggest)
	if err != nil {
		panic(err)
	}
}

// NewCleanApi provides "clean" API.
func NewCleanApi(opts ...client.Option) *clean.Api {
	return &clean.Api{
		Client: client.NewClient(endpointURL, opts...),
	}
}

// NewSuggestApi provides suggestion API.
func NewSuggestApi(opts ...client.Option) *suggest.Api {
	return &suggest.Api{
		Client: client.NewClient(endpointURLSuggest, opts...),
	}
}

// NewProfileApi provides profile related API.
func NewProfileApi(opts ...client.Option) *profile.Api {
	return &profile.Api{
		Client: client.NewClient(endpointURL, opts...),
	}
}

// NewStatApi provides statistic API.
func NewStatApi(opts ...client.Option) *stat.Api {
	return &stat.Api{
		Client: client.NewClient(endpointURL, opts...),
	}
}
