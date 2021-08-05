package client

import (
	"gomodules.xyz/go-sh"
)

const (
	BACKUP_DIR = "/home/hmsayem/Documents/backup/"
	TLS_DIR    = "/home/hmsayem/Documents/js-tls"
	TLS        = "--tlsca=ca.crt"
)

func Backup(stream string) {
	session := sh.NewSession()
	session.SetDir(TLS_DIR)
	session.Command("nats", "stream", "backup", stream, BACKUP_DIR+stream, TLS).Run()
}
