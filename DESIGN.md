# Design

## Goals

- Place to store and share recipes ad-free
- Manage recipe for your profile
- Share recipes with URL
- Automatically calculate the calories for your recipe by serving size
- Standardize all calories by 1 cup

## Measurement Conversions

- Cup = 1
- 1/2 Cup = 0.5000
- 1/3 Cup = 0.3333
- 1/4 Cup = 0.2500
- 1/8 Cup = 0.1250
- Tablespoon = 0.0625
- Teaspoon = 0.0208
- 1/2 Teaspoon = 0.0104
- 1/4 Teaspoon = 0.0052
- Pinch = 0.0013

## Models

**User**

- ID
- Email
- Pin

**Food**

- ID
- Name
- Description
- CaloriesPerCup
- PhotoURL

**Recipe**

- ID
- UserID
- Name
- Description
- PhotoURL

**Ingredients**

- ID
- RecipeID
- FoodID
- Quantity (1, 2, 0.5, etc)

## Todo

- Setup ESLint
- UI: Sign-up / sign-in page
- UI: Landing page for user (search for recipe, list of recipes, delete recipe, create new recipe)
- UI: Recipe landing page (show recipe)
- UI: Edit Recipe Page
- API: POST /signin
- API: GET /signout
- API: POST /signup
- API: GET /recipes
- Web hosting for Go
- Web hosting Dolt
- SSL Encryption

## Done

- Design Data Model
- Setup Dolt Repo
- Import CSV into Dolt and Commit
- Setup Nuxt UI