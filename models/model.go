package models

// api 1
type News struct {
	Title    string   `json:"title"`
	URL      string   `json:"url"`
	Source   string   `json:"source"`
	Date     string   `json:"date"`
	Keywords []string `json:"keywords"`
}

// api 2
type Music struct {
	ThumbURL                   string        `json:"thumb_url"`
	Mbid                       string        `json:"mbid"`
	FacebookPageURL            string        `json:"facebook_page_url"`
	ImageURL                   string        `json:"image_url"`
	TrackerCount               int64         `json:"tracker_count"`
	Tracking                   []interface{} `json:"tracking"`
	UpcomingEventCount         int64         `json:"upcoming_event_count"`
	URL                        string        `json:"url"`
	SupportURL                 string        `json:"support_url"`
	ShowMultiTicket            bool          `json:"show_multi_ticket"`
	Name                       string        `json:"name"`
	Options                    Options       `json:"options"`
	Links                      []Link        `json:"links"`
	ArtistOptinShowPhoneNumber bool          `json:"artist_optin_show_phone_number"`
	ID                         string        `json:"id"`
}

type Link struct {
	Type string  `json:"type"`
	URL  *string `json:"url,omitempty"`
}

type Options struct {
	DisplayListenUnit bool `json:"display_listen_unit"`
}

// api 3
type Canonical struct {
	Version     string   `json:"version"`
	GeneratedAt string   `json:"generated_at"`
	Count       int64    `json:"count"`
	Entities    []Entity `json:"entities"`
}

type Entity struct {
	ID         string   `json:"id"`
	Label      string   `json:"label"`
	Aliases    []string `json:"aliases"`
	Type       Type     `json:"type"`
	VideoCount int64    `json:"videoCount"`
	UpdatedAt  string   `json:"updatedAt"`
}

// api 4

type Type string

const (
	Asset   Type = "asset"
	Concept Type = "concept"
	Country Type = "country"
	Org     Type = "org"
	Person  Type = "person"
)

type Breweries struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	BreweryType   BreweryType `json:"brewery_type"`
	Address1      *string     `json:"address_1"`
	Address2      *string     `json:"address_2"`
	Address3      interface{} `json:"address_3"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	Longitude     *float64    `json:"longitude"`
	Latitude      *float64    `json:"latitude"`
	Phone         *string     `json:"phone"`
	WebsiteURL    *string     `json:"website_url"`
	State         string      `json:"state"`
	Street        *string     `json:"street"`
}

type BreweryType string

const (
	Brewpub    BreweryType = "brewpub"
	Closed     BreweryType = "closed"
	Large      BreweryType = "large"
	Micro      BreweryType = "micro"
	Proprietor BreweryType = "proprietor"
)

// api 5

type Whater struct {
	IcaoID      string      `json:"icaoId"`
	ReceiptTime string      `json:"receiptTime"`
	ObsTime     int64       `json:"obsTime"`
	ReportTime  string      `json:"reportTime"`
	Temp        float64     `json:"temp"`
	Dewp        float64     `json:"dewp"`
	Wdir        interface{} `json:"wdir"`
	Wspd        int64       `json:"wspd"`
	Visib       interface{} `json:"visib"`
	Altim       float64     `json:"altim"`
	Slp         float64     `json:"slp"`
	QcField     int64       `json:"qcField"`
	MetarType   string      `json:"metarType"`
	RawOb       string      `json:"rawOb"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"lon"`
	Elev        int64       `json:"elev"`
	Name        string      `json:"name"`
	Cover       string      `json:"cover"`
	Clouds      []Cloud     `json:"clouds"`
	FltCat      string      `json:"fltCat"`
	WxString    *string     `json:"wxString,omitempty"`
	Precip      *float64    `json:"precip,omitempty"`
	Wgst        *int64      `json:"wgst,omitempty"`
}

type Cloud struct {
	Cover string `json:"cover"`
	Base  int64  `json:"base"`
}

// --- Nuevo modelo para regresar los 5 objetos principales ---

type DashboardData struct {
	NewsList   []News      `json:"news"`
	MusicList  Music       `json:"music"`
	Canonical  Canonical   `json:"canonical"`
	Breweries  []Breweries `json:"breweries"`
	WhaterList []Whater    `json:"whater"`
}

type Result[T any] struct {
	Data    T      `json:"data"`
	Error   error  `json:"-"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type TaskResult struct {
	Source string
	Data   any
	Error  error
}
