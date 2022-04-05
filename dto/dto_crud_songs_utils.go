package dto

import "go.mongodb.org/mongo-driver/mongo"

var (
  DtoSongsParameters = map[string]string{
    "database": "songsdb",
    "collection": "songs",
  }
)

var  songsCollectionClient *mongo.Collection = nil

func GetSongsCollectionClient() *mongo.Collection{
  if songsCollectionClient == nil {
    songsCollectionClient = GetCollectionClient(DtoSongsParameters["database"], DtoSongsParameters["collection"])
  }
  return songsCollectionClient
}