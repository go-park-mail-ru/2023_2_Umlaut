package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type AdminService struct {
	repoAdmin repository.Admin
}

func NewAdminService(repoAdmin repository.Admin) *AdminService {
	return &AdminService{repoAdmin: repoAdmin}
}

func (s *AdminService) GetStatistic(ctx context.Context) (int, error) {
	return 0, nil
}

func (s *AdminService) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	return s.repoAdmin.CreateRecommendation(ctx, rec)
}

func (s *AdminService) CreateFeedFeedback(ctx context.Context, rec model.Recommendation) (int, error) {
	return s.repoAdmin.CreateFeedFeedback(ctx, rec)
}

func (s *AdminService) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	return s.repoAdmin.CreateFeedback(ctx, stat)
}

func (s *AdminService) GetRecommendationsStatistics(ctx context.Context) (model.RecommendationStatistic, error) {
	recommendations, err := s.repoAdmin.GetRecommendations(ctx)
	var recommendationsStat model.RecommendationStatistic
	var counter [11]int32
	sum := 0
	for _, recommendation := range recommendations {
		if recommendation.Recommend != nil {
			counter[*recommendation.Recommend] += 1
			sum += *recommendation.Recommend
		}
	}
	recommendationsStat.AvgRecommend = float32(sum) / float32(len(recommendations))
	recommendationsStat.NPS = 5
	return recommendationsStat, err
}

func (s *AdminService) GetFeedbackStatistics(ctx context.Context) (model.FeedbackStatistic, error) {
	feedbacks, err := s.repoAdmin.GetFeedbacks(ctx)
	if err != nil {
		return model.FeedbackStatistic{}, err
	}

	return getFeedbackStatistic(feedbacks), nil
}

func getFeedbackStatistic(feedbacks []model.Feedback) model.FeedbackStatistic {
	likedMap := make(map[string]int32)
	needFixMap := make(map[string]model.NeedFixObject)
	var ratingCount [11]int32
	comment := []string{}
	sum := 0
	for _, feedback := range feedbacks {
		if feedback.Liked != nil {
			likedMap[*feedback.Liked] += 1
		}
		if feedback.NeedFix != nil {
			tmp := needFixMap[*feedback.NeedFix]
			tmp.Count += 1
			if feedback.CommentFix != nil {
				tmp.CommentFix = append(tmp.CommentFix, *feedback.CommentFix)
			}
			needFixMap[*feedback.NeedFix] = tmp
		}
		if feedback.Comment != nil {
			comment = append(comment, *feedback.Comment)
		}
		ratingCount[*feedback.Rating] += 1
		sum += *feedback.Rating
	}
	return model.FeedbackStatistic{
		AvgRating:   float32(sum) / float32(len(feedbacks)),
		RatingCount: ratingCount[:],
		LikedMap:    likedMap,
		NeedFixMap:  needFixMap,
		Comments:    comment,
	}
}
