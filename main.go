package main

import "wizzy/core"

// todo: edit logic
// core logic
func main() {

	core.Run()

	/*
		_, err := reader.ReadTemplate("./test/")
		if err != nil {
			fmt.Println("WIZZY ERROR, e:", err.Error())
			return
		}

		params := []model.Param{
			{Id: "feature", Value: "user"},
			{Id: "name", Value: "Login"},
			{Id: "type", Value: "GET"},
			{Id: "objectOut", Value: "User"},
			{Id: "params", Value: "username, password"},
		}
		rule := model.Rule{
			Rule:        "file",
			Origin:      "{{name}}Activity.java",
			Destination: "./test/features/{{feature}}",
			Path:        "./test",
		}

		err = writter.CanWriteFile(rule, params)
		if err != nil {
			fmt.Println("WIZZY ERROR, verifying conditions for execution e:", err.Error())
			return
		}

		err = writter.WriteFile(rule, params)
		if err != nil {
			fmt.Println("WIZZY ERROR, creating file e:", err.Error())
			return
		}

	*/
}
