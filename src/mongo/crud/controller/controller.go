package controller

import (
	"encoding/json"
	"fmt"
	"gofinal/src/src/mongo/crud/model"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoController struct {
	session *mgo.Session
}

func GetMongoController(s *mgo.Session) *MongoController {
	return &MongoController{session: s}
}
func (m MongoController) Get(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
		<header>
			<meta charset="UTF-8">
			<title>Index</title>
		</header>
		<body>
			<a href="/users/766"> GO TO </a>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html; charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
}
func (m MongoController) Create(rs http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{}
	u.Id = bson.NewObjectId()
	u.Age, _ = strconv.Atoi(r.FormValue("Age"))
	u.Name = r.FormValue("Name")
	u.Gender = r.FormValue("Gender")
	//err := json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("hiii", r.FormValue("Name"))
	fmt.Println("hiii", u)
	m.session.DB("goweb").C("users").Insert(u)

	uj, _ := json.MarshalIndent(u, "", "\n")
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusCreated)
	rs.Write(append([]byte("user created:\n"), uj...))
}

func (m MongoController) Update(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
		<header>
			<meta charset="UTF-8">
			<title>Index</title>
		</header>
		<body>
			<form action="/create" method="POST" enctype='application/json'>
			names: <input type="text" name="Name">
			genders: <input type="text" name="Gender">
			<input type="submit">
			</form>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html; charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
	fmt.Println("hey")
}

func (m MongoController) Delete(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
		<header>
			<meta charset="UTF-8">
			<title>Index</title>
		</header>
		<body>
			<a href="/users/766"> GO TO </a>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html; charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
}
func (m MongoController) Index(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
		<header>
			<meta charset="UTF-8">
			<title>Index</title>
		</header>
		<body>
			<form action="/create"  method="post" enctype='application/json'>
				name : <input type="text" name="Name">
				gender : <input type="text" name="Gender">
				age : <input type="number" name="Age">
				<input value="Submit" type="submit">
			</form>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html;application/json charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
}
