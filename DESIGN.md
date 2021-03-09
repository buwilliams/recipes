# Design

## Goals

- Ad-free recipes
- Create recipe for your profile
- Calculate the calories for recipe
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