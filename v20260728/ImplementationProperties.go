package v20260728

type ImplementationProperties struct {
	Description Cursor    `json:"description"`
	Icons       Audience  `json:"icons"`
	Name        Cursor    `json:"name"`
	Title       Cursor    `json:"title"`
	Version     Cursor    `json:"version"`
	WebsiteURL  BlobClass `json:"websiteUrl"`
}
