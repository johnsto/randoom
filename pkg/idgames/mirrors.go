package idgames

type Mirror struct {
	Name string
	CountryCode string
	URL string
}

var Mirrors = []Mirror{
	{"Gamers.org", "US", "http://www.gamers.org/pub/idgames/"},
	{"Syringa Networks", "US", "ftp://mirrors.syringanetworks.net/idgames/"},
	{"Mancubus", "US", "http://ftp.mancubus.net/pub/idgames/"},   
	{"youfailit", "US", "http://youfailit.net/pub/idgames/"},
	{"Quaddicted", "DE", "https://www.quaddicted.com/files/idgames/"},
	{"University of Athens", "GR", "http://ftp.ntua.gr/mirror/idgames/idstuff/"},
}
