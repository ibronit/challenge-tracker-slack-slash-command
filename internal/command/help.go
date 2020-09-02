package command

type Help struct {
	Text string `json:"text"`
}

func GetHelp() Help {
	return Help{Text: "help"}
}
