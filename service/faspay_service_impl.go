package service

import (
	usecase "faspay/faspay_services/usecase"
	vmFaspay "faspay/faspay_services/view_model"
	"time"

	raven "github.com/getsentry/raven-go"
	logging "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
)

type FaspayServiceImpl struct {
	usecaseFaspay usecase.FaspayUsecase
	logger        logging.Logger
}

func NewFaspayServiceImpl(usecaseFaspay usecase.FaspayUsecase, logger logging.Logger) FaspayService {
	return &FaspayServiceImpl{usecaseFaspay, logger}
}

/*
//Token
func (s FaspayServiceImpl) TokenHandler(req vmFaspay.EmptyVM) (*vmFaspay.FaspayResponse, error) {
	level.Info(s.logger).Log("function", "NewFaspayServiceImpl TokenHandler", "result", "Entry")

	res, err := s.usecaseFaspay.TokenHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewFaspayServiceImpl TokenHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewFaspayServiceImpl TokenHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewFaspayServiceImpl TokenHandler", "result", "Exit")
	return res, nil
}
*/

//Register
func (s FaspayServiceImpl) RegisterHandler(req vmFaspay.FaspayRegisterRequest) (*vmFaspay.FaspayResponse, error) {
	level.Info(s.logger).Log("function", "NewFaspayServiceImpl RegisterHandler", "result", "Entry")

	res, err := s.usecaseFaspay.RegisterHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewFaspayServiceImpl RegisterHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewFaspayServiceImpl RegisterHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewFaspayServiceImpl RegisterHandler", "result", "Exit")
	return res, nil
}

//Confirm Register
func (s FaspayServiceImpl) ConfirmRegisterHandler(req vmFaspay.FaspayConfirmRegisterRequest) (*vmFaspay.FaspayResponse, error) {
	level.Info(s.logger).Log("function", "NewFaspayServiceImpl ConfirmRegisterHandler", "result", "Entry")

	res, err := s.usecaseFaspay.ConfirmRegisterHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewFaspayServiceImpl ConfirmRegisterHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewFaspayServiceImpl ConfirmRegisterHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewFaspayServiceImpl ConfirmRegisterHandler", "result", "Exit")
	return res, nil
}

func (s FaspayServiceImpl) TransferHandler(req vmFaspay.FaspayTransferRequest) (*vmFaspay.FaspayResponse, error) {
	level.Info(s.logger).Log("function", "NewFaspayServiceImpl TransferHandler", "result", "Entry")

	res, err := s.usecaseFaspay.TransferHandler(req)
	if err != nil {
		level.Error(s.logger).Log("function", "NewFaspayServiceImpl TransferHandler", "Error", err)
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"function", "NewFaspayServiceImpl TransferHandler",
			"took", time.Since(begin),
		)
	}(time.Now())

	level.Info(s.logger).Log("function", "NewFaspayServiceImpl TransferHandler", "result", "Exit")
	return res, nil
}
