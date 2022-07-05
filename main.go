package main

import (
	"fmt"
	//"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//var db *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Profile struct {
	gorm.Model
	User        User
	UserID      int
	Channel     Channel
	ChannelID   int
	Description string
}

type Channel struct {
	gorm.Model
	ChannelName string
	ContentType string
}

type Message struct {
	gorm.Model
	Message string
	UserID  int
	User    User
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Profile{}, &Channel{}, &Message{})
	seed(db)
}

func seed(db *gorm.DB) {
	channels := []Channel{{ChannelName: "Anitube", ContentType: "Anime Music Video"},
		{ChannelName: "Techzone", ContentType: "tech gyan"},
		{ChannelName: "Golang Ninja", ContentType: "Golang tutorial"}}

	db.Create(&channels)
	//var ch Channel
	//for _, channel := range channels {
	//fmt.Println("channel created -> ", i)

	//ch1 := db.Find(&ch,1)
	//fmt.Println(ch1)
	//}

	users := []User{{Name: "Sumit", Email: "sumit@gmail.com"},
		{Name: "Arti", Email: "arti@gmail.com"},
		{Name: "Mustii", Email: "mustii@gmail.com"}}

	//for _, user := range users {
	db.Create(&users)
	//}

	var anitube, techzone Channel
	db.First(&anitube, 1)
	db.First(&techzone, 2)

	var sumit, Mustii User
	db.First(&sumit, "Name = ?", "Sumit")
	db.First(&Mustii, "Name = ?", "Mustii")

	msgs := []Message{{Message: "what u doing?", User: sumit},
		{Message: "Coding", User: Mustii},
		{Message: "Good Luck", User: Mustii}}

	for i, msg := range msgs {
		db.Create(&msg)
		fmt.Println("mssg created ->", i, "msssg: ", msg)
	}
	var simpMsg Message
	db.First(&simpMsg, 1)

	profiles := []Profile{{User: sumit, Channel: anitube, Description: "sumit has anitube channel"},
		{User: Mustii, Channel: techzone, Description: "mustii has techzone channel"},
		{User: Mustii, Channel: anitube, Description: "msutii also has anitube channel"}}

	for _, profile := range profiles {
		db.Create(&profile)
	}

	var prfl Profile
	db.Find(&prfl)

}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	//defer db.Close()
	//db.LogMode(true)
	setup(db)
	var users []User
	db.First(&users, "Name = ?", "Mustii")
	for i, user := range users {
		fmt.Println("Name: ", user.Name, "Email: ", user.Email, "index -> ", i)
		//fmt.Println("Newuser -> ", i)
	}
	var msgs []Message
	db.Last(&msgs, Message{UserID: 3})
	for _, msg := range msgs {
		fmt.Println("Message: ", msg.Message)
	}
}

/*func doError(err *error){

}*/

// root:B32dp$@1@tcp(127.0.0.1:3306)go-crud?charset=utf8mb4
