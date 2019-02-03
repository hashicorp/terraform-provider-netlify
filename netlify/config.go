package netlify

import (
	"fmt"
	"net/url"

	"github.com/go-openapi/runtime"
	openapiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/netlify/open-api/go/plumbing"
)

type Config struct {
	Token   string
	BaseURL string
}

// Meta is the returned meta struct.
type Meta struct {
	Netlify  *plumbing.Netlify
	AuthInfo runtime.ClientAuthInfoWriter
}

// Client configures and returns a fully initialized NetlifyClient
func (c *Config) Client() (interface{}, error) {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error parsing base_url: %s", err)
	}

	if u.Scheme == "" {
		// Default to https if no protocol is specified
		u.Scheme = "https"
	}

	// Create the OpenAPI client with our custom roundtripper.
	client := openapiClient.NewWithClient(
		u.Host, u.Path, []string{u.Scheme},
		cleanhttp.DefaultClient())
	client.Transport = logging.NewTransport("Netlify", client.Transport)

	// Setup our auth
	authInfo := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("User-Agent", "Terraform")
		r.SetHeaderParam("Authorization", "Bearer "+c.Token)
		return nil
	})

	return &Meta{
		Netlify:  plumbing.New(client, strfmt.Default),
		AuthInfo: authInfo,
	}, nil
}
