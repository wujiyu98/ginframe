package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Cart struct {
	ID  uint
	Qty uint
}

type Meta struct {
	Title       string
	Keywords    string
	Description string
}

type ArticleCategory struct {
	Model
	Name      string
	Pathname  string
	ParentID  uint
	SortOrder uint
	Image     string
	Summary   string
	Meta      Meta `gorm:"embedded;embeddedPrefix:meta_"`
}

type Address struct {
	Model
	UserID    uint
	CountryID uint
	StateID   uint
	CityID    uint
	Firstname string
	Lastname  string
	Company   string
	Postcode  string
	Address1  string
	Address2  string
}

type Article struct {
	Model
	LanguageID        uint
	ArticleCategoryID uint
	Title             string
	Pathname          string
	SortOrder         uint
	Showed            *byte `gorm:"default:1"`
	Summary           string
	Image             string
	Author            string
	Content           string
	Meta              Meta `gorm:"embedded;embeddedPrefix:meta_"`
}

type Banner struct {
	Model
	Name    string
	Image   string
	Url     string
	Title   string
	Summary string
}

type Category struct {
	Model
	Name      string
	Pathname  string
	ParentID  uint
	SortOrder uint
	Qty       uint
	Image     string
	Summary   string
	Meta      Meta `gorm:"embedded;embeddedPrefix:meta_"`
}

type Manufacturer struct {
	Model
	Name      string
	Pathname  string
	SortOrder uint
	Qty       uint
	Image     string
	Summary   string
	Meta      Meta `gorm:"embedded;embeddedPrefix:meta_"`
}

type CategoryManufacturer struct {
	CategoryID     uint
	ManufacturerID uint
}

type Language struct {
	Model
	Name   string
	Domain string
	Code   string
	Image  string
}

type Message struct {
	Model
	Name        string
	Email       string
	Country     string
	MobilePhone string
	Company     string
	Comment     string
}

type Enquiry struct {
	Model
	Name        string
	Email       string
	Country     string
	MobilePhone string
	Company     string
	Comment     string
	Products    []EnquiryProduct `gorm:"type:json;serializer:json"`
}

type EnquiryProduct struct {
	Title        string
	Manufacturer string
	Summary      string
	Qty          uint
	Price        float64
}

type Order struct {
	Model
	UserID            uint
	OrderNumber       string
	Address           string
	TransactionNumber string
	PaypalFee         float64
	Total             float64
	Payment           byte
	Status            byte
	Freight           float64
	Express           string
	TrackingNumber    string
	CreatedAt         time.Time
}

type OrderProduct struct {
	Model
	OrderID      uint
	ProductID    uint
	URL          string
	Title        string
	Image        string
	Manufacturer string
	Summary      string
	Qty          uint
	Price        float64
}

type Product struct {
	Model
	LanguageID        uint
	CategoryID        uint
	ManufacturerID    uint
	Title             string
	Pathname          string
	SortOrder         uint
	Showed            *byte `gorm:"default:1"`
	Summary           string
	Image             string
	Images            []string `gorm:"type:json;serializer:json"`
	Content           string
	Stock             uint
	Hot               byte
	New               byte
	Special           byte
	Price             float64
	Prices            []map[string]string `gorm:"type:json;serializer:json"`
	Meta              Meta                `gorm:"embedded;embeddedPrefix:meta_"`
	ProductAttributes []ProductAttribute
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	if len(p.ProductAttributes) > 0 {
		var items []ProductAttribute
		for _, v := range p.ProductAttributes {
			v.ProductID = p.ID
			items = append(items, v)
		}

		tx.Save(&items)

	}
	return
}

type Attribute struct {
	Model
	CategoryID uint
	Name       string
}

type ProductAttribute struct {
	ProductID   uint
	AttributeID uint
	Text        string
}

type User struct {
	Model
	Email           string
	Name            string
	Country         string
	Company         string
	Password        string
	RememberToken   string
	MobilePhone     string
	EmailVerifiedAt time.Time
}

type Seo struct {
	Model
	LanguageID uint
	Name       string
	Pathname   string
	Meta       Meta `gorm:"embedded;embeddedPrefix:meta_"`
}

type SiteInfo struct {
	Model
	LanguageID  uint
	SiteName    string
	Company     string
	Contact     string
	Phone       string
	Phone2      string
	MobilePhone string
	Skype       string
	QQ          string
	Whatsapp    string
	Address     string
}

type City struct {
	Model
	Name        string
	StateID     string
	StateCode   string
	CountryID   string
	CountryCode string
	Latitude    float64
	Longitude   float64
	Flag        bool
	WikiDataID  string
}

// Countries [...]
type Country struct {
	Model
	Name           string
	Iso3           string
	NumericCode    string
	Iso2           string
	Phonecode      string
	Capital        string
	Currency       string
	CurrencyName   string
	CurrencySymbol string
	Tld            string
	Native         string
	Region         string
	Subregion      string
	Timezones      string
	Translations   string
	Latitude       float64
	Longitude      float64
	Emoji          string
	EmojiU         string
	Flag           bool
	WikiDataID     string
}

// States [...]
type State struct {
	Model
	Name        string
	CountryID   uint
	CountryCode string
	FipsCode    string
	Iso2        string
	Type        string
	Latitude    float64
	Longitude   float64
	Flag        bool
	WikiDataID  string
}
