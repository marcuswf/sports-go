package db

const (
	eventsList = "list"
)

func getEventQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				name, 
				type, 
				location, 
				visible, 
				advertised_start_time 
			FROM sport_events
		`,
	}
}
