package router

import (
	"encoding/json"
	"net/http"
	
	"m/models"
)

type Video struct{
	ID   uint32
	Name string
	Desc string
	Size uint32
}
//db
var DBVid []Video


func TheRouter(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case "GET":
			GetHanlder(w, r)
		case "POST":
			PostVid(w, r)
		case "DELETE":
			DeleteVid(w, r)
		case "PUT":
			EditVid(w, r)

	}
}


//handle get requests
func GetHanlder(w http.ResponseWriter, r *http.Request){
	if r.URL.Query()["name"] == nil{
		GetVids(w, r)
	}else {
		GetAVid(w, r)
	}
	
}

func GetVids(w http.ResponseWriter, r *http.Request){
	results := models.GetAllVids()
	json.NewEncoder(w).Encode(results)
}

func GetAVid(w http.ResponseWriter, r *http.Request){
	vidName := r.URL.Query()["name"][0]
	
	results := models.GetSpecificVid(vidName)
	json.NewEncoder(w).Encode(results)
}



func PostVid(w http.ResponseWriter, r *http.Request){
	//get the json and decode it
	var vid models.Video
	json.NewDecoder(r.Body).Decode(&vid)

	models.InsertAVid(&vid)
	//return what you added
	w.Write([]byte("Video added successfully"))
	
}

func DeleteVid(w http.ResponseWriter, r *http.Request){
	if r.URL.Query()["name"] == nil{
		w.Write([]byte("Please provide an Video name"))
	}else {
		vidName := r.URL.Query()["name"][0]
		models.DeleteAVid(vidName)

		w.Write([]byte("Video deleted"))
		

	}

}

func EditVid(w http.ResponseWriter, r *http.Request){
		//encode the body
		var newvid models.Video
		json.NewDecoder(r.Body).Decode(&newvid)
		models.UpdateVid(newvid.Name, newvid.Desc, newvid.Size)

}
