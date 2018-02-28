package services

import (
	"ec/models"
	"net/http"
	"log"
	"bytes"
	"fmt"
	"io/ioutil"
)

func PostECResultsToMaster (url string, ecResults []models.ECResult) string {

	var jsonStr = []byte(models.ToJson(ecResults))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("Error posting to master %v",err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
