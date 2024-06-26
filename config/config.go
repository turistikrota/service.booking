package config

type MongoBooking struct {
	Host       string `env:"MONGO_BOOKING_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_BOOKING_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_BOOKING_USERNAME" envDefault:""`
	Password   string `env:"MONGO_BOOKING_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_BOOKING_DATABASE" envDefault:"empty"`
	Collection string `env:"MONGO_BOOKING_COLLECTION" envDefault:"empties"`
	Query      string `env:"MONGO_BOOKING_QUERY" envDefault:""`
}

type MongoInvite struct {
	Collection string `env:"MONGO_INVITE_COLLECTION" envDefault:"empties"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type Rpc struct {
	ListingHost    string `env:"RPC_LISTING_HOST" envDefault:"localhost:3001"`
	ListingUsesSsl bool   `env:"RPC_LISTING_USES_SSL" envDefault:"localhost:3001"`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Http struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type Topics struct {
	Booking BookingTopics
	Notify  NotifyTopics
}
type RSA struct {
	PrivateKeyFile string `env:"RSA_PRIVATE_KEY"`
	PublicKeyFile  string `env:"RSA_PUBLIC_KEY"`
}

type BookingTopics struct {
	PayPending   string `env:"STREAMING_TOPIC_BOOKING_PAY_PENDING"`
	PaySuccess   string `env:"STREAMING_TOPIC_BOOKING_PAY_SUCCESS"`
	PayTimeout   string `env:"STREAMING_TOPIC_BOOKING_PAY_TIMEOUT"`
	PayCancelled string `env:"STREAMING_TOPIC_BOOKING_PAY_CANCELLED"`
	Cancelled    string `env:"STREAMING_TOPIC_BOOKING_CANCELLED"`

	ValidationSuccess string `env:"STREAMING_TOPIC_BOOKING_VALIDATION_SUCCESS"`
	ValidationFail    string `env:"STREAMING_TOPIC_BOOKING_VALIDATION_FAIL"`
	ValidationStart   string `env:"STREAMING_TOPIC_BOOKING_VALIDATION_START"`
}

type NotifyTopics struct {
	SendMail string `env:"STREAMING_TOPIC_NOTIFY_EMAIL"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Booking MongoBooking
		Invite  MongoInvite
	}
	Rpc         Rpc
	Http        Http
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Session     Session
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	CacheRedis  CacheRedis
	RSA         RSA
}
