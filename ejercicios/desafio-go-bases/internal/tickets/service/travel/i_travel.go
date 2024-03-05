package travel

type TravelService interface {
	GetTotalTickets(destination string) (int, error)
	GetCountByPeriod(time string) (int, error)
	AverageDestination(destination string) (float64, error)
}
