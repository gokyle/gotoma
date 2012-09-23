package main

import "os/exec"

func speak(text string) (err error) {
	var path string
	var cmd *exec.Cmd

	path, err = exec.LookPath("say")
	if err != nil {
		return
	}

	cmd = exec.Command(path, text)
	err = cmd.Run()
	return
}

func speaker(text string) (err error) {
	if shouldSpeak {
		err = speak(text)
	}
	return
}
