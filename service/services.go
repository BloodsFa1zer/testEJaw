package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"test/database"
)

type SellerService struct {
	DbSeller database.SellerDatabaseInterface
	validate *validator.Validate
	//	RedisClient redisDB.ClientRedisInterface
}

func NewSellerService(DbSeller database.SellerDatabaseInterface, validate *validator.Validate) *SellerService {
	return &SellerService{DbSeller: DbSeller, validate: validate}
}

var ErrSellerNotFound = errors.New("there is no such seller")

type SellerServiceInterface interface {
	GetAllSellers() (*[]database.Seller, error)
	GetByIDSeller(id int) (*database.Seller, error)
	CreateSeller(seller database.Seller) (int, error)
	UpdateSeller(seller database.Seller) error
	DeleteSeller(id int) error
}

func (ss *SellerService) GetAllSellers() (*[]database.Seller, error) {
	sellers, err := ss.DbSeller.SelectAll()
	if err != nil {
		log.Warn().Err(err).Msg("cannot select sellers")
		return nil, ErrSellerNotFound
	}

	return sellers, nil
}

func (ss *SellerService) GetByIDSeller(id int) (*database.Seller, error) {
	seller, err := ss.DbSeller.SelectByID(id)
	if err != nil {
		log.Warn().Err(err).Msg("cannot select sellers")
		return nil, ErrSellerNotFound
	}

	return seller, nil
}

func (ss *SellerService) CreateSeller(seller database.Seller) (int, error) {
	newSeller := database.Seller{Name: seller.Name, Phone: seller.Phone}
	id, err := ss.DbSeller.Insert(newSeller)
	if err != nil {
		log.Warn().Err(err).Msg("cannot create user")
		return 0, err
	}
	log.Info().Msg("user successfully created")

	return id, nil
}

func (ss *SellerService) UpdateSeller(seller database.Seller) error {
	updatedSeller := database.Seller{ID: seller.ID, Name: seller.Name, Phone: seller.Phone}

	_, err := ss.DbSeller.SelectByID(seller.ID)
	if err != nil {
		return ErrSellerNotFound
	}

	err = ss.DbSeller.Update(updatedSeller)
	if err != nil {
		log.Warn().Err(err).Msg("cannot update user")
		return err
	}
	log.Info().Msg("user successfully updated")

	return nil
}

func (ss *SellerService) DeleteSeller(id int) error {
	err := ss.DbSeller.Delete(id)
	if err != nil {
		log.Warn().Err(err).Msg("cannot delete user")
		return err
	}
	log.Info().Msg("user successfully deleted")

	return nil
}
