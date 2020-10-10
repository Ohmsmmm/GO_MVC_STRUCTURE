package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
)

type DataHandler struct {
}

func (h *DataHandler) EndPointIOAccount(c *gin.Context) {

	var data Hashtag
	//Send to struct
	e := c.BindJSON(&data)
	if e != nil {
		fmt.Println(e)
	}

	v := reflect.ValueOf(data)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	s := make([]string, len(values))
	for i, v := range values {
		s[i] = fmt.Sprint(v)
	}
	//check Character Input
	checkLetter := h.IsLetter(s)
	if !checkLetter {
		c.JSON(http.StatusBadRequest,"Please Enter Character A-Z")
	}

	result, err:= h.genMsg(s)
	if err != "" {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, result)
}

func (h *DataHandler) IsLetter(s []string) bool {
	for _, r := range s {
		for _, x := range r{
			if (x < 'a' || x > 'z') && (x < 'A' || x > 'Z') {
				return false
			}
		}
	}
	return true
}

func (h *DataHandler) genRandCharacter(n int) string{
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func (h *DataHandler) genMsg(s []string) (result []AccountIO, err string) {
	for i:=1; i<=926; i++ {
		// random Hashtag
		randomIndex := rand.Intn(len(s))
		pickHastag := s[randomIndex]
		//random Message
		minMsg := 30
		maxMsg := 140
		max := maxMsg - len(pickHastag)
		min := minMsg - len(pickHastag)
		numMsg := rand.Intn(max-min) + min
		randMessage := h.genRandCharacter(numMsg)
		// compose Message
		msg := fmt.Sprintf("%s#%s", randMessage, pickHastag)
		fmt.Println(msg)
		isIO := h.scanIO(msg)
		if(isIO >= 2) {
			Account := AccountIO{
				AccountNo: i,
				Msg:       msg,
				CountIO: isIO,
			}
			result = append(result, Account)
		}
	}
	return
}
func (h *DataHandler) scanIO(s string) int{
		// Displaying strings
		fmt.Println("Scanning IO : ", s)
		// Counting the elements of the strings
		// Using Count() function
		countI := strings.Count(s, "I")
		countO := strings.Count(s, "O")
		countIo := countI+countO
	return countIo
}