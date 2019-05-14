package servapi

import (
	"context"
	"fmt"
	"github.com/HtLord/servmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var dbName = "test1"
var collName = "keeper"

func CreateKeeper(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseForm()
	if ErrorHandler(w, r, err) {
		fmt.Fprint(w, "403 form parsing")
		return
	}

	current := time.Now().Format(time.RFC3339Nano)
	r.PostForm.Add("ctime", current)
	r.PostForm.Add("utime", current)

	keeper := servmodel.KeeperWrapper(r.PostForm)
	fmt.Println(keeper)
	_, err = GetColl(dbName, collName).InsertOne(context.TODO(), keeper)
	if ErrorHandler(w, r, err) {
		fmt.Fprint(w, "503 db operation")
		return
	}
	fmt.Fprint(w, "200ok")
}

func ReadOneKeeper(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	filter := servmodel.KeeperFilterWrapper(r.PostForm)

	var keeper servmodel.Keeper
	err := GetColl(dbName, collName).FindOne(context.TODO(), filter).Decode(&keeper)
	if ErrorHandler(w, r, err) {
		return
	}

	err = JsonResponseWrapper(w, keeper)
	if ErrorHandler(w, r, err) {
		return
	}

}

func ReadOneKeeperById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var idFilter bson.M
	id, err := primitive.ObjectIDFromHex(r.PostForm.Get("_id"))
	if err == nil {
		idFilter = bson.M{"_id": id}
	}

	var keeper servmodel.Keeper
	err = GetColl(dbName, collName).FindOne(context.TODO(), idFilter).Decode(&keeper)
	if ErrorHandler(w, r, err) {
		return
	}

	err = JsonResponseWrapper(w, keeper)
	if ErrorHandler(w, r, err) {
		return
	}
}

func ReadMultiKeeper(w http.ResponseWriter, r *http.Request) {

	filter := bson.D{{"_id", "123"}}

	cursor, err := GetColl(dbName, collName).Find(context.TODO(), filter)

	var result []servmodel.Keeper
	for cursor.Next(context.TODO()) {
		var news servmodel.Keeper
		err = cursor.Decode(&news)
		if ErrorHandler(w, r, err) {
			return
		}
		result = append(result, news)
	}
}

func UpdateKeeper() {

}

func DeleteKeeper() {

}
