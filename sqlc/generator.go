package sqlc

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/shutej/sqlc/meta"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var boolean = regexp.MustCompile("BOOLEAN")
var integer = regexp.MustCompile("INT")
var int_64 = regexp.MustCompile("INTEGER|BIGINT")
var float_32 = regexp.MustCompile("FLOAT")
var float_64 = regexp.MustCompile("DOUBLE PRECISION|NUMERIC")
var varchar = regexp.MustCompile("VARCHAR|CHARACTER VARYING|TINYTEXT|TEXT|MEDIUMTEXT|LONGTEXT|CHAR")
var datetime = regexp.MustCompile("TIMESTAMP|DATETIME")
var date = regexp.MustCompile("DATE")
var time_ = regexp.MustCompile("TIME")
var blob = regexp.MustCompile("INET|TINYBLOB|BLOB|MEDIUMBLOB|LONGBLOB")
var dbType = regexp.MustCompile("mysql|postgres|sqlite")

type Provenance struct {
	Version   string
	Timestamp time.Time
}

type TableMeta struct {
	Name   string
	Fields []FieldMeta
}

type FieldMeta struct {
	Name string
	Type string
}

type Options struct {
	File    string `short:"f" long:"file" description:"The path to the sqlite file"`
	Url     string `short:"u" long:"url" description:"The DB URL"`
	Output  string `short:"o" long:"output" description:"The path to save the generated objects to" required:"true"`
	Package string `short:"p" long:"package" description:"The package to put the generated objects into" required:"true"`
	Type    string `short:"t" long:"type" description:"The type of the DB (mysql,postgres,sqlite)" required:"true"`
	Schema  string `short:"s" long:"schema" description:"The target DB schema (required for MySQL and Postgres)"`
	Version func() `short:"V" long:"version" description:"Print sqlc version and exit"`
	Dialect Dialect
}

func (o *Options) DbType() (Dialect, error) {
	switch o.Type {
	case "sqlite":
		return Sqlite, nil
	case "mysql":
		return MySQL, nil
	case "postgres":
		return Postgres, nil
	default:
		return Sqlite, errors.New("Invalid Db type")
	}
}

func (o *Options) Validate() error {

	if !dbType.MatchString(o.Type) {
		return errors.New("Invalid DB type")
	}

	d, err := o.DbType()
	if err != nil {
		return err
	}

	switch d {
	case MySQL, Postgres:
		if o.Schema == "" {
			return errors.New("Must specify a target schema")
		}
	}

	if o.File == "" && o.Url == "" {
		return errors.New("Must specify EITHER file path for sqlite OR url to DB")
	}

	if o.File != "" && o.Url != "" {
		return errors.New("Cannot specify BOTH file path for sqlite AND url to DB")
	}
	return nil
}

func Generate(db *sql.DB, version string, opts *Options) error {

	tables, err := opts.Dialect.metadata(opts.Schema, db)
	if err != nil {
		return err
	}

	provenance := Provenance{
		Version:   version,
		Timestamp: time.Now(),
	}

	params := make(map[string]interface{})
	params["Schema"] = opts.Schema
	params["Tables"] = tables
	params["Package"] = opts.Package
	params["Types"] = meta.Types
	params["Provenance"] = provenance

	m := template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
	}

	schemaBin, _ := sqlc_tmpl_schema_tmpl()
	t := template.Must(template.New("schema.tmpl").Funcs(m).Parse(string(schemaBin)))

	var b bytes.Buffer
	t.Execute(&b, params)

	if err := ioutil.WriteFile(opts.Output, b.Bytes(), os.ModePerm); err != nil {
		log.Fatalf("Could not write templated file: %s", err)
		return err
	}

	return nil
}

func (d Dialect) metadata(schema string, db *sql.DB) ([]TableMeta, error) {
	switch d {
	case Sqlite:
		return sqlite(db)
	case MySQL:
		return infoSchema(MySQL, schema, db)
	case Postgres:
		return infoSchema(Postgres, schema, db)
	default:
		return nil, errors.New("Unsupported dialect")
	}
}

