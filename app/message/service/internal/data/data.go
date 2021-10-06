package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gopkg.in/gomail.v2"
	"microservice/app/message/service/internal/biz"
	"microservice/app/message/service/internal/conf"
	"sync"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewEmailRepo,
	NewEmailSender,
)

// Data .
type Data struct {
	log *log.Helper
	ms  *biz.EmailSender
}

// NewData .
func NewData(logger log.Logger, emailSender *biz.EmailSender) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "message-service/data/email"))
	d := &Data{log: l, ms: emailSender}
	return d, func() {}, nil
}

func NewEmailSender(conf *conf.Email, logger log.Logger) (*biz.EmailSender, func(), error) {
	log1 := log.NewHelper(log.With(logger, "module", "message-service/data/email"))

	d := gomail.NewDialer(conf.Smtp, int(conf.Port), conf.Account, conf.Password)
	dialerPool := &sync.Pool{
		New: func() interface{} {
			sc, err := d.Dial()
			if err != nil {
				log1.Warnf("cannot get a connection from smtp server. error: %v\n", err)
				return nil
			}
			return sc
		},
	}
	cleanup := func() {
		//_ = sendCloser.Close()
		// 暂时没有好的清理方法，暂时让pool垃圾回收
	}
	return &biz.EmailSender{DialerPool: dialerPool, From: conf.Account}, cleanup, nil
}
