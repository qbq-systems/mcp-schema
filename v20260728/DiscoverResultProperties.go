package v20260728

type DiscoverResultProperties struct {
	Meta              AdditionalProperties `json:"_meta"`
	CacheScope        LoggingLevel         `json:"cacheScope"`
	Capabilities      ClientResult         `json:"capabilities"`
	Instructions      Cursor               `json:"instructions"`
	ResultType        Cursor               `json:"resultType"`
	SupportedVersions SupportedVersions    `json:"supportedVersions"`
	TTLMS             Priority             `json:"ttlMs"`
}
