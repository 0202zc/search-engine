// Constants is where you put your global const
package constants

const (
	//worker id
	WorkerID = 1

	//https
	certCrtPath = "crtPath" //should be modify
	certKeyPath = "KeyPath" //should be modify

	//keys
	SecretKey   = "123" //Should be modify
	IdentityKey = "id"

	/* json key
	 * search
	 */
	Time      = "time"
	Total     = "total"
	PageCount = "pagecount"
	Page      = "page"
	Limit     = "limit"
	Contents  = "contents"

	RelatedTexts = "relatedtexts"

	/*
	 * collection
	 */
	Name    = "name"
	Entry   = "entry"
	ColltID = "colltid"

	//micro service
	EtcdAddress           = "127.0.0.1:2379"
	UserServiceName       = "userModel"
	SearchServiceName     = "SearchApi"
	CollectionServiceName = "collectionModel"

	//mysql basic setup
	MySQLDefaultDSN     = "root:root@tcp(localhost:3310)/db_search?charset=utf8&parseTime=True&loc=Local"
	UserTableName       = "user"
	KeywordTableName    = "keyword"
	RecordTableName     = "record"
	CollectionTableName = "colletion"
)
