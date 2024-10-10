package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	
	repoKeuangan := "D:/smartlogy/smartkeuangan"
	repoDashboard := "D:/smartlogy/smartdashboard"

	fmt.Println("Memperbarui Smartkeuangan dan Smartdashboard...")

	fmt.Println("Menghapus semua unstagged files...")
	err := runCommand("git", "clean", "-fd")
	if err != nil {
		log.Fatalf("Error menghapus unstagged files: %v", err)
	}

	if err := updateRepo(repoKeuangan, "master"); err != nil {
		log.Fatalf("Gagal memperbarui SmartKeuangan: %v", err)
	}

	
	if err := updateRepo(repoDashboard, "main"); err != nil {
		log.Fatalf("Gagal memperbarui SmartDashboard: %v", err)
	}

	fmt.Println("SmartKeuangan dan SmartDashboard berhasil diperbarui.")

	fmt.Println("Tekan 'Enter' untuk keluar...")
	fmt.Scanln()
}


func updateRepo(repoPath, branch string) error {
	
	if err := changeDirectory(repoPath); err != nil {
		return fmt.Errorf("gagal mengubah direktori ke %s: %v", repoPath, err)
	}

	
	if err := gitPull(branch); err != nil {
		return fmt.Errorf("gagal melakukan git pull di %s: %v", repoPath, err)
	}

	return nil
}


func changeDirectory(repoPath string) error {
	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan path absolut: %v", err)
	}
	return os.Chdir(absPath)
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()

	return cmd.Run()
}


func gitPull(branch string) error {
	var stderr bytes.Buffer
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Stdout = log.Writer() 
	cmd.Stderr = &stderr      

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull gagal: %v, %s", err, stderr.String())
	}

	return nil
}
