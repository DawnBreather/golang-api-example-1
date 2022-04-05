package config

import (
  "src/utils"
)

const (
  SHARD_INDEX_PLACEHOLDER = "{#}"
)

var configObject *config

func Get() *config{
  if configObject == nil {
    configObject = &config{
      Api: ApiConfig{
        Port: utils.GetEnv(PORT, "8080"),
        Socket: ":" + utils.GetEnv(PORT, "8080"),
      },
      Dto: MongoDbDtoConfig{
        Primary: PrimaryMongoDbDtoConfig{
          TimeoutSeconds:                utils.MustParseStringToInt32(utils.GetEnv(DTO_MONGODB_TIMEOUT_SECONDS, "10")),
          Schema:                        utils.GetEnv(DTO_MONGODB_CONNECTION_SCHEMA, "mongodb://"),
          SearchIndexNameForSongsSearch: utils.GetEnv(DTO_MONGODB_SEARCH_INDEX_NAME_FOR_SONGS_SEARCH, "artists_and_title"),
          DatabaseName:                  utils.GetEnv(DTO_MONGODB_DATABASE_NAME, "songsdb"),
          CollectionName:                utils.GetEnv(DTO_MONGODB_COLLECTION_NAME, "songs"),
          Username:                      utils.GetEnv(DTO_MONGODB_USERNAME, ""),
          Password:                      utils.GetEnv(DTO_MONGODB_PASSWORD, ""),
          ShardsNamesTemplate:           utils.GetEnv(DTO_MONGODB_SHARDS_NAMES_TEMPLATE, "cluster0-shard-00-"+SHARD_INDEX_PLACEHOLDER+".rusxw.mongodb.net:27017"),
          ShardsQuantity:                utils.MustParseStringToInt32(utils.GetEnv(DTO_MONGODB_SHARDS_QUANTITY, "3")),
        },
        Optional: OptionalMongoDbDtoConfig{
          RetryWrites: utils.GetEnv(DTO_MONGODB_CONNECTION_RETRY_WRITES, "true"),
          W:           utils.GetEnv(DTO_MONGODB_CONNECTION_W, "majority"),
          ReplicaSet:  utils.GetEnv(DTO_MONGODB_CONNECTION_REPLICA_SET, "atlas-zhqegh-shard-0"),
          Tls:         utils.GetEnv(DTO_MONGODB_CONNECTION_TLS, "true"),
        },
      },
    }
  }
  return configObject
}