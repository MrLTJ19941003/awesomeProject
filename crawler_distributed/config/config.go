package config

const (
	//Parser names
	ParserCityList = "ParseCityList"
	ParserCity     = "ParserCity"
	ParserProfile  = "ParserProfile"
	NilParser      = "NilParser"

	//Service Port
	ItemSaverPort = 8990
	ClawPort0     = 8991

	//ElasticSearch
	ElasticIndex     = "dating_profile"
	ElasticIndexTest = "dating_test"

	//RPC Endpoints
	ItemSaverRpc = "ItemSaverService.Save"
	CrawRpc      = "CrawlService.Process"
)
