package models

type (
    // Report struct
    Report struct {
        Name   		string  `json:"firstName"`
        Instances int     `json:"instances"`
        Gender		string  `json:"gender"`
        Accuracy 	string  `json:"accuracy"`
    }
)
