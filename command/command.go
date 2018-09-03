package command

import (
	"fmt"
	"github.com/dutchakdev/clif/category"
	"github.com/dutchakdev/clif/db"
	"github.com/dutchakdev/clif/helpers"
	"github.com/dutchakdev/clif/process"
	"github.com/dutchakdev/clif/promt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"strings"
)

type cmd struct {
	id int
	name string
	command string
	category int
}

type cmdList []cmd

func ActionAddCmd(c *cli.Context) error {
	var cmd = strings.Join(os.Args[3:], " ")
	cats := append([]string{"Create new"}, category.GetCategories()...)
	selectedCategory := promt.ShowPromt("Select category", cats)
	var catId int
	var catName string

	if (selectedCategory == "Create new") {
		catId, catName = category.CreateCategory()
	} else {
		catId, catName = category.GetCategoryId(selectedCategory)
	}
	if (catId > 0) {
		prompt := promptui.Prompt{
			Label: "Enter command shortcut name",
			Default: cmd,
		}
		commandName, err := prompt.Run()
		helpers.CheckErr(err)

		if (len(commandName) > 0) {
			SaveCmd(commandName, cmd, catId)
			d := color.New(color.FgGreen, color.Bold)
			d.Printf("Command with name %v successfully added to category %v\n", commandName, catName)
			return nil
		}
	}
	return nil
}

func SaveCmd(name string, command string, category int) {
	database, _ := db.DbConnection()
	statement, _ := database.Prepare("INSERT INTO commands (name, command, category) VALUES (?, ?, ?)")
	_, err := statement.Exec(name, command, category)
	helpers.CheckErr(err)
}

func RunCmdByPath(path string, args []string){
	p := strings.Split(path, "/")
	if (len(p) == 2) {
		RunCmd(p[0], p[1], args)
	} else if (len(p) == 1) {
		FindCmdInCategories(p[0])
	}
}

var commandlistView []string

func FindCmdInCategories(selectedCatName string)  {
	var catId int
	if (len(selectedCatName) < 1) {
		selectedCatName = promt.ShowPromt("Select category", category.GetCategories())
	}
	catId, _ = category.GetCategoryId(selectedCatName)

	for _, v := range FindCmdByCat(catId) {
		commandlistView = append(commandlistView, "{"+strconv.Itoa(v.id)+"} [" + selectedCatName + "/" + v.name + "]->  " + v.command)
	}
	selectedCmd := promt.ShowPromt("Run command", commandlistView)
	var cmdId, _ = strconv.Atoi(strings.Replace(strings.Split(selectedCmd, "}")[0], "{", "", -1))

	var args = strings.Split(GetCommandById(cmdId), " ");
	fmt.Println(args)

	if proc, err := process.StartProcess(args); err == nil {
		proc.Wait()
	}
}

func RunCmd(categoryName string, cmdName string, addArgs []string) {
	catId, _ := category.GetCategoryId(categoryName)
	var args = strings.Split(GetCommand(catId, cmdName), " ");
	if proc, err := process.StartProcess(args); err == nil {
		proc.Wait()
	}

	return
}

func FindCmdByCat(catId int) cmdList {
	var id int
	var name string
	var command string
	var categoryId int

	database, _ := db.DbConnection()
	rows, _ := database.Query("SELECT id, name, command, category FROM commands WHERE category = ?", catId)
	cols, _ := rows.Columns()
	pointers := make([]interface{}, len(cols))
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	var itemList cmdList
	for rows.Next() {
		rows.Scan(&id, &name, &command, &categoryId)
		item := cmd{id: id, name: name, command: command, category: categoryId}
		itemList = append(itemList, item)
	}
	return itemList
}

func GetCommand(catId int, cmdName string) (string) {
	var command string
	database, _ := db.DbConnection()
	err := database.QueryRow("SELECT command FROM commands WHERE name = ? AND category = ?", cmdName, catId).Scan(&command)
	helpers.CheckErr(err)
	return command
}

func GetCommandById(cmdId int) (string) {
	var command string
	database, _ := db.DbConnection()
	err := database.QueryRow("SELECT command FROM commands WHERE id = ?", cmdId).Scan(&command)
	helpers.CheckErr(err)
	return command
}