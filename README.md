## Stack
* Golang 1.18+
* MongoDB Atlas

## Endpoints
### GET /songs
```shell
# heroku
curl -XGET "https://yousician-test.herokuapp.com/songs?page=1&limit=10 -v"

# localhost
curl -XGET "http://localhost:8081/songs?page=1&limit=10 -v"
```

### GET /songs/search
```shell
# heroku
curl -XGET -d "yousicians kennel" "https://yousician-test.herokuapp.com/songs/search?relevanceScoreBaseline=0.5 -v"

# localhost
curl -XGET -d "yousicians kennel" "http://localhost:8081/songs/search?relevanceScoreBaseline=0.5 -v"
```
### GET /songs/avg/difficulty
```shell
# heroku
curl -XGET "https://yousician-test.herokuapp.com/songs/avg/difficulty?level=13 -v"

# localhost
curl -XGET "http://localhost:8081/songs/avg/difficulty?level=13 -v"
```

### POST /song/rating
```shell
# heroku
curl -XPOST "https://yousician-test.herokuapp.com/song/rating?song_id={song_id}&rating=5 -v"

# localhost
curl -XPOST "http://localhost:8081/song/rating?song_id={song_id}&rating=5 -v"

```

### GET /song/rating/:song_id
```shell
# heroku
curl -XGET "https://yousician-test.herokuapp.com/song/rating/{song_id} -v"
# localhost
curl -XGET "http://localhost:8081/song/rating/{song_id} -v"
```

## Deployment steps
### Heroku
```shell
heroku login
heroku container:login
heroku container:push web -a yousician-test
heroku container:release web -a yousician-test
heroku logs --tail
```