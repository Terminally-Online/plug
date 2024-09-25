package intent

type ProtocolModel struct {
	Slug   string // The protocol slug
	Url    string // The web-based URL the protocol is available
	Forked string // The protocol slug of the forked protocol -- This does not impact logic handling
}

/*
Protocols represents the list of supported protocols. Inside the file definition for each 
protocol, the actions are locally defined and implemented.
*/
var Protocols = []ProtocolModel{
	{
		Slug: "aave_v2",
		Url:  "https://aave.com",
	},
	{
		Slug: "aave_v3",
		Url:  "https://aave.com",
	},
	{
		Slug: "compound_v2",
		Url:  "https://compound.finance/",
	},
	{
		Slug:   "venus",
		Url:    "https://venus.io/",
		Forked: "compound_v2",
	},
}
