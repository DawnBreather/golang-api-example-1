package dto

type SongsWithPaginationResult struct {
  Songs []map[string]interface{}
  Pagination songsWithPaginationPaginationStruct
}

type songsWithPaginationPaginationStruct struct{
  PageNumber int64
  PerPageLimit int64
}

// SEARCH SONGS WITH PAGINATION
// GET /songs
///////////////////////////////
func GetAllSongsWithPagination(pageNumber, perPageLimit int64) (songs SongsWithPaginationResult, ok bool){
  if songsData, ok := ReadDocumentsWithPagination(GetSongsCollectionClient(), pageNumber, perPageLimit); ok {
    return SongsWithPaginationResult{
      Songs: songsData,
      Pagination: songsWithPaginationPaginationStruct{
        pageNumber,
        perPageLimit,
      },
    }, ok
  }

  return
}