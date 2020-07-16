package models

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Sess  *mgo.Session
	Users *mgo.Collection
	Sessions *mgo.Collection
	Tmpl	*template.Template
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request){
	items := []*Account{}

	err := h.Users.Find(bson.M{}).All(&items)
	if err != nil{
		log.Fatal(err)
	}

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Items	[]*Account
	}{
		Items: items,
	})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request){
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

}

func (h *Handler) AddUsers(w http.ResponseWriter,r *http.Request){
	newItem := bson.M{
		"_id":		bson.NewObjectId(),
		"username":	r.FormValue("username"),
		"password":	r.FormValue("password"),
	}
	err := h.Users.Insert(newItem)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Insert - Lastinserted Id:", newItem["_id"])
	http.Redirect(w,r,"/",http.StatusFound)
}

func (h *Handler) EditUsers(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]){
		http.Error(w, "bad id",500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])

	post := &Account{}
	err := h.Users.Find(bson.M{"_id": id}).One(&post)
	if err != nil{
		log.Fatal(err)
	}
	err = h.Tmpl.ExecuteTemplate(w,"edit.html",post)
	if err != nil{
		log.Fatal(err)
	}

}

func (h *Handler) UpdateUsers(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]){
		http.Error(w,"bad id",500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])
	post := Account{}
	err := h.Users.Find(bson.M{"_id": id}).One(&post)
	if err !=nil{
		fmt.Println(err.Error())
	}
	post.Username = r.FormValue("username")
	post.Password = r.FormValue("password")

	err = h.Users.Update(bson.M{
		"_id": id,
	}, bson.M{
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	})
	affected := 1
	if err == mgo.ErrNotFound{
		affected = 0
	}else if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println("Updated - Rows Affected", affected)
	http.Redirect(w,r,"/",http.StatusFound)
}

func (h *Handler) DeleteUsers(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]){
		http.Error(w,"bad id", 500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])

	err := h.Users.Remove(bson.M{"_id": id})
	affected := 1
	if err == mgo.ErrNotFound{
		affected =0
	} else if err != nil{
		fmt.Println(err.Error())
	}
	w.Header().Set("Content-Type","application-json")
	resp := `{"affected":}` + strconv.Itoa(int(affected)) + `}`
	w.Write([]byte(resp))
}

func (h *Handler) ListSessions(w http.ResponseWriter, r *http.Request) {
	items := []*Session{}
	err := h.Sessions.Find(bson.M{}).All(&items)
	if err != nil {
		log.Fatal(err)
	}

	err = h.Tmpl.ExecuteTemplate(w, "indexSessions.html", struct {
		Items	[]*Session
	}{
		Items: items,
	})
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}