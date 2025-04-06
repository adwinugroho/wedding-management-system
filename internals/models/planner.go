package models

import "time"

type (
	Budget struct {
		ID             string          `json:"id"`
		BudgetType     string          `json:"budget_type"` // contract/engagement/reception
		TotalGuest     int             `json:"total_guest"`
		Document       []Document      `json:"document"`
		Venue          *Venue          `json:"venue"`
		Catering       *Catering       `json:"catering"`
		Entertainment  *Entertainment  `json:"entertainment"`
		Decoration     *Decoration     `json:"decoration"`
		Documentation  *Documentation  `json:"documentation"`
		Attire         *Attire         `json:"attire"`
		Transportation *Transportation `json:"transportation"`
		MakeupArtist   *MakeupArtist   `json:"makeup_artist"`
		Mahar          []Mahar         `json:"mahar"`
		Other          []Other         `json:"other"`
		TotalPrice     int             `json:"total_price"`    // total price
		PaymentStatus  string          `json:"payment_status"` // paid/unpaid/pending
		CreatedAt      time.Time       `json:"created_at"`
		UpdatedAt      time.Time       `json:"updated_at"`
	}

	Venue struct {
		VenueName       string         `json:"venue_name"`
		VenueType       string         `json:"venue_type"` // hotel/restaurant/other
		Location        string         `json:"location"`
		Capacity        int            `json:"capacity"`
		Price           int            `json:"price"`
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
		IncludeSetup    map[string]any `json:"include_setup"`
	}

	Catering struct {
		Pax             int            `json:"pax"`
		Buffet          int            `json:"buffet"`           // prasmanan
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
		Menu            []Menu         `json:"menu"`
		IncludeSetup    map[string]any `json:"include_setup"`
	}

	Menu struct {
		Name      string `json:"name"`
		ItemPrice string `json:"item_price"` // price per pax
	}

	Entertainment struct {
		Performer       []Performer    `json:"performer"`
		SoundSystem     bool           `json:"sound_system"`
		MasterCeremony  bool           `json:"mc"`
		Duration        int            `json:"duration"`         // in hour
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
		IncludeSetup    map[string]any `json:"include_setup"`
	}

	Performer struct {
		Name            string `json:"name"`
		TypePerformer   int    `json:"type_performer"`   // acoustic/electro/other art
		PaymentStatus   string `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string `json:"payment_deadline"` // 2025-01-01
	}

	Decoration struct {
		DecorationType  string         `json:"decoration_type"` // floral/lighting/other art
		Size            int            `json:"size"`            // in square meter
		IncludeSetup    map[string]any `json:"include_setup"`
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Documentation struct {
		IsFotographer   bool           `json:"is_fotographer"`
		IsVideographer  bool           `json:"is_videographer"`
		IncludeSetup    map[string]any `json:"include_setup"`
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Attire struct {
		IsBrideAttire   bool           `json:"is_bride_attire"`  // pakaian pagar ayu
		IsGroomAttire   bool           `json:"is_groom_attire"`  // pakaian mempelai pria dan wanita
		IsFamilyAttire  bool           `json:"is_family_attire"` // pakaian keluarga
		IncludeSetup    map[string]any `json:"include_setup"`
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Transportation struct {
		TransportationType string         `json:"transportation_type"` // shuttle/car/other
		IncludeSetup       map[string]any `json:"include_setup"`
		Price              int            `json:"price"`            // total price
		PaymentStatus      string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline    string         `json:"payment_deadline"` // 2025-01-01
	}

	MakeupArtist struct {
		IsMakeupArtist  bool           `json:"is_makeup_artist"`
		IncludeSetup    map[string]any `json:"include_setup"`
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Other struct {
		OtherName       string         `json:"other_name"`  // atk, suvenir, undangan cetak, undangan digital, kamar hotel dll
		OtherPrice      int            `json:"other_price"` // price per type
		IncludeSetup    map[string]any `json:"include_setup"`
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Mahar struct {
		MaharName       string         `json:"mahar_name"`   // mahar, wedding ring, seserahan, kotak seserahan
		WeddingRing     string         `json:"wedding_ring"` // emas/perak/platinum
		IncludeSetup    map[string]any `json:"include_setup"`
		Price           int            `json:"price"`            // total price
		PaymentStatus   string         `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string         `json:"payment_deadline"` // 2025-01-01
	}

	Document struct {
		DocumentName    string                `json:"document_name"` // Perlengkapan untuk KUA/Izin gedung dll.
		DocumentType    string                `json:"document_type"` // pdf/image/video
		DocumentURL     string                `json:"document_url"`
		Price           int                   `json:"price"`            // total price
		PaymentStatus   string                `json:"payment_status"`   // paid/unpaid/pending
		PaymentDeadline string                `json:"payment_deadline"` // 2025-01-01
		Requirement     []DocumentRequirement `json:"requirement"`      // KUA/Izin gedung dll.
	}

	DocumentRequirement struct {
		Name     string `json:"name"`
		Price    int    `json:"price"`
		Quantity int    `json:"quantity"`
	}
)
