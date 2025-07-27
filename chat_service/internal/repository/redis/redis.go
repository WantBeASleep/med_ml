package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"chat_service/internal/config"

	"github.com/redis/go-redis/v9"
)

const (
	doctorKey  = "doctor"
	patientKey = "patient"
)

var (
	ErrDoctorNotFound  = errors.New("doctor not found")
	ErrPatientNotFound = errors.New("patient not found")
)

type Redis struct {
	client *redis.Client
}

func New(cfg config.Redis) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) SetDoctor(ctx context.Context, doctor DoctorDTO, ttl time.Duration) error {
	key := key(doctorKey, doctor.ID)

	pipe := r.client.Pipeline()

	_ = pipe.HSet(ctx, key, doctor)
	_ = pipe.Expire(ctx, key, ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("set doctor to redis: %w", err)
	}

	return nil
}

func (r *Redis) GetDoctor(ctx context.Context, doctorID string) (DoctorDTO, error) {
	pipe := r.client.Pipeline()

	dataCmd := pipe.HGetAll(ctx, key(doctorKey, doctorID))

	if _, err := pipe.Exec(ctx); err != nil {
		return DoctorDTO{}, fmt.Errorf("redis execute pipeline: %w", err)
	}

	if len(dataCmd.Val()) == 0 {
		return DoctorDTO{}, ErrDoctorNotFound
	}

	var doctor DoctorDTO
	if err := dataCmd.Scan(&doctor); err != nil {
		return DoctorDTO{}, fmt.Errorf("scan: %w", err)
	}

	return doctor, nil
}

func (r *Redis) SetPatient(ctx context.Context, patient PatientDTO, ttl time.Duration) error {
	key := key(patientKey, patient.ID)

	pipe := r.client.Pipeline()

	_ = pipe.HSet(ctx, key, patient)
	_ = pipe.Expire(ctx, key, ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return fmt.Errorf("set patient to redis: %w", err)
	}

	return nil
}

func (r *Redis) GetPatient(ctx context.Context, patientID string) (PatientDTO, error) {
	key := key(patientKey, patientID)

	pipe := r.client.Pipeline()

	dataCmd := pipe.HGetAll(ctx, key)

	if _, err := pipe.Exec(ctx); err != nil {
		return PatientDTO{}, fmt.Errorf("redis execute pipeline: %w", err)
	}

	if len(dataCmd.Val()) == 0 {
		return PatientDTO{}, ErrPatientNotFound
	}

	var patient PatientDTO
	if err := dataCmd.Scan(&patient); err != nil {
		return PatientDTO{}, fmt.Errorf("scan: %w", err)
	}

	return patient, nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func key(prefix string, id string) string {
	return fmt.Sprintf("%s:%s", prefix, id)
}
