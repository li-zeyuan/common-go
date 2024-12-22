package utils

import (
	"context"
	"errors"
	"net/url"
	"path"

	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

func Url2ObjectKey(ctx context.Context, rawURL string) (string, error) {
	if len(rawURL) == 0 {
		mylogger.Error(ctx, "empty raw url")
		return "", errors.New("empty raw url")
	}

	uri, err := url.Parse(rawURL)
	if err != nil {
		mylogger.Error(ctx, "parse raw url fail", zap.Error(err))
		return "", err
	}

	return path.Base(uri.Path), nil
}

func Url2ObjectKeyList(ctx context.Context, rawURLs []string) ([]string, error) {
	if len(rawURLs) == 0 {
		return []string{}, nil
	}

	res := make([]string, 0, len(rawURLs))
	for _, rUrl := range rawURLs {
		uri, err := url.Parse(rUrl)
		if err != nil {
			mylogger.Error(ctx, "parse raw url fail", zap.Error(err))
			return nil, err
		}

		res = append(res, path.Base(uri.Path))
	}


	return res,nil
}