package main

import 
    ( //"Components"
    "github.com/gorilla/mux"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
    "strconv"
    "sort"
   "fmt"
   "io/ioutil"
    )

var db *sql.DB
var err error
type Post struct{
       ID int `json:"id"`
       Title string `json:"title"`
       PosterPath string  `json:"poster_path"`
       
}
type Reviews struct{
  User string `json:"user"`
  Review string `json:"review"`
}
type NewReviews struct{
  ID int `json:"id"`
  User string `json:"user"`
  Review string `json:"review"`
  Rating float64 `json:"rating"`

}
type WholePost struct{
       ID int `json:"id"`
       Title string `json:"title"`
       PosterPath string  `json:"poster_path"`
       ReleaseDate string   `json:"release_date"`
       Overview string    `json:"overview"`
       Key string `json:"key"`
       Reviews []Reviews `json:"reviews"`
       Rating float64 `json:"rating`
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

func getMovies(w http.ResponseWriter, r *http.Request) {
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
    result, err := db.Query("SELECT id , title ,PosterPath FROM Movies WHERE id BETWEEN ? AND ?",i*12-11,i*12)
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath)
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
    result, err := db.Query("SELECT id , title ,PosterPath FROM Movies ")

  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath)
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

func getMovie(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    params := mux.Vars(r)
   // i, err := strconv.Atoi(params["id"])
   db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/Picture_Perfect")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()



     result, err := db.Query("SELECT * FROM Movies WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var post WholePost
  for result.Next() {
     err := result.Scan(&post.ID, &post.Title,&post.PosterPath,&post.ReleaseDate,&post.Overview,&post.Key)
    if err != nil {
      panic(err.Error())
    }
  }
    co:=0.0
   po:=0.0
   results, err := db.Query("SELECT Name,Review,Rating FROM MovieReviews WHERE MovieID = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer results.Close()
  for results.Next() {
     var s Reviews
     var t float64
     err := results.Scan(&s.User,&s.Review,&t)
    if err != nil {
      panic(err.Error())
    }
    post.Reviews=append(post.Reviews,s)
     po+=t
     co++; 
  }
  
   
  if(co!=0){
     r:=(po/co)
    post.Rating=r;
   }

     json.NewEncoder(w).Encode(post)
    }


func getTVs(w http.ResponseWriter, r *http.Request) {
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
    result, err := db.Query("SELECT id , title ,PosterPath FROM TVseries WHERE id BETWEEN ? AND ?",i*12-11,i*12)
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath)
    if err != nil {
      panic(err.Error())
    }
    posts = append(posts, post)
  }
  
  json.NewEncoder(w).Encode(posts)
}

func getTVQuery(w http.ResponseWriter, r *http.Request) {
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
    result, err := db.Query("SELECT id , title ,PosterPath FROM TVseries ")

  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Title,&post.PosterPath)
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

func getTV(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    params := mux.Vars(r)
   // i, err := strconv.Atoi(params["id"])
   db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/Picture_Perfect")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()



     result, err := db.Query("SELECT * FROM TVseries WHERE id = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  var post WholePost
  for result.Next() {
     err := result.Scan(&post.ID, &post.Title,&post.PosterPath,&post.ReleaseDate,&post.Overview,&post.Key)
    if err != nil {
      panic(err.Error())
    }
  }

   co:=0.0
   po:=0.0
   results, err := db.Query("SELECT Name,Review,Rating FROM TVReviews WHERE MovieID = ?", params["id"])
  if err != nil {
    panic(err.Error())
  }
  defer results.Close()
  for results.Next() {
     var s Reviews
     var t float64
     err := results.Scan(&s.User,&s.Review,&t)
    if err != nil {
      panic(err.Error())
    }
    post.Reviews=append(post.Reviews,s)
     po+=t
     co++; 
  }
  
   
  if(co!=0){
     r:=(po/co)
    post.Rating=r;
   }
     json.NewEncoder(w).Encode(post)
    }  

func createPost(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
   w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
  var Review = NewReviews{}
  post, _ := ioutil.ReadAll(r.Body)
  err := json.Unmarshal(post, &Review)
    if err == nil {
                 fmt.Println(Review.User)
                 db, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/Picture_Perfect")
                 if err != nil {
                 panic(err.Error())
                  }
                  defer db.Close()
                  stmt, err := db.Prepare("INSERT INTO MovieReviews(MovieID, Username , Name,Review,Rating) VALUES( ?, ?, ?, ? ,? )")
                 if err != nil {
                  panic(err.Error())
                  }
                 defer stmt.Close() 
                  if _, err := stmt.Exec(Review.ID,"KRG",Review.User,Review.Review,Review.Rating); err != nil {
                  panic(err.Error())
    }
 
 }else {
       fmt.Printf( "Error in Decoding ")
    }
    json.NewEncoder(w).Encode(post)
}  

func main(){
    	

    	router := mux.NewRouter()
    router.HandleFunc("/movies/{id}", getMovies).Methods("GET")
      router.HandleFunc("/movies/{query}/{id}", getQuery).Methods("GET")
       router.HandleFunc("/movie/{id}", getMovie).Methods("GET")
       router.HandleFunc("/tvs/{id}", getTVs).Methods("GET")
      router.HandleFunc("/tvs/{query}/{id}", getTVQuery).Methods("GET")
       router.HandleFunc("/tv/{id}", getTV).Methods("GET")
        router.HandleFunc("/data", createPost).Methods("POST")


        http.ListenAndServe(":8000", router)


    }
