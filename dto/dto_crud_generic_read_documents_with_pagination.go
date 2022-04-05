package dto

import (
  "context"
  "fmt"
  "github.com/gobeam/mongo-go-pagination"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
)

func ReadDocumentsWithPagination(collectionClient *mongo.Collection, pageNumber int64, documentsPerPageLimit int64) (documentList []map[string]interface{}, ok bool) {
  paginatedData, err := mongopagination.New(collectionClient).Context(context.TODO()).Limit(documentsPerPageLimit).Page(pageNumber).Aggregate()
  if err != nil {
   fmt.Printf("Unable to execute search query: %v\n", err)
   return
  }

  var documentsList []map[string]interface{}
  if paginatedData != nil {
   for _, raw := range paginatedData.Data {
     var document map[string]interface{}
     if marshallErr := bson.Unmarshal(raw, &document); marshallErr == nil {
       documentsList = append(documentsList, document)
     }
   }
  }

  return documentsList, true
}