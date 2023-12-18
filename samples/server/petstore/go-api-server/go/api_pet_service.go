/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"context"
	"net/http"
	"errors"
	"time"
	"os"
)

// PetAPIService is a service that implements the logic for the PetAPIServicer
// This service should implement the business logic for every endpoint for the PetAPI API.
// Include any external packages or services that will be required by this service.
type PetAPIService struct {
}

// NewPetAPIService creates a default api service
func NewPetAPIService() PetAPIServicer {
	return &PetAPIService{}
}

// AddPet - Add a new pet to the store
func (s *PetAPIService) AddPet(ctx context.Context, pet Pet) (ImplResponse, error) {
	// TODO - update AddPet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	// return Response(200, Pet{}), nil

	// TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	// return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddPet method not implemented")
}

// DeletePet - Deletes a pet
func (s *PetAPIService) DeletePet(ctx context.Context, petId int64, apiKey string) (ImplResponse, error) {
	// TODO - update DeletePet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePet method not implemented")
}

// FilterPetsByCategory - Finds Pets
func (s *PetAPIService) FilterPetsByCategory(ctx context.Context, gender Gender, species Species, notSpecies []Species) (ImplResponse, error) {
	// TODO - update FilterPetsByCategory with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	// return Response(200, []Pet{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FilterPetsByCategory method not implemented")
}

// FindPetsByStatus - Finds Pets by status
func (s *PetAPIService) FindPetsByStatus(ctx context.Context, status []string, inlineEnumPath string, inlineEnum string) (ImplResponse, error) {
	// TODO - update FindPetsByStatus with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	// return Response(200, []Pet{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindPetsByStatus method not implemented")
}

// FindPetsByTags - Finds Pets by tags
// Deprecated
func (s *PetAPIService) FindPetsByTags(ctx context.Context, tags []string, bornAfter time.Time, bornBefore time.Time) (ImplResponse, error) {
	// TODO - update FindPetsByTags with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	// return Response(200, []Pet{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindPetsByTags method not implemented")
}

// GetPetById - Find pet by ID
func (s *PetAPIService) GetPetById(ctx context.Context, petId int64) (ImplResponse, error) {
	// TODO - update GetPetById with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	// return Response(200, Pet{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPetById method not implemented")
}

// GetPetImageById - Returns the image for the Pet that has been previously uploaded
func (s *PetAPIService) GetPetImageById(ctx context.Context, petId int64) (ImplResponse, error) {
	// TODO - update GetPetImageById with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, *os.File{}) or use other options such as http.Ok ...
	// return Response(200, *os.File{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPetImageById method not implemented")
}

// GetPetsByTime - Get the pets by time
func (s *PetAPIService) GetPetsByTime(ctx context.Context, createdTime time.Time) (ImplResponse, error) {
	// TODO - update GetPetsByTime with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ApiResponse{}) or use other options such as http.Ok ...
	// return Response(200, ApiResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPetsByTime method not implemented")
}

// GetPetsUsingBooleanQueryParameters - Get the pets by only using boolean query parameters
func (s *PetAPIService) GetPetsUsingBooleanQueryParameters(ctx context.Context, expr bool, grouping bool, inactive bool) (ImplResponse, error) {
	// TODO - update GetPetsUsingBooleanQueryParameters with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ApiResponse{}) or use other options such as http.Ok ...
	// return Response(200, ApiResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPetsUsingBooleanQueryParameters method not implemented")
}

// UpdatePet - Update an existing pet
func (s *PetAPIService) UpdatePet(ctx context.Context, pet Pet) (ImplResponse, error) {
	// TODO - update UpdatePet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	// return Response(200, Pet{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	// return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePet method not implemented")
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (s *PetAPIService) UpdatePetWithForm(ctx context.Context, petId int64, name string, status string) (ImplResponse, error) {
	// TODO - update UpdatePetWithForm with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	// return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePetWithForm method not implemented")
}

// UploadFile - uploads an image
func (s *PetAPIService) UploadFile(ctx context.Context, petId int64, additionalMetadata string, extraOptionalMetadata []string, file *os.File) (ImplResponse, error) {
	// TODO - update UploadFile with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ApiResponse{}) or use other options such as http.Ok ...
	// return Response(200, ApiResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UploadFile method not implemented")
}

// UploadFileArrayOfFiles - uploads images (array of files)
func (s *PetAPIService) UploadFileArrayOfFiles(ctx context.Context, petId int64, additionalMetadata string, files []*os.File) (ImplResponse, error) {
	// TODO - update UploadFileArrayOfFiles with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ApiResponse{}) or use other options such as http.Ok ...
	// return Response(200, ApiResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UploadFileArrayOfFiles method not implemented")
}
