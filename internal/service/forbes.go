package service

import (
	"go-crawler-distributed/internal/dao"
	"go-crawler-distributed/pkg/app"
)

/**
* @Author: super
* @Date: 2020-12-30 13:35
* @Description:
**/

type Forbes struct {
	NameEn         string `json:"name_en"`
	Wealth         int    `json:"wealth"`
	SourceOfWealth string `json:"source_of_wealth"`
	Region         string `json:"region"`
	ModifiedOn     string `json:"modified_on"`
	ID             string `json:"id"`
	Rank           int    `json:"rank"`
	Name           string `json:"name"`
}

type IForbesService interface {
	GetForbes() ([]*Forbes, error)
	GetForbesList(pager *app.Pager) ([]*Forbes, error)
}

type ForbesService struct {
	forbesDao dao.IForbes
}

func (s *ForbesService) GetForbes() ([]*Forbes, error) {
	result, err := s.forbesDao.SelectAll()
	if err != nil {
		return nil, err
	}
	forbes := make([]*Forbes, len(result))
	for i, f := range result {
		forbes[i] = &Forbes{
			ID:             f.ID,
			Rank:           f.Rank,
			Name:           f.Name,
			NameEn:         f.NameEn,
			Wealth:         f.Wealth,
			SourceOfWealth: f.SourceOfWealth,
			Region:         f.Region,
			ModifiedOn:     f.ModifiedOn,
		}
	}
	return forbes, nil
}

func (s *ForbesService) GetForbesList(pager *app.Pager) ([]*Forbes, error) {
	result, err := s.forbesDao.SelectList(pager.Page, pager.PageSize)
	if err != nil {
		return nil, err
	}
	forbes := make([]*Forbes, len(result))
	for i, f := range result {
		forbes[i] = &Forbes{
			ID:             f.ID,
			Rank:           f.Rank,
			Name:           f.Name,
			NameEn:         f.NameEn,
			Wealth:         f.Wealth,
			SourceOfWealth: f.SourceOfWealth,
			Region:         f.Region,
			ModifiedOn:     f.ModifiedOn,
		}
	}
	return forbes, nil
}

func NewForbesService(forbesDao dao.IForbes) IForbesService {
	return &ForbesService{forbesDao: forbesDao}
}
