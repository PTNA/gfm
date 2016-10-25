package gfm

import "encoding/xml"

type FMRSError struct {
	XMLName xml.Name `xml:"error"`
	Code    int      `xml:"code,attr"`
}

type FMRSProduct struct {
	XMLName xml.Name `xml:"product"`
	Build   string   `xml:"build,attr"`
	Name    string   `xml:"name,attr"`
	Version string   `xml:"version,attr"`
}

type FMRSDataSource struct {
	XMLName         xml.Name `xml:"datasource"`
	Database        string   `xml:"database,attr"`
	DateFormat      string   `xml:"date-format,attr"`
	Layout          string   `xml:"layout,attr"`
	Table           string   `xml:"table,attr"`
	TimeFormat      string   `xml:"time-format,attr"`
	TimeStampFormat string   `xml:"timestamp-format,attr"`
	TotalCount      int      `xml:"total-count,attr"`
}

type FMRSMetadata struct {
	XMLName          xml.Name              `xml:"metadata"`
	FieldDefinitions []FMRSFieldDefinition `xml:"field-definition"`
}

type FMRSFieldDefinition struct {
	AutoEnter     string `xml:"auto-enter,attr"`
	FourDigitYear string `xml:"four-digit-year,attr"`
	Global        string `xml:"global,attr"`
	MaxRepeat     int    `xml:"max-repeat,attr"`
	Name          string `xml:"name,attr"`
	NotEmpty      string `xml:"not-empty,attr"`
	NumericOnly   string `xml:"numeric-only,attr"`
	Result        string `xml:"result,attr"`
	TimeOfDay     string `xml:"time-of-day,attr"`
	Type          string `xml:"type,attr"`
	// TODO: improve types to be booleans where appropriate
}

type FMRS struct {
	XMLName   xml.Name     `xml:"resultset"`
	Count     int          `xml:"count,attr"`
	FetchSize int          `xml:"fetch-size,attr"`
	Records   []FMRSRecord `xml:"record"`
}

type FMRSRecord struct {
	RecordID int         `xml:"record-id,attr"`
	ModID    int         `xml:"mod-id,attr"`
	Fields   []FMRSField `xml:"field"`
}

type FMRSField struct {
	Name string `xml:"name,attr"`
	Data string `xml:"data"`
}

type FMResultSet struct {
	XMLName    xml.Name `xml:"fmresultset"`
	Error      FMRSError
	Product    FMRSProduct
	Metadata   FMRSMetadata
	DataSource FMRSDataSource
	ResultSet  FMRS
}
