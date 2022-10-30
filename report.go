package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("headless", true))
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Minute*3)
	defer cancel()

	f, err := os.OpenFile("./report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)

	if err := chromedp.Run(ctx, chromedp.Tasks{
		report("https://wxapp.nwpu.edu.cn/uc/api/oauth/index?redirect=https://yqtb.nwpu.edu.cn/wx/common/metaWeiXin_new.jsp&appid=200200204192458714&state=1"),
	}); err != nil {
		log.Fatalln(err)
	}

	log.Println("report success")
}

func report(url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#app`, chromedp.ByID),
		chromedp.SetValue(`[type=text]`, os.Getenv("student_id"), chromedp.ByQuery),
		chromedp.SetValue(`[type=password]`, os.Getenv("password"), chromedp.ByQuery),
		chromedp.Click(`.btn`, chromedp.ByQuery),
		chromedp.WaitVisible(`#form1`, chromedp.ByID),
		chromedp.Sleep(10 * time.Second),
		chromedp.Evaluate(`go_subfx();document.getElementById("brcn").checked = true;savefx();`, nil),
		chromedp.Sleep(3 * time.Second),
	}
}
