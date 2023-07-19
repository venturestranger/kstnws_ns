package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Notification struct {
	Id			int `json:"id" db:"id"`
	IdUser		int `json:"id_user" db:"id_user"`
	Status		int `json:"status" db:"status"`
	Content		string `json:"content" db:"content"`
	Date		string `json:"date" db:"date"`
}

func GetHandler(c *gin.Context) {
	ch := make(chan bool)

	go func() {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Println(err)
			SendStatus(http.StatusBadRequest, c)
			ch <- true
			return
		}
		defer db.Close()

		statement := "select * from notifications "

		if c.Query("id") != "" || c.Query("id_user") != "" {
			var andFlag bool = false

			if c.Query("id") != "" {
				statement += " where id = " + c.Query("id")
				andFlag = true
			}
			if c.Query("id_user") != "" {
				if andFlag {
					statement += " and "
				} else {
					statement += " where "
				}
				statement += " id_user = " + c.Query("id_user")
			}
		}
		if c.Query("order_way_") != "" {
			statement += " order by date " + c.Query("order_way_")
		}
		if c.Query("offset_") != "" {
			statement += " offset " + c.Query("offset_")
		}
		if c.Query("limit_") != "" {
			statement += " limit " + c.Query("limit_")
		}

		var notifications []Notification
		db.Select(&notifications, statement)

		SendData(notifications, c)
		ch <- true
	}()
	<- ch
}

func DeleteHandler(c *gin.Context) {
	ch := make(chan bool)

	go func() {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Println(err)
			SendStatus(http.StatusBadRequest, c)
			ch <- true
			return
		}
		defer db.Close()

		id := c.Query("id")
		_, err = db.Exec(fmt.Sprintf("delete from notifications where id = %s", id))

		if err != nil {
			log.Println(err)
			SendStatus(http.StatusInternalServerError, c)
			ch <- true
			return
		}

		SendStatus(http.StatusOK, c)
		ch <- true
	}()
	<- ch
}

func PutHandler(c *gin.Context) {
	ch := make(chan bool)

	go func() {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Println(err)
			SendStatus(http.StatusBadRequest, c)
			ch <- true
			return
		}

		id := c.Query("id")
		var notification Notification
		payload, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(payload, &notification)
		_, err = db.NamedExec(fmt.Sprintf("update notifications set status = :status, content = :content, id_user = :id_user, date = :date where id = %s", id), notification)

		SendStatus(http.StatusOK, c)
		ch <- true
	}()
	<- ch
}

func PostHandler(c *gin.Context) {
	ch := make(chan bool)

	go func() {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Println(err)
			SendStatus(http.StatusBadRequest, c)
			ch <- true
			return
		}
		defer db.Close()

		var notification Notification
		payload, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(payload, &notification)
		_, err = db.NamedExec("insert into notifications(status, content, id_user, date) values(:status, :content, :id_user, :date)", notification)

		if err != nil {
			log.Println(err)
			SendStatus(http.StatusBadRequest, c)
			ch <- true
			return
		}

		SendStatus(http.StatusOK, c)
		ch <- true
	}()
	<- ch
}
