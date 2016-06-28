package siadclient

// SiadClient is used to connect to siad
type SiadClient struct {
	siadurl string
}

// NewSiadClient creates a new SiadClient given a 'host:port' connectionstring
func NewSiadClient(connectionstring string) *SiadClient {
	s := SiadClient{}
	s.siadurl = "http://" + connectionstring
	return &s
}
