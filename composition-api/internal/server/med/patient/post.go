package patient

import (
	"context"

	api "composition-api/internal/generated/http/api"
	"composition-api/internal/services/patient"

	"github.com/AlekSi/pointer"
)

func (h *handler) MedPatientPost(ctx context.Context, req *api.MedPatientPostReq) (api.MedPatientPostRes, error) {
	id, err := h.services.PatientService.CreatePatient(ctx, patient.CreatePatientArg{
		Fullname:   req.Fullname,
		Email:      req.Email,
		Policy:     req.Policy,
		Active:     req.Active,
		Malignancy: req.Malignancy,
		BirthDate:  req.BirthDate,
	})
	if err != nil {
		return nil, err
	}

	return pointer.To(api.SimpleUuid{ID: id}), nil
}
