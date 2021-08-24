package client

import (
	"fmt"
	"gomodules.xyz/go-sh"
)

type Auth struct {
	TLS      bool
	CA       string
	Cert     string
	Key      string
	User     string
	Password string
	Creds    string
}
type Stream struct {
	Name string
	Dir  string
}

func Backup(stream Stream, auth Auth) {
	session := sh.NewSession()
	if auth.TLS == true {
		err := session.Command("nats", "stream", "backup", stream.Name, stream.Dir+stream.Name+auth.User, "--tlscert", auth.Cert, "--tlskey", auth.Key, "--tlsca", auth.CA, "--user", auth.User, "--password", auth.Password, "--creds", auth.Creds).Run()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := session.Command("nats", "stream", "backup", stream.Name, stream.Dir+stream.Name+auth.User, "--user", auth.User, "--password", auth.Password, "--creds", auth.Creds).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
func Restore(stream Stream, auth Auth) {
	session := sh.NewSession()
	if auth.TLS == true {
		err := session.Command("nats", "stream", "restore", stream.Name, stream.Dir+stream.Name+auth.User, "--tlscert", auth.Cert, "--tlskey", auth.Key, "--tlsca", auth.CA, "--user", auth.User, "--password", auth.Password, "--creds", auth.Creds).Run()
		if err != nil {
			fmt.Println()
		}
	} else {
		err := session.Command("nats", "stream", "restore", stream.Name, stream.Dir+stream.Name+auth.User, "--user", auth.User, "--password", auth.Password, "--creds", auth.Creds).Run()
		if err != nil {
			fmt.Println()
		}
	}
}
