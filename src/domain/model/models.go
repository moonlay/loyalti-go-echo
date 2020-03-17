package model

import (
	"time"
)

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

type AccountMerchant struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}

type ResponseProcess struct {
	Sub        string `json:"sub"`
	Upn        string `json:"upn"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	UserId     string `json:"user_id"`
}

type Response struct {
	Access_Token  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
	Scope         string `json:"scope"`
	Id_Token      string `json:"id_token"`
	Token_Type    string `json:"token_type"`
	Expires_In    int    `json:"expires_in"`
}

type User struct {
	Username string
	Password string
}

type NewMerchantTest struct {
	Id            int    `json:"id"`
	MerchantEmail string `json:"merchant_email"`
}

type NewMerchantCommand struct {
	Id                    int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created               time.Time  `json:"created"`
	CreatedBy             string     `json:"created_by"`
	Modified              time.Time  `json:"modified"`
	ModifiedBy            string     `json:"modified_by"`
	Active                bool       `json:"active"`
	IsDeleted             bool       `json:"is_deleted"`
	Deleted               *time.Time `json:"deleted"`
	Deleted_by            string     `json:"deleted_by"`
	MerchantName          string     `json:"merchant_name"`
	MerchantEmail         string     `json:"merchant_email"`
	MerchantPhoneNumber   string     `json:"merchant_phone_number"`
	MerchantProvince      string     `json:"merchant_province"`
	MerchantCity          string     `json:"merchant_city"`
	MerchantAddress       string     `json:"merchant_address"`
	MerchantPostalCode    string     `json:"merchant_postal_code"`
	MerchantCategoryId    int        `json:"merchant_category_id"`
	MerchantWebsite       string     `json:"merchant_website"`
	MerchantMediaSocialId int        `json:"merchant_media_social_id"`
	MerchantDescription   string     `json:"merchant_description"`
	MerchantImageProfile  string     `json:"merchant_image_profile"`
	MerchantGallery       string     `json:"merchant_gallery"`
	MerchantPassword      string     `json:"merchant_password"`
}

type Merchant struct {
	//GeneralModels
	Id                    int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created               time.Time  `json:"created"`
	CreatedBy             string     `json:"created_by"`
	Modified              time.Time  `json:"modified"`
	ModifiedBy            string     `json:"modified_by"`
	Active                bool       `json:"active"`
	IsDeleted             bool       `json:"is_deleted"`
	Deleted               *time.Time `json:"deleted"`
	Deleted_by            string     `json:"deleted_by"`
	MerchantName          string     `json:"merchant_name"`
	MerchantEmail         string     `json:"merchant_email"`
	MerchantPhoneNumber   string     `json:"merchant_phone_number"`
	MerchantProvince      string     `json:"merchant_province"`
	MerchantCity          string     `json:"merchant_city"`
	MerchantAddress       string     `json:"merchant_address"`
	MerchantPostalCode    string     `json:"merchant_postal_code"`
	MerchantCategoryId    int        `json:"merchant_category_id"`
	MerchantWebsite       string     `json:"merchant_website"`
	MerchantMediaSocialId int        `json:"merchant_media_social_id"`
	MerchantDescription   string     `json:"merchant_description"`
	MerchantImageProfile  string     `json:"merchant_image_profile"`
	MerchantGallery       string     `json:"merchant_gallery"`
}

type MerchantCategory struct {
	Id           int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time  `json:"created"`
	CreatedBy    string     `json:"created_by"`
	Modified     time.Time  `json:"modified"`
	ModifiedBy   string     `json:"modified_by"`
	Active       bool       `json:"active"`
	IsDeleted    bool       `json:"is_deleted"`
	Deleted      *time.Time `json:"deleted"`
	Deleted_by   string     `json:"deleted_by"`
	CategoryName string     `json:"category_name"`
	ImageUrl     string     `json:"image_url"`
}

type MerchantSocialMedia struct {
	Id                  int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created             time.Time  `json:"created"`
	CreatedBy           string     `json:"created_by"`
	Modified            time.Time  `json:"modified"`
	ModifiedBy          string     `json:"modified_by"`
	Active              bool       `json:"active"`
	IsDeleted           bool       `json:"is_deleted"`
	Deleted             *time.Time `json:"deleted"`
	Deleted_by          string     `json:"deleted_by"`
	SocialMediaName     string     `json:"social_media_name"`
	SocialMediaImageUrl string     `json:"social_media_image_url"`
}

type CardType struct {
	Id           int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time  `json:"created"`
	CreatedBy    string     `json:"created_by"`
	Modified     time.Time  `json:"modified"`
	ModifiedBy   string     `json:"modified_by"`
	Active       bool       `json:"active"`
	IsDeleted    bool       `json:"is_deleted"`
	Deleted      *time.Time `json:"deleted"`
	Deleted_by   string     `json:"deleted_by"`
	CardTypeName string     `json:"card_type_name"`
}

type Outlet struct {
	Id               int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created          time.Time  `json:"created"`
	CreatedBy        string     `json:"created_by"`
	Modified         time.Time  `json:"modified"`
	ModifiedBy       string     `json:"modified_by"`
	Active           bool       `json:"active"`
	IsDeleted        bool       `json:"is_deleted"`
	Deleted          *time.Time `json:"deleted"`
	Deleted_by       string     `json:"deleted_by"`
	OutletName       string     `json:"outlet_name"`
	OutletAddress    string     `json:"outlet_address"`
	OutletPhone      string     `json:"outlet_phone"`
	OutletCity       string     `json:"outlet_city"`
	OutletProvince   string     `json:"outlet_province"`
	OutletPostalCode string     `json:"outlet_postal_code"`
	OutletLongitude  string     `json:"outlet_longitude"`
	OutletLatitude   string     `json:"outlet_latitude"`
	OutletDay        time.Time  `json:"outlet_day"`
	OutletHour       time.Time  `json:"outlet_hour"`
	MerchantId       int        `json:"merchant_id"`
	Timezone         string     `json:"timezone"`
}

type Program struct {
	Id                 int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created            time.Time  `json:"created"`
	CreatedBy          string     `json:"created_by"`
	Modified           time.Time  `json:"modified"`
	ModifiedBy         string     `json:"modified_by"`
	Active             bool       `json:"active"`
	IsDeleted          bool       `json:"is_deleted"`
	Deleted            *time.Time `json:"deleted"`
	Deleted_by         string     `json:"deleted_by"`
	ProgramName        string     `json:"program_name"`
	ProgramImage       string     `json:"program_image"`
	ProgramStartDate   time.Time  `json:"program_start_date"`
	ProgramEndDate     time.Time  `json:"program_end_date"`
	ProgramDescription string     `json:"program_description"`
	Card               string     `json:"card"`
	OutletID           string     `json:"outlet_id"`
	MerchantId         int        `json:"merchant_id"`
	//MerchantName       string     `json:"merchant_name"`
	CategoryId            int      `json:"category_id"`
	Benefit               *string  `json:"benefit"`
	TermsAndCondition     *string  `json:"terms_and_condition"`
	Tier                  *string  `json:"tier"`
	RedeemRules           *string  `json:"redeem_rules"`
	RewardTarget          *float64 `json:"reward_target"`
	QRCodeId              *string  `json:"qr_code_id"`
	ProgramPoint          *int     `json:"program_point"`
	MinPayment            *int     `json:"min_payment"`
	IsReqBillNumber       bool     `json:"is_req_bill_number"`
	IsReqTotalTransaction bool     `json:"is_req_total_transaction"`
	IsPushNotification    bool     `json:"is_push_notification"`
	IsLendCard            bool     `json:"is_lend_card"`
	IsGiveCard            bool     `json:"is_give_card"`
	IsWelcomeBonus        bool     `json:"is_welcome_bonus"`
}

type SpecialProgram struct {
	Id                 int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created            time.Time  `json:"created"`
	CreatedBy          string     `json:"created_by"`
	Modified           time.Time  `json:"modified"`
	ModifiedBy         string     `json:"modified_by"`
	Active             bool       `json:"active"`
	IsDeleted          bool       `json:"is_deleted"`
	Deleted            *time.Time `json:"deleted"`
	Deleted_by         string     `json:"deleted_by"`
	ProgramName        string     `json:"program_name"`
	ProgramImage       string     `json:"program_image"`
	ProgramStartDate   time.Time  `json:"program_start_date"`
	ProgramEndDate     time.Time  `json:"program_end_date"`
	ProgramDescription string     `json:"program_description"`
	Card               string     `json:"card"`
	OutletID           string     `json:"outlet_id"`
	MerchantId         int        `json:"merchant_id"`
	//MerchantName       string     `json:"merchant_name"`
	CategoryId        int      `json:"category_id"`
	Benefit           *string  `json:"benefit"`
	TermsAndCondition *string  `json:"terms_and_condition"`
	Tier              *string  `json:"tier"`
	RedeemRules       *string  `json:"redeem_rules"`
	RewardTarget      *float64 `json:"reward_target"`
	QRCodeId          *string  `json:"qr_code_id"`
}

type Product struct {
	Id                string     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created           time.Time  `json:"created"`
	CreatedBy         string     `json:"created_by"`
	Modified          time.Time  `json:"modified"`
	ModifiedBy        string     `json:"modified_by"`
	Active            bool       `json:"active"`
	IsDeleted         bool       `json:"is_deleted"`
	Deleted           *time.Time `json:"deleted"`
	Deleted_by        string     `json:"deleted_by"`
	ProductName       string     `json:"product_name"`
	ProductDesc       string     `json:"product_desc"`
	MerchantId        int        `gorm:"FOREIGNKEY"; json:"merchant_id"`
	ProductCategoryId int        `gorm:"FOREIGNKEY"; json:"product_category_id"`
}

type ProductCategory struct {
	Id                  string     `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created             time.Time  `json:"created"`
	CreatedBy           string     `json:"created_by"`
	Modified            time.Time  `json:"modified"`
	ModifiedBy          string     `json:"modified_by"`
	Active              bool       `json:"active"`
	IsDeleted           bool       `json:"is_deleted"`
	Deleted             *time.Time `json:"deleted"`
	Deleted_by          string     `json:"deleted_by"`
	ProductCategoryDesc string     `json:"product_category_desc"`
}

