package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sut63/team05/controllers"
	_ "github.com/sut63/team05/docs"
	"github.com/sut63/team05/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Hospital struct input data
type Hospitals struct {
	Hospital []Hospital
}

// Hospital struct
type Hospital struct {
	HospitalName string
}

// Member struct input data
type Members struct {
	Member []Member
}

// Member struct
type Member struct {
	MemberEmail    string
	MemberName     string
	MemberPassword string
}

// Genders struct input data
type Genders struct {
	Gender []Gender
}

// Gender struct
type Gender struct {
	GenderName string
}

// GroupOfAges struct input data
type GroupOfAges struct {
	GroupOfAge []GroupOfAge
}

// GroupOfAge struct
type GroupOfAge struct {
	GroupOfAgeName string
	GroupOfAgeAge  string
}

// Officers struct input data
type Officers struct {
	Officer []Officer
}

// Officer struct
type Officer struct {
	OfficerEmail    string
	OfficerName     string
	OfficerPassword string
}

// MoneyTransfers struct input data
type MoneyTransfers struct {
	MoneyTransfer []MoneyTransfer
}

// MoneyTransfer struct
type MoneyTransfer struct {
	MoneytransferType string
}

// Banks struct input data
type Banks struct {
	Bank []Bank
}

// Bank struct
type Bank struct {
	BankType string
}

// Amountpaids struct input data
type Amountpaids struct {
	Amountpaid []Amountpaid
}

// Amountpaid struct
type Amountpaid struct {
	AmountpaidMoney int
}

// Category struct input data
type Categorys struct {
	Category []Category
}

// Gender struct
type Category struct {
	CategoryName string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:Health.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewMemberController(v1, client)
	controllers.NewInsuranceController(v1, client)
	controllers.NewHospitalController(v1, client)
	controllers.NewGenderController(v1, client)
	controllers.NewGroupOfAgeController(v1, client)
	controllers.NewOfficerController(v1, client)
	controllers.NewProductController(v1, client)
	controllers.NewMoneyTransferController(v1, client)
	controllers.NewAmountpaidController(v1, client)
	controllers.NewcategoryController(v1, client)

	// Set Members Data
	members := Members{
		Member: []Member{
			Member{"b6115296@g.sut.ac.th", "Teerapat Saiprom", "123456789"},
			Member{"b6132552@g.sut.ac.th", "Teerasuk Supawaha", "123456789"},
		},
	}

	for _, m := range members.Member {
		client.Member.
			Create().
			SetMemberEmail(m.MemberEmail).
			SetMemberName(m.MemberName).
			SetMemberPassword(m.MemberPassword).
			Save(context.Background())
	}

	// Set Hospitals Data
	hospitals := Hospitals{
		Hospital: []Hospital{
			Hospital{"Suranaree"},
			Hospital{"Maharat (Nakhon ratchasima)"},
		},
	}

	for _, h := range hospitals.Hospital {
		client.Hospital.
			Create().
			SetHospitalName(h.HospitalName).
			Save(context.Background())
	}

	// Set Offices Data
	officers := Officers{
		Officer: []Officer{
			Officer{"gamse0505@gmail.com", "Somchai Ngaosri", "Aa123"},
			Officer{"Panyaporn@gmail.com", "Panyaporn Ngaosri", "Bb123"},
		},
	}

	for _, ofc := range officers.Officer {
		client.Officer.
			Create().
			SetOfficerEmail(ofc.OfficerEmail).
			SetOfficerName(ofc.OfficerName).
			SetOfficerPassword(ofc.OfficerPassword).
			Save(context.Background())
	}

	// Set GroupOfAges Data
	groupofages := GroupOfAges{
		GroupOfAge: []GroupOfAge{
			GroupOfAge{"เด็ก", "10 - 15 ปี"},
			GroupOfAge{"วัยรุ่น", "16 - 29 ปี"},
			GroupOfAge{"ผู้ใหญ่", "30 - 59 ปี"},
			GroupOfAge{"ผู้สูงอายุ", "60 ปีขึ้นไป"},
		},
	}

	for _, goa := range groupofages.GroupOfAge {

		client.GroupOfAge.
			Create().
			SetGroupOfAgeName(goa.GroupOfAgeName).
			SetGroupOfAgeAge(goa.GroupOfAgeAge).
			Save(context.Background())
	}

	// Set Genders Data
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, gd := range genders.Gender {

		client.Gender.
			Create().
			SetGenderName(gd.GenderName).
			Save(context.Background())
	}

	// Set MoneyTransfers Data
	moneytransfers := MoneyTransfers{
		MoneyTransfer: []MoneyTransfer{
			MoneyTransfer{"Internet banking"},
			MoneyTransfer{"Moblie banking"},
			MoneyTransfer{"ATM"},
		},
	}

	for _, mn := range moneytransfers.MoneyTransfer {

		client.MoneyTransfer.
			Create().
			SetMoneytransferType(mn.MoneytransferType).
			Save(context.Background())
	}

	// Set Banks Data
	banks := Banks{
		Bank: []Bank{
			Bank{"ธนาคารกสิกรไทย"},
			Bank{"ธนาคารกรุงเทพ"},
			Bank{"ธนาคารกรุงศรีอยุธยา"},
			Bank{"ธนาคารไทยพาณิชย์"},
			Bank{"ธนาคารกรุงไทย"},
			Bank{"ธนาคารทหารไทย"},
			Bank{"ธนาคารธนาชาต"},
			Bank{"ธนาคารออมสิน"},
		},
	}

	for _, b := range banks.Bank {

		client.Bank.
			Create().
			SetBankType(b.BankType).
			Save(context.Background())
	}

	// Set Amountpaid Data
	amountpaids := Amountpaids{
		Amountpaid: []Amountpaid{
			Amountpaid{35000.},
			Amountpaid{57000},
			Amountpaid{82500},
			Amountpaid{90000},
			Amountpaid{110000},
		},
	}

	for _, ap := range amountpaids.Amountpaid {

		client.Amountpaid.
			Create().
			SetAmountpaidMoney(ap.AmountpaidMoney).
			Save(context.Background())
	}

	// Set Categorys Data
	categorys := Categorys{
		Category: []Category{
			Category{"สนใจผลิตภัณฑ์ประกันสุขภาพ"},
			Category{"สอบถามข้อมูลผลตอบแทน"},
			Category{"สอบถามการชำระเบี้ยประกัน"},
		},
	}

	for _, cg := range categorys.Category {
		client.Category.
			Create().
			SetCategoryName(cg.CategoryName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
