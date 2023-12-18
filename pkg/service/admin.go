package service

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type AdminService struct {
	repoAdmin repository.Admin
	repoUser  repository.User
}

func NewAdminService(repoAdmin repository.Admin, repoUser repository.User) *AdminService {
	return &AdminService{repoAdmin: repoAdmin, repoUser: repoUser}
}

func (s *AdminService) CreateRecommendation(ctx context.Context, rec core.Recommendation) (int, error) {
	return s.repoAdmin.CreateRecommendation(ctx, rec)
}

func (s *AdminService) CreateFeedFeedback(ctx context.Context, rec core.Recommendation) (int, error) {
	return s.repoAdmin.CreateFeedFeedback(ctx, rec)
}

func (s *AdminService) CreateFeedback(ctx context.Context, stat core.Feedback) (int, error) {
	return s.repoAdmin.CreateFeedback(ctx, stat)
}

func (s *AdminService) GetCSATType(ctx context.Context, userId int) (int, error) {
	ok, err := s.repoUser.ShowCSAT(ctx, userId)
	if err != nil {
		return 0, fmt.Errorf("GetCSATType error: %v", err)
	}
	if !ok {
		return 0, nil
	}

	ok, err = s.repoAdmin.ShowFeedback(ctx, userId)
	if err != nil {
		return 0, fmt.Errorf("GetCSATType error: %v", err)
	}
	if ok {
		return 1, nil
	}

	ok, err = s.repoAdmin.ShowRecommendation(ctx, userId)
	if err != nil {
		return 0, fmt.Errorf("GetCSATType error: %v", err)
	}
	if ok {
		return 2, nil
	}

	return 0, nil
}

func (s *AdminService) GetRecommendationsStatistics(ctx context.Context) (core.RecommendationStatistic, error) {
	recommendations, err := s.repoAdmin.GetRecommendations(ctx)
	var recommendationsStat core.RecommendationStatistic
	var counter [11]int32
	sum := 0
	for _, recommendation := range recommendations {
		if recommendation.Rating != nil {
			counter[*recommendation.Rating] += 1
			sum += *recommendation.Rating
		}
	}
	recommendationsStat.AvgRecommend = float32(sum) / float32(len(recommendations))
	recommendationsStat.NPS = 5
	return recommendationsStat, err
}

func (s *AdminService) GetFeedbackStatistics(ctx context.Context) (core.FeedbackStatistic, error) {
	feedbacks, err := s.repoAdmin.GetFeedbacks(ctx)
	if err != nil {
		return core.FeedbackStatistic{}, err
	}

	return getFeedbackStatistic(feedbacks), nil
}

func getFeedbackStatistic(feedbacks []core.Feedback) core.FeedbackStatistic {
	likedMap := make(map[string]int32)
	needFixMap := make(map[string]core.NeedFixObject)
	var ratingCount [11]int32
	var comment []string
	sum := 0
	for _, feedback := range feedbacks {
		if feedback.Liked != nil {
			likedMap[*feedback.Liked] += 1
		}
		if feedback.NeedFix != nil {
			tmp := needFixMap[*feedback.NeedFix]
			tmp.Count += 1
			if feedback.Comment != nil {
				tmp.CommentFix = append(tmp.CommentFix, *feedback.Comment)
			}
			needFixMap[*feedback.NeedFix] = tmp
		}
		if feedback.Comment != nil {
			comment = append(comment, *feedback.Comment)
		}
		ratingCount[*feedback.Rating] += 1
		sum += *feedback.Rating
	}
	return core.FeedbackStatistic{
		AvgRating:   float32(sum) / float32(len(feedbacks)),
		RatingCount: ratingCount[:],
		LikedMap:    likedMap,
		NeedFixMap:  needFixMap,
		Comments:    comment,
	}
}
