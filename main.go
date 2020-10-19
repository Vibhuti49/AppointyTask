package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "HomePage")
}

type Meeting struct{
    Id int `json: "id"`
    Title string `json:"title"`
    Participants int `json:"participants"`
    StartTime time `json: "starttime"`
    EndTime time `json: "endtime"`
    TimeStamp time `json: "timestamp"`
}
type Participant struct{
    Name string `json: "name"`
    Email string `json: "email"`
    RSVP string `json: "rsvp"`
}
var Meetings []Meeting
var Attendees []Participant

func returnAllMeetings(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(Meetings)
}
func returnAllAttendees(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(Attendees)
}
func handleMeetings() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/Meetings", returnAllMeetings)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleParticipants() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/attendees", returnAllAttendees)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
func main(){
     Meetings = []Meeting{
        Meeting{Id:1, Title: "Introduction",Participants:10,StartTime:(time.Date(2020, 10, 20, 10, 29, 59, 651387237, time.UTC)),EndTime:(time.Date(2020, 10, 20, 10, 29, 59, 651387237, time.UTC)),TimeStamp:(time.Now())},
        Meeting{Id:2, Title: "Brainstorming",Participants:10,StartTime:(time.Date(2020, 10, 21, 10, 29, 59, 651387237, time.UTC)),EndTime:(time.Date(2020, 10, 21, 10, 29, 59, 651387237, time.UTC)),TimeStamp:(time.Now())},
    }
    Attendees= []Participant{
        Participant{Name:"Abc",Email:"a@b.com",RSVP:"Yes"},
        Participant{Name:"Xyz",Email:"c@d.com",RSVP:"Yes"},
    }
    handleMeetings()
    handleParticipants()
}