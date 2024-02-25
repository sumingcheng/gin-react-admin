package logger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogEntry 定义日志条目的结构。
type LogEntry struct {
	Time    string `json:"time"`    // 时间
	Level   string `json:"level"`   // 日志级别
	Message string `json:"message"` // 消息内容
}

// Logger 提供日志记录功能。
type Logger struct {
	mu      sync.Mutex  // 互斥锁，用于同步
	info    *log.Logger // 信息级别的日志记录器
	error   *log.Logger // 错误级别的日志记录器
	logFile *os.File    // 当前日志文件
}

var instance *Logger // Logger的单例实例
var once sync.Once   // 确保单例只被初始化一次

// New 初始化或返回Logger的单例实例。
func New() *Logger {
	once.Do(func() {
		instance = &Logger{}
		instance.rotateLogFile()
		go instance.cleanupOldLogs()
	})
	return instance
}

// rotateLogFile 处理日志文件的轮转。
func (l *Logger) rotateLogFile() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.logFile != nil {
		_ = l.logFile.Close() // 尝试关闭当前日志文件，忽略错误
	}

	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("获取执行文件路径失败: %v", err)
	}
	execDir := filepath.Dir(execPath)
	logDir := filepath.Join(execDir, "logger") // 使用执行文件所在目录下的logger文件夹

	// 确保日志目录存在
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			log.Fatalf("创建日志目录失败: %v", err)
		}
	}

	// 创建新的日志文件
	filename := time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(filepath.Join(logDir, filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}
	l.logFile = file

	l.info = log.New(l.logFile, "INFO: ", log.LstdFlags)
	l.error = log.New(l.logFile, "ERROR: ", log.LstdFlags)
}

// Info 记录信息级别的日志。
func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

// Error 记录错误级别的日志。
func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

// log 处理不同级别的日志记录。
func (l *Logger) log(level, message string) {
	entry := LogEntry{
		Time:    time.Now().UTC().Format(time.RFC3339),
		Level:   level,
		Message: message,
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	entryData, err := json.Marshal(entry)
	if err != nil {
		log.Printf("序列化日志条目失败: %v", err)
		return
	}

	if level == "INFO" {
		l.info.Println(string(entryData))
	} else {
		l.error.Println(string(entryData))
	}
}

// cleanupOldLogs 定期清理旧的日志文件。
func (l *Logger) cleanupOldLogs() {
	ticker := time.NewTicker(24 * time.Hour)
	for range ticker.C {
		execPath, _ := os.Executable()
		execDir := filepath.Dir(execPath)
		logDir := filepath.Join(execDir, "logger")

		files, err := filepath.Glob(filepath.Join(logDir, "*.log"))
		if err != nil {
			log.Printf("列出日志文件失败: %v", err)
			continue
		}

		for _, file := range files {
			info, err := os.Stat(file)
			if err != nil {
				log.Printf("获取日志文件状态失败: %v", err)
				continue
			}
			if time.Since(info.ModTime()).Hours() > 15*24 {
				if err := os.Remove(file); err != nil {
					log.Printf("删除日志文件失败: %v", err)
				}
			}
		}
	}
}
