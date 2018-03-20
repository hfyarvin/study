package main

import (
	"os"
	"os/exec"
)

func main() {
	arv := []string{fmt.Sprintf("--host=%s", "127.0.0.1"), fmt.Sprintf("--port=%d", p.Port), "-uroot", "-p1234567890", "u103", "<", ".sql"}
	cmd := exec.Command("mysql", arv...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(out.String())
}
