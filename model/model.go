package model

type EventCheckin struct {
	Phone 		string `dynamodbav:"phone"`
	EventCode 	string `dynamodbav:"event_code"`
	Email 		string `dynamodbav:"email"`
	Name 		string `dynamodbav:"name"`
	EventVersion string `dynamodbav:"event_version"`
}

type Result struct {
	Name string
	Count int
}