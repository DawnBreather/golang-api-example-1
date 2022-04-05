package config

const (
  PORT = "PORT"
)

const (
  DTO_MONGODB_TIMEOUT_SECONDS                     = "DTO_MONGODB_TIMEOUT_SECONDS"
  DTO_MONGODB_CONNECTION_SCHEMA                   = "DTO_MONGODB_CONNECTION_SCHEMA"
  DTO_MONGODB_SEARCH_INDEX_NAME_FOR_SONGS_SEARCH  = "DTO_MONGODB_SEARCH_INDEX_NAME_FOR_SONGS_SEARCH"
  DTO_MONGODB_DATABASE_NAME                       = "DTO_MONGODB_DATABASE_NAME"
  DTO_MONGODB_COLLECTION_NAME                     = "DTO_MONGODB_COLLECTION_NAME"
  DTO_MONGODB_USERNAME                            = "DTO_MONGODB_USERNAME"
  DTO_MONGODB_PASSWORD                            = "DTO_MONGODB_PASSWORD"
  DTO_MONGODB_SHARDS_NAMES_TEMPLATE               = "DTO_MONGODB_SHARDS_NAMES_TEMPLATE"
  DTO_MONGODB_SHARDS_QUANTITY                     = "DTO_MONGODB_SHARDS_QUANTITY"
  DTO_MONGODB_CONNECTION_RETRY_WRITES             = "DTO_MONGODB_CONNECTION_RETRY_WRITES"
  DTO_MONGODB_CONNECTION_W                        = "DTO_MONGODB_CONNECTION_W"
  DTO_MONGODB_CONNECTION_REPLICA_SET              = "DTO_MONGODB_CONNECTION_REPLICA_SET"
  DTO_MONGODB_CONNECTION_TLS                      = "DTO_MONGODB_CONNECTION_TLS"
)

type config struct {
  Api ApiConfig
  Dto MongoDbDtoConfig
}

type ApiConfig struct {
  Port string `yaml:"PORT"`
  Socket string `yaml:"ApiSocket"`
}

type MongoDbDtoConfig struct {
  Primary  PrimaryMongoDbDtoConfig
  Optional OptionalMongoDbDtoConfig
}

type PrimaryMongoDbDtoConfig struct {
  TimeoutSeconds                int    `yaml:"DTO_MONGODB_TIMEOUT_SECONDS"`
  Schema                        string `yaml:"DTO_MONGODB_CONNECTION_SCHEMA"`
  SearchIndexNameForSongsSearch string `yaml:"DTO_MONGODB_SEARCH_INDEX_NAME_FOR_SONGS_SEARCH"`
  DatabaseName                  string `yaml:"DTO_MONGODB_DATABASE_NAME"`
  CollectionName                string `yaml:"DTO_MONGODB_COLLECTION_NAME"`
  Username                      string `yaml:"DTO_MONGODB_USERNAME"`
  Password                      string `yaml:"DTO_MONGODB_PASSWORD"`
  ShardsNamesTemplate           string `yaml:"DTO_MONGODB_SHARDS_NAMES_TEMPLATE"`
  ShardsQuantity                int    `yaml:"DTO_MONGODB_SHARDS_QUANTITY"`
}


type OptionalMongoDbDtoConfig struct {
  RetryWrites string `yaml:"DTO_MONGODB_CONNECTION_RETRY_WRITES"`
  W           string `yaml:"DTO_MONGODB_CONNECTION_W"`
  ReplicaSet  string `yaml:"DTO_MONGODB_CONNECTION_REPLICA_SET"`
  Tls         string `yaml:"DTO_MONGODB_CONNECTION_TLS"`
}
