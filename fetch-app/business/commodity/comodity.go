package commodity

//Commodity product commodity that available to sell
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

//CommodityReport report commodity that access for admin
type CommodityReport struct {
	Province string
	Min      float64
	Max      float64
	Median   int
	Average  float64
}
