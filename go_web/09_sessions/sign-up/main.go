package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("helloworld"), bcrypt.MinCost)
	dbUsers["khoi"] = user{"khoi", bs}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.html", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedin(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.html", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedin(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	var u user
	//process form submission
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(res, "username already taken", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un

		//store users in dbUsers\
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs}
		dbUsers[un] = u

		//redirect
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.html", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedin(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {

		//get form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		//username not matched?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "username and/or password do not match", http.StatusForbidden)
			return
		}

		//password not matched
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "username and/or password do not match", http.StatusForbidden)
			return
		}
		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
		dbSessions[c.Value] = un
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedin(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	//delete session
	delete(dbSessions, c.Value)
	//remove cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func getUser(res http.ResponseWriter, req *http.Request) user {
	//get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(res, c)

	//if user exist already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedin(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
