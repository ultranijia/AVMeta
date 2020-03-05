package util

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// WalkDir 遍历目录
func WalkDir(dirPath, success, fail string) ([]string, error) {
	// 定义文件列表
	var files []string

	// 遍历目录
	err := filepath.Walk(dirPath, func(filePath string, f os.FileInfo, err error) error {
		// 错误
		if f == nil {
			return err
		}
		// 忽略目录
		if f.IsDir() {
			return nil
		}

		// 检测是否为过滤目录
		if strings.Contains(
			strings.ToUpper(filePath),
			strings.ToUpper(success)) ||
			strings.Contains(
				strings.ToUpper(filePath),
				strings.ToUpper(fail)) {
			return nil
		}

		// 隐藏文件正则
		rHidden := regexp.MustCompile(`^\.(.)*`)
		// 检测是否为隐藏文件
		if rHidden.MatchString(f.Name()) {
			return nil
		}

		// 获取后缀并转换为小写
		ext := strings.ToLower(path.Ext(filePath))

		// 验证是否存在于后缀扩展名中
		if _, ok := videoExts[ext]; ok {
			// 存在则加入扩展
			files = append(files, filePath)
		}

		return nil
	})

	return files, err
}

// GetRunPath 获取当前程序执行路径
func GetRunPath() string {
	// 获取当前执行路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// 检查错误
	if err != nil {
		return "."
	}

	return dir
}