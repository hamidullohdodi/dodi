package postgres

import (
	pb "auth_service/genproto"
	"log"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}

	rep := &pb.RegisterUserReq{
		Username: "fsesff1s11dfrffdfsd1f",
		Password: "12134r11eerfe1f56",
		Email:    "e111rf@gmai1l.com",
		FullName: "rf111e1",
		UserType: "er111ff1ffr",
		Bio:      "dod111i1rtgrtfffrefegtrg",
	}
	resp, err := UserRepo.RegisterUser(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}
func TestLoginUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}

	rep := &pb.LoginReq{
		Email:    "erf@gmail.com",
		Password: "1234reerfef56",
	}

	resp, err := UserRepo.LoginUser(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}

func TestGetUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}
	rep := &pb.ById{
		Id: "ed414140-dab8-4815-bd71-0b740ae1414e",
	}
	resp, err := UserRepo.GetUser(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}
	rep := &pb.UpdateUserReq{
		Username: "dodi",
		Password: "123456",
	}
	resp, err := UserRepo.UpdateUser(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}

func TestUpdateUserType(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}

	rep := &pb.UserTypeReq{
		UserType: "customer",
		Id:       "f6b2aac0-9965-4c54-9ea4-058710e6f6ce",
	}
	resp, err := UserRepo.UpdateUserType(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}

func TestGetAllUsers(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}
	rep := &pb.PageLimit{
		Page:  1,
		Limit: 3,
	}
	resp, err := UserRepo.GetAllUsers(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
		t.Errorf("%v", resp)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	UserRepo := UserAuthRepo{db: db}
	rep := &pb.ById{
		Id: "ed414140-dab8-4815-bd71-0b740ae1414e",
	}
	resp, err := UserRepo.DeleteUser(rep)
	if err != nil {
		t.Errorf("%v", err)
	}
	if resp == nil {
	}
}
