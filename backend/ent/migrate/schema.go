// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AmountpaidsColumns holds the columns for the "amountpaids" table.
	AmountpaidsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amountpaid_money", Type: field.TypeInt, Unique: true},
	}
	// AmountpaidsTable holds the schema information for the "amountpaids" table.
	AmountpaidsTable = &schema.Table{
		Name:        "amountpaids",
		Columns:     AmountpaidsColumns,
		PrimaryKey:  []*schema.Column{AmountpaidsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BanksColumns holds the columns for the "banks" table.
	BanksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bank_type", Type: field.TypeString},
	}
	// BanksTable holds the schema information for the "banks" table.
	BanksTable = &schema.Table{
		Name:        "banks",
		Columns:     BanksColumns,
		PrimaryKey:  []*schema.Column{BanksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "category_name", Type: field.TypeString, Unique: true},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:        "categories",
		Columns:     CategoriesColumns,
		PrimaryKey:  []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender_name", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:        "genders",
		Columns:     GendersColumns,
		PrimaryKey:  []*schema.Column{GendersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupOfAgesColumns holds the columns for the "group_of_ages" table.
	GroupOfAgesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "group_of_age_name", Type: field.TypeString},
		{Name: "group_of_age_age", Type: field.TypeString},
	}
	// GroupOfAgesTable holds the schema information for the "group_of_ages" table.
	GroupOfAgesTable = &schema.Table{
		Name:        "group_of_ages",
		Columns:     GroupOfAgesColumns,
		PrimaryKey:  []*schema.Column{GroupOfAgesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// HospitalsColumns holds the columns for the "hospitals" table.
	HospitalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hospital_name", Type: field.TypeString, Unique: true},
	}
	// HospitalsTable holds the schema information for the "hospitals" table.
	HospitalsTable = &schema.Table{
		Name:        "hospitals",
		Columns:     HospitalsColumns,
		PrimaryKey:  []*schema.Column{HospitalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// InquiriesColumns holds the columns for the "inquiries" table.
	InquiriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "inquiry_messages", Type: field.TypeString},
		{Name: "inquiry_phone_messages", Type: field.TypeString},
		{Name: "inquiry_time_messages", Type: field.TypeTime},
		{Name: "category_id", Type: field.TypeInt, Nullable: true},
		{Name: "member_id", Type: field.TypeInt, Nullable: true},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// InquiriesTable holds the schema information for the "inquiries" table.
	InquiriesTable = &schema.Table{
		Name:       "inquiries",
		Columns:    InquiriesColumns,
		PrimaryKey: []*schema.Column{InquiriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "inquiries_categories_category_inquiry",
				Columns: []*schema.Column{InquiriesColumns[4]},

				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "inquiries_members_member_inquiry",
				Columns: []*schema.Column{InquiriesColumns[5]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "inquiries_officers_officer_inquiry",
				Columns: []*schema.Column{InquiriesColumns[6]},

				RefColumns: []*schema.Column{OfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "inquiries_products_product_inquiry",
				Columns: []*schema.Column{InquiriesColumns[7]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// InsurancesColumns holds the columns for the "insurances" table.
	InsurancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "insurance_identification", Type: field.TypeString, Size: 13},
		{Name: "insurance_insurer", Type: field.TypeString},
		{Name: "insurance_address", Type: field.TypeString},
		{Name: "insurance_time_buy", Type: field.TypeTime},
		{Name: "insurance_time_firstpay", Type: field.TypeTime},
		{Name: "hospital_id", Type: field.TypeInt, Nullable: true},
		{Name: "member_id", Type: field.TypeInt, Nullable: true},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// InsurancesTable holds the schema information for the "insurances" table.
	InsurancesTable = &schema.Table{
		Name:       "insurances",
		Columns:    InsurancesColumns,
		PrimaryKey: []*schema.Column{InsurancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "insurances_hospitals_hospital_insurance",
				Columns: []*schema.Column{InsurancesColumns[6]},

				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "insurances_members_member_insurance",
				Columns: []*schema.Column{InsurancesColumns[7]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "insurances_officers_officer_insurance",
				Columns: []*schema.Column{InsurancesColumns[8]},

				RefColumns: []*schema.Column{OfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "insurances_products_product_insurance",
				Columns: []*schema.Column{InsurancesColumns[9]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MembersColumns holds the columns for the "members" table.
	MembersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "member_email", Type: field.TypeString, Unique: true},
		{Name: "member_name", Type: field.TypeString, Unique: true},
		{Name: "member_password", Type: field.TypeString},
		{Name: "position_id", Type: field.TypeInt, Nullable: true},
	}
	// MembersTable holds the schema information for the "members" table.
	MembersTable = &schema.Table{
		Name:       "members",
		Columns:    MembersColumns,
		PrimaryKey: []*schema.Column{MembersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "members_positions_members",
				Columns: []*schema.Column{MembersColumns[4]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MoneytransfersColumns holds the columns for the "moneytransfers" table.
	MoneytransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "moneytransfer_type", Type: field.TypeString},
	}
	// MoneytransfersTable holds the schema information for the "moneytransfers" table.
	MoneytransfersTable = &schema.Table{
		Name:        "moneytransfers",
		Columns:     MoneytransfersColumns,
		PrimaryKey:  []*schema.Column{MoneytransfersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OfficersColumns holds the columns for the "officers" table.
	OfficersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "officer_email", Type: field.TypeString, Unique: true},
		{Name: "officer_name", Type: field.TypeString, Unique: true},
		{Name: "officer_password", Type: field.TypeString},
		{Name: "position_id", Type: field.TypeInt, Nullable: true},
	}
	// OfficersTable holds the schema information for the "officers" table.
	OfficersTable = &schema.Table{
		Name:       "officers",
		Columns:    OfficersColumns,
		PrimaryKey: []*schema.Column{OfficersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "officers_positions_officers",
				Columns: []*schema.Column{OfficersColumns[4]},

				RefColumns: []*schema.Column{PositionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaybacksColumns holds the columns for the "paybacks" table.
	PaybacksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "payback_accountnumber", Type: field.TypeString},
		{Name: "payback_transfertime", Type: field.TypeTime},
		{Name: "bank_id", Type: field.TypeInt, Nullable: true},
		{Name: "member_id", Type: field.TypeInt, Nullable: true},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// PaybacksTable holds the schema information for the "paybacks" table.
	PaybacksTable = &schema.Table{
		Name:       "paybacks",
		Columns:    PaybacksColumns,
		PrimaryKey: []*schema.Column{PaybacksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "paybacks_banks_bank_payback",
				Columns: []*schema.Column{PaybacksColumns[3]},

				RefColumns: []*schema.Column{BanksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "paybacks_members_member_payback",
				Columns: []*schema.Column{PaybacksColumns[4]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "paybacks_officers_officer_payback",
				Columns: []*schema.Column{PaybacksColumns[5]},

				RefColumns: []*schema.Column{OfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "paybacks_products_product_payback",
				Columns: []*schema.Column{PaybacksColumns[6]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account_name", Type: field.TypeString},
		{Name: "account_number", Type: field.TypeString, Size: 10},
		{Name: "phone_number", Type: field.TypeString, Size: 10},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "transfer_time", Type: field.TypeTime},
		{Name: "bank_id", Type: field.TypeInt, Nullable: true},
		{Name: "insurance_id", Type: field.TypeInt, Nullable: true},
		{Name: "member_id", Type: field.TypeInt, Nullable: true},
		{Name: "moneytransfer_id", Type: field.TypeInt, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "payments_banks_bank_payment",
				Columns: []*schema.Column{PaymentsColumns[6]},

				RefColumns: []*schema.Column{BanksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_insurances_insurance_payment",
				Columns: []*schema.Column{PaymentsColumns[7]},

				RefColumns: []*schema.Column{InsurancesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_members_member_payment",
				Columns: []*schema.Column{PaymentsColumns[8]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "payments_moneytransfers_moneytransfer_payment",
				Columns: []*schema.Column{PaymentsColumns[9]},

				RefColumns: []*schema.Column{MoneytransfersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PositionsColumns holds the columns for the "positions" table.
	PositionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "position_name", Type: field.TypeString, Unique: true},
	}
	// PositionsTable holds the schema information for the "positions" table.
	PositionsTable = &schema.Table{
		Name:        "positions",
		Columns:     PositionsColumns,
		PrimaryKey:  []*schema.Column{PositionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "product_name", Type: field.TypeString},
		{Name: "product_price", Type: field.TypeInt},
		{Name: "product_time", Type: field.TypeInt},
		{Name: "product_payment_of_year", Type: field.TypeInt},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "group_of_age_id", Type: field.TypeInt, Nullable: true},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "products_genders_gender_product",
				Columns: []*schema.Column{ProductsColumns[5]},

				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "products_group_of_ages_groupofage_product",
				Columns: []*schema.Column{ProductsColumns[6]},

				RefColumns: []*schema.Column{GroupOfAgesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "products_officers_officer_product",
				Columns: []*schema.Column{ProductsColumns[7]},

				RefColumns: []*schema.Column{OfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RecordinsurancesColumns holds the columns for the "recordinsurances" table.
	RecordinsurancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number_of_days_of_treat", Type: field.TypeInt},
		{Name: "recordinsurance_contact", Type: field.TypeString, Size: 10},
		{Name: "recordinsurance_address", Type: field.TypeString},
		{Name: "recordinsurance_time", Type: field.TypeTime},
		{Name: "amountpaid_id", Type: field.TypeInt, Nullable: true},
		{Name: "hospital_id", Type: field.TypeInt, Nullable: true},
		{Name: "member_id", Type: field.TypeInt, Nullable: true},
		{Name: "officer_id", Type: field.TypeInt, Nullable: true},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// RecordinsurancesTable holds the schema information for the "recordinsurances" table.
	RecordinsurancesTable = &schema.Table{
		Name:       "recordinsurances",
		Columns:    RecordinsurancesColumns,
		PrimaryKey: []*schema.Column{RecordinsurancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "recordinsurances_amountpaids_amountpaid_recordinsurance",
				Columns: []*schema.Column{RecordinsurancesColumns[5]},

				RefColumns: []*schema.Column{AmountpaidsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "recordinsurances_hospitals_hospital_recordinsurance",
				Columns: []*schema.Column{RecordinsurancesColumns[6]},

				RefColumns: []*schema.Column{HospitalsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "recordinsurances_members_member_recordinsurance",
				Columns: []*schema.Column{RecordinsurancesColumns[7]},

				RefColumns: []*schema.Column{MembersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "recordinsurances_officers_officer_recordinsurance",
				Columns: []*schema.Column{RecordinsurancesColumns[8]},

				RefColumns: []*schema.Column{OfficersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "recordinsurances_products_product_recordinsurance",
				Columns: []*schema.Column{RecordinsurancesColumns[9]},

				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AmountpaidsTable,
		BanksTable,
		CategoriesTable,
		GendersTable,
		GroupOfAgesTable,
		HospitalsTable,
		InquiriesTable,
		InsurancesTable,
		MembersTable,
		MoneytransfersTable,
		OfficersTable,
		PaybacksTable,
		PaymentsTable,
		PositionsTable,
		ProductsTable,
		RecordinsurancesTable,
	}
)

func init() {
	InquiriesTable.ForeignKeys[0].RefTable = CategoriesTable
	InquiriesTable.ForeignKeys[1].RefTable = MembersTable
	InquiriesTable.ForeignKeys[2].RefTable = OfficersTable
	InquiriesTable.ForeignKeys[3].RefTable = ProductsTable
	InsurancesTable.ForeignKeys[0].RefTable = HospitalsTable
	InsurancesTable.ForeignKeys[1].RefTable = MembersTable
	InsurancesTable.ForeignKeys[2].RefTable = OfficersTable
	InsurancesTable.ForeignKeys[3].RefTable = ProductsTable
	MembersTable.ForeignKeys[0].RefTable = PositionsTable
	OfficersTable.ForeignKeys[0].RefTable = PositionsTable
	PaybacksTable.ForeignKeys[0].RefTable = BanksTable
	PaybacksTable.ForeignKeys[1].RefTable = MembersTable
	PaybacksTable.ForeignKeys[2].RefTable = OfficersTable
	PaybacksTable.ForeignKeys[3].RefTable = ProductsTable
	PaymentsTable.ForeignKeys[0].RefTable = BanksTable
	PaymentsTable.ForeignKeys[1].RefTable = InsurancesTable
	PaymentsTable.ForeignKeys[2].RefTable = MembersTable
	PaymentsTable.ForeignKeys[3].RefTable = MoneytransfersTable
	ProductsTable.ForeignKeys[0].RefTable = GendersTable
	ProductsTable.ForeignKeys[1].RefTable = GroupOfAgesTable
	ProductsTable.ForeignKeys[2].RefTable = OfficersTable
	RecordinsurancesTable.ForeignKeys[0].RefTable = AmountpaidsTable
	RecordinsurancesTable.ForeignKeys[1].RefTable = HospitalsTable
	RecordinsurancesTable.ForeignKeys[2].RefTable = MembersTable
	RecordinsurancesTable.ForeignKeys[3].RefTable = OfficersTable
	RecordinsurancesTable.ForeignKeys[4].RefTable = ProductsTable
}
