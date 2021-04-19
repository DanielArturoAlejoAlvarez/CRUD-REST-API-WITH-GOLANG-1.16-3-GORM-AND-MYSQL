# CRUD REST API GOLANG GORM AND MYSQL

## Description

This repository is a Software of Application with Golang, Mux, GORM (ORM) and MySQL.

## Installation

Using Go 1.16 preferably.

## DataBase

Using MySQL preferably.

!EYE Configure DB info through the create of file .env

## Apps

Using Postman, Insomnia, Talend API Tester, etc.

## Usage

```html
$ git clone https://github.com/DanielArturoAlejoAlvarez/CRUD-REST-API-WITH-GOLANG-1.16-3-GORM-AND-MYSQL.git
[NAME APP]

$ go build

$ ./[name-app] [ENTER]
```

Follow the following steps and you're good to go! Important:

![alt text](https://steemitimages.com/p/7258xSVeJbKmPTruw3bvBgohFbX4k8MUsxTzCCqqN9vKTXsrpi8KPrgiEGLnGE3Dds94YijqrjwSUo3RdhyYDD49rFA5fehFe7yWuy47k5b9oNMzXxeWwiCGU9SB7nUvCdbREyKVKPcJH?format=match&mode=fit)

## Coding

### Config

```go
...
var DB *gorm.DB

var err error

var DNS = "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Can't connect to DB!")
	}

	DB.AutoMigrate(&User{})
}
...
```

### Routes

```go
...
func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("api/users", GetUsers).Methods("GET")
	r.HandleFunc("api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("api/users", CreateUser).Methods("POST")
	r.HandleFunc("api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("api/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}
...
```

### Controllers

```go
...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
...
```

### Models

```go 
...
type User struct {
	gorm.Model
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Age         int    `json:"age"`
	Avatar      string `json:"avatar"`
	Flag        bool   `json:"flag"`
}
...
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/DanielArturoAlejoAlvarez/CRUD-REST-API-WITH-GOLANG-1.16-3-GORM-AND-MYSQL. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
````