package main

func test(){
  fmt.Println(dto.GetAllSongsWithPagination(1, 10))
  fmt.Println(dto.GetSongsAverageDifficultyByLevel(-1))
  fmt.Println(dto.GetSongsByMessage("yousicians kennel", 0.5))
  fmt.Println(dto.AppendRatingToSongById("6248567a523f53b20a48f2b1", 5))
  fmt.Println(dto.GetSongRatingById("6248567a523f53b20a48f2b1"))
}