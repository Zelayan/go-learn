package watch

import (
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)
var sqlite_path = ""
type User struct {
	ID int `gorm:"primary_key"`
	Name string
	Age int
	Birthday time.Time
}

type DB struct {
	db *gorm.DB
}

func (d *DB)InitTable() error {
	 return d.db.AutoMigrate(User{})
}

func (d *DB) Create(user *User) error{
	return d.db.Create(user).Error
}

type Sync struct {

}

func (s *Sync) NewDBSync() *Sync {
	return &Sync{}
}


type WatchConfig struct {
	BindAddr string `yaml:"bind_addr,omitempty"`
	BindPort string `yaml:"bind_port,omitempty"`

	NoBackup bool `yaml:"no_backup,omitempty"`

	UseSSL        bool   `yaml:"use_ssl,omitempty"`
	SSLKeyFile    string `yaml:"ssl_key_file,omitempty"`
	SSLCertFile   string `yaml:"ssl_cert_file,omitempty"`
	SkipSSLVerify bool   `yaml:"skip_ssl_verify,omitempty"`

	AuthKey string `yaml:"auth_key,omitempty"`

	SyncFile   string `yaml:"sync_file,omitempty"`
	RemoteConn string `yaml:"remote_conn,omitempty"`

	SyncInterval int64 `yaml:"sync_interval,omitempty"`
}

func sync(addr string, db string, options WatchConfig) {
	//var pull_url string
	var download_url string

	sql_backup_path := fmt.Sprintf("%s.new.sql", db)
	backup_path := fmt.Sprintf("%s.old", db)

	//done := make(chan struct{})
	download := make(chan struct{}, 1)
	client := &http.Client{}

	go func() {
		for {
			<-download

			req, err := http.NewRequest("GET", download_url, nil)
			resp, err := client.Do(req)

			if err != nil {
				fmt.Printf("unable to download latest DB from upstream, retrying in 5s: %s", err)

				go func() {
					time.Sleep(time.Duration(5) * time.Second)
					select {
					case download <- struct{}{}:

					default:
						// already a download in queue, don't add another one
					}
				}()

				continue
			}

			_ = os.Remove(sql_backup_path)
			out, err := os.Create(sql_backup_path)
			if err != nil {
				fmt.Printf("unable to write latest DB to file: %s", err)
				continue
			}

			defer out.Close()
			io.Copy(out, resp.Body)

			if dbexists, _ := exists(db); dbexists {
				err = copyFileContents(db, backup_path)
				if err != nil {
					fmt.Printf("unable to back up current sqlite database: %s", err)
					continue
				}

				err = os.Chmod(db, 0600)
				if err != nil {
					fmt.Printf("unable to change permissions on existing DB prior to import, err: %s", err)
					continue
				}
			}

			drop_sql_command := "PRAGMA writable_schema = 1; delete from sqlite_master where type in ('table', 'index', 'trigger'); PRAGMA writable_schema = 0; VACUUM; PRAGMA INTEGRITY_CHECK;"
			drop_sqlite_out, err := exec.Command(sqlite_path, db, drop_sql_command).CombinedOutput()
			if err != nil {
				fmt.Printf("unable to drop existing DB prior to import, output: %s", string(drop_sqlite_out))
				continue
			}

			sql_command := fmt.Sprintf(".read %s", sql_backup_path)
			sqlite_out, err := exec.Command(sqlite_path, db, sql_command).CombinedOutput()
			if err != nil {
				fmt.Printf("unable to import newly downloaded DB from upstream, output: %s", string(sqlite_out))
				continue
			}

			err = os.Chmod(db, 0400)
			if err != nil {
			//	log.Error("unable to change permissions on DB following import, err: %s", err)
				continue
			}

		//log.Info("updated DB on disk with latest")

			_ = os.Remove(sql_backup_path)
		}
	}()

}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}