SHELL:=/bin/bash

internal/model/mock/mock_product_repository.go:
	mockgen -destination=internal/model/mock/mock_product_repository.go -package=mock github.com/fajarachmadyusup13/product-management/internal/model ProductRepository


internal/model/mock/mock_product_usecase.go:
	mockgen -destination=internal/model/mock/mock_product_usecase.go -package=mock github.com/fajarachmadyusup13/product-management/internal/model ProductUsecase

mockgen: internal/model/mock/mock_product_repository.go \
	internal/model/mock/mock_product_usecase.go