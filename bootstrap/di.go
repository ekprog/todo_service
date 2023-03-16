package bootstrap

import (
	"go.uber.org/dig"
	"microservice/app"
	"microservice/delivery"
	"microservice/domain"
	"microservice/interactors"
	"microservice/repos"
	"microservice/services"
)

func initDependencies(di *dig.Container) error {

	// Repos
	di.Provide(repos.NewUsersRepo, dig.As(new(domain.UsersRepository)))
	di.Provide(repos.NewProjectsRepo, dig.As(new(domain.ProjectsRepository)))
	di.Provide(repos.NewTasksRepo, dig.As(new(domain.TasksRepository)))
	di.Provide(repos.NewSmartTasksRepo, dig.As(new(domain.SmartTasksRepository)))

	// Services
	di.Provide(services.NewPeriodMatcherService, dig.As(new(domain.PeriodMatcher)))

	// Use Cases
	di.Provide(interactors.NewUsersInteractor, dig.As(new(domain.UsersInteractor)))
	di.Provide(interactors.NewProjectsInteractor, dig.As(new(domain.ProjectsInteractor)))
	di.Provide(interactors.NewTasksInteractor, dig.As(new(domain.TasksInteractor)))
	di.Provide(interactors.NewSmartTasksInteractor, dig.As(new(domain.SmartTasksInteractor)))

	di.Provide(delivery.NewToDoDeliveryService)

	// DELIVERY
	deliveryInit := func(todo *delivery.ToDoDeliveryService) error {
		if err := app.InitDelivery(todo); err == nil {
			return err
		}
		return nil
	}

	err := di.Invoke(deliveryInit)
	if err != nil {
		return err
	}

	return nil
}
