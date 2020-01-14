package newrelic

const (
	// v2RootUrl is the root url for v2 of the NewRelic api
	v2RootURL = "https://api.newrelic.com/v2"
)

// NewRelic the root NewRelic object
type NewRelic struct {
	apiKey  string
	version string
}

// NewService creates a new instance of NewRelic
func New(apiKey string, version string) *NewRelic {
	nr := &NewRelic{
		apiKey:  apiKey,
		version: version,
	}

	return nr
}
