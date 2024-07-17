package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msgsage string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteJson(c *gin.Context, data interface{}) {
	c.JSON(200, Response{Msgsage: "Success", Data: data})
}
func Error(err string) error {
	return errors.New(err)
}

func WriteError(c *gin.Context, msg string, err error) {
	if err == nil {
		c.JSON(500, Response{Msgsage: "Error", Data: map[string]interface{}{
			"Error-Message": msg}})
		return
	}
	c.JSON(500, Response{Msgsage: "Error", Data: map[string]interface{}{
		"Error-Message": msg,
		"Error":         err.Error(),
	}})
}

func Atoi(value string, defaultValue int) int {
	v, e := strconv.Atoi(value)
	if e != nil {
		return defaultValue
	}
	return v
}

func Itoa(value int, defaultValue string) string {
	//
	return ""
}

func PanicIfError(err error, format string, v ...any) {
	if err != nil {
		v = append(v, err.Error())
		log.Printf(format+"\nerror: %s", v...)
		// Sprintf(format+"\nerror: %s", v...)
		panic("")
	}
}

func FatalIfError(err error, format string, v ...any) {
	if err != nil {
		v = append(v, err.Error())
		log.Fatalf(format+"\nerror: %s", v...)
	}
}

func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func Println(a ...any) {
	fmt.Println(a...)
}

func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func LogFatal(v ...any) {
	log.Fatal(v...)
}

func LogFatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}

func LogFatalln(v ...any) {
	log.Fatalln(v...)
}

func LogPrintln(v ...any) {
	log.Println(v...)
}

func LogPrintf(format string, v ...any) {
	log.Printf(format, v...)
}
func LogPanic(v ...any) {
	log.Panic(v...)
}

func LogPanicf(format string, v ...any) {
	log.Panicf(format, v...)
}

func Marshal(pointer interface{}) ([]byte, error) {
	return json.Marshal(pointer)
}

func Unmarshal(data []byte, pointer interface{}) error {
	return json.Unmarshal(data, pointer)
}
