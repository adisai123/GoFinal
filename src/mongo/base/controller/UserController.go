package controller

import (
	"encoding/json"
	"gofinal/src/src/mongo/base/model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
}

func GetController() *UserController {
	return &UserController{}
}

func (uc UserController) Create(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := model.User{}
	json.NewDecoder(r.Body).Decode(&u)
	u.Id = "007"
	uj, _ := json.Marshal(u)
	rs.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rs.WriteHeader(http.StatusCreated)
	rs.Write([]byte(uj))
}
func (uc UserController) GetUser(rs http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := model.User{
		Name:   "aDITYA",
		Age:    25,
		Gender: "MALE",
		Id:     p.ByName("id"),
	}
	uj, _ := json.Marshal(u)
	rs.Header().Set("Content-Type", "application/json")
	rs.WriteHeader(http.StatusOK)
	rs.Write([]byte(uj))
}
func (uc UserController) Index(rs http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
