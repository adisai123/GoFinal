package controller

import (
	"encoding/json"
	"fmt"
	"gofinal/src/src/mongo/crud/model"
	"net/http"

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
func (m MongoController) Create(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("err", err)
	fmt.Println("hiii", u.Name)
	u.Id = bson.NewObjectId()
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
			<form action="/create" method="post">
			name: <input type="text" name="name">
			gender: <input type="text" name="gender">
			<input type="submit">
			</form>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html; charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
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
			<form action="/create" method="post">
				name: <input type="text" name="Name">
				gender: <input type="text" name="Gender">
				age: <input type="number" name="Age">
				<input type="submit">
			</form>
		</body>	
	</html>
	`
	rs.Header().Set("Content-Type", "text/html; charset=UPF-8")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(s))
}
