package v20260728

type CacheableResultProperties struct {
	Meta       AdditionalProperties `json:"_meta"`
	CacheScope LoggingLevel         `json:"cacheScope"`
	ResultType Cursor               `json:"resultType"`
	TTLMS      Priority             `json:"ttlMs"`
	Contents   *Contents            `json:"contents,omitempty"`
}
