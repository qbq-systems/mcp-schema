package vdraft

type ImplementationProperties struct {
	Icons      Audience  `json:"icons"`
	Name       Cursor    `json:"name"`
	Title      Cursor    `json:"title"`
	Version    Default   `json:"version"`
	WebsiteURL BlobClass `json:"websiteUrl"`
}
