package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func getProfile(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	// defer repo.conn.Close(c.Request.Context())

	prof := &Profile{ID: c.GetString("uid")}

	log.Printf("[updateProf] | %v", prof.ID)

	err := repo.conn.QueryRow(context.Background(), selectProfileByID, prof.ID).Scan(&prof.ID, &prof.Name, &prof.Age, &prof.Weight, &prof.LowTarget, &prof.HighTarget)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, &prof)
}

func updateProfile(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	// defer repo.conn.Close(c.Request.Context())

	var prof *Profile
	err := c.Bind(&prof)
	if err != nil {
		c.JSON(501, err)
		return
	}

	prof.ID = c.GetString("uid")
	log.Printf("[updateProf] | %v", prof)

	_, err = repo.conn.Exec(context.Background(), updateProfilebyID, &prof.ID, &prof.Name, &prof.Age, &prof.Weight, &prof.LowTarget, &prof.HighTarget)
	if err != nil {
		log.Printf("[Error] [updateProf] | %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, &prof)
}

func getData(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	rows, err := repo.conn.Query(context.Background(), selectData, c.GetString("uid"))
	if err != nil {
		log.Printf("[GET DATA] %v", err)
		c.JSON(500, err)
		return
	}

	dataList := make([]*Data, 0)

	for rows.Next() {
		data := &Data{}
		err = rows.Scan(
			&data.ID,
			&data.Time,
			&data.BGLevel,
		)
		if err != nil {
			log.Printf("[GET TASKS 2] %v", err)
			c.JSON(501, err)
			return
		}
		dataList = append(dataList, data)
	}

	c.JSON(200, &dataList)
}

func postData(c *gin.Context) {
	repo := NewRepo(sqlConnString)
	defer repo.conn.Close(context.Background())
	var data *Data
	err := c.Bind(&data)
	if err != nil {
		log.Printf("[POST DATA] %v", err)
		c.JSON(501, err)
		return
	}

	data.ID = c.GetString("uid")

	_, err = repo.conn.Exec(context.Background(), postDataStmt,
		&data.ID,
		&data.Time,
		&data.BGLevel,
	)
	if err != nil {
		log.Printf("[POST DATA] %v", err)
		c.JSON(500, err)
		return
	}

	c.JSON(200, true)
}

const (
	selectProfileByID = "SELECT * FROM profiles WHERE uid = $1"
	updateProfilebyID = "INSERT INTO profiles (uid, name, age, weight, low, high) VALUES($1, $2, $3, $4, $5, $6) ON CONFLICT (uid) DO UPDATE SET (name, age, weight, low, high) = ($2, $3, $4, $5, $6);"
	selectData        = "SELECT * FROM data WHERE uid = $1 ORDER BY time"
	postDataStmt      = "INSERT INTO data (uid, time, bg) VALUES ($1, $2, $3)"
)
