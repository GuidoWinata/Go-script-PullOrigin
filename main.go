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
	
	repoPath := "D:/Sinau-Kang/Learn-Elysia/absensi"

	if err := changeDirectory(repoPath); err != nil {
		log.Fatalf("Gagal mengubah direktori: %v", err)
	}

	
	if err := gitPull(); err != nil {
		log.Fatalf("Gagal melakukan git pull: Tidak ada koneksi internet")
	}

	fmt.Println("SmartKeuangan berhasil diperbarui.")

	fmt.Println("Tekan 'Enter' untuk keluar...")
	fmt.Scanln()
}


func changeDirectory(repoPath string) error {
	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return fmt.Errorf("gagal mendapatkan path absolut: %v", err)
	}
	
	return os.Chdir(absPath)
}


func gitPull() error {
	var stderr bytes.Buffer
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Stdout = log.Writer()
	cmd.Stderr = &stderr
 
	if err := cmd.Run(); err != nil {
		return fmt.Errorf(stderr.String())
	}
	return nil
}