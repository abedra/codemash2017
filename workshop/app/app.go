package main

import (
        "log"
        "fmt"
        "github.com/gorilla/mux"
        "github.com/gorilla/handlers"
        "github.com/gorilla/context"
        "html/template"
        "net/http"
	"net/url"
        "os"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Page struct {
        Title string
        Error string
        Repsheet bool
}

type GResponse struct {
	Success            bool
	ChallengeTimestamp string `json:challenge_ts`
	Hostname           string
	ErrorCodes         []string
}

var templates = template.Must(template.ParseFiles("index.html", "admin.html"))

func verifyRecaptcha(gResponse string) bool {
        data := url.Values{}
        data.Set("secret", "6LcuXBAUAAAAAAFcwv--LwXc1mU5C_yYfZICZDCM")
        data.Set("response", gResponse)

        request, _ := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", strings.NewReader(data.Encode()))
        request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        transport := http.Transport{}
        response, _ := transport.RoundTrip(request)
        defer response.Body.Close()

	bodyString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var parsedResponse GResponse
	json.Unmarshal(bodyString, &parsedResponse)

	if !parsedResponse.Success {
		log.Println("Recaptcha validation failed")
	}

	return parsedResponse.Success
}

func repsheetHandler(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                if r.Header["X-Repsheet"] != nil {
                        context.Set(r, "repsheet", true)
                }
                next.ServeHTTP(w, r)
        })
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
        if (request.Method == "GET") {
                response.Header().Set("Content-type", "text/html")
                err := request.ParseForm()
                if err != nil {
                        http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
                }

                page := Page{Title: "Login"}

                if context.Get(request, "repsheet") != nil {
                        page.Repsheet = true
                }

                templates.ExecuteTemplate(response, "index.html", page)
        } else if (request.Method == "POST") {
                request.ParseForm()
                username := request.PostFormValue("inputEmail")
                password := request.PostFormValue("inputPassword")

                recaptcha := request.PostFormValue("g-recaptcha-response")
		verifyRecaptcha(recaptcha)

                if (username == "admin@example.com" && password == "P4$$w0rd!") {
                        log.Println("Successfull login for admin@example.com")
                        http.Redirect(response, request, "/admin", 302)
                } else {
                        log.Println("Login failed for admin@example.com")
                        page := Page{Title: "Login", Error: "Username or Password Invalid"}
                        templates.ExecuteTemplate(response, "index.html", page)
                }
        }
}

func AdminHandler(response http.ResponseWriter, request *http.Request) {
        response.Header().Set("Content-type", "text/html")
        err := request.ParseForm()
        if err != nil {
                http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
        }
        templates.ExecuteTemplate(response, "admin.html", Page{Title: "Admin"})
}

func main() {
        logFile, err := os.OpenFile("logs/app.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
        if err != nil {
                fmt.Println("Error accessing log file:", err)
                os.Exit(1)
        }

        r := mux.NewRouter()
        r.Handle("/", handlers.ProxyHeaders(handlers.LoggingHandler(logFile, repsheetHandler(http.HandlerFunc(LoginHandler)))))
        r.Handle("/admin", handlers.ProxyHeaders(handlers.LoggingHandler(logFile, repsheetHandler(http.HandlerFunc(AdminHandler)))))
        r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
        http.Handle("/", r)

        err = http.ListenAndServe("localhost:8080", r)
        if err != nil {
                log.Fatal("Error starting server: ", err)
        }
}
