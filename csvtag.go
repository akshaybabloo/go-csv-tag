package csvtag

// CsvOptions - options when loading or dumping csv.
type CsvOptions struct {
	Separator rune
	Header    []string
	TagKey    string
}

const DefaultTagKey string = "csv"
