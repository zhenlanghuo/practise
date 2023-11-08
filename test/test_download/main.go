package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/ledongthuc/pdf"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var year = flag.Int("year", 2003, "year")
var month = flag.Int("month", 1, "month")
var day = flag.Int("day", 1, "day")
var hour = flag.Int("hour", 1, "hour")
var min = flag.Int("min", 0, "min")
var sec = flag.Int("sec", 0, "second")
var duration = flag.Int("duration", 3, "duration min")
var timeout = flag.Int("timeout", 2000, "timeout ms")
var keyword = flag.String("keyword", "September 14", "")
var concurrent = flag.Int("concurrent", 5, "")

func main() {

	flag.Parse()

	l, _ := time.LoadLocation("Asia/Shanghai")
	ts := time.Date(*year, time.Month(*month), *day, *hour, *min, *sec, 0, l).Unix()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(ts-time.Now().Unix()))
	<-ctx.Done()

	ctx, _ = context.WithTimeout(context.Background(), time.Minute*time.Duration(*duration))

	wg := sync.WaitGroup{}

	for i := 0; i < *concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				download()
			}
		}()
	}

	wg.Wait()
	fmt.Println("finish")

}

func download() {
	start := time.Now()

	imgUrl := "https://www.dol.gov/ui/data.pdf"

	//imgUrl := "https://adpemploymentreport.com/"

	// Get the data
	http.DefaultClient.Timeout = time.Millisecond * time.Duration(*timeout)
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println(time.Now(), err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(time.Now(), err)
		return
	}
	//fmt.Println(time.Now(), len(data), time.Since(start))

	content, _ := readPdfBytes(data)

	//fmt.Println(time.Now(), time.Since(start))
	fmt.Println(fmt.Sprintf("%v, %v, %v", time.Now(), strings.Contains(content, *keyword), time.Since(start)))

}

func readPdfBytes(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	r, err := pdf.NewReader(reader, reader.Size())
	// remember close file
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
