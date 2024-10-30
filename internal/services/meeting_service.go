package services

import (
	"context"
	"log/slog"
	"social-network-app/internal/domain"
	"social-network-app/internal/repositories"
)

type MeetingService struct {
	meetingRepository *repositories.MeetingRepository
	log               *slog.Logger
}

func BuildMeetingService(meetingRepository *repositories.MeetingRepository, log *slog.Logger) *MeetingService {
	return &MeetingService{meetingRepository: meetingRepository, log: log}
}

func (ms *MeetingService) CreateMeeting(ctx context.Context, meeting domain.Meeting) (string, error) {
	id, err := ms.meetingRepository.CreateMeeting(ctx, meeting)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (ms *MeetingService) GetMeetings(ctx context.Context) ([]domain.Meeting, error) {
	meetings, err := ms.meetingRepository.GetMeetings(ctx)

	if err != nil {
		return nil, err
	}

	return meetings, nil
}
