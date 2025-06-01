package med

import (
	"context"
	"errors"
	"fmt"
	"time"

	"chat_service/internal/entity"
	medpb "chat_service/internal/repository/pb/med/v1"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrDoctorNotFound  = errors.New("doctor not found")
	ErrPatientNotFound = errors.New("patient not found")
)

type Client struct {
	client medpb.MedSrvClient
	conn   *grpc.ClientConn
}

func New(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to med service: %w", err)
	}

	return &Client{
		client: medpb.NewMedSrvClient(conn),
		conn:   conn,
	}, nil
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		client: medpb.NewMedSrvClient(conn),
		conn:   conn,
	}
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}

func (c *Client) GetDoctor(ctx context.Context, doctorID string) (entity.Doctor, error) {
	req := &medpb.GetDoctorIn{
		Id: doctorID,
	}

	resp, err := c.client.GetDoctor(ctx, req)
	if err != nil {
		return entity.Doctor{}, fmt.Errorf("get doctor from med service: %w", err)
	}

	if resp.Doctor == nil {
		return entity.Doctor{}, ErrDoctorNotFound
	}

	var description string
	if resp.Doctor.Description != nil && *resp.Doctor.Description != "" {
		description = *resp.Doctor.Description
	}

	doctor, err := entity.NewDoctor(
		resp.Doctor.Id,
		resp.Doctor.Fullname,
		resp.Doctor.Org,
		resp.Doctor.Job,
		description,
	)
	if err != nil {
		return entity.Doctor{}, fmt.Errorf("create doctor entity: %w", err)
	}

	log.Debug().
		Str("doctorID", doctorID).
		Str("fullname", doctor.Fullname).
		Msg("doctor retrieved from med service")

	return doctor, nil
}

func (c *Client) GetPatient(ctx context.Context, patientID string) (entity.Patient, error) {
	req := &medpb.GetPatientIn{
		Id: patientID,
	}

	resp, err := c.client.GetPatient(ctx, req)
	if err != nil {
		return entity.Patient{}, fmt.Errorf("get patient from med service: %w", err)
	}

	if resp.Patient == nil {
		return entity.Patient{}, ErrPatientNotFound
	}

	birthDate, err := time.Parse(time.RFC3339, resp.Patient.BirthDate)
	if err != nil {
		return entity.Patient{}, fmt.Errorf("parse birth date: %w", err)
	}

	var lastUziDate time.Time

	if resp.Patient.LastUziDate != nil && *resp.Patient.LastUziDate != "" {
		parsed, err := time.Parse(time.RFC3339, *resp.Patient.LastUziDate)
		if err != nil {
			return entity.Patient{}, fmt.Errorf("parse last uzi date: %w", err)
		}

		lastUziDate = parsed
	}

	patient, err := entity.NewPatient(
		resp.Patient.Id,
		resp.Patient.Fullname,
		resp.Patient.Email,
		resp.Patient.Policy,
		resp.Patient.Active,
		resp.Patient.Malignancy,
		birthDate,
		lastUziDate,
	)
	if err != nil {
		return entity.Patient{}, fmt.Errorf("create patient entity: %w", err)
	}

	log.Debug().
		Str("patientID", patientID).
		Str("fullname", patient.Fullname).
		Msg("patient retrieved from med service")

	return patient, nil
}
