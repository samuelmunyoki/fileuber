package server

import (
	"fmt"

	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func ServeFileWithSpeed(c *gin.Context, filePath string, fileName string, done chan bool) {
	file, err := os.Open(filePath)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "File Error", "message": err.Error()})
		done <- true
		return
	}
	defer file.Close()

	attachmentName := fmt.Sprintf("attachment; filename=%s", fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", attachmentName)

	buffer := make([]byte, 1024)
	startTime := time.Now()
	bytesSent := 0
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			elapsedTime := time.Since(startTime)
			speed := float64(bytesSent) / elapsedTime.Seconds() / 1024/1024 // Speed in MB/s
			AddDownload(fileName, speed)
		}
	}()

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			break
		}
		bytesSent += bytesRead
		c.Writer.Write(buffer[:bytesRead])
		c.Writer.Flush()
	}
	done <- true
}

func GetHandler(c *gin.Context) {
	listDirectories, err := GetDirectories(ShareCfg.Path)
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": "File Error", "message": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Main website", "path": ShareCfg.Path, "directories": listDirectories})
}

func PostTHandler(c *gin.Context) {
	directory := c.PostForm("directory")
	isfolder := c.PostForm("isfolder")
	joinedPath := filepath.Join(ShareCfg.Path, directory)

	if isfolder != "true" {
		done := make(chan bool)
		go ServeFileWithSpeed(c, joinedPath, directory, done)
		<-done
	} else {
		ShareCfg.UpdateShareConfig(joinedPath)
		listDirectories, err := GetDirectories(ShareCfg.Path)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{"error": "File Error", "message": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Main website", "path": ShareCfg.Path, "directories": listDirectories})
	}
}

func HomeHandler(c *gin.Context) {
	ShareCfg.ResetShareConfig()
	listDirectories, err := GetDirectories(ShareCfg.InitialPath)
	if err != nil {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{"error": "File Error", "message": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "Main website", "path": ShareCfg.Path, "directories": listDirectories})
}