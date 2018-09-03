package helpers

import "github.com/manifoldco/promptui"

func ShowSelect(label string, items []string) (string)  {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	CheckErr(err)

	return result
}
