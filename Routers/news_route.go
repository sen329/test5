package Routers

import (
	newsinfo "test5/Controller/news_info"
	middleware "test5/Middleware"

	"github.com/gorilla/mux"
)

func RouteNews(r *mux.Router) *mux.Router {

	route_news := r.PathPrefix("/news").Subrouter()
	route_news.Use(middleware.Middleware, middleware.CheckRoleNews)

	route_news.HandleFunc("/add", newsinfo.AddNews).Methods("POST")
	route_news.HandleFunc("/getAll", newsinfo.GetAllNews).Methods("GET")
	route_news.HandleFunc("/get", newsinfo.GetNews).Methods("GET")
	route_news.HandleFunc("/update", newsinfo.UpdateNews).Methods("PUT")
	route_news.HandleFunc("/delete", newsinfo.DeleteNews).Methods("DELETE")

	route_news.HandleFunc("/addDetail", newsinfo.AddNewsDetail).Methods("POST")
	route_news.HandleFunc("/getDetails", newsinfo.GetNewsDetails).Methods("GET")
	route_news.HandleFunc("/getDetail", newsinfo.GetNewsDetail).Methods("GET")
	route_news.HandleFunc("/updateDetailTitle", newsinfo.UpdateNewsTitle).Methods("PUT")
	route_news.HandleFunc("/updateDetailBanner", newsinfo.UpdateNewsBanner).Methods("PUT")
	route_news.HandleFunc("/updateDetailContent", newsinfo.UpdateNewsContent).Methods("PUT")
	route_news.HandleFunc("/deleteDetail", newsinfo.DeleteNewsDetail).Methods("DELETE")

	route_news.HandleFunc("/addDetailINTL", newsinfo.AddNewsDetailINTL).Methods("POST")
	route_news.HandleFunc("/getDetailsINTL", newsinfo.GetNewsDetailsINTL).Methods("GET")
	route_news.HandleFunc("/getDetailINTL", newsinfo.GetNewsDetailINTL).Methods("GET")
	route_news.HandleFunc("/updateDetailTitleINTL", newsinfo.UpdateNewsTitleINTL).Methods("PUT")
	route_news.HandleFunc("/updateDetailBannerINTL", newsinfo.UpdateNewsBannerINTL).Methods("PUT")
	route_news.HandleFunc("/updateDetailContentINTL", newsinfo.UpdateNewsContentINTL).Methods("PUT")
	route_news.HandleFunc("/deleteDetailINTL", newsinfo.DeleteNewsDetailINTL).Methods("DELETE")

	route_news.HandleFunc("/addImage", newsinfo.AddImage).Methods("POST")
	route_news.HandleFunc("/getImages", newsinfo.GetImages).Methods("GET")
	route_news.HandleFunc("/getImage", newsinfo.GetImage).Methods("GET")
	route_news.HandleFunc("/getFavImages", newsinfo.GetyourFavImages).Methods("GET")
	route_news.HandleFunc("/updateImages", newsinfo.UpdateImage).Methods("PUT")
	route_news.HandleFunc("/deleteImage", newsinfo.DeleteImage).Methods("DELETE")

	route_news.HandleFunc("/getTypes", newsinfo.GetNewsTypes).Methods("GET")
	route_news.HandleFunc("/getType", newsinfo.GetNewsType).Methods("GET")

	return r
}
