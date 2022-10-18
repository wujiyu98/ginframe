package entity

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primarykey;type:int(11)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Meta struct {
	Title       string `gorm:"size:255;default:''"`
	Keywords    string `gorm:"size:120;default:''"`
	Description string `gorm:"size:255;default:''"`
}

type ArticleCategory struct {
	Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	ParentID  uint   `gorm:"type:int(11);not null;default:0"`
	SortOrder uint   `gorm:"type:int(11);default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

type Address struct {
	Model
	UserID    uint   `gorm:"type:int(11);not null"`
	CountryID uint   `gorm:"type:int(11);not null"`
	StateID   uint   `gorm:"type:int(11);not null"`
	CityID    uint   `gorm:"type:int(11);not null"`
	Firstname string `gorm:"type:varchar(20);not null"`
	Lastname  string `gorm:"type:varchar(20);not null"`
	Company   string `gorm:"type:varchar(255);default:''"`
	Postcode  string `gorm:"type:varchar(30);default:''"`
	Address1  string `gorm:"type:varchar(255);default:''"`
	Address2  string `gorm:"type:varchar(255);default:''"`
}

type Article struct {
	Model
	LanguageID        uint   `gorm:"type:int(8);not null"`
	ArticleCategoryID uint   `gorm:"type:int(11);not null"`
	Title             string `gorm:"size:255;not null;unique"`
	Pathname          string `gorm:"size:255;not null;unique"`
	SortOrder         uint   `gorm:"default:0"`
	Showed            byte   `gorm:"default:1"`
	Summary           string `gorm:"type:text;"`
	Image             string `gorm:"size:255;default:''"`
	Author            string `gorm:"size:60;default:''"`
	Content           string `gorm:"type:longtext"`
	Meta              Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

type Banner struct {
	Model
	Name    string `gorm:"type:varchar(255);default:''"`
	Image   string `gorm:"type:varchar(255);default:''"`
	Url     string `gorm:"type:varchar(255);default:''"`
	Title   string `gorm:"type:varchar(255);default:''"`
	Summary string `gorm:"type:varchar(255);default:''"`
}

type Category struct {
	Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	ParentID  uint   `gorm:"type:int(11);not null;default:0"`
	SortOrder uint   `gorm:"type:int(11);default:0"`
	Qty       uint   `gorm:"default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

type Manufacturer struct {
	Model
	Name      string `gorm:"size:255;not null"`
	Pathname  string `gorm:"size:255;not null;unique"`
	SortOrder uint   `gorm:"type:int(11);default:0"`
	Qty       uint   `gorm:"default:0"`
	Image     string `gorm:"size:255;default:''"`
	Summary   string `gorm:"type:text;"`
	Meta      Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

type CategoryManufacturer struct {
	CategoryID     uint `gorm:"type:int(11);primaryKey;autoIncrement:false"`
	ManufacturerID uint `gorm:"type:int(11);primaryKey;autoIncrement:false"`
}

type Language struct {
	Model
	Name   string `gorm:"size:60;not null"`
	Domain string `gorm:"size:120;not null"`
	Code   string `gorm:"size:10;not null"`
	Image  string `gorm:"size:255;default:''"`
}

type Message struct {
	Model
	Name        string `gorm:"varchar(60);not null"`
	Email       string `gorm:"varchar(60);not null"`
	Country     string `gorm:"varchar(60);not null"`
	MobilePhone string `gorm:"varchar(30);not null"`
	Company     string `gorm:"varchar(120);default:''"`
	Comment     string `gorm:"varchar(255);not null"`
}

type Enquiry struct {
	Model
	Name        string `gorm:"varchar(60);not null"`
	Email       string `gorm:"varchar(60);not null"`
	Country     string `gorm:"varchar(60);not null"`
	MobilePhone string `gorm:"varchar(30);not null"`
	Company     string `gorm:"varchar(120);default:''"`
	Comment     string `gorm:"varchar(255);not null"`
}

type EnquiryProduct struct {
	Title        string  `gorm:"type:varchar(255);not null"`
	Manufacturer string  `gorm:"type:varchar(255);not null;default:''"`
	Summary      string  `gorm:"type:varchar(255);not null;default:''"`
	Qty          uint    `gorm:"type:int(11);not null;default:0"`
	Price        float64 `gorm:"decimal(10,2);not null;default:0"`
}

type Order struct {
	Model
	UserID            uint    `gorm:"type:int(11);not null"`
	OrderNumber       string  `gorm:"type:char(15);not null"`
	Address           string  `gorm:"size:500;not null"`
	TransactionNumber string  `gorm:"size:255;not null;default ''"`
	PaypalFee         float64 `gorm:"type:decimal(10,2);not null;default:0"`
	Total             float64 `gorm:"type:decimal(10,2);not null"`
	Payment           byte    `gorm:"not null;default:0"`
	Status            byte    `gorm:"not null;default:0"`
	Freight           float64 `gorm:"type:decimal(10,2);not null;default:35"`
	Express           string  `gorm:"size:255;not null;default:''"`
	TrackingNumber    string  `gorm:"size:255;not null;default:''"`
	CreatedAt         time.Time
}

type OrderProduct struct {
	Model
	OrderID      uint    `gorm:"not null"`
	ProductID    uint    `gorm:"type:int(11);not null"`
	URL          string  `gorm:"not null;default:''"`
	Title        string  `gorm:"type:varchar(255);not null"`
	Image        string  `gorm:"type:varchar(255);not null;default:''"`
	Manufacturer string  `gorm:"type:varchar(255);not null;default:''"`
	Summary      string  `gorm:"type:varchar(255);not null;default:''"`
	Qty          uint    `gorm:"type:int(11);not null"`
	Price        float64 `gorm:"decimal(10,2);not null"`
}

type Product struct {
	Model
	LanguageID     uint                `gorm:"type:int(8);not null"`
	CategoryID     uint                `gorm:"type:int(11);not null"`
	ManufacturerID uint                `gorm:"type:int(11);not null"`
	Title          string              `gorm:"size:255;not null;unique"`
	Pathname       string              `gorm:"size:255;not null;unique"`
	SortOrder      uint                `gorm:"type:int(11);default:0"`
	Showed         *byte               `gorm:"default:1"`
	Summary        string              `gorm:"type:text;"`
	Image          string              `gorm:"size:255;default:''"`
	Images         []string            `gorm:"type:json;serializer:json"`
	Content        string              `gorm:"type:longtext"`
	Stock          uint                `gorm:"type:int(11);default:0"`
	Hot            byte                `gorm:"default:0"`
	New            byte                `gorm:"default:0"`
	Special        byte                `gorm:"default:0"`
	Price          float64             `gorm:"type:decimal(10,4)"`
	Prices         []map[string]string `gorm:"type:json;serializer:json"`
	Meta           Meta                `gorm:"embedded;embeddedPrefix:meta_"`
}

type Attribute struct {
	Model
	CategoryID uint
	Name       string
}

type ProductAttribute struct {
	ProductID   uint   `gorm:"type:int(11);primaryKey;autoIncrement:false"`
	AttributeID uint   `gorm:"type:int(11);primaryKey;autoIncrement:false"`
	Text        string `gorm:"type:varchar(255);default:''"`
}

type User struct {
	Model
	Email           string `gorm:"type:varchar(255)"`
	Name            string `gorm:"type:varchar(60)"`
	Country         string `gorm:"type:varchar(60)"`
	Company         string `gorm:"type:varchar(255)"`
	Password        string `gorm:"type:varchar(255)"`
	RememberToken   string `gorm:"type:varchar(255)"`
	MobilePhone     string `gorm:"type:varchar(20)"`
	EmailVerifiedAt time.Time
}

type Seo struct {
	Model
	LanguageID uint   `gorm:"type:int(8);not null"`
	Name       string `gorm:"size:255"`
	Pathname   string `gorm:"size:255"`
	Meta       Meta   `gorm:"embedded;embeddedPrefix:meta_"`
}

type SiteInfo struct {
	Model
	LanguageID  uint   `gorm:"type:int(8);not null"`
	SiteName    string `gorm:"size:255;default:''"`
	Company     string `gorm:"size:255;default:''"`
	Contact     string `gorm:"size:255;default:''"`
	Phone       string `gorm:"size:20;default:''"`
	Phone2      string `gorm:"size:20;default:''"`
	MobilePhone string `gorm:"size:20;default:''"`
	Skype       string `gorm:"size:60;default:''"`
	QQ          string `gorm:"size:20;default:''"`
	Whatsapp    string `gorm:"size:20;default:''"`
	Address     string `gorm:"size:255;default:''"`
}

type City struct {
	Model
	Name        string  `gorm:"column:name;type:varchar(255);not null"`
	StateID     string  `gorm:"index:cities_test_ibfk_1;column:state_id;type:mediumint(8) unsigned;not null"`
	StateCode   string  `gorm:"column:state_code;type:varchar(255);not null"`
	CountryID   string  `gorm:"index:cities_test_ibfk_2;column:country_id;type:mediumint(8) unsigned;not null"`
	CountryCode string  `gorm:"column:country_code;type:char(2);not null"`
	Latitude    float64 `gorm:"column:latitude;type:decimal(10,8);not null"`
	Longitude   float64 `gorm:"column:longitude;type:decimal(11,8);not null"`
	Flag        bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID  string  `gorm:"column:wikiDataId;type:varchar(255)"`
}

// Countries [...]
type Country struct {
	Model
	Name           string  `gorm:"column:name;type:varchar(100);not null"`
	Iso3           string  `gorm:"column:iso3;type:char(3)"`
	NumericCode    string  `gorm:"column:numeric_code;type:char(3)"`
	Iso2           string  `gorm:"column:iso2;type:char(2)"`
	Phonecode      string  `gorm:"column:phonecode;type:varchar(255)"`
	Capital        string  `gorm:"column:capital;type:varchar(255)"`
	Currency       string  `gorm:"column:currency;type:varchar(255)"`
	CurrencyName   string  `gorm:"column:currency_name;type:varchar(255)"`
	CurrencySymbol string  `gorm:"column:currency_symbol;type:varchar(255)"`
	Tld            string  `gorm:"column:tld;type:varchar(255)"`
	Native         string  `gorm:"column:native;type:varchar(255)"`
	Region         string  `gorm:"column:region;type:varchar(255)"`
	Subregion      string  `gorm:"column:subregion;type:varchar(255)"`
	Timezones      string  `gorm:"column:timezones;type:text"`
	Translations   string  `gorm:"column:translations;type:text"`
	Latitude       float64 `gorm:"column:latitude;type:decimal(10,8)"`
	Longitude      float64 `gorm:"column:longitude;type:decimal(11,8)"`
	Emoji          string  `gorm:"column:emoji;type:varchar(191)"`
	EmojiU         string  `gorm:"column:emojiU;type:varchar(191)"`
	Flag           bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID     string  `gorm:"column:wikiDataId;type:varchar(255)"`
}

// States [...]
type State struct {
	Model
	Name        string  `gorm:"column:name;type:varchar(255);not null"`
	CountryID   uint    `gorm:"index:country_region;column:country_id;type:mediumint(8) unsigned;not null"`
	CountryCode string  `gorm:"column:country_code;type:char(2);not null"`
	FipsCode    string  `gorm:"column:fips_code;type:varchar(255)"`
	Iso2        string  `gorm:"column:iso2;type:varchar(255)"`
	Type        string  `gorm:"column:type;type:varchar(191)"`
	Latitude    float64 `gorm:"column:latitude;type:decimal(10,8)"`
	Longitude   float64 `gorm:"column:longitude;type:decimal(11,8)"`
	Flag        bool    `gorm:"column:flag;type:tinyint(1);not null;default:1"`
	WikiDataID  string  `gorm:"column:wikiDataId;type:varchar(255)"`
}
