package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type MessageRepository struct {
	pool *Postgres
	qb   squirrel.StatementBuilderType
}

func NewMessageRepository(pool *Postgres) *MessageRepository {
	return &MessageRepository{
		pool: pool,
		qb:   squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *MessageRepository) SaveMessage(ctx context.Context, message MessageDTO) error {
	query, args, err := r.qb.Insert("chat.messages").
		Columns("id", "chat_id", "sender_id", "content", "type", "content_type", "created_at").
		Values(
			message.ID,
			message.ChatID,
			message.SenderID,
			message.Content,
			message.Type,
			message.ContentType,
			message.CreatedAt,
		).
		ToSql()
	if err != nil {
		return fmt.Errorf("build insert query: %w", err)
	}

	if _, err := r.pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("insert message: %w", err)
	}

	return nil
}

func (r *MessageRepository) GetMessages(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]MessageDTO, error) {
	if limit <= 0 {
		limit = 50
	}

	query, args, err := r.qb.Select("id", "chat_id", "sender_id", "content", "type", "content_type", "created_at").
		From("chat.messages").
		Where(squirrel.Eq{"chat_id": chatID}).
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build select query: %w", err)
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query messages: %w", err)
	}

	defer rows.Close()

	var messages []MessageDTO

	for rows.Next() {
		var messageDTO MessageDTO

		err := rows.Scan(
			&messageDTO.ID,
			&messageDTO.ChatID,
			&messageDTO.SenderID,
			&messageDTO.Content,
			&messageDTO.Type,
			&messageDTO.ContentType,
			&messageDTO.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan message: %w", err)
		}

		messages = append(messages, messageDTO)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate messages: %w", err)
	}

	return messages, nil
}
