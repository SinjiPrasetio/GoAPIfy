package core

import (
	"GoAPIfy/tools/core/color"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Docker(s string) {
	switch s {
	case "up":
		PrintLogo()
		DockerUp()
	case "down":
		PrintLogo()
		DockerDown()
	default:
		fmt.Println(color.Colorize(color.Red, "Error argument, must be 'up' or 'down'."))
		os.Exit(0)
	}
}

func DockerUp() {
	fmt.Println(color.Colorize(color.Magenta, "Configuring Docker..."))
	os.Remove("./docker-compose.yml")
	var path string
	production := os.Getenv("APP_PRODUCTION")
	if production == "true" {
		path = "./tools/templates/docker/docker-compose.yml.production"
	} else {
		path = "./tools/templates/docker/docker-compose.yml.development"
	}
	src, err := os.Open(path)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, "GoAPIfy is corrupted, core files is missing!"))
		os.Exit(0)
	}
	srcByte, err := ioutil.ReadAll(src)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}

	out, err := os.Create("./docker-compose.yml")
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}
	defer out.Close()

	_, err = out.Write(srcByte)
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
		os.Exit(0)
	}

	fmt.Println(color.Colorize(color.Magenta, "Launching Docker..."))
	cmd := exec.Command("docker-compose", "up", "-d")

	cmd.Stdout = nil
	cmd.Stderr = nil

	err = cmd.Run()
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
	}

	fmt.Println(color.Colorize(color.Magenta, "Docker Launched!"))

}

func DockerDown() {
	fmt.Println(color.Colorize(color.Magenta, "Stopping Docker..."))
	cmd := exec.Command("docker-compose", "down")

	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Run()
	if err != nil {
		fmt.Println(color.Colorize(color.Red, err.Error()))
	}

	fmt.Println(color.Colorize(color.Magenta, "Docker Removed!"))
}
