package promt

import (
	"github.com/dutchakdev/clif/helpers"
	"github.com/manifoldco/promptui"
)

func ShowPromt(label string, items []string) (string)  {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	helpers.CheckErr(err)

	return result
}