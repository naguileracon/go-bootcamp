package importer

type ProductsImporter interface {
	Import(filePath string) (err error)
}
