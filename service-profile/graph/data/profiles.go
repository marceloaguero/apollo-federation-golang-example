package data

import "github.com/marceloaguero/apollo-federation-golang-example/service-profile/graph/model"

var Profiles = []*model.Profile{
	{
		UserID: "1",
		Age:    "20",
		Phone:  "00-0000-0001",
		Job:    "front-end developer",
	},
	{
		UserID: "2",
		Age:    "21",
		Phone:  "00-0000-0002",
		Job:    "back-end developer",
	},
	{
		UserID: "3",
		Age:    "20",
		Phone:  "00-0000-0003",
		Job:    "ios developer",
	},
}
