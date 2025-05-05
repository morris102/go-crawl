package vocabulary

type VocalbularyRepository interface {
	CreateMany(items []Word) error
}

type VocalbularyRepositoryImpl struct{}

func NewVocalbularyRepository() (VocalbularyRepository, error) {
	return &VocalbularyRepositoryImpl{}, nil
}

func (inst *VocalbularyRepositoryImpl) CreateMany(items []Word) error {
	return nil
}
