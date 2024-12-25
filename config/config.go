package config

type Config struct {
	LogLevel         string `env:"LOG_LEVEL"`
	Environment      string `env:"ENVIRONMENT"`
	ServiceName      string `env:"SERVICE_NAME"`
	ServiceVersion   string `env:"SERVICE_VERSION"`
	BaseURL          string `env:"BASE_URL"`
	HTTPPort         int    `env:"HTTP_PORT"`
	GRPCPort         int    `env:"GRPC_PORT"`
	MetricPort       int    `env:"METRIC_PORT"`
	InternalAPIKey   string `env:"INTERNAL_API_KEY"`
	OMSWebServiceURL string `env:"OMS_WEB_SERVICE_URL"`

	Redis    *Redis
}

type Redis struct {
	ConnectionURL       string   `env:"REDIS_CONNECTION_URL"`
	Password            string   `env:"REDIS_PASSWORD" json:"-"`
	DB                  int      `env:"REDIS_DB"`
	PoolSize            int      `env:"REDIS_POOL_SIZE"`
	DialTimeoutSeconds  int      `env:"REDIS_DIAL_TIMEOUT_SECONDS"`
	ReadTimeoutSeconds  int      `env:"REDIS_READ_TIMEOUT_SECONDS"`
	WriteTimeoutSeconds int      `env:"REDIS_WRITE_TIMEOUT_SECONDS"`
	IdleTimeoutSeconds  int      `env:"REDIS_IDLE_TIMEOUT_SECONDS"`
	MaxIdleConn         int      `env:"REDIS_MAX_IDLE_CONN_NUMBER"`
	MaxActiveConn       int      `env:"REDIS_MAX_ACTIVE_CONN_NUMBER"`
}

type Security struct {
	AesKey string `env:"SECURITY_AES_KEY" json:"-"`
}

type ES struct {
	Addrs                       []string `env:"ES_ADDRS"`
	User                        string   `env:"ES_USER"`
	Password                    string   `env:"ES_PASSWORD" json:"-"`
	IndexSearchOrder            string   `env:"ES_INDEX_SEARCH_ORDER"`
	IndexSearchReturnOrder      string   `env:"ES_INDEX_SEARCH_RETURN_ORDER"`
	IndexSearchShipment         string   `env:"ES_INDEX_SEARCH_SHIPMENT"`
	IndexSearchTicket           string   `env:"ES_INDEX_SEARCH_TICKET"`
	IndexSearchReturnShipment   string   `env:"ES_INDEX_SEARCH_RETURN_SHIPMENT"`
	IndexReportShipmentSKU      string   `env:"ES_INDEX_REPORT_SHIPMENT_SKU"`
	IndexSearchPayment          string   `env:"ES_INDEX_SEARCH_PAYMENT"`
	IndexSearchTrackingComplain string   `env:"ES_INDEX_SEARCH_TRACKING_COMPLAIN"`
}
