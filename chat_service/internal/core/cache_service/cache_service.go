package cache_service

import (
	"context"
	"fmt"
	"time"

	"chat_service/internal/entity"
	"chat_service/internal/repository/redis"

	"github.com/rs/zerolog/log"
)

type MedClient interface {
	GetDoctor(ctx context.Context, doctorID string) (entity.Doctor, error)
	GetPatient(ctx context.Context, patientID string) (entity.Patient, error)
}

type RedisRepository interface {
	GetDoctor(ctx context.Context, doctorID string) (redis.DoctorDTO, error)
	SetDoctor(ctx context.Context, doctor redis.DoctorDTO, ttl time.Duration) error
	GetPatient(ctx context.Context, patientID string) (redis.PatientDTO, error)
	SetPatient(ctx context.Context, patient redis.PatientDTO, ttl time.Duration) error
}

type CacheService struct {
	redisRepo RedisRepository
	medClient MedClient
	cacheTTL  time.Duration
}

func NewCacheService(redisRepo RedisRepository, medClient MedClient, cacheTTL time.Duration) *CacheService {
	return &CacheService{
		redisRepo: redisRepo,
		medClient: medClient,
		cacheTTL:  cacheTTL,
	}
}

func (s *CacheService) GetDoctor(ctx context.Context, doctorID string) (entity.Doctor, error) {
	doctorDTO, err := s.redisRepo.GetDoctor(ctx, doctorID)
	if err == nil {
		doctor, err := entity.NewDoctor(
			doctorDTO.ID,
			doctorDTO.Fullname,
			doctorDTO.Org,
			doctorDTO.Job,
			doctorDTO.Description,
		)
		if err == nil {
			log.Debug().
				Str("doctorID", doctorID).
				Msg("doctor found in cache")
			return doctor, nil
		}
	}

	doctor, err := s.medClient.GetDoctor(ctx, doctorID)
	if err != nil {
		return entity.Doctor{}, fmt.Errorf("get doctor from med service: %w", err)
	}

	doctorDTO = redis.DoctorDTO{
		ID:          doctor.ID.String(),
		Fullname:    doctor.Fullname,
		Org:         doctor.Org,
		Job:         doctor.Job,
		Description: doctor.Description,
	}

	if err := s.redisRepo.SetDoctor(ctx, doctorDTO, s.cacheTTL); err != nil {
		log.Warn().
			Err(err).
			Str("doctorID", doctorID).
			Msg("cache doctor")
	}

	log.Debug().
		Str("doctorID", doctorID).
		Str("fullname", doctor.Fullname).
		Msg("doctor retrieved from med service and cached")

	return doctor, nil
}

func (s *CacheService) GetPatient(ctx context.Context, patientID string) (entity.Patient, error) {
	patientDTO, err := s.redisRepo.GetPatient(ctx, patientID)
	if err == nil {
		patient, err := entity.NewPatient(
			patientDTO.ID,
			patientDTO.Fullname,
			patientDTO.Email,
			patientDTO.Policy,
			patientDTO.Active,
			patientDTO.Malignancy,
			patientDTO.BirthDate,
			patientDTO.LastUSG,
		)
		if err == nil {
			log.Debug().
				Str("patientID", patientID).
				Msg("patient found in cache")
			return patient, nil
		}
	}

	patient, err := s.medClient.GetPatient(ctx, patientID)
	if err != nil {
		return entity.Patient{}, fmt.Errorf("get patient from med service: %w", err)
	}

	patientDTO = redis.PatientDTO{
		ID:         patient.ID.String(),
		Fullname:   patient.Fullname,
		Email:      patient.Email,
		Policy:     patient.Policy,
		Active:     patient.Active,
		Malignancy: patient.Malignancy,
		BirthDate:  patient.BirthDate,
		LastUSG:    patient.LastUziDate,
	}

	if err := s.redisRepo.SetPatient(ctx, patientDTO, s.cacheTTL); err != nil {
		log.Warn().
			Err(err).
			Str("patientID", patientID).
			Msg("cache patient")
	}

	log.Debug().
		Str("patientID", patientID).
		Str("fullname", patient.Fullname).
		Msg("patient retrieved from med service and cached")

	return patient, nil
}
