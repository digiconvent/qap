package story

const (
	masterOfBagEnd = "Master of Bag End"
	bagEnd         = "Bag End"
	council        = "The Council of Elrond"
	fellowship     = "The Fellowship"
	court          = "The Court of Edoras"
	ents           = "The Ents of Fangorn"
	KingsAuthority = "The King's authority"
	ringBearers    = "The Ring-bearers"
	parley         = "The parley at Orthanc"
)

var StartingMemberships = map[string][]string{
	"bilbo":      {masterOfBagEnd, bagEnd, ringBearers},
	"frodo":      {bagEnd, council, fellowship, ringBearers},
	"samwise":    {bagEnd, fellowship, ringBearers},
	"elrond":     {council},
	"gandalf":    {council, fellowship, parley},
	"aragorn":    {council, fellowship, parley},
	"legolas":    {council, fellowship},
	"gimli":      {council, fellowship},
	"boromir":    {council, fellowship},
	"merry":      {fellowship},
	"pippin":     {fellowship},
	"theoden":    {court, KingsAuthority, parley},
	"saruman":    {parley},
	"eomer":      {court},
	"eowyn":      {court},
	"wormtongue": {court},
	"treebeard":  {ents},
	"quickbeam":  {ents},
	"leaflock":   {ents},
	"skinbark":   {ents},
	"beechbone":  {ents},
}
