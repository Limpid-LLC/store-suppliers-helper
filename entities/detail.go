package entities

type Detail struct {
	ID               any                    `json:"id"`
	Title            string                 `json:"title"`
	Article          string                 `json:"article"`
	BrandName        string                 `json:"brand_name"`
	Prices           []DetailPrice          `json:"prices"`
	StoragesResidues []DetailStorageResidue `json:"storages_residues"`
}

type DetailPrice struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type DetailStorageResidue struct {
	StorageID   any    `json:"storage_id"`
	StorageName string `json:"storage_name"`
	Residue     any    `json:"residue"`
}
