package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"gin_server/server/db"
	"gin_server/server/model"
	"gin_server/server/service"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	database    db.Database = db.Newdatabase()
	client, _               = database.DBinstance()
	DbtoConnect             = database.OpenDatabase(client, "students")
)

func StudentHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err := c.SaveUploadedFile(file, "./assets/"+file.Filename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	f, err := excelize.OpenFile("./assets/" + file.Filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, _ := f.GetRows("Sheet1")
	students := service.GetStructFromData(rows)
	StudentCollection := database.OpenCollection(client, DbtoConnect, "students_data")
	items := make([]interface{}, len(students))
	for i, v := range students {
		items[i] = v
	}
	result, err := StudentCollection.InsertMany(context.TODO(), items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Students with this data already exists.",
		})
		return
	}
	c.JSON(200, result)

}

func DownloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "application/octet-stream")
	c.File("./assets/" + filename)
}

func GetStudents(c *gin.Context) {
	StudentCollection := database.OpenCollection(client, DbtoConnect, "students_data")
	cursor, err := StudentCollection.Find(context.TODO(), bson.D{})
	if err != nil {

	}
	var results []*model.Student
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, results)
}
