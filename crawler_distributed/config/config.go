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
	//shijijiayuan获取用户的接口url
	GetUserUrl = "http://www.jiayuan.com/ajax/interested.php?r=0.32553824808613996&ad_param[]=pid%3Apersonalmatch_profile_new%7Ccount%3A6%7C"//cachesql%3A3600%7Csim_uid%3A187086135

)
