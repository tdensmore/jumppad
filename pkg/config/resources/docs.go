package resources

import "github.com/jumppad-labs/hclconfig/types"

// TypeDocs is the resource string for a Docs resource
const TypeDocs string = "docs"

// Docs allows the running of a Docusaurus container which can be used for
// online tutorials or documentation
type Docs struct {
	types.ResourceMetadata `hcl:",remain"`

	Networks []NetworkAttachment `hcl:"network,block" json:"networks,omitempty"` // Attach to the correct network // only when Image is specified

	Image *Image `hcl:"image,block" json:"image,omitempty"` // image to use for the container

	Logo   Logo   `hcl:"logo,block" json:"logo,omitempty"`
	Assets string `hcl:"assets,optional" json:"assets,omitempty"`

	Content []string `hcl:"content" json:"content"`

	Port          int  `hcl:"port,optional" json:"port"`
	OpenInBrowser bool `hcl:"open_in_browser,optional" json:"open_in_browser"` // When a host port is defined open the location in a browser

	// Output parameters

	// FQRN is the fully qualified resource name for the container, this can be used
	// to access the container from other sources
	FQRN string `hcl:"fqdn,optional" json:"fqdn,omitempty"`
}

type Logo struct {
	URL    string `hcl:"url" json:"url"`
	Width  int    `hcl:"width" json:"width"`
	Height int    `hcl:"height" json:"height"`
}

func (d *Docs) Process() error {
	// if port not set set port to 80
	if d.Port == 0 {
		d.Port = 80
	}

	// do we have an existing resource in the state?
	// if so we need to set any computed resources for dependents
	cfg, err := LoadState()
	if err == nil {
		// try and find the resource in the state
		r, _ := cfg.FindResource(d.ID)
		if r != nil {
			kstate := r.(*Docs)
			d.FQRN = kstate.FQRN
		}
	}

	return nil
}
