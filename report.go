package main

import (
	"context"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	f, err := os.OpenFile("./report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	if err := chromedp.Run(ctx, chromedp.Tasks{
		login("https://uis.nwpu.edu.cn/cas/login"),
		report("https://yqtb.nwpu.edu.cn"),
	}); err != nil {
		log.Println(err)
	}

	log.Println("report success.")
}

// login
func login(loginUrl string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.WaitVisible(`.card`, chromedp.ByQuery),
		chromedp.SetValue(`#username`, os.Getenv("student_id"), chromedp.ByID),
		chromedp.SetValue(`#password`, os.Getenv("password"), chromedp.ByID),
		chromedp.Evaluate(`submitFm1()`, nil),
		chromedp.WaitVisible(`.services`, chromedp.ByQuery),
	}
}

// report
func report(reportUrl string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(reportUrl),
		chromedp.WaitVisible(`.page`, chromedp.ByQuery),
		chromedp.Evaluate(`go_subfx();document.getElementById("brcn").checked = true;savefx();`, nil),
	}
}
