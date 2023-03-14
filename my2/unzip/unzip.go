package unzip

import (
	//"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"my2/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/yeka/zip"

	simplifiedchinese "golang.org/x/text/encoding/simplifiedchinese"
	transform "golang.org/x/text/transform"
)

func NoPasswdUnzip(path string, OutfilePath string) error {
	// password := ""
	// for _, p := range config.Config.Passwords {
	// 	if p != "" {
	// 		password = p
	// 		break
	// 	}
	// }

	// 1、使用zip.OpenReader打开zip文件
	archive, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer func(archive *zip.ReadCloser) {
		err := archive.Close()
		if err != nil {

		}
	}(archive)
	// 2、循环访问 zip 中的文件 zip.File 切片
	for _, f := range archive.File {
		// if f.IsEncrypted() {
		// 	f.SetPassword(password)
		// }
		filePath := filepath.Join(OutfilePath, cn(f.Name))
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(OutfilePath)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return nil
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		// 3、使用 zip.File.Open 方法读取 zip 中文件的内容
		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		// 4、使用 io.Copy 或 io.Writer.Write 保存解压后的文件内容
		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		// 5、使用 zip.Reader.Close 关闭 zip 文件
		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}
	return nil
}

func ExtractZip(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("读取目录失败：", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			ExtractZip(filepath.Join(path, file.Name()))
		} else if strings.HasSuffix(file.Name(), ".zip") {
			zipPath := filepath.Join(path, file.Name())
			fmt.Println(zipPath)

			err := NoPasswdUnzip(zipPath, *(config.Config.OutfilePath))
			if err != nil {
				fmt.Printf("解压 %s 失败：%s\n", zipPath, err)
			} else {
				fmt.Printf("解压 %s 成功\n", zipPath)
			}
		}
	}
}

func Unzip(zipPath, s string) {
	panic("unimplemented")
}

func cn(str string) string {
	// 将GB2312编码转换为UTF-8编码
	reader := transform.NewReader(strings.NewReader(str), simplifiedchinese.GB18030.NewDecoder())
	buf, err := io.ReadAll(reader)
	if err != nil {
		return str
	}
	return string(buf)
}
