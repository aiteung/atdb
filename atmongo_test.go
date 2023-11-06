package atdb

import (
	"fmt"
	"testing"
	"time"
)

type PertemuanBimbingan struct {
	Tipe    string    `bson:"tipe_bimbingan" json:"Tipe"`
	Penilai string    `bson:"penilai" json:"Penilai"`
	Log     string    `bson:"Log,omitempty" json:"Log"`
	Tanggal time.Time `bson:"tanggal" json:"Tanggal"`
}

func TestGetRandomDoc(t *testing.T) {

	db := MongoConnect(DBInfo{DBString: "MongoString", DBName: "bimbingan"})
	result, err := GetRandomDoc[PertemuanBimbingan](db, "pertemuan", 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", result)
}
