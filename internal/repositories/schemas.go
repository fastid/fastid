package repositories

type KeysSchema struct {
	UnpackingKey []byte `json:"unpacking_key"`
	SignatureKey []byte `json:"signature_key"`
}
