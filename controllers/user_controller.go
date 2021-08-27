package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Authentication1/models"
	"github.com/Authentication1/utils"

	"strconv"

	"github.com/go-ldap/ldap"
	"github.com/gorilla/mux"
)

type User struct {
	givenName, sn, cn, passwordUser string
}

const (
	BindUsername = "cn=admin,dc=arqsoft,dc=unal,dc=edu,dc=co"
	BindPassword = "admin"
	BaseDN       = "dc=arqsoft,dc=unal,dc=edu,dc=co"
	Filter       = "(objectClass=inetOrgPerson)"

	user     = "Valentina  Viafara"
	password = "123"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	//defer db.Close()

	db.Find(&user, id)

	if user.User_id > 0 {
		j, _ := json.Marshal(user)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
	utils.CloseConnection(db)
}

func GetUserLdap(w http.ResponseWriter, r *http.Request) {

	user := mux.Vars(r)["idUser"]
	password := mux.Vars(r)["password"]

	l := connectServer()

	value := authenticated(l, user, password)
	if value {

		utils.SendResponse(w, http.StatusOK, []byte("true"))
	} else {
		utils.SendResponse(w, http.StatusOK, []byte("false"))
	}

	fmt.Printf("\nESTO FUNCIONO%s\n", user)
	fmt.Printf("\nESTO FUNCIONO%s\n", password)

	l.Close()

}

func UserLogin(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	userdb := models.User{}
	db := utils.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	db.Where("email = ?", user.Email).Find(&userdb)
	if (userdb.User_id > 0) && (user.User_password == userdb.User_password) {
		j, _ := json.Marshal(userdb)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
	utils.CloseConnection(db)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := []models.User{}
	db := utils.GetConnection()
	db.Find(&users)
	j, _ := json.Marshal(users)
	utils.SendResponse(w, http.StatusOK, j)
	utils.CloseConnection(db)
}

func SetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	db := utils.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	user.Reg_date = time.Now()
	err = db.Create(&user).Error
	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	j, _ := json.Marshal(user)
	utils.SendResponse(w, http.StatusCreated, j)
	utils.CloseConnection(db)

	l := connectServer()

	email := user.Email
	password := user.User_password
	givenName := user.Firstname
	lastName := user.Lastname
	newUser := User{givenName, lastName, email, password}
	if err := addUser(l, &newUser); err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}

	fmt.Printf("user: %v\n", user.Firstname)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userFind := models.User{}
	userData := models.User{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()

	db.Find(&userFind, id)

	if userFind.User_id > 0 {
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			utils.SendErr(w, http.StatusBadRequest)
			return
		}
		db.Model(&userFind).Updates(userData)
		j, _ := json.Marshal(userFind)
		utils.SendResponse(w, http.StatusOK, j)
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
	utils.CloseConnection(db)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	id := mux.Vars(r)["id"]
	db := utils.GetConnection()
	db.Find(&user, id)
	if user.User_id > 0 {
		db.Where("user_id = ?", id).Delete(&user)

		utils.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
	utils.CloseConnection(db)
}

func connectServer() (l *ldap.Conn) {
	ldapURL := "ldap://localhost:389"
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		return nil
	} else {
		fmt.Println("Connection with the server sucessfully!")
	}

	return l
}

func authenticated(l *ldap.Conn, user string, password string) bool {
	l.Bind(BindUsername, BindPassword)

	//filter : = ("cn=%s", ldap.EscapeFilter(user))
	result, err := searchUser(l, user)
	if err != nil {
		fmt.Printf("Couldn't fetch search entries: %s", err)
		return false
	} else {
		fmt.Print("DN : ", result)
	}

	if err := l.Bind(result, password); err != nil {
		fmt.Printf("\n\nFailed to auth. %s", err)
		return false
	} else {
		fmt.Printf("\nAuthenticated successfuly!\n")
	}

	return true
}

func searchUser(l *ldap.Conn, user string) (string, error) {
	l.Bind(BindUsername, BindPassword)
	filter := fmt.Sprintf("(cn=%s)", user)
	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		[]string{"cn", "sn", "uid", "postalAddress", "telephoneNumber", "uidNumber"},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return "", fmt.Errorf("Search Error: %s", err)
	} else {
		fmt.Printf("User search successfuly!\n--------------------------------------------------\n")
		result.Entries[0].Print()

		return result.Entries[0].DN, nil
	}
}

func addUser(l *ldap.Conn, newUser *User) error {

	l.Bind(BindUsername, BindPassword)

	result, err := searchAllUser(l)
	if err != nil {
		return err
	}

	uidNumberNew := fmt.Sprintf("%v", getMaxUidNumber(result))

	dn := fmt.Sprintf("cn=%s,ou=viwit,dc=arqsoft,dc=unal,dc=edu,dc=co", newUser.cn)
	directory := fmt.Sprintf("/home/users/%s", newUser.cn)

	a := ldap.NewAddRequest(dn, []ldap.Control{})
	a.Attribute("objectClass", []string{"inetOrgPerson", "posixAccount", "top"})
	a.Attribute("cn", []string{newUser.cn})
	a.Attribute("gidNumber", []string{"500"})
	a.Attribute("givenName", []string{newUser.givenName})
	a.Attribute("homeDirectory", []string{directory})
	a.Attribute("uid", []string{newUser.cn})
	a.Attribute("uidNumber", []string{uidNumberNew})
	a.Attribute("sn", []string{newUser.sn})
	a.Attribute("userPassword", []string{newUser.passwordUser})
	add(a, l)

	return nil
}

func add(addRequest *ldap.AddRequest, l *ldap.Conn) {
	err := l.Add(addRequest)
	if err != nil {
		fmt.Println("Fail in request", err)
	} else {
		fmt.Println("\nUser DONE\n", err)
	}
}

func searchAllUser(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.Bind(BindUsername, BindPassword)

	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{"cn", "sn", "uid", "postalAddress", "telephoneNumber", "uidNumber"},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		result.Print()
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

func getMaxUidNumber(result *ldap.SearchResult) int {

	if len(result.Entries) == 0 {
		return 1000
	}

	uidNum := 0
	for _, s := range result.Entries {
		uidNumtemp := s.GetAttributeValue("uidNumber")
		uidNumtempInt, _ := strconv.Atoi(uidNumtemp)
		if uidNumtempInt > uidNum {
			uidNum = uidNumtempInt
		}
	}

	return uidNum + 1

}
