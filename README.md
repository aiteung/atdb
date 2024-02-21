# atdb

mongo db helper example

```go
var dbmongoinfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "iteung",
}
var mongocon = atdb.MongoConnect(dbmongoinfo)
atdb.InsertOneDoc(mongocon, "skt", surat1)
```
get data
```go
filter := bson.M{"username": email}
usercred := atdb.GetOneDoc[UserCred](Awanggamongoconn, "user", filter)
```

## Dev

set env

GOPROXY=proxy.golang.org

```sh
go get go.mongodb.org/mongo-driver

git tag v0.1.1
git push origin --tags
go list -m github.com/aiteung/atdb@v0.1.1
```

```go
filter := bson.M{"_id": bson.M{"$eq": objectId}}
	update := bson.M{
		"$set": bson.M{
			"rating":   rating.Rating,
			"komentar": rating.Komentar,
		},
	}
	atdb.UpdateDoc(config.HRISmongoconn, "uxlaporan", filter, update)
```