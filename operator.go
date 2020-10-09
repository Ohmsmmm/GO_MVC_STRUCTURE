package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
type DataHandler struct {
}

func (h *DataHandler) EndPointDataCleansing(c *gin.Context) {
	
	var data DataCleansingRequest
	//data.InputString = c.Param("input_string")
	e := c.BindJSON(&data)
	if e != nil {
		fmt.Println(e)
	}

	result, err := h.cleansing(data)
	if err != "" {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, result)
}

func (h *DataHandler) cleansing(data DataCleansingRequest) (result []DataCleansingResponse, err string) {

	fmt.Println("inputString : ", data.InputString)
	// 1. รับข้อความอันมีความยาวไม่ตํา่ กว่า 20 ตัวอักษร และไม่มากกวา่ 255 ตัวอักษร (นับช่องวา่ ง)
	l := len(data.InputString)
	fmt.Print("Write ", l, " as ")
	if l < 20 {
		err = fmt.Sprintf("ข้อความอันมีความยาวไม่ต่ำกว่า 20 ตัวอักษร : \"%s\" ", data.InputString)
		return
	}
	if l > 255 {
		err = fmt.Sprintf("ข้อความอันมีความยาวไม่มากกว่า 255 ตัวอักษร : \"%s\" ", data.InputString)
		return
	}

	// 2. ข้อความต้องผ่านกระบวนการทําความสะอาดข้อมลู หรือDataCleansingก่อนด้วยการแปลงให้
	// ตัวอักษรพิมพ์ใหญ่ภายในข้อความให้เปน็ ตัวอักษรพมิพ์เล็กทั้งหมด
	data.InputString = strings.ToLower(data.InputString)

	// 3. ทําความสะอาดข้อความเพิ่มเติมด้วยการนําเครื่องหมาย “?” “!” “,” ออกจากข้อความ
	data.InputString = strings.ReplaceAll(data.InputString, "?", "")
	data.InputString = strings.ReplaceAll(data.InputString, "!", "")
	data.InputString = strings.ReplaceAll(data.InputString, ",", "")

	// 4. ตัดข้อความโดยใช้ “ “ (ช่องวา่ ง หรือ Whitespace) ภายในข้อความ
	wordAsArray := strings.Split(data.InputString, " ")

	// 5. แสดงผลในลักษณะของตารางแจกแจงความถี่ของคําแต่ละคําในข้อความออกมา (ต้องการเพียงแคค่ ํา
	// แต่ละคํามคี วามถี่เป็นเทา่ ใด)
	// map tabel คล้าย array แบบไม่มีโครงสร้าง
	MapWordCount := make(map[string]*DataCleansingResponse)
	for _, word := range wordAsArray {

		if valueInMap, ok := MapWordCount[word]; ok {
			valueInMap.Count++
		} else {

			MapWordCount[word] = &DataCleansingResponse{
				Word:  word,
				Count: 1, //conut begin 1
			}
		}
	}

	fmt.Println("inputString : ", data.InputString)

	for _, val := range MapWordCount {
		result = append(result, *val)
	}

	return
}