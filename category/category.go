package category

import (
	"fmt"
	"github.com/dutchakdev/clif/db"
	"github.com/dutchakdev/clif/globals"
	"github.com/dutchakdev/clif/helpers"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func CreateCategory() (int, string) {
	if (len(globals.CategoryName) < 1) {
		prompt := promptui.Prompt{
			Label: "Enter category name:",
		}

		categoryName, err := prompt.Run()
		helpers.CheckErr(err)
		fmt.Println(categoryName)
	}

	database, _ := db.DbConnection()
	statement, _ := database.Prepare("INSERT INTO categories (name) VALUES (?)")
	res, err := statement.Exec(globals.CategoryName)
	helpers.CheckErr(err)

	id, err := res.LastInsertId();
	helpers.CheckErr(err)

	d := color.New(color.FgGreen, color.Bold)
	d.Printf("Category %v (id: %v) successfully created.\n", globals.CategoryName, id)

	return int(id), globals.CategoryName
}

func RemoveCategory(categoryName string){
	if (globals.Interactive == "false") {
		prompt := promptui.Prompt{
			Label:     "Are you sure to delete this item?:",
			IsConfirm: true,
		}
		_, err := prompt.Run()
		helpers.CheckErr(err)
	}

	database, _ := db.DbConnection()
	statement, _ := database.Prepare("DELETE FROM categories WHERE name = ?")
	_, err := statement.Exec(categoryName)
	helpers.CheckErr(err)

	d := color.New(color.FgGreen, color.Bold)
	d.Printf("Category %v successfully removed.\n", categoryName)
	return
}

func GetCategoryId(name string) (int, string)  {
	var id int;
	database, _ := db.DbConnection()
	err := database.QueryRow("SELECT id FROM categories WHERE name = ?", name).Scan(&id)
	helpers.CheckErr(err)
	return id, name
}

func GetCategories() []string {
	var result []string
	var id int
	var name string

	database, _ := db.DbConnection()
	rows, _ := database.Query("select id, name from categories")
	cols, _ := rows.Columns()
	pointers := make([]interface{}, len(cols))
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	for rows.Next() {
		rows.Scan(&id, &name)
		result = append(result, name)
	}
	return result
}
