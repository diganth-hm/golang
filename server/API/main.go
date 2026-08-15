package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"time"
	"strconv"
	"log"


)

//model file for this api
type Course struct{
    CourseId string      `json:"couseid"`
	CourseName string    `json:"coursename"`
	CoursePrice float64  `json:"courseprice"`
    Author *Author        `json:"author"`
}
type Author struct{
	Fullname string  `json:"fullname"`
	Website  string  `json:"website"`
}

//fake db
 var courses []Course 

 //middleware or helper

 func (c *Course) IsEmpty() bool{
	return  c.CourseName==""
 }


func main(){
   fmt.Println("Server runnig on http://localhost:5193")
   r:=mux.NewRouter()


   //seeding
   courses=append(courses,Course{CourseId:"3",
                                CourseName:"Golang development",
							    CoursePrice:354.55,
							    Author:&Author{Fullname:"Harkirat",
							    Website:"100xdev.com"}})
   courses=append(courses,Course{CourseId:"1",
                                CourseName:"web development",
							    CoursePrice:890.55,
							    Author:&Author{Fullname:"shraddha",
							    Website:"Apnacollege.com"}})

	//routing
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOnecourse).Methods("GET")
	r.HandleFunc("/course/new",createonecourse).Methods("POST")
	r.HandleFunc("/course/up/{id}",update).Methods("PUT")
	r.HandleFunc("/course/del/{id}",deleteonecourse).Methods("DELETE")

	//listening to a port
	log.Fatal(http.ListenAndServe(":5193",r))

}



 //controller file

 //serv home
func servHome(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Welcomme to home page of a build API</h1>"))
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
      if course.CourseId == params["id"]{
		json.NewEncoder(w).Encode(course)
		return
	  }
	  json.NewEncoder(w).Encode("No course found in the data base")
	}
}

func createonecourse(w http.ResponseWriter,r *http.Request){
    fmt.Println("Create one course ")
	w.Header().Set("content-type","application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please enter some data")
		return
	}

	var course Course
	_ =json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("no data inside the json file")
		return
	}
	//now gen a unique Id and appending it in to slices

	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses=append(courses,course)
	json.NewEncoder(w).Encode(course)
	return

}

func update(w http.ResponseWriter,r * http.Request){
	fmt.Println("updating a course")
	w.Header().Set("content-type","application/json")
    params:= mux.Vars(r)
    // for loop 
	for  index,course := range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No course found with the id you gave")
		return
		
	}

}

func deleteonecourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("delete one course")
	w.Header().Set("content-type","application/json")
    params:=mux.Vars(r)
	//looping
	for index,course := range courses{
		if course.CourseId==params["id"]{
		 courses=append(courses[:index],courses[index+1:]...)
		 json.NewEncoder(w).Encode("This course has been deleted")
		 break
			
		}
	}
}