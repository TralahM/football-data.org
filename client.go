package football_data

import (
	"log"
	"net/http"
)

// A Client sends http.Requests and returns http.Responses or errors in case of
// failure.
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc is a function tyoe that implements the Client Interface.
type ClientFunc func(*http.Request) (*http.Response, error)

func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// A Decorator wraps a Client with extra behaviour

type Decorator func(Client) Client

// Decorate decorates a Client c with all the given Decorators in order.
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c

	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func Logging(l *log.Logger) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

// Authorization returns a Decorator that authorizes every Client request with
// the given token.
func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

// Header returns a Decorator that adds the given HTTP header to every request
// done by a Client
func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

// RawString is a raw encoded JSON object. it impements Marshaler and
// Unmarshaler and can
// be used to delay JSON decoding or precompute a JSON encoding.

// Base URL
var baseUrl string = "https://api.football-data.org"
var responseHeaders map[string]interface{} = map[string]interface{}{
	"X-API-Version":               string("v2"),   // indicates the version you are using
	"X-Authenticated-Client":      string("User"), // Shows the detected API-client or 'anonymous'
	"X-RequestCounter-Reset":      23,             // Defines the seconds left to reset your request counter.
	"X-Requests-Available-Minute": 28,             // Shows the remaining requests before being blocked.
}

func NewAPIClient(XAuthToken string) APIClient {
	client := Decorate(http.DefaultClient,
		Header("Accept", "application/json;charset=utf8"),
		Header("X-Auth-Token", XAuthToken),
	)
	return APIClient{client: client}
}
