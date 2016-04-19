package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joeshaw/envdecode"
	"github.com/mailgun/mailgun-go"
)

var (
	url     = flag.String("url", "", "Url of the email message")
	to      = flag.String("to", "", "Email recipient to send")
	from    = flag.String("from", "testing@test.com", "From email")
	subject = flag.String("subject", "Default subject", "Email subject")
	cfg     Config
)

// Config holds environment settings, mostly secrets
type Config struct {
	APIKey string `env:"MAILGUN_API_KEY,default=asdf"`
	Domain string `env:"MAILGUN_DOMAIN,default=asdf"`
}

func main() {

	flag.Parse()
	err := envdecode.Decode(&cfg)

	if err != nil {
		log.Fatal(err)
	}

	if *url == "" || *to == "" {
		log.Fatal("At least must provide a url and a to dest")
	}

	content := getContent(*url)
	mg := mailgun.NewMailgun(cfg.Domain, cfg.APIKey, "")
	m := mailgun.NewMessage(*from, *subject, "test", *to)
	m.SetHtml(content)
	_, id, err := mg.Send(m)
	if err != nil {
		log.Fatal("Error sending", err)
	}

	log.Printf("Message sent %s\n", id)

}

func getContent(url string) string {
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("Unable to fetch url content %s", err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read body %s", err)
	}
	return string(body)
}
