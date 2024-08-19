package goal

import (
	pb "budgeting/genproto/goal"
	"budgeting/storage/mongodb"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"
)

func TestCreateGoal(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := &pb.CreateGoalReq{
		Name:         "dodi",
		TargetAmount: "11111",
		Deadline:     "222222222",
	}
	res, err := budgeting.CreateGoal(req)
	if err != nil {
		t.Fatalf("Failed to create goal: %v", err)
	}
	fmt.Println(res)
}

func TestUpdateGoal(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := &pb.UpdateGoalReq{
		UserId:            "66ba96c7ca8deaabc9b94f24",
		Name:          "dodi1",
		TargetAmount:  "11115551",
		Deadline:      "222225552222",
		CurrentAmount: "4444444",
	}
	res, err := budgeting.UpdateGoal(req)
	if err != nil {
		t.Fatalf("Failed to update goal: %v", err)
	}
	fmt.Println(res)
}

func TestDeleteGoal(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := &pb.DeleteGoalReq{
		UserId: "66ba96c7ca8deaabc9b94f24",
	}
	res, err := budgeting.DeleteGoal(req)
	if err != nil {
		t.Fatalf("Failed to delete goal: %v", err)
	}
	fmt.Println(res)

}

func TestListGoals(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}

	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	log := slog.New(slog.NewTextHandler(log.Writer(), nil))
	repo := NewGoalRepo(mdb, log)

	req := &pb.ListGoalsReq{
		Limit:  "10",
		Offset: "0",
	}

	res, err := repo.ListGoals(req)
	if err != nil {
		t.Fatalf("Failed to list budgets: %v", err)
	}

	for _, budget := range res.Goals {
		fmt.Printf("Budget: %+v\n", budget)
	}
	if len(res.Goals) == 0 {
		t.Errorf("Expected budgets but got none")
	}
}

func TestGetUserSpending(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := pb.GetUserMoneyRequest{
		UserId:    "...",
		StartingTime: "time",
		EndingTime:   "time",
	}
	resp, err := budgeting.GetUserSpending(&req)
	if err != nil {
		t.Fatalf("Failed to get budgets: %v", err)
	}
	fmt.Println(resp)

}

func TestGetGoal(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := &pb.GetGoalReq{
		UserId: "66ba82c775b7c5315d3e7694",
	}
	resp, err := budgeting.GetGoal(req)
	if err != nil {
		t.Fatalf("Failed to get budgets: %v", err)
	}
	fmt.Println(resp)
}

func TestGetGoalReportProgress(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	req := pb.GoalProgressRequest{
		UserId:    "..",
		StartingTime: "time",
		EndingTime:   "time",
	}
	resp, err := budgeting.GetGoalReportProgress(&req)
	if err != nil {
		t.Fatalf("Failed to get budgets: %v", err)
	}
	fmt.Println(resp)
}

func TestGetBudgetSummary(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	rep := pb.UserId{
		UserId: "..",
	}
	resp, err := budgeting.GetBudgetSummary(&rep)
	if err != nil {
		t.Fatalf("Failed to get budgets: %v", err)
	}
	fmt.Println(resp)
}

func TestGetUserIncome(t *testing.T) {
	if _, err := os.Stat("logger/app.log"); os.IsNotExist(err) {
		err := os.MkdirAll("logger", os.ModePerm)
		if err != nil {
			t.Fatalf("Failed to create log directory: %v", err)
		}
		_, err = os.Create("logger/app.log")
		if err != nil {
			t.Fatalf("Failed to create log file: %v", err)
		}
	}
	mdb, err := mongodb.Connect(context.Background())
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	log := slog.Logger{}
	budgeting := NewGoalRepo(mdb, &log)
	rep := pb.GetUserMoneyRequest{
		UserId:    "..",
		StartingTime: "time",
		EndingTime:   "time",
	}
	resp, err := budgeting.GetUserIncome(&rep)
	if err != nil {
		t.Fatalf("Failed to get budgets: %v", err)
	}
	fmt.Println(resp)
}
