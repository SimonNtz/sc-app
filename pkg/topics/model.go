package topics

/* Matched to a handle model */

type Topic struct {
	Id 			string
	Lead 		string
	Corpus 		string
	OptsModel	Option
	/*Add other meta from the handle*/
}

/* Matched to a handle model are stored to 'Consents' User type field */

type Option struct {
	Radio 		bool `json:"radio"`
	Text 		bool `json:"text"`
	Rater 		bool `json:"rater"`
}
