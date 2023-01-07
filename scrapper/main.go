package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
var baseURL string = "https://www.saramin.co.kr/zf_user/salaries/total-salary/list"

// https://www.saramin.co.kr/zf_user/salaries/total-salary/view?csn=1078653665

type extractedJob struct {
	id    string
	title string
	csn   string
	info  string
}

func main() {
	var jobs []extractedJob
	fmt.Println("Scrapper Start")
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	writejobs(jobs)
	fmt.Println("Done, Extracted")
}

func writejobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "Title", "Csn", "Info"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{
			"https://www.saramin.co.kr/zf_user/salaries/total-salary/view?csn=" + job.csn,
			job.title,
			job.csn,
			job.info,
		}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "?page=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".company_info").Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Find(".link_tit").Attr("href")

	//rgx, _ := regexp.Compile("(?<=(csn=))([0-9]+)") 왜 안되는지 이유가 모르곘네 후방탐색 지원이 안되나??
	r, _ := regexp.Compile("(csn=[0-9]+)")
	csn := r.FindString(id)
	csn = strings.Split(csn, "=")[1]

	title, _ := card.Find(".link_tit").Attr("title")
	//fmt.Println(title, csn)

	info := cleanString(card.Find(".info_item>dd").Text())
	//fmt.Println(info)

	return extractedJob{
		id:    id,
		title: title,
		csn:   csn,
		info:  info,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}