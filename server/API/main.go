package main
import (
	"fmt"
	"net/http"
	"encodeibg/json"
	"github.com/gorilla/mux"

)

//model file for this api
type Course struct{
    CourseId string      `json:"couseid"`
	CourseName string    `json:"coursename"`
	CoursePrice float64  `json:"courseprice"`
    Author *Author        `jaon:"author"`
}
type Author struct{
	Fullname string  `json:"fullname"`
	Website  string  `json:"website"`
}

//fake db
 var courses []Course 

 //middleware or helper

 func (c *Course) IsEmpty() bool{
	return c.CouseID=="" && c,CourseName==""
 }


func mian(){
   


}



 //controller file
func servHome(w http.ResponseWriter,r *http.Request){
	w.write([]byte("<h1>Welcomme to home page of a build API</h1>"))
 }
func getAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("Get all course deatails")
	w.Header().Set("content-typr","application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOnecourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("get one course")
	w.Header().Set("content-type","application/json")
	//using gorilla mux
	params := mux.Vars(r)
	//for loop of courses to find the matching couse id
	for _,course := range courses{
      if course.CourseId == params["ID"]{
		json.NewEncoder(w).Encode(course)
		return
	  }
	  json.NewEncoder(w).Encode("No course found in the data base")
	}
}
