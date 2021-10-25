package model

type Class string

const (
	ClassBeast   Class = "Beast"
	ClassAqua    Class = "Aquatic"
	ClassPlant   Class = "Plant"
	ClassBird    Class = "Bird"
	ClassBug     Class = "Bug"
	ClassReptile Class = "Reptile"
	ClassMech    Class = "Mech"
	ClassDawn    Class = "Dawn"
	ClassDusk    Class = "Dusk"
)

var MapSpecialPartID = map[string]string{
	// Bionic
	"horn-p4r451t3":    "horn-parasite",
	"horn-5h04l-5t4r":  "horn-shoal-star",
	"back-1nd14n-5t4r": "back-indian-star",

	// Mystic
	"eyes-calico-zeal":         "eyes-zeal",
	"ears-pointy-nyan":         "ears-nyan",
	"back-hasagi":              "back-ronin",
	"horn-winter-branch":       "horn-little-branch",
	"tail-cottontail":          "tail-sakura-cottontail",
	"mouth-skull-cracker":      "mouth-nut-cracker",
	"mouth-feasting-mosquito":  "mouth-mosquito",
	"horn-laggingggggg":        "horn-lagging",
	"tail-fire-ant":            "tail-ant",
	"back-starry-shell":        "back-snail-shell",
	"ears-vector":              "ears-larva",
	"eyes-broken-bookworm":     "eyes-bookworm",
	"eyes-insomnia":            "eyes-sleepless",
	"mouth-lam-handsome":       "mouth-lam",
	"horn-candy-babylonia":     "horn-babylonia",
	"ears-red-nimo":            "ears-nimo",
	"tail-kuro-koi":            "tail-koi",
	"back-crystal-hermit":      "back-hermit",
	"ears-heart-cheek":         "ears-pink-cheek",
	"tail-snowy-swallow":       "tail-swallow",
	"back-starry-balloon":      "back-balloon",
	"horn-golden-eggshell":     "horn-eggshell",
	"mouth-mr-doubletalk":      "mouth-doubletalk",
	"eyes-sky-mavis":           "eyes-mavis",
	"eyes-crimson-gecko":       "eyes-gecko",
	"mouth-venom-bite":         "mouth-toothless-bite",
	"ears-deadly-pogona":       "ears-pogona",
	"back-rugged-sail":         "back-bone-sail",
	"tail-escaped-gecko":       "tail-wall-gecko",
	"horn-pinku-unko":          "horn-unko",
	"tail-namek-carrot":        "tail-carrot",
	"mouth-humorless":          "mouth-serious",
	"eyes-dreamy-papi":         "eyes-papi",
	"ears-the-last-leaf":       "ears-leafy",
	"back-pink-turnip":         "back-turnip",
	"horn-golden-bamboo-shoot": "horn-bamboo-shoot",

	// Xmas
	"eyes-snowflakes":        "eyes-little-peas",
	"ears-merry-lamb":        "ears-innocent-lamb",
	"back-candy-canes":       "back-garish-worm",
	"horn-spruce-spear":      "horn-feather-spear",
	"tail-december-surprise": "tail-snake-jar",
	"mouth-rudolph":          "mouth-zigzag",

	// Japan
	"back-hamaya":     "back-risky-beast",
	"horn-umaibo":     "horn-pocky",
	"horn-kendama":    "horn-imp",
	"mouth-kawaii":    "mouth-cute-bunny",
	"tail-maki":       "tail-fish-snack",
	"ears-mon":        "ears-earwing",
	"eyes-yen":        "eyes-sleepless",
	"mouth-geisha":    "mouth-piranha",
	"tail-koinobori":  "tail-koi",
	"ears-karimata":   "ears-risky-bird",
	"tail-omatsuri":   "tail-granmas-fan",
	"back-origami":    "back-cupid",
	"eyes-kabuki":     "eyes-topaz",
	"eyes-dokuganryu": "eyes-scar",
	"mouth-dango":     "mouth-tiny-turtle",
	"ears-maiko":      "ears-sakura",
	"back-yakitori":   "back-shiitake",
	"horn-yorishiro":  "horn-beech",
}

type Figure struct {
	Atlas    string `json:"atlas,omitempty"`
	Model    string `json:"model,omitempty"`
	Image    string `json:"image,omitempty"`
	Typename string `json:"__typename,omitempty"`
}
type Parts struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Class        Class         `json:"class,omitempty"`
	Type         string        `json:"type,omitempty"`
	SpecialGenes string        `json:"specialGenes,omitempty"`
	Stage        int           `json:"stage,omitempty"`
	Abilities    []interface{} `json:"abilities,omitempty"`
	Typename     string        `json:"__typename,omitempty"`
}
type Stats struct {
	Hp       int    `json:"hp,omitempty"`
	Speed    int    `json:"speed,omitempty"`
	Skill    int    `json:"skill,omitempty"`
	Morale   int    `json:"morale,omitempty"`
	Typename string `json:"__typename,omitempty"`
}

type OwnerProfile struct {
	Name     string `json:"name,omitempty"`
	Typename string `json:"__typename,omitempty"`
}
type BattleInfo struct {
	Banned   *bool  `json:"banned,omitempty"`
	BanUntil int    `json:"banUntil,omitempty"`
	Level    int    `json:"level,omitempty"`
	Typename string `json:"__typename,omitempty"`
}
type Children struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Class    Class  `json:"class,omitempty"`
	Image    string `json:"image,omitempty"`
	Title    string `json:"title,omitempty"`
	Stage    int    `json:"stage,omitempty"`
	Typename string `json:"__typename,omitempty"`
}
type Axie struct {
	ID           string       `json:"id,omitempty"`
	Image        string       `json:"image,omitempty"`
	Class        Class        `json:"class,omitempty"`
	Chain        string       `json:"chain,omitempty"`
	Name         string       `json:"name,omitempty"`
	Genes        string       `json:"genes,omitempty"`
	Owner        string       `json:"owner,omitempty"`
	BirthDate    int          `json:"birthDate,omitempty"`
	BodyShape    string       `json:"bodyShape,omitempty"`
	SireID       int          `json:"sireId,omitempty"`
	SireClass    string       `json:"sireClass,omitempty"`
	MatronID     int          `json:"matronId,omitempty"`
	MatronClass  string       `json:"matronClass,omitempty"`
	Stage        int          `json:"stage,omitempty"`
	Title        string       `json:"title,omitempty"`
	BreedCount   int          `json:"breedCount,omitempty"`
	Level        int          `json:"level,omitempty"`
	Figure       Figure       `json:"figure,omitempty"`
	Parts        []Parts      `json:"parts,omitempty"`
	Stats        Stats        `json:"stats,omitempty"`
	Auction      Auction      `json:"auction,omitempty"`
	OwnerProfile OwnerProfile `json:"ownerProfile,omitempty"`
	BattleInfo   BattleInfo   `json:"battleInfo,omitempty"`
	Children     []Children   `json:"children,omitempty"`
	Typename     string       `json:"__typename,omitempty"`
}

func (a Axie) IsEmpty() bool {
	return a.ID == ""
}
