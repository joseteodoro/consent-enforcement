package consent_manager

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateDataType(ctx context.Context, display string) (*DataType, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateGroup(ctx context.Context, display string) (*Group, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePatient(ctx context.Context, display string, groups []string, dataTypes []string) (*Patient, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePractioner(ctx context.Context, display string, groups []string) (*Practioner, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateConsentEnforcementRule(ctx context.Context, display string, readerGroups []string, targetGroups []string, expiration int) (*ConsentEnforcement, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateContract(ctx context.Context, display string, clauses []string) (*Contract, error) {
	panic("not implemented")
}
func (r *mutationResolver) GiveConsent(ctx context.Context, patientID string, contractID string, expiration int) (*Signature, error) {
	panic("not implemented")
}
func (r *mutationResolver) PatientResetPassphase(ctx context.Context, patientID string, restorePhrase string) (*Patient, error) {
	panic("not implemented")
}
func (r *mutationResolver) PractionerResetPassphase(ctx context.Context, practionerID string, restorePhrase string) (*Practioner, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) DataTypes(ctx context.Context, limit *int, offset *int) ([]*DataType, error) {
	panic("not implemented")
}
func (r *queryResolver) PublicGroups(ctx context.Context, display *string, limit *int, offset *int) ([]*Group, error) {
	panic("not implemented")
}
func (r *queryResolver) Consent(ctx context.Context, patientID string, practionerID string, dataTypeID *string) (*AccessToken, error) {
	panic("not implemented")
}
func (r *queryResolver) ConsentEnforcement(ctx context.Context, patientID string, practionerID string, reason string) (*AccessToken, error) {
	panic("not implemented")
}

// NewRootResolvers works a root resolver for other ones
func NewRootResolvers() Config {
	c := Config{Resolvers: &Resolver{}}
	return c
}
