package admin

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//InsertAdminSpec create Admin spec
type InsertAdminSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Password  string `validate:"required"`
}

type UpdateAdminSpec struct {
	Email     string `validate:"required"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
}

type UpdateAdminPasswordSpec struct {
	NewPassword string `validate:"required"`
	OldPassword string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct Admin service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//InsertAdmin Create new Admin and store into database
func (s *service) InsertAdmin(insertAdminSpec InsertAdminSpec) error {
	err := validator.GetValidator().Struct(insertAdminSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	admin, _ := s.repository.FindAdminByEmail(insertAdminSpec.Email)
	if admin != nil {
		return business.ErrDataExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(insertAdminSpec.Password), bcrypt.DefaultCost)
	if err != nil {
		return business.ErrInvalidSpec
	}
	var newuuid = uuid.New().String()
	Admin := NewAdmin(
		newuuid,
		insertAdminSpec.Email,
		insertAdminSpec.FirstName,
		insertAdminSpec.LastName,
		string(hashedPassword),
		newuuid,
		time.Now(),
	)

	err = s.repository.InsertAdmin(Admin)
	if err != nil {
		return err
	}

	return nil
}

//FindAdminByAdminnameAndPassword If data not found will return nil
func (s *service) FindAdminByEmailAndPassword(email string, password string) (*Admin, error) {
	return s.repository.FindAdminByEmailAndPassword(email, password)
}

//FindAdminByAdminnameAndPassword If data not found will return nil
func (s *service) FindAdminByEmail(email string) (*Admin, error) {
	return s.repository.FindAdminByEmail(email)
}

//FindAdminByID If data not found will return nil without error
func (s *service) FindAdminByID(id string) (*Admin, error) {
	return s.repository.FindAdminByID(id)
}

//UpdateAdminPaasword if data not found or old password wrong will return error
func (s *service) UpdateAdminPassword(id string, newpassword, oldPassword string, modifier string) error {

	admin, err := s.repository.FindAdminByID(id)
	if err != nil {
		return err
	} else if admin == nil {
		return business.ErrNotFound
	} else if admin.DeletedBy != "" {
		return business.ErrNotFound
	} else {
		_, err := s.repository.FindAdminByEmailAndPassword(admin.Email, oldPassword)
		if err != nil {
			return business.ErrPasswordMisMatch
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newpassword), bcrypt.DefaultCost)
	if err != nil {
		return business.ErrInvalidSpec
	}
	modifiedAdmin := admin.ModifyAdminPassword(
		string(hashedPassword),
		modifier,
		time.Now(),
	)

	return s.repository.UpdateAdminPassword(modifiedAdmin)
}

//UpdateAdmin if data not found will return error
func (s *service) UpdateAdmin(id string, updateAdminSpec UpdateAdminSpec, modifier string) error {
	admin, err := s.repository.FindAdminByID(id)
	if err != nil {
		return err
	} else if admin == nil {
		return business.ErrNotFound
	} else if admin.DeletedBy != "" {
		return business.ErrNotFound
	}

	modifiedAdmin := admin.ModifyAdmin(
		updateAdminSpec.FirstName,
		updateAdminSpec.LastName,
		time.Now(),
		modifier,
	)

	return s.repository.UpdateAdmin(modifiedAdmin)
}

//DeleteAdmin if data not found will return error
func (s *service) DeleteAdmin(id string, modifier string) error {
	admin, err := s.repository.FindAdminByID(id)
	if err != nil {
		return err
	} else if admin == nil {
		return business.ErrNotFound
	}

	deleteAdmin := admin.DeleteAdmin(
		time.Now(),
		modifier,
	)

	return s.repository.DeleteAdmin(deleteAdmin)
}
