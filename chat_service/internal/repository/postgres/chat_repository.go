package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrChatNotFound = errors.New("chat not found")
)

type ChatRepository struct {
	pool *Postgres
	qb   squirrel.StatementBuilderType
}

func NewChatRepository(pool *Postgres) *ChatRepository {
	return &ChatRepository{
		pool: pool,
		qb:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *ChatRepository) CreateChatWithParticipants(ctx context.Context, chat ChatDTO, participantIDs []uuid.UUID) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	query, args, err := r.qb.Insert("chat.chats").
		Columns("id", "name", "description", "patient_id", "last_activity", "created_at").
		Values(chat.ID, chat.Name, chat.Description, chat.PatientID, chat.LastActivity, chat.CreatedAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("build insert chat query: %w", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("insert chat: %w", err)
	}

	if len(participantIDs) > 0 {
		insertParticipants := r.qb.Insert("chat.chat_participants").
			Columns("chat_id", "doctor_id")

		for _, participantID := range participantIDs {
			insertParticipants = insertParticipants.Values(chat.ID, participantID)
		}

		query, args, err = insertParticipants.ToSql()
		if err != nil {
			return fmt.Errorf("build insert participants query: %w", err)
		}

		if _, err := tx.Exec(ctx, query, args...); err != nil {
			return fmt.Errorf("insert participants: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *ChatRepository) CreateChat(ctx context.Context, chat ChatDTO) error {
	query, args, err := r.qb.Insert("chat.chats").
		Columns("id", "name", "description", "patient_id", "last_activity", "created_at").
		Values(chat.ID, chat.Name, chat.Description, chat.PatientID, chat.LastActivity, chat.CreatedAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("build insert query: %w", err)
	}

	_, err = r.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("insert chat: %w", err)
	}

	return nil
}

func (r *ChatRepository) GetChat(ctx context.Context, chatID uuid.UUID) (ChatDTO, error) {
	query, args, err := r.qb.Select("id", "name", "description", "patient_id", "last_activity", "created_at").
		From("chat.chats").
		Where(squirrel.Eq{"id": chatID}).
		ToSql()
	if err != nil {
		return ChatDTO{}, fmt.Errorf("build select query: %w", err)
	}

	var chatDTO ChatDTO

	if err := r.pool.QueryRow(ctx, query, args...).Scan(
		&chatDTO.ID,
		&chatDTO.Name,
		&chatDTO.Description,
		&chatDTO.PatientID,
		&chatDTO.LastActivity,
		&chatDTO.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ChatDTO{}, ErrChatNotFound
		}

		return ChatDTO{}, fmt.Errorf("scan chat: %w", err)
	}

	return chatDTO, nil
}

func (r *ChatRepository) GetChatParticipants(ctx context.Context, chatID uuid.UUID) ([]uuid.UUID, error) {
	query, args, err := r.qb.Select("doctor_id").
		From("chat.chat_participants").
		Where(squirrel.Eq{"chat_id": chatID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build select participants query: %w", err)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query participants: %w", err)
	}

	defer rows.Close()

	var participants []uuid.UUID

	for rows.Next() {
		var doctorID uuid.UUID
		if err := rows.Scan(&doctorID); err != nil {
			return nil, fmt.Errorf("scan participant: %w", err)
		}

		participants = append(participants, doctorID)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate participants: %w", err)
	}

	return participants, nil
}

func (r *ChatRepository) ListChatsByDoctorID(ctx context.Context, doctorID uuid.UUID) ([]ChatDTO, error) {
	query, args, err := r.qb.Select(
		"c.id",
		"c.name",
		"c.description",
		"c.patient_id",
		"c.last_activity",
		"c.created_at",
	).
		From("chat.chats c").
		Join("chat.chat_participants cp ON c.id = cp.chat_id").
		Where(squirrel.Eq{"cp.doctor_id": doctorID}).
		OrderBy("c.last_activity DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build select query: %w", err)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query chats: %w", err)
	}

	defer rows.Close()

	var chats []ChatDTO

	for rows.Next() {
		var chatDTO ChatDTO
		err := rows.Scan(
			&chatDTO.ID,
			&chatDTO.Name,
			&chatDTO.Description,
			&chatDTO.PatientID,
			&chatDTO.LastActivity,
			&chatDTO.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan chat: %w", err)
		}

		chats = append(chats, chatDTO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate chats: %w", err)
	}

	return chats, nil
}

func (r *ChatRepository) UpdateChatLastActivity(ctx context.Context, chatID uuid.UUID) error {
	query, args, err := r.qb.Update("chat.chats").
		Set("last_activity", time.Now()).
		Where(squirrel.Eq{"id": chatID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("build update chat last activity query: %w", err)
	}

	if _, err := r.pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("update chat last activity: %w", err)
	}

	return nil
}
