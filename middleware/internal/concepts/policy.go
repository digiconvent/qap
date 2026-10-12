package concepts

type Policy struct {
	Id                string
	Name              string
	Groups            []string
	FavourThreshold   int
	AgainstThreshold  *int
	Quorum            *int
	DeadlineSeconds   *int
	ExtensionSeconds  *int
	WithholdPermitted bool
	CommentsInAudit   bool
	Anonymous         bool
}

type PolicyReference struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
