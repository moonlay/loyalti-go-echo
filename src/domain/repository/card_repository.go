package repository

import (
	"database/sql"
	"fmt"
	"github.com/beevik/guid"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

const (
	MemberTypeCard 	= "Member"
	GoldTier 		= "Gold"
	SilverTier 		= "Silver"
	PlatinumTier 	= "Platinum"
)

type MemberTier string

const (
	Silver MemberTier = "1"
	Gold MemberTier = "2"
	Platinum MemberTier = "3"
)

type CardRepository interface {
	CreateCard (newcard *model.Card) error
	DeleteCard (newcard *model.Card) error
	UpdateCard (newcard *model.Card) error
}

type card_repo struct {
	DB *gorm.DB
}

func (p *card_repo) CreateCard (newcard *model.Card) error {
	kartu := model.Card{
		Id:                guid.NewString(),
		Created:           time.Now(),
		CreatedBy:         "System",
		Modified:          time.Now(),
		ModifiedBy:        "System",
		Active:            true,
		IsDeleted:         false,
		Deleted:           nil,
		DeletedBy:         "",
		Title:             newcard.Title,
		Description:       newcard.Description,
		FontColor:         newcard.FontColor,
		TemplateColor:     newcard.TemplateColor,
		IconImage:         newcard.IconImage,
		TermsAndCondition: newcard.TermsAndCondition,
		Benefit:           newcard.Benefit,
		ValidUntil:        time.Time{},
		CurrentPoint:      newcard.CurrentPoint,
		IsValid:           true,
		ProgramId:         newcard.ProgramId,
		CardType:          newcard.CardType,
		IconImageStamp:    newcard.IconImageStamp,
		MerchantId:        newcard.MerchantId,
		Tier:              "",
	}
	db := database.ConnectionDB()
	err := db.Create(&kartu).Error
	if err == nil {
		fmt.Println("Error")
	}
	return err
}

func CreateCardRepository (db *gorm.DB) CardRepository {
	return &card_repo{
		DB:db,
	}	
}

func (p *card_repo) DeleteCard (newcard *model.Card) error {
	db := database.ConnectionDB()
	err := db.Model(&newcard).Where("id = ?", newcard.Id).Update("active", false).Error
	if err == nil {
		fmt.Println("tidak ada error")
	}
	return err
}

func (p *card_repo) UpdateCard (newcard *model.Card) error {
	db := database.ConnectionDB()
	err := db.Model(&newcard).Where("id = ?", newcard.Id).Update(&newcard).Error

	return err
}

//func CreateCardMerchant(card *model.Card) error {
//	db := database.ConnectionDB()
//
//	for i := 0; i <= 2; i++ {
//		cards := model.Card{
//			Id:                guid.NewString(),
//			Created:           time.Now(),
//			CreatedBy:         "",
//			Modified:          time.Now(),
//			ModifiedBy:        "",
//			Active:            true,
//			IsDeleted:         false,
//			Deleted:           nil,
//			DeletedBy:         "",
//			Title:             card.Title,
//			Description:       card.Description,
//			FontColor:         card.FontColor,
//			TemplateColor:     card.TemplateColor,
//			IconImage:         card.IconImage,
//			TermsAndCondition: card.TermsAndCondition,
//			Benefit:           card.Benefit,
//			ValidUntil:        time.Now(),
//			CurrentPoint:      card.CurrentPoint,
//			IsValid:           card.IsValid,
//			ProgramId:         card.ProgramId,
//			CardType:          card.CardType,
//			IconImageStamp:    card.IconImageStamp,
//			MerchantId:        card.MerchantId,
//		}
//		if(i == 0){
//			fmt.Println("masuk ke if == 0")
//			fmt.Println("isi enum silver", domain.EnumMember.Silver)
//			cards.Tier = domain.EnumMember.Silver
//			db.Create(&cards)
//		}else if(i == 1){
//			fmt.Println("masuk ke if == 1")
//			fmt.Println("isi enum silver", domain.EnumMember.Gold)
//			cards.Tier = domain.EnumMember.Gold
//			db.Create(&cards)
//		}else {
//			fmt.Println("masuk ke if == 2")
//			fmt.Println("isi enum silver", domain.EnumMember.Platinum)
//			cards.Tier = domain.EnumMember.Platinum
//			db.Create(&cards)
//		}
//	}
//	return card.Description
//}

func CreateCardMerchant(card *model.Card) string {
	db := database.ConnectionDB()

	for i := 0; i <= 2; i++ {
		cards := model.Card{
			Id:                guid.NewString(),
			Created:           time.Now(),
			CreatedBy:         "",
			Modified:          time.Now(),
			ModifiedBy:        "",
			Active:            true,
			IsDeleted:         false,
			Deleted:           nil,
			DeletedBy:         "",
			Title:             card.Title,
			Description:       card.Description,
			FontColor:         card.FontColor,
			TemplateColor:     card.TemplateColor,
			IconImage:         card.IconImage,
			TermsAndCondition: card.TermsAndCondition,
			Benefit:           card.Benefit,
			ValidUntil:        time.Now(),
			CurrentPoint:      card.CurrentPoint,
			IsValid:           card.IsValid,
			ProgramId:         card.ProgramId,
			CardType:          card.CardType,
			IconImageStamp:    card.IconImageStamp,
			MerchantId:        card.MerchantId,
		}
		if(i == 0){
			fmt.Println("masuk ke if == 0")
			fmt.Println("isi enum silver", domain.EnumMember.Silver)
			cards.Tier = domain.EnumMember.Silver
			db.Create(&cards)
		}else if(i == 1){
			fmt.Println("masuk ke if == 1")
			fmt.Println("isi enum silver", domain.EnumMember.Gold)
			cards.Tier = domain.EnumMember.Gold
			db.Create(&cards)
		}else {
			fmt.Println("masuk ke if == 2")
			fmt.Println("isi enum silver", domain.EnumMember.Platinum)
			cards.Tier = domain.EnumMember.Platinum
			db.Create(&cards)
		}
	}
	return card.Description
}

func GetCardMerchant(page *int, size *int, id *int, card_type *string) []model.Card {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var kartu []model.Card
	var rows *sql.Rows
	var err error
	var total int

	if id != nil {

		if page != nil && size != nil {

			rows, err = db.Where("program_id = ?", id).Find(&kartu).Order("title").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}

			if card_type != nil {
				if page != nil && size != nil && id != nil {
					fmt.Println("card type dapat ga")
					rows, err = db.Where("program_id = ? AND card_type = ?", id, card_type).Find(&kartu).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
					if err != nil {
						panic(err)
					}
				}
			}
		}  else {
				rows, err = db.Find(&kartu).Rows()
				if err != nil {
					panic(err)
				}
			}
		}


	result := make([]model.Card, 0)

	for rows.Next() {
		o := new(model.Card)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.DeletedBy,
			&o.Title,
			&o.Description,
			&o.FontColor,
			&o.TemplateColor,
			&o.IconImage,
			&o.TermsAndCondition,
			&o.Benefit,
			&o.ValidUntil,
			&o.CurrentPoint,
			&o.IsValid,
			&o.ProgramId,
			&o.CardType,
			&o.IconImageStamp,
			&o.MerchantId,
		)
		//add tier
		if o.CardType == MemberTypeCard{
			if o.CurrentPoint >= 0 && o.CurrentPoint <= 100{
				o.Tier = SilverTier
			}else if o.CurrentPoint > 100 && o.CurrentPoint <= 230{
				o.Tier = GoldTier
			}else if o.CurrentPoint > 230 && o.CurrentPoint <= 400{
				o.Tier = PlatinumTier
			}
		}else{
			o.Tier = "Ga ada tier"
		}

		result = append(result, *o)
	}

	db.Close()
	return result
}


func GetCardMember(program_id int )[]model.Card{
	db := database.ConnectionDB()
	card := &model.Card{}
	db.Model("card").Where("program_id = ?", program_id).Order("tier asc").First(&card)

	tier := make([]model.Card, 0)
	tier = append(tier, *card)

	fmt.Println("tier : ", tier)

	return tier
}