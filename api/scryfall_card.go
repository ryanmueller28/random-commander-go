package api

type ScryfallCard struct {
	Object          string            `json:"object"`
	ID              string            `json:"id"`
	OracleID        string            `json:"oracle_id"`
	MultiverseIDs   []uint64          `json:"multiverse_ids"` // In Go, `uuid` could be treated as a string
	MtgoID          uint64            `json:"mtgo_id"`
	TcgplayerID     uint64            `json:"tcgplayer_id"`
	CardmarketID    uint64            `json:"cardmarket_id"`
	Name            string            `json:"name"`
	Lang            string            `json:"lang"`
	ReleasedAt      string            `json:"released_at"`
	URI             string            `json:"uri"`
	ScryfallURI     string            `json:"scryfall_uri"`
	Layout          string            `json:"layout"`
	HighresImage    bool              `json:"highres_image"`
	ImageStatus     string            `json:"image_status"`
	ImageUris       ImageUris         `json:"image_uris"`
	ManaCost        string            `json:"mana_cost"`
	CMC             float64           `json:"cmc"`
	TypeLine        string            `json:"type_line"`
	OracleText      string            `json:"oracle_text"`
	Colors          []string          `json:"colors"`
	ColorIdentity   []string          `json:"color_identity"`
	Keywords        []string          `json:"keywords"`
	Legalities      map[string]string `json:"legalities"`
	Games           []string          `json:"games"`
	Reserved        bool              `json:"reserved"`
	Foil            bool              `json:"foil"`
	Nonfoil         bool              `json:"nonfoil"`
	Finishes        []string          `json:"finishes"`
	Oversized       bool              `json:"oversized"`
	Promo           bool              `json:"promo"`
	Reprint         bool              `json:"reprint"`
	Variation       bool              `json:"variation"`
	SetID           string            `json:"set_id"`
	Set             string            `json:"set"`
	SetName         string            `json:"set_name"`
	SetType         string            `json:"set_type"`
	SetURI          string            `json:"set_uri"`
	SetSearchURI    string            `json:"set_search_uri"`
	ScryfallSetURI  string            `json:"scryfall_set_uri"`
	RulingsURI      string            `json:"rulings_uri"`
	PrintsSearchURI string            `json:"prints_search_uri"`
	CollectorNumber string            `json:"collector_number"`
	Digital         bool              `json:"digital"`
	Rarity          string            `json:"rarity"`
	FlavorText      *string           `json:"flavor_text,omitempty"`
	CardBackID      string            `json:"card_back_id"`
	Artist          string            `json:"artist"`
	ArtistIDs       []string          `json:"artist_ids"`
	IllustrationID  string            `json:"illustration_id"`
	BorderColor     string            `json:"border_color"`
	Frame           string            `json:"frame"`
	FrameEffects    []string          `json:"frame_effects"`
	SecurityStamp   *string           `json:"security_stamp,omitempty"`
	FullArt         bool              `json:"full_art"`
	Textless        bool              `json:"textless"`
	Booster         bool              `json:"booster"`
	StorySpotlight  bool              `json:"story_spotlight"`
	EdhrecRank      *uint64           `json:"edhrec_rank,omitempty"`
	Prices          Prices            `json:"prices"`
	RelatedUris     RelatedUris       `json:"related_uris"`
	PurchaseUris    PurchaseUris      `json:"purchase_uris"`
}

type ImageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Prices struct {
	USD       *string `json:"usd,omitempty"`
	USDFoil   *string `json:"usd_foil,omitempty"`
	USDEtched *string `json:"usd_etched,omitempty"`
	EUR       *string `json:"eur,omitempty"`
	EURFoil   *string `json:"eur_foil,omitempty"`
	Tix       *string `json:"tix,omitempty"`
}

type RelatedUris struct {
	Gatherer                  *string `json:"gatherer,omitempty"`
	TcgplayerInfiniteArticles *string `json:"tcgplayer_infinite_articles,omitempty"`
	TcgplayerInfiniteDecks    *string `json:"tcgplayer_infinite_decks,omitempty"`
	Edhrec                    *string `json:"edhrec,omitempty"`
}

type PurchaseUris struct {
	Tcgplayer   *string `json:"tcgplayer,omitempty"`
	Cardmarket  *string `json:"cardmarket,omitempty"`
	Cardhoarder *string `json:"cardhoarder,omitempty"`
}
