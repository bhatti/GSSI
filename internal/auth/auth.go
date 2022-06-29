package auth

import (
	"context"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type authorizer struct {
	enforcer *casbin.Enforcer
}

type subjectContextKey struct{}

func Subject(ctx context.Context) string {
	v := ctx.Value(subjectContextKey{})
	if v == nil {
		return ""
	}
	return v.(string)
}

func Authenticate(ctx context.Context) (context.Context, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"ctx": ctx,
		}).Warn("authenticator couldn't find peer info")
		return ctx, status.New(
			codes.Unknown,
			"authenticator couldn't find peer info",
		).Err()
	}

	if peer.AuthInfo == nil {
		logrus.WithFields(logrus.Fields{
			"ctx": ctx,
			"peer": peer,
		}).Warn("authenticator couldn't find peer auth info")
		return context.WithValue(ctx, subjectContextKey{}, ""), nil
	}

	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	subject := tlsInfo.State.VerifiedChains[0][0].Subject.CommonName
	return context.WithValue(ctx, subjectContextKey{}, subject), nil
}

type Authorizer interface {
	Authorize(subject string, object string, action string) error
}

func New(model string, policy string) Authorizer {
	return &authorizer{
		enforcer: casbin.NewEnforcer(model, policy),
	}
}

func (a *authorizer) Authorize(
	subject string,
	object string,
	action string) error {
	valid := a.enforcer.Enforce(subject, object, action)
	logrus.WithFields(logrus.Fields{
		"subject": subject,
		"object": object,
		"action": action,
		"valid": valid,
	}).Info("authenticator authorizing")
	if !valid {
		return status.New(
			codes.PermissionDenied,
			fmt.Sprintf("%s not permitted to %s to %s",
				subject, action, object)).Err()
	}
	return nil
}