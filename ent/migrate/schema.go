// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChiiPersonFieldsColumns holds the columns for the "chii_person_fields" table.
	ChiiPersonFieldsColumns = []*schema.Column{
		{Name: "prsn_id", Type: field.TypeUint8, Increment: true},
		{Name: "prsn_cat", Type: field.TypeEnum, Enums: []string{"prsn", "crt"}},
		{Name: "gender", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "bloodtype", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "birth_year", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "year(4)"}},
		{Name: "birth_mon", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
		{Name: "birth_day", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(2)"}},
	}
	// ChiiPersonFieldsTable holds the schema information for the "chii_person_fields" table.
	ChiiPersonFieldsTable = &schema.Table{
		Name:       "chii_person_fields",
		Columns:    ChiiPersonFieldsColumns,
		PrimaryKey: []*schema.Column{ChiiPersonFieldsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "characterfields_prsn_id",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonFieldsColumns[0]},
			},
			{
				Name:    "characterfields_prsn_cat_prsn_id",
				Unique:  true,
				Columns: []*schema.Column{ChiiPersonFieldsColumns[1], ChiiPersonFieldsColumns[0]},
			},
		},
	}
	// ChiiPersonsColumns holds the columns for the "chii_persons" table.
	ChiiPersonsColumns = []*schema.Column{
		{Name: "prsn_id", Type: field.TypeInt, Increment: true},
		{Name: "prsn_name", Type: field.TypeString, Size: 255},
		{Name: "prsn_type", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "prsn_infobox", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "prsn_producer", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "prsn_mangaka", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "prsn_artist", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "prsn_seiyu", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "prsn_writer", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "prsn_illustrator", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "prsn_actor", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(1)"}},
		{Name: "prsn_summary", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "prsn_img", Type: field.TypeString, Size: 255, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "prsn_img_anidb", Type: field.TypeString, Size: 255, SchemaType: map[string]string{"mysql": "varchar(255)"}},
		{Name: "prsn_comment", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "mediumint(9)"}},
		{Name: "prsn_collects", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "mediumint(8)"}},
		{Name: "prsn_dateline", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "prsn_lastpost", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int(11)"}},
		{Name: "prsn_lock", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "prsn_anidb_id", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumint(8)"}},
		{Name: "prsn_ban", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(3)"}},
		{Name: "prsn_redirect", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "int(10)"}},
		{Name: "prsn_nsfw", Type: field.TypeBool},
	}
	// ChiiPersonsTable holds the schema information for the "chii_persons" table.
	ChiiPersonsTable = &schema.Table{
		Name:       "chii_persons",
		Columns:    ChiiPersonsColumns,
		PrimaryKey: []*schema.Column{ChiiPersonsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "person_prsn_type",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[2]},
			},
			{
				Name:    "person_prsn_producer",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[4]},
			},
			{
				Name:    "person_prsn_mangaka",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[5]},
			},
			{
				Name:    "person_prsn_artist",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[6]},
			},
			{
				Name:    "person_prsn_seiyu",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[7]},
			},
			{
				Name:    "person_prsn_writer",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[8]},
			},
			{
				Name:    "person_prsn_illustrator",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[9]},
			},
			{
				Name:    "person_prsn_lock",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[18]},
			},
			{
				Name:    "person_prsn_ban",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[20]},
			},
			{
				Name:    "person_prsn_actor",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonsColumns[10]},
			},
		},
	}
	// ChiiPersonCsIndexColumns holds the columns for the "chii_person_cs_index" table.
	ChiiPersonCsIndexColumns = []*schema.Column{
		{Name: "prsn_id", Type: field.TypeUint8, Increment: true, SchemaType: map[string]string{"mysql": "mediumint(9)"}},
		{Name: "prsn_type", Type: field.TypeEnum, Enums: []string{"prsn", "crt"}},
		{Name: "prsn_position", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "smallint(5)"}},
		{Name: "subject_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "mediumint(9)"}},
		{Name: "subject_type_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "tinyint(4)"}},
		{Name: "summary", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumtext"}},
		{Name: "prsn_appear_eps", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"mysql": "mediumtext"}},
	}
	// ChiiPersonCsIndexTable holds the schema information for the "chii_person_cs_index" table.
	ChiiPersonCsIndexTable = &schema.Table{
		Name:       "chii_person_cs_index",
		Columns:    ChiiPersonCsIndexColumns,
		PrimaryKey: []*schema.Column{ChiiPersonCsIndexColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "personcsindex_prsn_type_prsn_id_subject_id_prsn_position",
				Unique:  true,
				Columns: []*schema.Column{ChiiPersonCsIndexColumns[1], ChiiPersonCsIndexColumns[0], ChiiPersonCsIndexColumns[3], ChiiPersonCsIndexColumns[2]},
			},
			{
				Name:    "personcsindex_prsn_id",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonCsIndexColumns[0]},
			},
			{
				Name:    "personcsindex_prsn_position",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonCsIndexColumns[2]},
			},
			{
				Name:    "personcsindex_subject_type_id",
				Unique:  false,
				Columns: []*schema.Column{ChiiPersonCsIndexColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChiiPersonFieldsTable,
		ChiiPersonsTable,
		ChiiPersonCsIndexTable,
	}
)

func init() {
	ChiiPersonFieldsTable.Annotation = &entsql.Annotation{
		Table: "chii_person_fields",
	}
	ChiiPersonsTable.Annotation = &entsql.Annotation{
		Table: "chii_persons",
	}
	ChiiPersonCsIndexTable.Annotation = &entsql.Annotation{
		Table: "chii_person_cs_index",
	}
}