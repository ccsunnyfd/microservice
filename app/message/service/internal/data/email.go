package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
	"microservice/app/message/service/internal/biz"
)

var _ biz.EmailRepo = (*emailRepo)(nil)

type emailRepo struct {
	data *Data
	log  *log.Helper
}

func (er *emailRepo) SendMessage(ctx context.Context, mail *biz.Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", er.data.ms.From)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", "Register Confirm Mail!")
	m.SetBody("text/html", "Hello <b>" + mail.To + "</b><br/><i>" + mail.Message + "</i>!")
	s := er.data.ms.DialerPool.Get()
	if s == nil {
		return errors.New(503, "mail dialer unavailable", "cannot get a valid dialer for sending mail")
	}
	ss, ok := s.(gomail.SendCloser)
	if !ok {
		return errors.New(503, "mail dialer unavailable", "exception occurred when trying to map a valid dialer")
	}
	defer er.data.ms.DialerPool.Put(s)
	return gomail.Send(ss, m)
}

func NewEmailRepo(data *Data, logger log.Logger) biz.EmailRepo {
	return &emailRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/message/email")),
	}
}
