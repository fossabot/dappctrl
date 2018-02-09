package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pxctrl/serv"
	"pxctrl/util"
)

type config struct {
	LogFile    string
	ServerAddr string
	ServerTLS  bool
	SessionDir string
}

func newConfig() *config {
	conf := serv.NewConfig()
	return &config{
		LogFile:    "fxtrig.log",
		ServerAddr: conf.Addr,
		ServerTLS:  conf.TLS,
		SessionDir: "sessions",
	}
}

const (
	logPerm  = 0644
	sessPerm = 0644
)

func main() {
	conf := newConfig()
	name := util.ExeDirJoin("pxtrig.config.json")
	if err := util.ReadJSONFile(name, &conf); err != nil {
		log.Fatalf("failed to read configuration: %s\n", err)
	}

	if len(conf.LogFile) != 0 {
		out, err := os.OpenFile(conf.LogFile,
			os.O_RDWR|os.O_CREATE|os.O_APPEND, logPerm)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer out.Close()

		log.SetOutput(out)
	}

	switch os.Getenv("script_type") {
	case "user-pass-verify":
		handleAuth(conf)
	case "client-connect":
		handleConnect(conf)
	case "client-disconnect":
		handleDisconnect(conf)
	default:
		log.Fatalf("unknown script_type")
	}
}

func getUsername() string {
	return os.Getenv("username")
}

func getCreds() (string, string) {
	user := getUsername()
	pass := os.Getenv("password")

	if len(user) != 0 && len(pass) != 0 {
		return user, pass
	}

	if len(os.Args) < 2 {
		log.Fatalf("no filename passed to read credentials")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("failed to open file with credentials: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	user = scanner.Text()
	scanner.Scan()
	pass = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal("failed to read file with credentials: %s", err)
	}

	return user, pass
}

func post(conf *config, path string, req interface{}, rep interface{}) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("failed to encode request: %s", err)
	}

	var proto = "http"
	if conf.ServerTLS {
		proto += "s"
	}

	resp, err := http.Post(proto+"://"+conf.ServerAddr+path,
		"application/json", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("failed to post request: %s", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(rep); err != nil {
		log.Fatalf("failed to decode reply: %s", err)
	}
}

func storeSession(conf *config, user, session string) {
	name := filepath.Join(conf.SessionDir, user)
	err := ioutil.WriteFile(name, []byte(session), sessPerm)
	if err != nil {
		log.Fatalf("failed to store session: %s", err)
	}
}

func handleAuth(conf *config) {
	user, pass := getCreds()

	req := serv.AuthRequest{PaymentID: user, PaymentPassword: pass}

	var rep serv.AuthReply
	post(conf, serv.PathAuthenticate, req, &rep)
	if len(rep.Error) != 0 {
		log.Fatalf("failed to authenticate %s: %s", user, rep.Error)
	}

	storeSession(conf, user, rep.SessionID)
}

func handleConnect(conf *config) {

}

func handleDisconnect(conf *config) {

}
