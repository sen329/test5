package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	controller "github.com/sen329/test5/Controller"
	middleware "github.com/sen329/test5/Middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var db *sql.DB

	controller.Open()

	defer db.Close()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router := mux.NewRouter()

	router.HandleFunc("/register", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")

	// ---- normal route ---- //
	route := router.PathPrefix("/").Subrouter()
	route.Use(middleware.Middleware)
	route.HandleFunc("/test", controller.Test).Methods("GET")

	//	icon_frame
	route.HandleFunc("/addIconFrame", controller.AddiconFrame).Methods("POST")
	route.HandleFunc("/getIconFrames", controller.GeticonFrames).Methods("GET")
	route.HandleFunc("/getIconFrame", controller.GeticonFrame).Methods("GET")
	route.HandleFunc("/updateIconFrame", controller.UpdateiconFrame).Methods("PUT")
	route.HandleFunc("/deleteIconFrame", controller.DeleteiconFrame).Methods("DELETE")

	//	icon_avatar
	route.HandleFunc("/addIconAvatar", controller.AddiconAvatar).Methods("POST")
	route.HandleFunc("/getIconAvatars", controller.GeticonAvatars).Methods("GET")
	route.HandleFunc("/getIconAvatar", controller.GeticonAvatar).Methods("GET")
	route.HandleFunc("/updateIconAvatar", controller.UpdateiconAvatar).Methods("PUT")
	route.HandleFunc("/deleteIconAvatar", controller.DeleteiconAvatar).Methods("DELETE")

	//	ksatriya
	route.HandleFunc("/addKsatriya", controller.AddnewKsatriya).Methods("POST")
	route.HandleFunc("/getKsatriyas", controller.GetKsatriyas).Methods("GET")
	route.HandleFunc("/getKsatriya", controller.GetKsatriya).Methods("GET")
	route.HandleFunc("/updateKsatriya", controller.UpdateKsatriya).Methods("PUT")
	route.HandleFunc("/deleteKsatriya", controller.DeleteKsatriya).Methods("DELETE")

	//	ksatriya_fragment
	route.HandleFunc("/addKsatriyaFragment", controller.AddKsatriyaFragment).Methods("POST")
	route.HandleFunc("/getKsatriyaFragments", controller.GetKsatriyaFragments).Methods("GET")
	route.HandleFunc("/getKsatriyaFragment", controller.GetKsatriyaFragment).Methods("GET")
	route.HandleFunc("/updateKsatriyaFragment", controller.UpdateKsatriyaFragment).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaFragment", controller.DeleteKsatriyaFragment).Methods("DELETE")

	// ksatriya_skin
	route.HandleFunc("/addKsatriyaSkin", controller.AddKsatriyaSkin).Methods("POST")
	route.HandleFunc("/getAllKsatriyaSkin", controller.GetAllKsatriyaSkin).Methods("GET")
	route.HandleFunc("/getKsatriyaSkin", controller.GetKsatriyaSkin).Methods("GET")
	route.HandleFunc("/updateKsatriyaSkin", controller.UpdateKsatriyaSkin).Methods("PUT")
	route.HandleFunc("/deleteKsatriyaSkin", controller.DeleteKsatriyaSkin).Methods("DELETE")

	//	ksatriya_skin_part
	route.HandleFunc("/addKsatriyaSkinPart", controller.AddKsatriyaSkinPart).Methods("POST")
	route.HandleFunc("/getKsatriyaSkinParts", controller.GetKsatriyaSkinParts).Methods("GET")
	route.HandleFunc("/getKsatriyaSkinPart", controller.GetKsatriyaSkinPart).Methods("GET")
	route.HandleFunc("/deleteKsatriyaSkinPart", controller.DeleteKsatriyaSkinPart).Methods("DELETE")

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
	route_mail.HandleFunc("/send", controller.Sendmail).Methods("POST")
	route_mail.HandleFunc("/get", controller.Getmails).Methods("GET")

	route_mail.HandleFunc("/createTemplate", controller.Createtemplate).Methods("POST")
	route_mail.HandleFunc("/getTemplates", controller.Gettemplates).Methods("GET")
	route_mail.HandleFunc("/getTemplate", controller.Gettemplate).Methods("GET")
	route_mail.HandleFunc("/updateTemplate", controller.Updatetemplates).Methods("PUT")
	route_mail.HandleFunc("/deleteTemplate", controller.DeleteTemplates).Methods("DELETE")

	route_mail.HandleFunc("/createCustom", controller.Createcustommail).Methods("POST")
	route_mail.HandleFunc("/getCustoms", controller.Getcustommails).Methods("GET")
	route_mail.HandleFunc("/getCustom", controller.Getcustommail).Methods("GET")
	route_mail.HandleFunc("/updateCustom", controller.Updatecustommail).Methods("PUT")
	route_mail.HandleFunc("/deleteCustom", controller.Deletecustommail).Methods("DELETE")

	route_mail.HandleFunc("/attachItem", controller.Attachitem).Methods("POST")
	route_mail.HandleFunc("/getAttachments", controller.Getmailattachments).Methods("GET")
	route_mail.HandleFunc("/getAttachment", controller.Getmailattachment).Methods("GET")
	route_mail.HandleFunc("/updateAttachment", controller.Updatemailattachment).Methods("PUT")
	route_mail.HandleFunc("/deleteAttachment", controller.Removeitem).Methods("DELETE")

	// ---- Lotto Subroute ---- //
	route_lotto := router.PathPrefix("/lotto").Subrouter()
	route_lotto.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_lotto.HandleFunc("/addLottos", controller.AddnewLotto).Methods("POST")
	route_lotto.HandleFunc("/getLottos", controller.GetallLottos).Methods("GET")

	route_lotto.HandleFunc("/addFeature", controller.AddlottoFeature).Methods("POST")
	route_lotto.HandleFunc("/getFeatures", controller.GetlottoFeatures).Methods("GET")
	route_lotto.HandleFunc("/getFeature", controller.GetlottoFeature).Methods("GET")
	route_lotto.HandleFunc("/getFeatureOf", controller.GetlottoFeatureByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateFeature", controller.UpdatelottoFeature).Methods("PUT")
	route_lotto.HandleFunc("/deleteFeature", controller.DeletelottoFeature).Methods("DELETE")

	route_lotto.HandleFunc("/addItem", controller.AddlottoItem).Methods("POST")
	route_lotto.HandleFunc("/getItems", controller.GetlottoItems).Methods("GET")
	route_lotto.HandleFunc("/getItem", controller.GetlottoItem).Methods("GET")
	route_lotto.HandleFunc("/updateItem", controller.UpdatelottoItem).Methods("PUT")
	route_lotto.HandleFunc("/deleteItem", controller.DeletelottoItem).Methods("DELETE")

	route_lotto.HandleFunc("/addColor", controller.AddlottoColor).Methods("POST")
	route_lotto.HandleFunc("/getColors", controller.GetlottoColors).Methods("GET")
	route_lotto.HandleFunc("/getColor", controller.GetlottoColor).Methods("GET")
	route_lotto.HandleFunc("/updateColor", controller.UpdatelottoColor).Methods("PUT")
	route_lotto.HandleFunc("/deleteColor", controller.DeletelottoColor).Methods("DELETE")

	route_lotto.HandleFunc("/addLoot", controller.AddlottoLoot).Methods("POST")
	route_lotto.HandleFunc("/getLoots", controller.GetlottoLoots).Methods("GET")
	route_lotto.HandleFunc("/getLoot", controller.GetlottoLoot).Methods("GET")
	route_lotto.HandleFunc("/getLootOf", controller.GetlottoLootByLottoId).Methods("GET")
	route_lotto.HandleFunc("/updateLoot", controller.UpdatelottoLoot).Methods("PUT")
	route_lotto.HandleFunc("/deleteLoot", controller.DeletelottoLoot).Methods("DELETE")

	// ---- Gacha Subroute ---- //
	route_gacha := router.PathPrefix("/gacha").Subrouter()
	route_gacha.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_gacha.HandleFunc("/add", controller.AddGacha).Methods("POST")
	route_gacha.HandleFunc("/getAll", controller.GetAllGacha).Methods("GET")
	route_gacha.HandleFunc("/get", controller.GetGacha).Methods("GET")
	route_gacha.HandleFunc("/update", controller.UpdateGacha).Methods("PUT")
	route_gacha.HandleFunc("/delete", controller.DeleteGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addItem", controller.AddGachaItem).Methods("POST")
	route_gacha.HandleFunc("/getAllItem", controller.GetAllGachaItem).Methods("GET")
	route_gacha.HandleFunc("/getItem", controller.GetGachaItem).Methods("GET")
	route_gacha.HandleFunc("/updateItem", controller.UpdateGachaItem).Methods("PUT")
	route_gacha.HandleFunc("/deleteItem", controller.DeleteGachaItem).Methods("DELETE")

	route_gacha.HandleFunc("/addFeatured", controller.AddFeaturedGacha).Methods("POST")
	route_gacha.HandleFunc("/getAllFeatured", controller.GetAllFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/getFeatured", controller.GetFeaturedGacha).Methods("GET")
	route_gacha.HandleFunc("/updateFeatured", controller.UpdateFeaturedGacha).Methods("PUT")
	route_gacha.HandleFunc("/deleteFeatured", controller.DeleteFeaturedGacha).Methods("DELETE")

	route_gacha.HandleFunc("/addLoot", controller.AddGachaLoot).Methods("POST")
	route_gacha.HandleFunc("/getAllLoot", controller.GetAllGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/getLoot", controller.GetGachaLoot).Methods("GET")
	route_gacha.HandleFunc("/updateLoot", controller.UpdateGachaLoot).Methods("PUT")
	route_gacha.HandleFunc("/deleteLoot", controller.DeleteGachaLoot).Methods("DELETE")

	// ---- Shop Subroute ---- //
	route_shop := router.PathPrefix("/shop").Subrouter()
	route_shop.Use(middleware.Middleware, middleware.CheckRoleShop)

	route_shop.HandleFunc("/lotus/add", controller.AddLotus).Methods("POST")
	route_shop.HandleFunc("/lotus/getAll", controller.GetAllLotus).Methods("GET")
	route_shop.HandleFunc("/lotus/get", controller.GetLotus).Methods("GET")
	route_shop.HandleFunc("/lotus/update", controller.UpdateLotusShop).Methods("PUT")
	route_shop.HandleFunc("/lotus/delete", controller.DeleteLotusShop).Methods("DELETE")

	route_shop.HandleFunc("/lotus/addItem", controller.LotusAddNewItem).Methods("POST")
	route_shop.HandleFunc("/lotus/getAllItem", controller.LotusGetShopItems).Methods("GET")
	route_shop.HandleFunc("/lotus/getItem", controller.LotusGetShopItem).Methods("GET")
	route_shop.HandleFunc("/lotus/updateItem", controller.LotusUpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/lotus/deleteItem", controller.LotusDeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/lotus/addPeriod", controller.AddLotusPeriod).Methods("POST")
	route_shop.HandleFunc("/lotus/getAllPeriod", controller.LotusGetShopPeriods).Methods("GET")
	route_shop.HandleFunc("/lotus/getPeriod", controller.LotusGetShopPeriod).Methods("GET")
	route_shop.HandleFunc("/lotus/updatePeriod", controller.LotusUpdateShopPeriod).Methods("PUT")
	route_shop.HandleFunc("/lotus/deletePeriod", controller.LotusDeleteShopPeriod).Methods("DELETE")

	route_shop.HandleFunc("/addItem", controller.AddShopItem).Methods("POST")
	route_shop.HandleFunc("/getAllItems", controller.GetShopItems).Methods("GET")
	route_shop.HandleFunc("/getItem", controller.GetShopItem).Methods("GET")
	route_shop.HandleFunc("/updateItem", controller.UpdateShopItem).Methods("PUT")
	route_shop.HandleFunc("/deleteItem", controller.DeleteShopItem).Methods("DELETE")

	route_shop.HandleFunc("/addBundle", controller.AddShopBundle).Methods("POST")
	route_shop.HandleFunc("/getAllBundles", controller.GetShopBundles).Methods("GET")
	route_shop.HandleFunc("/getBundle", controller.GetShopBundle).Methods("GET")
	route_shop.HandleFunc("/updateBundle", controller.UpdateShopBundle).Methods("PUT")
	route_shop.HandleFunc("/deleteBundle", controller.DeleteShopBundle).Methods("DELETE")

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
