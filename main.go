package main

import 
    (
    "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
    "strconv"
    "sort"
    )

var db *sql.DB
var err error
type Post struct{
       ID int `json:"id"`
       Title string `json:"title"`
       PosterPath string  `json:"poster_path"`
       ReleaseDate string   `json:"release_date"`
       Overview string    `json:"overview"`
       Key string `json:"key"`
}

func filter(r string,text string) bool{   
  
    for i := 0; i < len(r)-len(text)+1; i++ {
       var c=true
                for j := 0; j < len(text); j++ {
                      if(r[i+j]!=text[j] && r[i+j]!=text[j]+32 && r[i+j]!=text[j]-32) {
                          c=false
                          break
                       }
             }
            if(c) {return true}
        }
  
  return false
}

func getMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    params := mux.Vars(r)
    i, err := strconv.Atoi(params["id"])
   db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/Picture_Perfect")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

   var posts []Post
    result, err := db.Query("SELECT * FROM Movies WHERE id BETWEEN ? AND ?",i*12-11,i*12)
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath,&post.ReleaseDate,&post.Overview,&post.Key)
    if err != nil {
      panic(err.Error())
    }
    posts = append(posts, post)
  }
  
  json.NewEncoder(w).Encode(posts)
}

func getQuery(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    params := mux.Vars(r)
    k, err := strconv.Atoi(params["id"])
    query := params["query"]

   db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/Picture_Perfect")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

   var posts []Post
    result, err := db.Query("SELECT * FROM Movies ")

  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath,&post.ReleaseDate,&post.Overview,&post.Key)
    if err != nil {
      panic(err.Error())
    }
    if(filter(post.Title,query)){
    posts = append(posts, post)
  }
  }
  sort.SliceStable(posts, func(i, j int) bool {
    return posts[i].Title < posts[j].Title
})
   var res []Post
  for i := k*12-12; i < k*12 && i<len(posts); i++ {
    res =append(res,posts[i])
  }
  json.NewEncoder(w).Encode(res)
}


func main(){
    	
      
    	router := mux.NewRouter()
    	router.HandleFunc("/Home/{id}", getMovie).Methods("GET")
      router.HandleFunc("/Home/{query}/{id}", getQuery).Methods("GET")

        http.ListenAndServe(":8000", router)


    }