func infoSchema(d Dialect, schema string, db *sql.DB) ([]TableMeta, error) {

	rows, err := db.Query(infoTableSQL(d), schema)
	if err != nil {
		return nil, err
	}

	tables := make([]TableMeta, 0)

	for rows.Next() {
		var t TableMeta
		rows.Scan(&t.Name)
		tables = append(tables, t)
	}

	for i, table := range tables {

		rows, err = db.Query(infoColumnsSQL(d), schema, table.Name)
		if err != nil {
			return nil, err
		}

		fields := make([]FieldMeta, 0)

		for rows.Next() {
			var colName, colType, colIsNullable sql.NullString
			err = rows.Scan(&colName, &colType, &colIsNullable)
			if err != nil {
				return nil, err
			}

			var fieldType string

			nullable := strings.ToUpper(colIsNullable.String) == "YES"

			if int_64.MatchString(colType.String) {
				fieldType = "Int64"
			} else if integer.MatchString(colType.String) {
				fieldType = "Int"
			} else if float_64.MatchString(colType.String) {
				fieldType = "Float64"
			} else if float_32.MatchString(colType.String) {
				fieldType = "Float32"
			} else if varchar.MatchString(colType.String) {
				fieldType = "String"
			} else if datetime.MatchString(colType.String) {
				fieldType = "Datetime"
			} else if date.MatchString(colType.String) {
				fieldType = "Date"
			} else if time_.MatchString(colType.String) {
				fieldType = "Time"
			} else if boolean.MatchString(colType.String) {
				fieldType = "Bool"
			} else if blob.MatchString(colType.String) {
				fieldType = "Blob"
			}

			if nullable {
				fieldType = "Null" + fieldType
			}

			field := FieldMeta{Name: colName.String, Type: fieldType}
			fields = append(fields, field)
		}
		tables[i].Fields = fields
	}

	return tables, nil
}

func sqlite(db *sql.DB) ([]TableMeta, error) {
	rows, err := db.Query("SELECT name FROM sqlite_master where type = 'table' and name NOT IN ('sqlite_sequence','schema_versions');")
	if err != nil {
		return nil, err
	}

	tables := make([]TableMeta, 0)

	for rows.Next() {
		var t TableMeta
		rows.Scan(&t.Name)
		tables = append(tables, t)
	}

	for i, table := range tables {
		pragma := fmt.Sprintf("PRAGMA table_info(%s);", table.Name)
		rows, err = db.Query(pragma)
		if err != nil {
			return nil, err
		}

		fields := make([]FieldMeta, 0)

		for rows.Next() {
			var notNull sql.NullBool
			var id, pk sql.NullInt64
			var colName, colType, defaultValue sql.NullString
			err = rows.Scan(&id, &colName, &colType, &notNull, &defaultValue, &pk)
			if err != nil {
				return nil, err
			}

			nullable := !notNull.Bool

			var fieldType string

			if int_64.MatchString(colType.String) {
				fieldType = "Int64"
			} else if integer.MatchString(colType.String) {
				fieldType = "Int"
			} else if float_64.MatchString(colType.String) {
				fieldType = "Float64"
			} else if float_32.MatchString(colType.String) {
				fieldType = "Float32"
			} else if varchar.MatchString(colType.String) {
				fieldType = "String"
			} else if datetime.MatchString(colType.String) {
				fieldType = "Time"
			} else if date.MatchString(colType.String) {
				fieldType = "Date"
			} else if time_.MatchString(colType.String) {
				fieldType = "Time"
			} else if boolean.MatchString(colType.String) {
				fieldType = "Bool"
			} else if blob.MatchString(colType.String) {
				fieldType = "Blob"
			}

			if nullable {
				fieldType = "Null" + fieldType
			}

			field := FieldMeta{Name: colName.String, Type: fieldType}
			//fmt.Printf("Field type: %s -> %s\n", fieldType, colType.String)
			fields = append(fields, field)
		}
		tables[i].Fields = fields
	}

	return tables, nil
}

func infoTableSQL(d Dialect) string {
	return fmt.Sprintf(infoTablesTmpl, d.renderPlaceholder(1))
}

func infoColumnsSQL(d Dialect) string {
	return fmt.Sprintf(infoColumnsTmpl, d.renderPlaceholder(1), d.renderPlaceholder(2))
}

const infoTablesTmpl = `
	select table_name
	from information_schema.tables
	where table_schema = %s AND table_name != 'schema_versions';
`

const infoColumnsTmpl = `
	SELECT column_name, UPPER(data_type), is_nullable
	FROM information_schema.columns
	WHERE table_schema = %s and table_name = %s;
`
