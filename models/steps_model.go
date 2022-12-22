package models

type Steps struct {
	ID               string  `json:"id,omitempty" bson:"_id,omitempty"`
	Steps            float64 `json:"steps"`
	StepsType        string  `json:"stepsType"`
	StepsDate        string  `json:"stepsDate"`
	StepsDescription string  `json:"stepsDescription"`
}
