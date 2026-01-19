package trainer

var ResourceCount int

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

var CategoryNahrung = []string{
	"Honig",
	"Kuchen",
	"Mehl",
	"Met",
	"Nahrung",
	"Wasser",
	"Weizen",
}

var CategoryBauwaren = []string{
	"Backstein",
	"Bruchstein",
	"Dachziegel",
	"Holz",
	"Lehm",
	"Marmor",
	"Steinblock",
}

var CategoryResourcen = []string{
	"Eisen",
	"Gold",
	"Kraut",
	"Leder",
	"Münze",
	"Öl",
	"Pilz",
	"Wolle",
}

var CategoryWaffen = []string{
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

var CategoryBonus = []string{
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

var CategorySonstiges = []string{
	"Nahrungsamulett",
	"Stärkeamulett",
	"Trefferamulett",
	"Verteidigungsamulett",
	"Wachbleibamulett",
	"Windamulett",
}

func init() {
	ResourceCount = len(StorageLocations)
	StorageIndexes = make(map[string]int, ResourceCount)

	for i, name := range StorageLocations {
		StorageIndexes[name] = i
	}
}
