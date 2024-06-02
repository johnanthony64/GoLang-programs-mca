package main

import (
	"encoding/json"
	"fmt"
)

type Restaurant struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Menu      Menu   `json:"menu"`
	OpenHour  int    `json:"openHour"`
	CloseHour int    `json:"closeHour"`
}

type Menu struct {
	Appetizers  []string `json:"appetizers"`
	MainCourses []string `json:"mainCourses"`
	Desserts    []string `json:"desserts"`
}

func main() {

	restaurant := Restaurant{
		Name:    "Shiro",
		Address: "14, Vittal Malya Road, Bengaluru -01",
		Menu: Menu{
			Appetizers:  []string{"Samosa", "Paneer Tikka"},
			MainCourses: []string{"Butter Chicken", "Palak Paneer"},
			Desserts:    []string{"Gulab Jamun", "Kulfi"},
		},
		OpenHour:  11,
		CloseHour: 22,
	}

	
	jsonData, err := json.Marshal(restaurant)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	fmt.Println("JSON data:", string(jsonData))


	var unmarshalledRestaurant Restaurant
	err = json.Unmarshal(jsonData, &unmarshalledRestaurant)
	if err != nil {
		fmt.Println("Error unmarshaling data:", err)
		return
	}

	fmt.Println("Unmarshalled Restaurant:", unmarshalledRestaurant)
}
