package main


type IpAPI struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	Countrycode string  `json:"countryCode"`
	Region      string  `json:"region"`
	Regionname  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}