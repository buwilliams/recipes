package main

import "database/sql"

// User defines shape of user table
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Pin   string `json:"pin"`
}

// Tweet defines shape of tweet table
type Tweet struct {
	ID      int    `json:"id"`
	UserID  string `json:"userId"`
	Created string `json:"created"`
	Content string `json:"content"`
}

// GetAllUsers will return []User
func (u User) GetAllUsers(db *sql.DB) (users []User, err error) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Email, &user.Pin)
		users = append(users, user)
	}
	defer rows.Close()
	return
}

// GetAllTweets will return []User
func (t Tweet) GetAllTweets(db *sql.DB) (tweets []Tweet, err error) {
	rows, err := db.Query("SELECT * FROM tweet")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tweet Tweet
		rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Created, &tweet.Content)
		tweets = append(tweets, tweet)
	}
	defer rows.Close()
	return
}
