package android

type VersionCode int

const (
	BASE VersionCode = iota + 1
	BASE_1_1
	CUPCAKE
	DONUT
	ECLAIR
	ECLAIR_0_1
	ECLAIR_MR1
	FROYO
	GINGERBREAD
	GINGERBREAD_MR1
	HONEYCOMB
	HONEYCOMB_MR1
	HONEYCOMB_MR2
	ICE_CREAM_SANDWICH
	ICE_CREAM_SANDWICH_MR1
	JELLY_BEAN
	JELLY_BEAN_MR1
	JELLY_BEAN_MR2
	KITKAT
	KITKAT_WATCH
	LOLLIPOP
	LOLLIPOP_MR1
	M
	N
	N_MR1
	O
	O_MR1
	P
	Q
)

func (v *VersionCode) GetReleases() []Release {
	return Releases[*v]
}

type Release string

var Releases = map[VersionCode][]Release{
	BASE:                   {"1"},
	BASE_1_1:               {"1.1"},
	CUPCAKE:                {"1.5"},
	DONUT:                  {"1.6"},
	ECLAIR:                 {"2.0"},
	ECLAIR_0_1:             {"2.0.1"},
	ECLAIR_MR1:             {"2.1"},
	FROYO:                  {"2.2", "2.2.1", "2.2.2", "2.2.3"},
	GINGERBREAD:            {"2.3", "2.3.1", "2.3.2"},
	GINGERBREAD_MR1:        {"2.3.3", "2.3.4", "2.3.5", "2.3.6", "2.3.7"},
	HONEYCOMB:              {"3.0"},
	HONEYCOMB_MR1:          {"3.1"},
	HONEYCOMB_MR2:          {"3.2", "3.2.1", "3.2.2", "3.2.3", "3.2.4", "3.2.5", "3.2.6"},
	ICE_CREAM_SANDWICH:     {"4.0", "4.0.1", "4.0.2"},
	ICE_CREAM_SANDWICH_MR1: {"4.0.3", "4.0.4"},
	JELLY_BEAN:             {"4.1", "4.1.1", "4.1.2"},
	JELLY_BEAN_MR1:         {"4.2", "4.2.1", "4.2.2"},
	JELLY_BEAN_MR2:         {"4.3", "4.3.1"},
	KITKAT:                 {"4.4", "4.4.1", "4.4.2", "4.4.3", "4.4.4"},
	KITKAT_WATCH:           {"4.4W", "4.4W.1", "4.4W.2"},
	LOLLIPOP:               {"5.0", "5.0.1", "5.0.2"},
	LOLLIPOP_MR1:           {"5.1", "5.1.1"},
	M:                      {"6.0", "6.0.1"},
	N:                      {"7.0"},
	N_MR1:                  {"7.1", "7.1.1", "7.1.2"},
	O:                      {"8.0"},
	O_MR1:                  {"8.1.0"},
	P:                      {"9"},
	Q:                      {"10"},
}
