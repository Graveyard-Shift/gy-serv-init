package core

import (
	"fmt"
	"gy-serv-init/src/files"
	"io/ioutil"
	"os"
	"os/exec"
)

func Main(pwd, projectName *string) {
	projectDir := fmt.Sprintf("%s/%s", *pwd, *projectName)

	mkdir(projectDir)

	cmd := exec.Command("go", "mod", "init", *projectName)
	cmd.Dir = projectDir
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(out)

	mkdir(fmt.Sprintf("%s/%s/src", *pwd, *projectName))
	mkdir(fmt.Sprintf("%s/%s/src/controllers", *pwd, *projectName))
	mkdir(fmt.Sprintf("%s/%s/src/router", *pwd, *projectName))

	touch(
		fmt.Sprintf("%s/%s/src/controllers/main.controllers.go", *pwd, *projectName),
		[]byte(files.ControllersMain),
	)

	touch(
		fmt.Sprintf("%s/%s/src/controllers/utils.go", *pwd, *projectName),
		[]byte(files.ControllersUtilities),
	)

	touch(
		fmt.Sprintf("%s/%s/src/router/router.go", *pwd, *projectName),
		[]byte(files.Router(projectName)),
	)

	touch(
		fmt.Sprintf("%s/%s/main.go", *pwd, *projectName),
		[]byte(files.Main(projectName)),
	)

	touch(
		fmt.Sprintf("%s/%s/.editorconfig", *pwd, *projectName),
		[]byte(files.EditorConfig),
	)

	touch(
		fmt.Sprintf("%s/%s/.gitignore", *pwd, *projectName),
		[]byte(files.Gitignore(projectName)),
	)
}

func touch(path string, data []byte) {
	err := ioutil.WriteFile(path, data, 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v\n", err)
	}
}

func mkdir(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		panic(err)
	}
}
