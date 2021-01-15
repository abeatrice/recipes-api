package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{}

	// create aws session & dynamodb service
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	// create scan request & send it
	req, output := svc.ScanRequest(&dynamodb.ScanInput{
		TableName: aws.String("Recipes"),
	})
	err := req.Send()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = "error"
		response.Message = fmt.Sprintf("failed to send dynamodb scan request: %v", err.Error())
		json, err := json.Marshal(response)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(json)
		return
	}

	// create slice of items from scan output
	recipes := make([]Recipe, 0)
	for _, item := range output.Items {
		recipes = append(recipes, Recipe{
			ID:          *item["ID"].S,
			Name:        *item["Name"].S,
			Description: *item["Description"].S,
		})
	}

	response.Status = "success"
	response.Data = recipes

	// create response json
	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// set status ok and write response
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
