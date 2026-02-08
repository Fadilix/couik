package cli

type Config struct {
	Mode           string `yaml:"mode"`            // default mode for quote typing
	DashboardAscii string `yaml:"dashboard_ascii"` // path for custom dashboard ascii
	QuoteType      string `yaml:"quote_type"`      // small, mid, thicc
	Time           int    `yaml:"time"`            // prefered time (30, 60, etc...)
}
