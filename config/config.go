package config

const (
	//Service ports
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"

	ParseCarDetail = "ParseCarDetail"
	ParseCarList   = "ParseCarList"
	ParseCarModel  = "ParseCarModel"

	NilParser = "NilParser"

	// ElasticSearch
	ElasticIndex_car = "car_profile"
	ElasticIndex     = "dating_profile"

	// Rate limiting
	Qps = 2

	//Rpc Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
