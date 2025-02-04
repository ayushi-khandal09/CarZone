package engine

import (
	"context"

	"github.com/ayushi-khandal09/carZone/models"
	"github.com/ayushi-khandal09/carZone/store"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func(s *EngineService) GetEngineById(ctx context.Context, id string) (*models.Engine, error) {
	engine, err := s.store.EngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	createEngine, err := s.store.EngineCreate(ctx, engineReq)
	if err != nil {
		return nil, err
	}
	return &createEngine, err
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	updateEngine, err := s.store.EngineUpdate(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &updateEngine, err
}

func (s *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	deleteEngine, err := s.store.EngineDelete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deleteEngine, err
}