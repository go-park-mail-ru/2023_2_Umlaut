package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
	"go.uber.org/zap"
	"log"
	"time"
)

type BackgroundService struct {
	repoUser repository.User
	repoLike repository.Like
}

func NewBackgroundService(repoUser repository.User, repoLike repository.Like) *BackgroundService {
	return &BackgroundService{repoUser: repoUser, repoLike: repoLike}
}

func (bs *BackgroundService) ResetDislike(ctx context.Context) error {
	return worker(ctx, bs.repoLike.ResetDislike, 24*time.Hour)
}

func (bs *BackgroundService) ResetLikeCounter(ctx context.Context) error {
	return worker(ctx, bs.repoUser.ResetLikeCounter, 24*time.Hour)
}

func worker(ctx context.Context, work func(context.Context) error, periodicity time.Duration) error {
	err := work(ctx)
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(periodicity)
		for {
			select {
			case <-ticker.C:
				err = work(ctx)
				if err != nil {
					logger, ok := ctx.Value(constants.KeyLogger).(*zap.Logger)
					if !ok {
						log.Println("Logger not found in context")
					} else {
						logger.Error("BackgroundService worker",
							zap.Error(err))
					}
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}
