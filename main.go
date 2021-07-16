package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	"github.com/sen329/test5/Controller/icon"
	"github.com/sen329/test5/Controller/ksatriya"
	"github.com/sen329/test5/Controller/mail"
	shop "github.com/sen329/test5/Controller/shop"
	"github.com/sen329/test5/Controller/shop/gacha"
	"github.com/sen329/test5/Controller/shop/lotto"
	"github.com/sen329/test5/Controller/shop/lotus"
	middleware "github.com/sen329/test5/Middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := controller.Open()

	defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	// ---- normal route ---- //
	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")

	//	icon_frame
	route.HandleFunc("/addIconFrame", icon.AddiconFrame).Methods("POST")
	route.HandleFunc("/getIconFrames", icon.GeticonFrames).Methods("GET")
	route.HandleFunc("/getIconFrame", icon.GeticonFrame).Methods("GET")
	route.HandleFunc("/updateIconFrame", icon.UpdateiconFrame).Methods("PUT")
	route.HandleFunc("/deleteIconFrame", icon.DeleteiconFrame).Methods("DELETE")

	//	icon_avatar
	route.HandleFunc("/addIconAvatar", icon.AddiconAvatar).Methods("POST")
	route.HandleFunc("/getIconAvatars", icon.GeticonAvatars).Methods("GET")
	route.HandleFunc("/getIconAvatar", icon.GeticonAvatar).Methods("GET")
	route.HandleFunc("/updateIconAvatar", icon.UpdateiconAvatar).Methods("PUT")
	route.HandleFunc("/deleteIconAvatar", icon.DeleteiconAvatar).Methods("DELETE")

	//	ksatriya
	route.HandleFunc("/addKsatriya", ksatriya.AddnewKsatriya).Methods("POST")
	route.HandleFunc("/getKsatriyas", ksatriya.GetKsatriyas).Methods("GET")
	route.HandleFunc("/getKsatriya", ksatriya.GetKsatriya).Methods("GET")
	route.HandleFunc("/updateKsatriya", ksatriya.UpdateKsatriya).Methods("PUT")
	route.HandleFunc("/deleteKsatriya", ksatriya.DeleteKsatriya).Methods("DELETE")

	//	ksatriya_fragment
	route.HandleFunc("/addKsatriyaFragment", ksatriya.AddKsatriyaFragment).Methods("POST")
	route.HandleFunc("/getKsatriyaFragments", ksatriya.GetKsatriyaFragments).Methods("GET")
	route.HandleFunc("/getKsatriyaFragment", ksatriya.GetKsatriyaFragment).Methods("GET")
	route.HandleFunc("/updateKsatriyaFragment", ksatriya.UpdateKsatriyaFragment).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaFragment", ksatriya.DeleteKsatriyaFragment).Methods("DELETE")

	// ksatriya_skin
	route.HandleFunc("/addKsatriyaSkin", ksatriya.AddKsatriyaSkin).Methods("POST")
	route.HandleFunc("/getAllKsatriyaSkin", ksatriya.GetAllKsatriyaSkin).Methods("GET")
	route.HandleFunc("/getKsatriyaSkin", ksatriya.GetKsatriyaSkin).Methods("GET")
	route.HandleFunc("/updateKsatriyaSkin", ksatriya.UpdateKsatriyaSkin).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaSkin", ksatriya.DeleteKsatriyaSkin).Methods("DELETE")

	//	ksatriya_skin_part
	route.HandleFunc("/addKsatriyaSkinPart", ksatriya.AddKsatriyaSkinPart).Methods("POST")
	route.HandleFunc("/getKsatriyaSkinParts", ksatriya.GetKsatriyaSkinParts).Methods("GET")
	route.HandleFunc("/getKsatriyaSkinPart", ksatriya.GetKsatriyaSkinPart).Methods("GET")
	route.HandleFunc("/deleteKsatriyaSkinPart", ksatriya.DeleteKsatriyaSkinPart).Methods("DELETE")

	//	rune
	route.HandleFunc("/addRune", controller.AddRune).Methods("POST")
	route.HandleFunc("/getRunes", controller.GetRunes).Methods("GET")
	route.HandleFunc("/getRune", controller.GetRune).Methods("GET")
	route.HandleFunc("/updateRune", controller.UpdateRune).Methods("PUT")
	route.HandleFunc("/deleteRune", controller.DeleteRune).Methods("DELETE")

	//	premium
	route.HandleFunc("/addPremium", controller.AddPremium).Methods("POST")
	route.HandleFunc("/getPremiums", controller.GetPremiums).Methods("GET")
	route.HandleFunc("/getPremium", controller.GetPremium).Methods("GET")
	route.HandleFunc("/updatePremium", controller.UpdatePremium).Methods("PUT")
	route.HandleFunc("/deletePremium", controller.DeletePremium).Methods("DELETE")

	//	energy
	route.HandleFunc("/addEnergy", controller.AddEnergy).Methods("POST")
	route.HandleFunc("/getEnergies", controller.GetEnergies).Methods("GET")
	route.HandleFunc("/getEnergy", controller.GetEnergy).Methods("GET")
	route.HandleFunc("/updateEnergy", controller.UpdateEnergy).Methods("PUT")
	route.HandleFunc("/deleteEnergy", controller.DeleteEnergy).Methods("DELETE")

	//currency type
	route.HandleFunc("/addCurrency", controller.AddCurrencyType).Methods("POST")
	route.HandleFunc("/getCurrencies", controller.GetAllCurrencyTypes).Methods("GET")
	route.HandleFunc("/getCurrency", controller.GetCurrencyType).Methods("GET")
	route.HandleFunc("/updateCurrency", controller.UpdateCurrencyType).Methods("PUT")
	route.HandleFunc("/deleteCurrency", controller.DeleteCurrencyType).Methods("DELETE")

	//box
	route.HandleFunc("/addBox", controller.AddBox).Methods("POST")
	route.HandleFunc("/getBoxes", controller.GetAllBox).Methods("GET")
	route.HandleFunc("/getBox", controller.GetBox).Methods("GET")
	route.HandleFunc("/updateBox", controller.UpdateBox).Methods("PUT")
	route.HandleFunc("/deleteBox", controller.DeleteBox).Methods("DELETE")

	route.HandleFunc("/addBoxLoot", controller.AddBoxLoot).Methods("POST")
	route.HandleFunc("/getBoxLoots", controller.GetAllBoxLoot).Methods("GET")
	route.HandleFunc("/getBoxLoot", controller.GetBoxLoot).Methods("GET")
	route.HandleFunc("/updateBoxLoot", controller.UpdateBoxLoot).Methods("PUT")
	route.HandleFunc("/deleteBoxLoot", controller.DeleteBoxLoot).Methods("DELETE")

	//chest
	route.HandleFunc("/addChest", controller.AddChest).Methods("POST")
	route.HandleFunc("/getChests", controller.GetAllChest).Methods("GET")
	route.HandleFunc("/getChest", controller.GetChest).Methods("GET")
	route.HandleFunc("/updateChest", controller.UpdateChest).Methods("PUT")
	route.HandleFunc("/deleteChest", controller.DeleteChest).Methods("DELETE")

	// ---- Mail Subroute ---- //
	route_mail := router.PathPrefix("/mail").Subrouter()
	route_mail.Use(middleware.Middleware, middleware.CheckRoleMail)
	route_mail.HandleFunc("/send", mail.Sendmail).Methods("POST")
	route_mail.HandleFunc("/get", mail.Getmails).Methods("GET")

	route_mail.HandleFunc("/createTemplate", mail.Createtemplate).Methods("POST")
	route_mail.HandleFunc("/getTemplates", mail.Gettemplates).Methods("GET")
	route_mail.HandleFunc("/getTemplate", mail.Gettemplate).Methods("GET")
	route_mail.HandleFunc("/updateTemplate", mail.Updatetemplates).Methods("PUT")
	route_mail.HandleFunc("/deleteTemplate", mail.DeleteTemplates).Methods("DELETE")

	route_mail.HandleFunc("/createCustom", mail.Createcustommail).Methods("POST")
	route_mail.HandleFunc("/getCustoms", mail.Getcustommails).Methods("GET")
	route_mail.HandleFunc("/getCustom", mail.Getcustommail).Methods("GET")
	route_mail.HandleFunc("/updateCustom", mail.Updatecustommail).Methods("PUT")
	route_mail.HandleFunc("/deleteCustom", mail.Deletecustommail).Methods("DELETE")

	route_mail.HandleFunc("/attachItem", mail.Attachitem).Methods("POST")
	route_mail.HandleFunc("/getAttachments", mail.Getmailattachments).Methods("GET")
	route_mail.HandleFunc("/getAttachment", mail.Getmailattachment).Methods("GET")
	route_mail.HandleFunc("/updateAttachment", mail.Updatemailattachment).Methods("PUT")
	route_mail.HandleFunc("/deleteAttachment", mail.Removeitem).Methods("DELETE")

	// ---- Lotto Subroute ---- //
	route_lotto := router.PathPrefix("/lotto").Subrouter()
	route_lotto.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_lotto.HandleFunc("/addLottos", lotto.AddnewLotto).Methods("POST")
	route_lotto.HandleFunc("/getLottos", lotto.GetallLottos).Methods("GET")

	route_lotto.HandleFunc("/addFeature", lotto.AddlottoFeature).Methods("POST")
	route_lotto.HandleFunc("/getFeatures", lotto.GetlottoFeatures).Methods("GET")
	route_lotto.HandleFunc("/getFeature", lotto.GetlottoFeature).Methods("GET")
	route_lotto.HandleFunc("/getFeatureOf", lotto.GetlottoFeatureByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateFeature", lotto.UpdatelottoFeature).Methods("PUT")
	route_lotto.HandleFunc("/deleteFeature", lotto.DeletelottoFeature).Methods("DELETE")

	route_lotto.HandleFunc("/addItem", lotto.AddlottoItem).Methods("POST")
	route_lotto.HandleFunc("/getItems", lotto.GetlottoItems).Methods("GET")
	route_lotto.HandleFunc("/getItem", lotto.GetlottoItem).Methods("GET")
	route_lotto.HandleFunc("/updateItem", lotto.UpdatelottoItem).Methods("PUT")
	route_lotto.HandleFunc("/deleteItem", lotto.DeletelottoItem).Methods("DELETE")

	route_lotto.HandleFunc("/addColor", lotto.AddlottoColor).Methods("POST")
	route_lotto.HandleFunc("/getColors", lotto.GetlottoColors).Methods("GET")
	route_lotto.HandleFunc("/getColor", lotto.GetlottoColor).Methods("GET")
	route_lotto.HandleFunc("/updateColor", lotto.UpdatelottoColor).Methods("PUT")
	route_lotto.HandleFunc("/deleteColor", lotto.DeletelottoColor).Methods("DELETE")

	route_lotto.HandleFunc("/addLoot", lotto.AddlottoLoot).Methods("POST")
	route_lotto.HandleFunc("/getLoots", lotto.GetlottoLoots).Methods("GET")
	route_lotto.HandleFunc("/getLoot", lotto.GetlottoLoot).Methods("GET")
	route_lotto.HandleFunc("/getLootOf", lotto.GetlottoLootByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateLoot", lotto.UpdatelottoLoot).Methods("PUT")
	route_lotto.HandleFunc("/deleteLoot", lotto.DeletelottoLoot).Methods("DELETE")

	// ---- Gacha Subroute ---- //
	route_gacha := router.PathPrefix("/gacha").Subrouter()
	route_gacha.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_gacha.HandleFunc("/add", gacha.AddGacha).Methods("POST")
	route_gacha.HandleFunc("/getAll", gacha.GetAllGacha).Methods("GET")
	route_gacha.HandleFunc("/get", gacha.GetGacha).Methods("GET")
	route_gacha.HandleFunc("/update", gacha.UpdateGacha).Methods("PUT")
	route_gacha.HandleFunc("/delete", gacha.DeleteGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addItem", gacha.AddGachaItem).Methods("POST")
	route_gacha.HandleFunc("/getAllItem", gacha.GetAllGachaItem).Methods("GET")
	route_gacha.HandleFunc("/getItem", gacha.GetGachaItem).Methods("GET")
	route_gacha.HandleFunc("/updateItem", gacha.UpdateGachaItem).Methods("PUT")
	route_gacha.HandleFunc("/deleteItem", gacha.DeleteGachaItem).Methods("DELETE")

	route_gacha.HandleFunc("/addFeatured", gacha.AddFeaturedGacha).Methods("POST")
	route_gacha.HandleFunc("/getAllFeatured", gacha.GetAllFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/getFeatured", gacha.GetFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/updateFeatured", gacha.UpdateFeaturedGacha).Methods("PUT")
	route_gacha.HandleFunc("/deleteFeatured", gacha.DeleteFeaturedGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addLoot", gacha.AddGachaLoot).Methods("POST")
	route_gacha.HandleFunc("/getAllLoot", gacha.GetAllGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/getLoot", gacha.GetGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/updateLoot", gacha.UpdateGachaLoot).Methods("PUT")
	route_gacha.HandleFunc("/deleteLoot", gacha.DeleteGachaLoot).Methods("DELETE")

	// ---- Shop Subroute ---- //
	route_shop := router.PathPrefix("/shop").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("/lotus/add", lotus.AddLotus).Methods("POST")
	route_shop.HandleFunc("/lotus/getAll", lotus.GetAllLotus).Methods("GET")
	route_shop.HandleFunc("/lotus/get", lotus.GetLotus).Methods("GET")
	route_shop.HandleFunc("/lotus/update", lotus.UpdateLotusShop).Methods("PUT")
	route_shop.HandleFunc("/lotus/delete", lotus.DeleteLotusShop).Methods("DELETE")

	route_shop.HandleFunc("/lotus/addItem", lotus.LotusAddNewItem).Methods("POST")
	route_shop.HandleFunc("/lotus/getAllItem", lotus.LotusGetShopItems).Methods("GET")
	route_shop.HandleFunc("/lotus/getItem", lotus.LotusGetShopItem).Methods("GET")
	route_shop.HandleFunc("/lotus/updateItem", lotus.LotusUpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/lotus/deleteItem", lotus.LotusDeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/lotus/addPeriod", lotus.AddLotusPeriod).Methods("POST")
	route_shop.HandleFunc("/lotus/getAllPeriod", lotus.LotusGetShopPeriods).Methods("GET")
	route_shop.HandleFunc("/lotus/getPeriod", lotus.LotusGetShopPeriod).Methods("GET")
	route_shop.HandleFunc("/lotus/updatePeriod", lotus.LotusUpdateShopPeriod).Methods("PUT")
	route_shop.HandleFunc("/lotus/deletePeriod", lotus.LotusDeleteShopPeriod).Methods("DELETE")

	route_shop.HandleFunc("/addItem", shop.AddShopItem).Methods("POST")
	route_shop.HandleFunc("/getAllItems", shop.GetShopItems).Methods("GET")
	route_shop.HandleFunc("/getItem", shop.GetShopItem).Methods("GET")
	route_shop.HandleFunc("/updateItem", shop.UpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", shop.DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", shop.AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", shop.GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", shop.GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", shop.UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", shop.DeleteShopBundle).Methods("DELETE")

	// ---- Role Subroute ---- //
	route_role := router.PathPrefix("/role").Subrouter()
	route_role.Use(middleware.Middleware, middleware.CheckRoleUser)

	//roles
	route_role.HandleFunc("/add", controller.AddRoles).Methods("POST")
	route_role.HandleFunc("/getAll", controller.GetAllRoles).Methods("GET")
	route_role.HandleFunc("/get", controller.GetRole).Methods("GET")
	route_role.HandleFunc("/update", controller.UpdateRole).Methods("PUT")
	route_role.HandleFunc("/delete", controller.DeleteRole).Methods("DELETE")

	//roles permission control
	route_role.HandleFunc("/addPermissionToRole", controller.AddNewPermissionToRole).Methods("POST")
	route_role.HandleFunc("/getAllRolesPermission", controller.GetAllRolesPermissions).Methods("GET")
	route_role.HandleFunc("/getRolePermission", controller.GetRolePermission).Methods("GET")
	route_role.HandleFunc("/removePermissionFromRole", controller.RemovePermissionFromRole).Methods("DELETE")
	route_role.HandleFunc("/getAllPermission", controller.GetAllPermissions).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
