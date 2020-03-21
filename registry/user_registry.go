package registry

import (
	"github.com/Shuto-san/4babe_api/interface/controller"
	ip "github.com/Shuto-san/4babe_api/interface/presenter"
	ir "github.com/Shuto-san/4babe_api/interface/repository"
	"github.com/Shuto-san/4babe_api/usecase/interactor"
	up "github.com/Shuto-san/4babe_api/usecase/presenter"
	ur "github.com/Shuto-san/4babe_api/usecase/repository"
)

func (r *registry) NewQuestionController() controller.QuestionController {
	return controller.NewQuestionController(r.NewQuestionInteractor())
}

func (r *registry) NewQuestionInteractor() interactor.QuestionInteractor {
	return interactor.NewQuestionInteractor(r.NewQuestionRepository(), r.NewQuestionPresenter())
}

func (r *registry) NewQuestionRepository() ur.QuestionRepository {
	return ir.NewQuestionRepository(r.db)
}

func (r *registry) NewQuestionPresenter() up.QuestionPresenter {
	return ip.NewQuestionPresenter()
}
