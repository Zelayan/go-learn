package watch

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os/exec"
	"testing"
	"time"
)

func TestInitTable(t *testing.T) {
	check := assert.New(t)
	db, err := gorm.Open(sqlite.Open("watch.db"), &gorm.Config{})
	check.Nil(err)

	d := &DB{db: db}
	err = d.InitTable()
	check.Nil(err)
	err = d.Create(&User{
		ID:       1,
		Name:     "xx",
		Age:      0,
		Birthday: time.Time{},
	})
	check.Nil(err)
}


func TestSync(t *testing.T) {

	output, err := exec.Command("/bin/bash", "-c", "sqlite3  new.db < tt.sql").Output()
	if err != nil {
		t.Error(err)
	}


	t.Log(string(output))
}


