# bcraft-test-task

Запуск  
```
docker compose up --build
```
------------------------------------------------------------
Получение списка всех рецептов  
GET localhost:8000/recipes
------------------------------------------------------------
Создание нового рецепта  
POST localhost:8000/recipe
```
{
"name": "NewName",
"description": "NewDdesc",
"ingredients": "NewIngr",
"cooking_steps": "NewCookingSteps"
}
```
------------------------------------------------------------
Получение информации о конкретном рецепте по ID  
GET localhost:8000/recipe?id={id}

------------------------------------------------------------
Редактирование рецепта по ID  
PUT localhost:8000/recipe
```
{
	"recipe_id": 3,
	"name": "UpdateName",
	"description": "UpdateDesc",
	"ingredients": "UpdateIngr",
	"cooking_steps": "UpdateCookingSteps"
}
```
------------------------------------------------------------
Удаление рецепта по ID  
DELETE localhost:8000/recipe?id={id}