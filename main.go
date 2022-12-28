package main

import "AMM/src/crawlers"

//"AMM/src"
//supfiles "AMM/sup_files"

func main() {
	//root_node := src.Read_HtmlFile()
	//result := src.SearchFirstNodeOccurrence(root_node, "ADD_DATE", "1663686573")
	//supfiles.Test()
	//supfiles.Test_readHtml()
	//supfiles.Test_reader()
	//supfiles.ReadMain()
	//routes.LoadRoutes()
	//http.ListenAndServe(":8080", nil)
	//src.GetAllTheMangasUrl()
	//supfiles.Mainrun()
	crawlers.GetLestEpUrl("https://asura.gg/manga/the-heavenly-demon-cant-live-a-normal-life/")
}
