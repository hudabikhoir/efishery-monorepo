package commodity

//Commodity product commodity that available to rent or sell
type Commodity struct {
	UUID         string
	Commodity    string
	Province     string
	City         string
	Size         string
	Price        string
	ConvertPrice string
	ParsedAt     string
	Timestamp    string
}
