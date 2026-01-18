package trainer

var StorageIndexes map[string]int

var StorageLocations = []string{
	"Nahrung",
	"Kuchen",
	"Met",
	"Wasser",
	"Lehm",
	"Bruchstein",
	"Weizen",
	"Holz",
	"Leder",
	"Wolle",
	"Eisen",
	"Gold",
	"Münze",
	"Mehl",
	"Honig",
	"Kraut",
	"Pilz",
	"Öl",
	"Backstein",
	"Dachziegel",
	"Steinblock",
	"Marmor",
	"Geschirr",
	"Möbel",
	"Schuhe",
	"Holzwerkzeug",
	"Eisenwerkzeug",
	"Waffenrock",
	"Lederrüstung",
	"Kettenhemd",
	"Plattenrüstung",
	"Kurzbogen",
	"Langbogen",
	"Holzspeer",
	"Eisenspeer",
	"Kurzschwert",
	"Langschwert",
	"Kleiner Nahrungstrank",
	"Großer Nahrungstrank",
	"Kleiner Wachtrank",
	"Großer Wachtrank",
	"Kleiner Heiltrank",
	"Großer Heiltrank",
	"Nahrungsamulett",
	"Wachbleibamulett",
	"Stärkeamulett",
	"Verteidigungsamulett",
	"Trefferamulett",
	"Windamulett",
}

var CategorieNahrung = []string{
	"Honig",
	"Kuchen",
	"Mehl",
	"Met",
	"Nahrung",
	"Wasser",
	"Weizen",
}

var CategorieBauwaren = []string{
	"Backstein",
	"Bruchstein",
	"Dachziegel",
	"Holz",
	"Lehm",
	"Marmor",
	"Steinblock",
}

var CategorieResourcen = []string{
	"Eisen",
	"Gold",
	"Kraut",
	"Leder",
	"Münze",
	"Öl",
	"Pilz",
	"Wolle",
}

var CategorieWaffen = []string{
	"Eisenspeer",
	"Holzspeer",
	"Kettenhemd",
	"Kurzbogen",
	"Kurzschwert",
	"Langbogen",
	"Langschwert",
	"Lederrüstung",
	"Plattenrüstung",
	"Waffenrock",
}

var CategorieBonus = []string{
	"Eisenwerkzeug",
	"Geschirr",
	"Großer Heiltrank",
	"Großer Nahrungstrank",
	"Großer Wachtrank",
	"Holzwerkzeug",
	"Kleiner Heiltrank",
	"Kleiner Nahrungstrank",
	"Kleiner Wachtrank",
	"Möbel",
	"Schuhe",
}

var CategorieSonstiges = []string{
	"Nahrungsamulett",
	"Stärkeamulett",
	"Trefferamulett",
	"Verteidigungsamulett",
	"Wachbleibamulett",
	"Windamulett",
}

func init() {
	StorageIndexes = make(map[string]int, len(StorageLocations))

	for i, name := range StorageLocations {
		StorageIndexes[name] = i
	}
}
