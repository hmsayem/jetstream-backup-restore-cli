package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "path/filepath"
  "strings"
  "github.com/nats-io/jwt/v2"

)

func main() {
 
  // Push x account information to account server
  xJwt, err := ioutil.ReadFile("./X.jwt")
  if err != nil {
    panic(err)
  }
  claim, err := jwt.DecodeAccountClaims(string(xJwt))
  if err != nil {
    panic(err)
  }
  if err := PushAccountToAccountServer(claim.Subject, string(xJwt)); err != nil {
    panic(err)
  }
}

func PushAccountToAccountServer(name, jwt string) error {
  ur, err := url.Parse("http://localhost:9090/jwt/v1/accounts/")
  if err != nil {
    return err
  }
  ur.Path = filepath.Join(ur.Path, name)
  resp, err := http.Post(ur.String(), "application/jwt", strings.NewReader(jwt))
  if err != nil {
    return err
  } else if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("could not post to <%q>: %v", ur.String(), resp.Status)
  }
  resp.Body.Close()

  return err
}


