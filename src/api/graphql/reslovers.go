package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
)

type Song model.Song

func MerchantResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	email, mail := p.Args["email"].(string)
	sort, tap := p.Args["sort"].(int)
	merchantid, id := p.Args["id"].(int)
	if ok && sip && tap{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		merchant := repository.GetMerchant(pages, sizes, sorts, nil, nil)
		fmt.Println(merchant)
		return merchant, nil

	} else if ok && sip && mail {
		var pages *int = &page
		var sizes *int = &size
		var emails *string = &email
		merchant := repository.GetMerchant(pages,sizes,nil,emails,nil)
		return merchant,nil

	} else if ok && sip && id{
		var pages *int = &page
		var sizes *int = &size
		var merchantid *int = &merchantid
		merchant := repository.GetMerchant(pages,sizes,nil,nil,merchantid)
		return merchant, nil

	} else if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		merchant := repository.GetMerchant(pages,sizes,nil,nil,nil)
		return merchant,nil
	}


	merchant := repository.GetMerchant(nil, nil, nil, nil, nil)
	fmt.Println(merchant)
	return merchant, nil
}

func MerchantCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,top := p.Args["sort"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		category := repository.GetCategory(pages, sizes, sorts)
		fmt.Println(category)
		return category,nil
	}

	category := repository.GetCategory(nil,nil, nil)

	return category, nil
}

func MerchantCardTypeResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)

	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
	card := repository.GetCardType(pages, sizes, sorts)
	fmt.Println(card)
	return card, nil
	}
	card := repository.GetCardType(nil,nil, nil)
	return card, nil
}

func SocialMediaResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,top := p.Args["sort"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		sosmed := repository.GetSocialMedia(pages, sizes, sorts)
		fmt.Println(sosmed)
		return sosmed,nil
	}
	sosmed := repository.GetSocialMedia(nil, nil, nil)
	return sosmed, nil
}

func ProvinceResolver (p graphql.ResolveParams) (interface{}, error){
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		province := repository.GetProvince(pages, sizes, sorts)
		fmt.Println(province)
		return province, nil
	}
	province := repository.GetProvince(nil, nil, nil)
	return province, nil
}

func CityResolver (p graphql.ResolveParams) (interface{}, error){
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		province := repository.GetCity(pages, sizes, sorts)
		fmt.Println(province)
		return province, nil
	}
	province := repository.GetCity(nil, nil, nil)
	return province, nil
}

func SpecialProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,deh := p.Args["sort"].(int)
	category,cat := p.Args["category"].(int)
	detail,id := p.Args["id"].(int)
	if ok && sip && deh && cat {
		fmt.Println("12345")
		var pagination *int = &page
		var sizing *int = &size
		var category *int = &category
		var sorting *int = &sort
		special := repository.GetSpecialProgram(pagination,sizing,sorting,category, nil)
		//fmt.Println(program)
		return special,nil
	}else if ok && sip && deh {
		fmt.Println("5678")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		//var cats *int = &category
		special := repository.GetSpecialProgram(pages, sizes, sorts,nil, nil)
		//fmt.Println(program)
		return special,nil
	} else if ok && sip {
		fmt.Println("7890")
		var paging *int = &page
		var sizing *int = &size
		special := repository.GetSpecialProgram(paging, sizing, nil, nil, nil)
		//fmt.Println(program)
		return special,nil
	} else if id {
		var detail *int = &detail
		special := repository.GetSpecialProgram(nil, nil,nil,nil, detail)
		return special, nil
	}

	special := repository.GetSpecialProgram(nil,nil, nil, nil, nil)

	return special, nil
}

//program function
func ProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,deh := p.Args["sort"].(int)
	category, cat := p.Args["category"].(int)
	detail, id := p.Args["id"].(int)
	if ok && sip && deh && cat {
		var pagination *int = &page
		var sizing *int = &size
		var category *int = &category
		var sorting *int = &sort
		program := repository.GetProgram(pagination,sizing,sorting,category,nil)
		//fmt.Println(program)
		return program,nil
	}else if ok && sip && deh {
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		//var cats *int = &category
		program := repository.GetProgram(pages, sizes, sorts,nil,nil)
		//fmt.Println(program)
		return program,nil
	} else if ok && sip {
		var paging *int = &page
		var sizing *int = &size
		program := repository.GetProgram(paging, sizing, nil, nil,nil)
		//fmt.Println(program)
		return program,nil
	} else if id {
		var detail *int = &detail
		program := repository.GetProgram(nil,nil,nil,nil, detail)
		return program,nil
	}

	program := repository.GetProgram(nil,nil, nil, nil,nil)

	return program, nil
}

func OutletResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	id, top := p.Args["id"].(int)
	email, mail := p.Args["email"].(string)
	if ok && sip && top && mail{
		var pages *int = &page
		var sizes *int = &size
		var merchant_id *int = &id
		var emails *string = &email
		outlet := repository.GetOutlet(pages, sizes, merchant_id, emails)
		fmt.Println(outlet)
		return outlet,nil
	} else if ok && sip && mail {
		var paging *int = &page
		var sizing *int = &size
		var emails *string = &email
		outlet := repository.GetOutlet(paging, sizing, nil, emails)
		return outlet, nil
	} else if ok && sip && top{
		var paging *int = &page
		var sizing *int = &size
		var merchant *int = &id
		outlet := repository.GetOutlet(paging, sizing, merchant, nil)
		return outlet,nil
	} else if ok && sip{
		var paging *int = &page
		var sizing *int = &size
		outlet := repository.GetOutlet(paging, sizing, nil, nil)
		return outlet,nil
	}
	outlet := repository.GetOutlet(nil,nil, nil, nil)
	return outlet, nil
}

func EmployeeResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort, top := p.Args["id"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		employee := repository.GetEmployee(pages, sizes, sorts)
		fmt.Println(employee)
		return employee,nil
	} else if ok && sip{
		var paging *int = &page
		var sizing *int = &size
		var sorting *int = &sort
		employee := repository.GetEmployee(paging, sizing, sorting)
		return employee,nil
	}
	employee := repository.GetEmployee(nil,nil, nil)
	return employee, nil
}

func TotalPointResolver(p graphql.ResolveParams)(interface{}, error){
	id := p.Args["id"].(int)
	pay := p.Args["pay"].(int)
	pin := p.Args["pin"].(string)
	outletid := p.Args["outletid"].(string)
	cardtype := p.Args["cardtype"].(string)
	var ids int = id
	var pays int = pay
	var pins string = pin
	var outletids string = outletid
	var cardtypes string = cardtype
	total := repository.TotalPoint(ids, pays, pins, outletids, cardtypes)
	return total, nil
}

func TotalChopResolver(p graphql.ResolveParams)(interface{}, error){
	id := p.Args["id"].(int)
	pay := p.Args["pay"].(int)
	pin := p.Args["pin"].(string)
	outletid := p.Args["outletid"].(string)
	cardtype := p.Args["cardtype"].(string)
	var ids int = id
	var pays int = pay
	var pins string = pin
	var outletids string = outletid
	var cardtypes string = cardtype
	total := repository.TotalChop(ids, pays, pins, outletids, cardtypes)
	return total, nil
}

//func MerchantCardMemberResolver(p graphql.ResolveParams)(interface{}, error){
//	fmt.Println("masuk ke resolver member")
//	program_id := p.Args["program_id"].(int)
//	var program_ids int = program_id
//	card := repository.GetCardMember(program_ids)
//	fmt.Println(card)
//	return card, nil
//}

func TransactionResolver (p graphql.ResolveParams)(interface{}, error){
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	outletid, tap := p.Args["outletid"].(string)
	if ok && sip && top && tap{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var outletids *string = &outletid
		transaction := repository.GetTransaction(pages, sizes, sorts, outletids)
		fmt.Println(transaction)
		return transaction, nil
	}

	transaction := repository.GetTransaction(nil, nil, nil, nil)

	return transaction, nil
}

func SongResolver(p graphql.ResolveParams) (interface{}, error) {
	users := []Song{
		Song{
			ID:       "1",
			Album:    "ts-fearless",
			Title:    "Fearless",
			Duration: "4:01",
			Type:     "song",
		},
		Song{
			ID:       "2",
			Album:    "ts-fearless",
			Title:    "Fifteen",
			Duration: "4:54",
			Type:     "song",
		},
	}
	return users, nil
	//return nil, nil
}

func CardResolver(p graphql.ResolveParams) (interface{}, error) {
	fmt.Println("masuk ke resolver card")
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	id, top := p.Args["id"].(int)
	card_type, tipe := p.Args["card_type"].(string)

	var outlet []model.Card

	if (card_type == "Member"){
		fmt.Println("masuk ke if member")
		outlet =  repository.GetCardMember(id)
	}

	fmt.Println("keluar dari if member")

	if ok && sip && top && tipe{
		fmt.Println("masuk ke if pertama")
		var pages *int = &page
		var sizes *int = &size
		var top *int = &id
		var types *string = &card_type
		outlet := repository.GetCardMerchant(pages, sizes, top, types)
		return outlet, nil
	} else if ok && sip && top{
		fmt.Println("masuk ke if kedua")
		var pages *int = &page
		var sizes *int = &size
		var card_id *int = &id
		outlet := repository.GetCardMerchant(pages, sizes, card_id,nil)

		fmt.Println(outlet)
		return outlet,nil
	} else if ok && sip{
		fmt.Println("masuk ke if ketiga")
		var paging *int = &page
		var sizing *int = &size
		outlet := repository.GetCardMerchant(paging, sizing, nil, nil)
		return outlet,nil
	}
	fmt.Println("lewat semua")
	//outlet := repository.GetCardMerchant(nil,nil, nil, nil)

	return outlet, nil
}

func VoucherResolver (p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)
	merchant_id, tap := p.Args["merchant_id"].(int)
	if ok && sip && top && tap{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		var merchant_ids *int = &merchant_id
		voucher := repository.GetVoucher(pages, sizes, sorts, merchant_ids)
		fmt.Println(voucher)
		return voucher, nil
	}

	voucher := repository.GetVoucher(nil, nil, nil, nil)

	return voucher, nil
}

//func MerchantResolve() {
//
//}