type MerchantStatus struct {
	Id                 string     `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created            time.Time  `json:"created"`
	CreatedBy          string     `json:"created_by"`
	Modified           time.Time  `json:"modified"`
	ModifiedBy         string     `json:"modified_by"`
	Active             bool       `json:"active"`
	IsDeleted          bool       `json:"is_deleted"`
	Deleted            *time.Time `json:"deleted"`
	Deleted_by         string     `json:"deleted_by"`
	MerchantStatusDesc string     `json:"merchant_status_desc"`
}

type Employee struct {
	Id            string     `gorm:"PRIMARY_KEY;NOT_NULL"; json:"id"`
	Created       time.Time  `json:"created"`
	CreatedBy     string     `json:"created_by"`
	Modified      time.Time  `json:"modified"`
	ModifiedBy    string     `json:"modified_by"`
	Active        bool       `json:"active"`
	IsDeleted     bool       `json:"is_deleted"`
	Deleted       *time.Time `json:"deleted"`
	Deleted_by    string     `json:"deleted_by"`
	EmployeeName  string     `json:"employee_name"`
	EmployeeEmail string     `json:"employee_email"`
	EmployeePin   string     `json:"employee_pin"`
	EmployeeRole  string     `json:"employee_role"`
	OutletId      string     `json:"outlet_id"`
	OutletName    string     `json:"outlet_name"`
}

type TotalPoint struct {
	Total int `json:"total_point"`
}

type TotalChop struct {
	Total int `json:"total_chop"`
}

type TransactionMerchant struct {
	Id               int        `gorm:"PRIMARY_KEY;NOT NUll"; json:"id"`
	Created          time.Time  `json:"created"`
	CreatedBy        string     `json:"created_by"`
	Modified         time.Time  `json:"modified"`
	ModifiedBy       string     `json:"modified_by"`
	Active           bool       `json:"active"`
	IsDeleted        bool       `json:"is_deleted"`
	Deleted          *time.Time `json:"deleted"`
	Deleted_by       string     `json:"deleted_by"`
	MerchantId       int        `json:"merchant_id"`
	OutletId         string     `json:"outlet_id"`
	TotalTransaction int        `json:"total_transaction"`
	PointTransaction int        `json:"point_transaction"`
	BillNumber       string     `json:"bill_number"`
}

type Card struct {
	Id                string     `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created           time.Time  `json:"created"`
	CreatedBy         string     `json:"created_by"`
	Modified          time.Time  `json:"modified"`
	ModifiedBy        string     `json:"modified_by"`
	Active            bool       `json:"active"`
	IsDeleted         bool       `json:"is_deleted"`
	Deleted           *time.Time `json:"deleted"`
	DeletedBy         string     `json:"deleted_by"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	FontColor         string     `json:"font_color"`
	TemplateColor     string     `json:"template_color"`
	IconImage         string     `json:"icon_image"`
	TermsAndCondition string     `json:"terms_and_condition"`
	Benefit           string     `json:"benefit"`
	ValidUntil        time.Time  `json:"valid_until"`
	CurrentPoint      int        `json:"current_point"`
	IsValid           bool       `json:"is_valid"`
	ProgramId         int        `json:"program_id"`
	CardType          string     `json:"card_type"`
	IconImageStamp    string     `json:"icon_image_stamp"`
	MerchantId        int        `json:"merchant_id"`
	Tier              string     `json:"tier"`
	TemplatePattern   string     `json:"template_pattern"`
}

type Member struct {
	Silver   string
	Gold     string
	Platinum string
}

type Voucher struct {
	Id                       string     `gorm:"PRIMARY_KEY;NOT NULL"; json:"id"`
	Created                  time.Time  `json:"created"`
	CreatedBy                string     `json:"created_by"`
	Modified                 time.Time  `json:"modified"`
	ModifiedBy               string     `json:"modified_by"`
	Active                   bool       `json:"active"`
	IsDeleted                bool       `json:"is_deleted"`
	Deleted                  *time.Time `json:"deleted"`
	DeletedBy                string     `json:"deleted_by"`
	VoucherName              string     `json:"voucher_name"`
	StartDate                time.Time  `json:"start_date"`
	EndDate                  time.Time  `json:"end_date"`
	VoucherDescription       string     `json:"voucher_description"`
	VoucherTermsAndCondition string     `json:"voucher_terms_and_condition"`
	IsPushNotification       bool       `json:"is_push_notification"`
	IsGiveVoucher            bool       `json:"is_give_voucher"`
	VoucherPeriod            string     `json:"voucher_period"`
	RewardTermsAndCondition  string     `json:"reward_terms_and_condition"`
	BackgroundVoucherPattern string     `json:"background_voucher"`
	BackgroundVoucherColour  string     `json:"background_voucher_colour"`
	MerchantId               string     `json:"merchant_id"`
	OutletId                 string     `json:"outlet_id"`
	ProgramId                int        `json:"program_id"`
}

type Province struct{
	IdProvince int `gorm:"PRIMARY_KEY;NOT NULL"; json:"id_province"`
	ProvinceName string `json:"province_name"`
}

type City struct {
	IdProvince int `json:"id_province"`
	IdCity int `gorm:"PRIMARY_KEY;NOT NULL"; json:"id_city"`
	CityName string `json:"city_name"`
}
