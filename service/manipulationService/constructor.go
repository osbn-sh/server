package manipulationService

//type IRepo = userRepository.AuthRepo
import (
	"ostadbun/repository/postgres/manipulationRepository"
	academicservice "ostadbun/service/academicService"
	"ostadbun/service/activityService"
)

type Manipulation struct {
	manipulationRepo manipulationRepository.DB
	academic         academicservice.Service
	activity         activityService.Activity
}

func New(activity activityService.Activity, academic academicservice.Service, manipulationRepo manipulationRepository.DB) Manipulation {
	return Manipulation{
		activity:         activity,
		manipulationRepo: manipulationRepo,
		academic:         academic,
	}
}
