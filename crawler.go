package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/kabukky/httpscerts"
	//"github.com/mjdilworth/stringutil"
)

func main() {
	fmt.Printf("Hi dude\n")
	//fmt.Printf(stringutil.Reverse("mjd"))
	// Check if the cert files are available.
	err := httpscerts.Check("test_cert.pem", "test_key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("test_cert.pem", "test_key.pem", "127.0.0.1:8000")
		if err != nil {
			log.Fatal("Error: Couldn't create https test certs.")
		}
	}
	//http.HandleFunc("/", handler)
	//http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	router := gin.Default()

	router.Use(secure.Secure(secure.Options{
		AllowedHosts:          []string{"localhost", "Michaels-MBP", "mjd.testdomain.com"},
		SSLRedirect:           true,
		SSLHost:               "mjd.testdomain.com",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            86400,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IsDevelopment:         true,
	}))

	// Create routes
	router.GET("/", DefaultLanding)
	router.GET("/url", UrlLanding)
	/*
		router.GET("/tasks", TasksList)
		router.POST("/tasks", TaskPost)
		router.GET("/tasks/:id", TasksDetail)
		router.GET("/health", Health)
		router.GET("/ping", PingPong)
	*/
	//	router.Run(":8000")
	router.RunTLS(":8000", "cert.pem", "key.pem")

}

//UrlLanding is our testing landing page whihc implements the functionality
func UrlLanding(c *gin.Context) {
	uri := c.DefaultQuery("uri", "www.capgemini.com")
	linkmap, err := GetLinks(uri)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	strLinks := linkmap.PrintLinks()
	c.String(http.StatusOK, "Simple Microservices Demo received uri as %s \n\n %s", uri, strLinks)

}

//DefaultLanding is the simple default route for the root. It does nothing but return text.
func DefaultLanding(c *gin.Context) {
	c.String(http.StatusOK, "mjd test")
}
