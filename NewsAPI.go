package servapi

import (
	"context"
	"fmt"
	"github.com/HtLord/servmodel"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func CreateNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	err := r.ParseForm()
	if ErrorHandler(w, r, err) {
		return
	}

	news := servmodel.NewsWrapper(r.PostForm)
	_, err = GetColl("test1", "news").InsertOne(context.TODO(), news)
	if ErrorHandler(w, r, err) {
		return
	}

	CountReq()
}

func ReadNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	keeperId := r.FormValue("KeeperId")
	filter := bson.D{{"keeperId", keeperId}}

	cursor, err := GetColl("test1", "news").Find(context.TODO(), filter)
	if ErrorHandler(w, r, err) {
		return
	}

	var result []servmodel.News
	for cursor.Next(context.TODO()) {
		var news servmodel.News
		err = cursor.Decode(&news)
		if ErrorHandler(w, r, err) {
			return
		}
		result = append(result, news)
	}

	fmt.Println(result)
}

func UpdateNews() {

}

func DeleteNews() {

}
