package graphQL

import (
	"github.com/graphql-go/graphql"
)

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"album": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var accountType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Account",
	Fields: graphql.Fields{
		"merchant_email": &graphql.Field{
			Type: graphql.Int,
		},
		"merchant_password": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var merchantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Merchant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_phone_number": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_email": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_address": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_city": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_province": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_website": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_category_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_media_social_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_description": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_image_profile": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_gallery": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var merchantCommandType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MerchantCommand",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active" : &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_deleted" : &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_phone_number": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_email": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_address": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_city": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_province": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_website": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_category_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_media_social_id": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_description": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_image_profile": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_gallery": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_password": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var socialmediaType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Card",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"social_media_name": &graphql.Field{
			Type: graphql.String,
		},
		"image_url": &graphql.Field{
			Type: graphql.String,
		},
		"tier": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var cardType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Card",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"card_type_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"category_name": &graphql.Field{
			Type: graphql.String,
		},
		"image_url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var programType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Program",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"program_name": &graphql.Field{
			Type: graphql.String,
		},
		"program_image": &graphql.Field{
			Type: graphql.String,
		},
		"program_start_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"program_end_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"program_description": &graphql.Field{
			Type: graphql.String,
		},
		"card": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_id": &graphql.Field{
			Type: graphql.Int,
		},
		"merchant_id": &graphql.Field{
			Type: graphql.Int,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"category_id": &graphql.Field{
			Type: graphql.Int,
		},
		"benefit": &graphql.Field{
			Type: graphql.String,
		},
		"terms_and_condition": &graphql.Field{
			Type: graphql.String,
		},
		"tier": &graphql.Field{
			Type: graphql.String,
		},
		"redeem_rules": &graphql.Field{
			Type: graphql.String,
		},
		"reward_target": &graphql.Field{
			Type: graphql.Float,
		},
		"qr_code_id": &graphql.Field{
			Type: graphql.String,
		},
		"program_point": &graphql.Field{
			Type: graphql.Int,
		},
		"min_payment": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var specialprogramType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpecialProgram",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"program_name": &graphql.Field{
			Type: graphql.String,
		},
		"program_image": &graphql.Field{
			Type: graphql.String,
		},
		"program_start_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"program_end_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"program_description": &graphql.Field{
			Type: graphql.String,
		},
		"card": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_id": &graphql.Field{
			Type: graphql.Int,
		},
		"merchant_id": &graphql.Field{
			Type: graphql.Int,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"category_id": &graphql.Field{
			Type: graphql.Int,
		},
		"program_title": &graphql.Field{
			Type: graphql.String,
		},
		"benefit": &graphql.Field{
			Type: graphql.String,
		},
		"terms_and_condition": &graphql.Field{
			Type: graphql.String,
		},
		"tier": &graphql.Field{
			Type: graphql.String,
		},
		"redeem_rules": &graphql.Field{
			Type: graphql.String,
		},
		"reward_target": &graphql.Field{
			Type: graphql.Float,
		},
		"qr_code_id": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var outletType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Outlet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_name": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_address": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_phone": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_city": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_province": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_longitude": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_latitude": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_id": &graphql.Field{
			Type: graphql.String,
		},
		"outlet_day": &graphql.Field{
			Type: graphql.DateTime,
		},
		"outlet_hour": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})

var employeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"employee_name": &graphql.Field{
			Type: graphql.String,
		},
		"employee_email": &graphql.Field{
			Type: graphql.String,
		},
		"employee_pin": &graphql.Field{
			Type: graphql.String,
		},
		"employee_role": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_id": &graphql.Field{
			Type: graphql.Int,
		},
		"outlet_name": &graphql.Field{
			Type:graphql.String,
		},
	},
})

var totalpointType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TotalPoint",
	Fields: graphql.Fields{
		"total_point": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var totalchopType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TotalChop",
	Fields: graphql.Fields{
		"total_chop": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var cardtier = graphql.NewObject(graphql.ObjectConfig{
	Name: "CardTier",
	Fields: graphql.Fields{
		"cardtier": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var transactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TotalPoint",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_id": &graphql.Field{
			Type: graphql.Int,
		},
		"outlet_id": &graphql.Field{
			Type: graphql.String,
		},
		"total_transaction": &graphql.Field{
			Type: graphql.Int,
		},
		"point_transaction": &graphql.Field{
			Type: graphql.Int,
		},
		"bill_number": &graphql.Field{
			Type : graphql.String,
		},
	},
})

var card = graphql.NewObject(graphql.ObjectConfig{
	Name: "Card",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"font_color": &graphql.Field{
			Type: graphql.String,
		},
		"template_color": &graphql.Field{
			Type: graphql.String,
		},
		"icon_image": &graphql.Field{
			Type: graphql.String,
		},
		"terms_and_condition": &graphql.Field{
			Type: graphql.String,
		},
		"benefit": &graphql.Field{
			Type: graphql.String,
		},
		"valid_until": &graphql.Field{
			Type: graphql.DateTime,
		},
		"current_point": &graphql.Field{
			Type: graphql.Int,
		},
		"is_valid": &graphql.Field{
			Type: graphql.Boolean,
		},
		"program_id": &graphql.Field{
			Type: graphql.Int,
		},
		"card_type": &graphql.Field{
			Type: graphql.String,
		},
		"icon_image_stamp": &graphql.Field{
			Type: graphql.String,
		},
		"merchant_id": &graphql.Field{
			Type: graphql.Int,
		},
		"tier": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var voucher = graphql.NewObject(graphql.ObjectConfig{
	Name: "Voucher",
	Fields:graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"created": &graphql.Field{
			Type: graphql.DateTime,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"deleted": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"voucher_name": &graphql.Field{
			Type:graphql.String,
		},
		"start_date": &graphql.Field{
			Type:graphql.DateTime,
		},
		"end_date": &graphql.Field{
			Type:graphql.DateTime,
		},
		"voucher_description": &graphql.Field{
			Type:graphql.String,
		},
		"voucher_terms_and_condition": &graphql.Field{
			Type:graphql.String,
		},
		"is_push_notification": &graphql.Field{
			Type:graphql.Boolean,
		},
		"is_give_voucher" :&graphql.Field{
			Type:graphql.Boolean,
		},
		"voucher_period" :&graphql.Field{
			Type:graphql.String,
		},
		"reward_terms_and_condition" :&graphql.Field{
			Type:graphql.String,
		},
		"background_voucher" : &graphql.Field{
			Type:graphql.String,
		},
		"background_voucher_colour": &graphql.Field{
			Type:graphql.String,
		},
		"merchant_id": &graphql.Field{
			Type:graphql.String,
		},
		"outlet_id": &graphql.Field{
			Type:graphql.String,
		},
		"program_id": &graphql.Field{
			Type:graphql.Int,
		},
	},
})

