package main


userAgent = []string{
	""
}

func RandomUserAgent() {

}

func crawl() {

}



func main(){
	worklist := make(chan []string)
	baseDomain := "https://www.theguardian.com/"
	go func(){worklist <- []string{"https://www.theguardian.com/"}}()

	seen := make(map[string]bool)

	list := worklist
}