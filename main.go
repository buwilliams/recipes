package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

func main() {
	db := Database{}
	db.Load("data")

	fmt.Println(db)

	fmt.Println("Users:")
	for _, u := range db.Users {
		fmt.Println("\t", strconv.Itoa(u.ID)+") ", u.Name)
	}

	fmt.Println("Recipes:")
	for _, r := range db.Recipes {
		fmt.Println("\t", strconv.Itoa(r.ID)+") ", r.Name)
	}
}

type Database struct {
	Users            []User
	Foods            []Food
	Recipes          []Recipe
	Ingredients      []Ingredient
	NextUserID       int
	NextFoodID       int
	NextIngredientID int
	NextRecipeID     int
}

func (db *Database) Load(path string) (err error) {
	err = LoadUsers(db, path)
	err = LoadFoods(db, path)
	err = LoadIngredients(db, path)
	err = LoadRecipes(db, path)
	return
}

func LoadUsers(db *Database, path string) (err error) {
	f, err := os.Open(path + "/user.csv")
	if err != nil {
		return
	}
	defer f.Close()

	gocsv.UnmarshalFile(f, &db.Users)

	db.NextUserID = len(db.Users) + 1

	return
}

func LoadFoods(db *Database, path string) (err error) {
	f, err := os.Open(path + "/food.csv")
	if err != nil {
		return
	}
	defer f.Close()

	gocsv.UnmarshalFile(f, &db.Foods)

	db.NextFoodID = len(db.Foods) + 1

	return
}

func LoadIngredients(db *Database, path string) (err error) {
	f, err := os.Open(path + "/ingredient.csv")
	if err != nil {
		return
	}
	defer f.Close()

	gocsv.UnmarshalFile(f, &db.Ingredients)

	db.NextIngredientID = len(db.Ingredients) + 1

	return
}

func LoadRecipes(db *Database, path string) (err error) {
	f, err := os.Open(path + "/recipe.csv")
	if err != nil {
		return
	}
	defer f.Close()

	gocsv.UnmarshalFile(f, &db.Recipes)

	db.NextRecipeID = len(db.Recipes) + 1

	return
}

type User struct {
	ID    int    `csv:"id"`
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Pin   string `csv:"pin"`
}

type Food struct {
	ID             int    `csv:"id"`
	Name           string `csv:"name"`
	Description    string `csv:"description"`
	CaloriesPerCup int    `csv:"caloriespercup"`
	PhotoURL       string `csv:"photourl"`
}

type Recipe struct {
	ID          int    `csv:"id"`
	UserID      int    `csv:"userid"`
	Name        string `csv:"name"`
	Description string `csv:"description"`
	PhotoURL    string `csv:"photourl"`
}

type Ingredient struct {
	ID       int `csv:"id"`
	RecipeID int `csv:"recipeid"`
	FoodID   int `csv:"foodid"`
	Quantity int `csv:"quantity"`
}
