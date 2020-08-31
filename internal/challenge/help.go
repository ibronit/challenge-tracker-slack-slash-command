package challenge

type Help struct {
	Text string
}

func GetHelp() Help {
	return Help{Text: "registration help"}
}
