package wad

// MD5 values from doomwiki.org
const (
	Doom_1_1_MD5                string = "981b03e6d1dc033301aa3095acc437ce"
	Doom_1_2_MD5                       = "792fd1fea023d61210857089a7c1e351"
	Doom_1_666_MD5                     = "54978d12de87f162b9bcc011676cb3c0"
	Doom_1_8_MD5                       = "11e1cd216801ea2657723abc86ecb01f"
	Doom_1_9_MD5                       = "1cd63c5ddff1bf8ce844237f580e9cf3"
	Doom_1_9ud_MD5                     = "c4fe9fd920207691a9f493668e0a2083"
	Doom_BFG_MD5                       = "fb35c4a5a9fd49ec29ab6e900572c524"
	Doom_Plutonia_MD5                  = "75c8cf89566741fa9d22447604053bd7"
	Doom_Plutonia_Anthology_MD5        = "3493be7e1e2588bc9c8b31eab2587a04"
	Doom_TNT_MD5                       = "4e158d9953c79ccf97bd0663244cc6b6"
	Doom_TNT_Anthology_MD5             = "1d39e405bf6ee3df69a8d2646c8d5c49"

	Doom2_1_666g_MD5 string = "d9153ced9fd5b898b36cc5844e35b520"
	Doom2_1_666_MD5         = "30e3c2d0350b67bfbf47271970b74b2f"
	Doom2_1_7_MD5           = "ea74a47a791fdef2e9f2ea8b8a9da13b"
	Doom2_1_7a_MD5          = "d7a07e5d3f4625074312bc299d7ed33f"
	Doom2_1_8_MD5           = "c236745bb01d89bbb866c8fed81b6f8c"
	Doom2_1_8f_MD5          = "3cb02349b3df649c86290907eed64e7b"
	Doom2_1_9_MD5           = "25e1459ca71d321525f84628f45ca8cd"
	Doom2_BFG_MD5           = "c3bea40570c23e511a7ed3ebcd9865f7"

	Heretic_1_0_MD5 string = "3117e399cdb4298eaa3941625f4b2923"
	Heretic_1_2_MD5        = "1e4cb4ef075ad344dd63971637307e04"
	Heretic_1_3_MD5        = "66d686b1ed6d35ff103f15dbd30e0341"

	Hexen_1_0_MD5     string = "b2543a03521365261d0a0f74d5dd90f0"
	Hexen_1_1_MD5            = "abb033caf81e26f12a2103e1fa25453f"
	Hexen_1_1_Mac_MD5        = "b68140a796f6fd7f3a5d3226a32b93be"

	Strife_1_1_MD5 string = "8f2d3a6a289f5d2f2f9c1eec02b47299"
	Strife_1_2_MD5        = "2fed2031a5b03892106e0f117f17901f"
)

// Builds is a list of all known commercial releases, their versions, and the
// hash of their IWAD file.
var IWADs = map[string]GameVersion{
	// Doom
	Doom_1_1_MD5:   {Doom, 1, "1.1"},
	Doom_1_2_MD5:   {Doom, 2, "1.2"},
	Doom_1_666_MD5: {Doom, 3, "1.666"},
	Doom_1_8_MD5:   {Doom, 4, "1.8"},
	Doom_1_9_MD5:   {Doom, 5, "1.9"},
	Doom_1_9ud_MD5: {Doom, 5, "1.9ud"},
	Doom_BFG_MD5:   {Doom, 7, "BFG"},
	// Doom 2
	Doom2_1_666g_MD5: {Doom2, 1, "1.666g"},
	Doom2_1_666_MD5:  {Doom2, 1, "1.666"},
	Doom2_1_7_MD5:    {Doom2, 3, "1.7"},
	Doom2_1_7a_MD5:   {Doom2, 3, "1.7a"},
	Doom2_1_8_MD5:    {Doom2, 5, "1.8"},
	Doom2_1_8f_MD5:   {Doom2, 5, "1.8f"},
	Doom2_1_9_MD5:    {Doom2, 6, "1.9"},
	Doom2_BFG_MD5:    {Doom2, 7, "BFG"},
	// Heretic
	Heretic_1_0_MD5: {Heretic, 1, "1.0"},
	Heretic_1_2_MD5: {Heretic, 2, "1.2"},
	Heretic_1_3_MD5: {Heretic, 3, "1.3"},
	// Hexen
	Hexen_1_0_MD5:     {Hexen, 1, "1.0"},
	Hexen_1_1_MD5:     {Hexen, 2, "1.1"},
	Hexen_1_1_Mac_MD5: {Hexen, 2, "1.1m"},
	// Strife
	Strife_1_1_MD5: {Strife, 1, "1.1"},
	Strife_1_2_MD5: {Strife, 2, "1.2"},
}
