package repositories

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"social-network-app/internal/domain"
)

type MeetingRepository struct {
	db *sqlx.DB
}

func BuildMeetingRepository(db *sqlx.DB) *MeetingRepository {
	return &MeetingRepository{db: db}
}

func (mr *MeetingRepository) CreateMeeting(ctx context.Context, meeting domain.Meeting) (string, error) {
	query, args, err := sq.Insert("meetings").
		Columns(
			"name",
			"place",
			"comment",
			"recipient_emails",
			"applicant_email",
			"start_date",
			"end_date",
			"is_full_day",
			"is_online",
			"author_email").
		Values(
			meeting.Name,
			meeting.Place,
			meeting.Comment,
			meeting.RecipientEmails,
			meeting.ApplicantEmail,
			meeting.StartDate,
			meeting.EndDate,
			meeting.IsFullDay,
			meeting.IsOnline,
			meeting.AuthorEmail).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var insertedID string
	err = mr.db.QueryRowxContext(ctx, query, args...).Scan(&insertedID)
	if err != nil {
		return "", err
	}

	return insertedID, nil
}

func (mr *MeetingRepository) GetMeetings(ctx context.Context) ([]domain.Meeting, error) {
	var meetings []domain.Meeting
	/*рассписать поля вместо звездочки*/
	builder := sq.Select("*").From("meetings")

	sql, args, err := builder.PlaceholderFormat(sq.Dollar).ToSql()
	err = mr.db.SelectContext(ctx, &meetings, sql, args...)

	if err != nil {
		return nil, err
	}

	return meetings, nil
}
