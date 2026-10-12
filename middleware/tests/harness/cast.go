package harness

const (
	Fellowship    = "The Fellowship"
	WhiteCouncil  = "The White Council"
	CourtOfEdoras = "The Court of Edoras"
	Istari        = "Istari"
)

func memberships() map[string][]string {
	return map[string][]string{
		"gandalf":    {Istari},
		"aragorn":    {WhiteCouncil},
		"legolas":    {WhiteCouncil},
		"samwise":    {Fellowship},
		"frodo":      {Fellowship},
		"boromir":    {Fellowship},
		"theoden":    {CourtOfEdoras},
		"eomer":      {CourtOfEdoras},
		"wormtongue": {CourtOfEdoras},
		"eowyn":      {CourtOfEdoras},
	}
}
