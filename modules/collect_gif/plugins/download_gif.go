package plugins

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func DownloadGifFromPaper(articleURL string) {
	// 配置参数
	targetDir := "./wechat_gifs" // 保存目录

	// 创建保存目录
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		log.Fatalf("创建目录失败: %v", err)
	}

	// 获取并解析公众号文章
	doc, err := fetchHTMLDocument(articleURL)
	if err != nil {
		log.Fatal(err)
	}

	// 查找所有包含 data-src 属性的元素
	doc.Find("[data-src]").Each(func(i int, s *goquery.Selection) {
		dataSrc, exists := s.Attr("data-src")
		if !exists {
			return
		}

		// 处理 URL（处理相对路径）
		imgURL, err := resolveURL(articleURL, dataSrc)
		if err != nil {
			log.Printf("URL解析失败: %s, 错误: %v", dataSrc, err)
			return
		}

		// 检查是否为 GIF 图片
		if !isGifURL(imgURL) {
			return
		}

		// 下载图片
		imgData, err := downloadImage(imgURL)
		if err != nil {
			log.Printf("下载失败: %s, 错误: %v", imgURL, err)
			return
		}

		// 生成唯一文件名
		filename := generateFilename(imgURL)
		savePath := filepath.Join(targetDir, filename)

		// 保存文件
		if err := os.WriteFile(savePath, imgData, 0644); err != nil {
			log.Printf("保存失败: %s, 错误: %v", savePath, err)
			return
		}

		log.Printf("成功保存: %s (%d bytes)", savePath, len(imgData))
	})
	log.Printf("提取结束啦小彭~")
}

// 获取并解析 HTML 文档
func fetchHTMLDocument(url string) (*goquery.Document, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 设置浏览器 User-Agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 错误: %s", resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}

// 解析完整 URL
func resolveURL(base, src string) (string, error) {
	u, err := url.Parse(src)
	if err != nil {
		return "", err
	}

	// 如果是完整 URL 直接返回
	if u.IsAbs() {
		return src, nil
	}

	// 处理相对路径
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	return baseURL.ResolveReference(u).String(), nil
}

// 检查是否为 GIF 图片 URL
func isGifURL(url string) bool {
	lowerURL := strings.ToLower(url)
	return strings.Contains(lowerURL, ".gif") ||
		strings.Contains(lowerURL, "wx_fmt=gif") ||
		strings.Contains(lowerURL, "format=gif")
}

// 下载图片
func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP 错误: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// 生成唯一文件名
func generateFilename(imgURL string) string {
	// 使用 MD5 哈希确保文件名唯一
	hasher := md5.New()
	hasher.Write([]byte(imgURL))
	hash := hex.EncodeToString(hasher.Sum(nil))

	// 提取原始文件名（如果有）
	u, err := url.Parse(imgURL)
	if err == nil {
		path := u.Path
		if ext := filepath.Ext(path); ext != "" {
			return hash[:8] + ext // 保留原始扩展名
		}
	}

	return hash[:8] + ".gif" // 默认使用 .gif 扩展名
}
