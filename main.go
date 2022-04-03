package main

import (
  "context"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
  "time"
)


func main() {
  serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
  clientOptions := options.Client().
    ApplyURI("mongodb://app:m8a5rITB8vIdsvnu@cluster0-shard-00-00.rusxw.mongodb.net:27017,cluster0-shard-00-01.rusxw.mongodb.net:27017,cluster0-shard-00-02.rusxw.mongodb.net:27017/?retryWrites=true&w=majority&replicaSet=atlas-zhqegh-shard-0&tls=true").
    SetServerAPIOptions(serverAPIOptions)
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    log.Fatal(err)
  }
  collection := client.Database("songsdb").Collection("songs")
  //match := bson.M{}


  // SEARCH SONGS WITH PAGINATION
  // GET /songs
  ///////////////////////////////
  //var limit int64 = 10
  //var page int64 = 1
  //
  //paginatedData, err := New(collection).Context(context.TODO()).Limit(limit).Page(page).Aggregate()
  //if err != nil {
  //  fmt.Printf("Unable to execute search query: %v\n", err)
  //}
  //
  //var aggProductList []map[string]interface{}
  //if paginatedData != nil {
  //  for _, raw := range paginatedData.Data {
  //    var product map[string]interface{}
  //    if marshallErr := bson.Unmarshal(raw, &product); marshallErr == nil {
  //      aggProductList = append(aggProductList, product)
  //    }
  //  }
  //}


  // SEARCH SONGS WITHOUT PAGINATION
  //
  ///////////////////////////////
  //paginatedData, err := New(collection).Context(context.TODO()).Aggregate()
  //if err != nil {
  //  fmt.Printf("Unable to execute search query: %v\n", err)
  //}
  //
  //var aggProductList []map[string]interface{}
  //if paginatedData != nil {
  //  for _, raw := range paginatedData.Data {
  //    var product map[string]interface{}
  //    if marshallErr := bson.Unmarshal(raw, &product); marshallErr == nil {
  //      aggProductList = append(aggProductList, product)
  //    }
  //  }
  //}

  // SEARCH ONLY LEVEL
  // GET /songs/avg/difficulty
  // https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
  ///////////////////////////////
  //opts := options.Find().SetProjection(bson.D{{"level", 1}, {"_id", 0}})
  //cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
  //if err != nil {
  //  fmt.Println(err)
  //}
  //var results[]bson.D
  //if err = cursor.All(context.TODO(), &results); err != nil {
  //  fmt.Println(err)
  //}
  //
  //var levelAverage float64 = 0.0
  //var sum int32 = 0
  //for _, result := range results{
  //  fmt.Println(result.Map()["level"])
  //  sum += result.Map()["level"].(int32)
  //}
  //levelAverage = float64(sum) / float64(len(results))
  //fmt.Printf("Average level: %.2f", levelAverage)


  // SEARCH SONGS BY MESSAGE
  // GET /songs/search
  // https://www.mongodb.com/docs/atlas/atlas-search/tutorial/run-query/#std-label-fts-tutorial-run-query
  // https://www.mongodb.com/basics/search-index
  // https://www.mongodb.com/docs/drivers/go/master/fundamentals/crud/read-operations/text/
  // https://www.mongodb.com/docs/manual/tutorial/text-search-in-aggregation/#text-score
  ///////////////////////////////
  //query := bson.D{
  //{"$search", bson.D{
  //{"index", "artists_and_title"},
  //{"text", bson.D{
  //{"query", "yousicians"},
  //{"path", bson.D{
  //  {"wildcard", "*"},
  //}},
  //}},
  //}},
  //}
  //
  //projection := bson.D{
  //{"$project", bson.D{
  //{"_id", 1},
  //{"title", 1},
  //{"artist", 1},
  //{"level", 1},
  //{"released", 1},
  //{"difficulty", 1},
  //{"ratings", 1},
  //{"score", bson.D {
  //  {"$meta", "searchScore"},
  //}},
  //}},
  //}
  //
  //match := bson.D{
  //{"$match", bson.D{
  //{"score", bson.D{
  //  {"$gt", 0.0},
  //}},
  //}},
  //}
  //
  //sort := bson.D{
  //{"$sort", bson.D{
  //{"score", -1},
  //{"artist", 1},
  //{"title", 1},
  //},
  //},
  //}
  //
  //var results[]bson.D
  //cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{query, projection, match, sort})
  //if err != nil {
  //fmt.Println(err)
  //}
  //cursor.All(context.TODO(), &results)
  //fmt.Println(results)



  //// ADD NEW EMPTY FIELD
  //// https://www.statology.org/mongodb-add-field/
  /////////////////////////////////////////////////
  //newField := bson.M{
  // "$set": bson.M{
  //   "ratings": []int{},
  // },
  //}
  //
  //updateResult, err := collection.UpdateMany(context.TODO(), bson.M{}, newField)
  //if err != nil {
  // fmt.Println(updateResult)
  //}
  //
  //fmt.Println(updateResult)







  //// ADD RATING
  //// https://stackoverflow.com/a/68219291/4265419
  //// https://stackoverflow.com/a/71270009/4265419
  //// https://www.statology.org/mongodb-add-field/
  /////////////////////////////////////////////////
  //objectId, _ := primitive.ObjectIDFromHex("6248567a523f53b20a48f2b1")
  //filterById := bson.M{
  //  "_id": objectId,
  //}
  //
  //push := bson.M{
  // "$push": bson.M{
  //   "ratings": 7,
  // },
  //}
  //
  //updateResult, err := collection.UpdateOne(context.TODO(), filterById, push, options.Update().SetUpsert(true))
  //if err != nil {
  // fmt.Println(err)
  //}
  //
  //fmt.Println(updateResult)



  //// SEARCH RATING METRICS BY ID
  //// GET /songs/search
  //// https://www.mongodb.com/docs/manual/reference/operator/aggregation/avg/
  //// https://www.mongodb.com/community/forums/t/mongodb-filter-an-array-by-id-using-aggregate/114182
  /////////////////////////////////
  //objectId, _ := primitive.ObjectIDFromHex("6248567a523f53b20a48f2b1")
  //match := bson.D{
  // {"$match", bson.D{
  //   {"_id", bson.D{
  //     {"$eq", objectId},
  //   }},
  // }},
  //}
  //
  //
  // projection := bson.D{
  // {"$project", bson.D{
  //   {"_id", 0},
  //   //{"title", 0},
  //   //{"artist", 0},
  //   //{"level", 0},
  //   //{"released", 0},
  //   //{"difficulty", 0},
  //   //{"ratings", 0},
  //   {"ratingsAvg", bson.D{
  //     {"$avg", "$ratings"},
  //   }},
  //   {"ratingsMin", bson.D{
  //     {"$min", "$ratings"},
  //   }},
  //   {"ratingsMax", bson.D{
  //     {"$max", "$ratings"},
  //   }},
  // }},
  //}
  //
  //var results[]bson.D
  //cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{match, projection})
  //if err != nil {
  // fmt.Println(err)
  //}
  //cursor.All(context.TODO(), &results)
  //fmt.Println(results)




}