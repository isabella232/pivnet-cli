// This file was generated by counterfeiter
package productfilefakes

import (
	"io"
	"os"
	"sync"

	"github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands/productfile"
)

type FakePivnetClient struct {
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	ProductFilesStub        func(productSlug string) ([]pivnet.ProductFile, error)
	productFilesMutex       sync.RWMutex
	productFilesArgsForCall []struct {
		productSlug string
	}
	productFilesReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFilesReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	ProductFilesForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.ProductFile, error)
	productFilesForReleaseMutex       sync.RWMutex
	productFilesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	productFilesForReleaseReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFilesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	ProductFileStub        func(productSlug string, productFileID int) (pivnet.ProductFile, error)
	productFileMutex       sync.RWMutex
	productFileArgsForCall []struct {
		productSlug   string
		productFileID int
	}
	productFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	productFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	ProductFileForReleaseStub        func(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error)
	productFileForReleaseMutex       sync.RWMutex
	productFileForReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	productFileForReleaseReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	productFileForReleaseReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	CreateProductFileStub        func(config pivnet.CreateProductFileConfig) (pivnet.ProductFile, error)
	createProductFileMutex       sync.RWMutex
	createProductFileArgsForCall []struct {
		config pivnet.CreateProductFileConfig
	}
	createProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	createProductFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	UpdateProductFileStub        func(productSlug string, productFile pivnet.ProductFile) (pivnet.ProductFile, error)
	updateProductFileMutex       sync.RWMutex
	updateProductFileArgsForCall []struct {
		productSlug string
		productFile pivnet.ProductFile
	}
	updateProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	updateProductFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AddProductFileToReleaseStub        func(productSlug string, releaseID int, productFileID int) error
	addProductFileToReleaseMutex       sync.RWMutex
	addProductFileToReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	addProductFileToReleaseReturns struct {
		result1 error
	}
	addProductFileToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveProductFileFromReleaseStub        func(productSlug string, releaseID int, productFileID int) error
	removeProductFileFromReleaseMutex       sync.RWMutex
	removeProductFileFromReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	removeProductFileFromReleaseReturns struct {
		result1 error
	}
	removeProductFileFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	AddProductFileToFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	addProductFileToFileGroupMutex       sync.RWMutex
	addProductFileToFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	addProductFileToFileGroupReturns struct {
		result1 error
	}
	addProductFileToFileGroupReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveProductFileFromFileGroupStub        func(productSlug string, fileGroupID int, productFileID int) error
	removeProductFileFromFileGroupMutex       sync.RWMutex
	removeProductFileFromFileGroupArgsForCall []struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}
	removeProductFileFromFileGroupReturns struct {
		result1 error
	}
	removeProductFileFromFileGroupReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteProductFileStub        func(productSlug string, releaseID int) (pivnet.ProductFile, error)
	deleteProductFileMutex       sync.RWMutex
	deleteProductFileArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	deleteProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	deleteProductFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	acceptEULAReturnsOnCall map[int]struct {
		result1 error
	}
	DownloadProductFileStub        func(location *os.File, productSlug string, releaseID int, productFileID int, progressWriter io.Writer) error
	downloadProductFileMutex       sync.RWMutex
	downloadProductFileArgsForCall []struct {
		location       *os.File
		productSlug    string
		releaseID      int
		productFileID  int
		progressWriter io.Writer
	}
	downloadProductFileReturns struct {
		result1 error
	}
	downloadProductFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFiles(productSlug string) ([]pivnet.ProductFile, error) {
	fake.productFilesMutex.Lock()
	ret, specificReturn := fake.productFilesReturnsOnCall[len(fake.productFilesArgsForCall)]
	fake.productFilesArgsForCall = append(fake.productFilesArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("ProductFiles", []interface{}{productSlug})
	fake.productFilesMutex.Unlock()
	if fake.ProductFilesStub != nil {
		return fake.ProductFilesStub(productSlug)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFilesReturns.result1, fake.productFilesReturns.result2
}

func (fake *FakePivnetClient) ProductFilesCallCount() int {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	return len(fake.productFilesArgsForCall)
}

func (fake *FakePivnetClient) ProductFilesArgsForCall(i int) string {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	return fake.productFilesArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) ProductFilesReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesStub = nil
	fake.productFilesReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesStub = nil
	if fake.productFilesReturnsOnCall == nil {
		fake.productFilesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFilesReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesForRelease(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	fake.productFilesForReleaseMutex.Lock()
	ret, specificReturn := fake.productFilesForReleaseReturnsOnCall[len(fake.productFilesForReleaseArgsForCall)]
	fake.productFilesForReleaseArgsForCall = append(fake.productFilesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ProductFilesForRelease", []interface{}{productSlug, releaseID})
	fake.productFilesForReleaseMutex.Unlock()
	if fake.ProductFilesForReleaseStub != nil {
		return fake.ProductFilesForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFilesForReleaseReturns.result1, fake.productFilesForReleaseReturns.result2
}

func (fake *FakePivnetClient) ProductFilesForReleaseCallCount() int {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return len(fake.productFilesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFilesForReleaseArgsForCall(i int) (string, int) {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return fake.productFilesForReleaseArgsForCall[i].productSlug, fake.productFilesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ProductFilesForReleaseReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesForReleaseStub = nil
	fake.productFilesForReleaseReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesForReleaseReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesForReleaseStub = nil
	if fake.productFilesForReleaseReturnsOnCall == nil {
		fake.productFilesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFilesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFile(productSlug string, productFileID int) (pivnet.ProductFile, error) {
	fake.productFileMutex.Lock()
	ret, specificReturn := fake.productFileReturnsOnCall[len(fake.productFileArgsForCall)]
	fake.productFileArgsForCall = append(fake.productFileArgsForCall, struct {
		productSlug   string
		productFileID int
	}{productSlug, productFileID})
	fake.recordInvocation("ProductFile", []interface{}{productSlug, productFileID})
	fake.productFileMutex.Unlock()
	if fake.ProductFileStub != nil {
		return fake.ProductFileStub(productSlug, productFileID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFileReturns.result1, fake.productFileReturns.result2
}

func (fake *FakePivnetClient) ProductFileCallCount() int {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	return len(fake.productFileArgsForCall)
}

func (fake *FakePivnetClient) ProductFileArgsForCall(i int) (string, int) {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	return fake.productFileArgsForCall[i].productSlug, fake.productFileArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) ProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileStub = nil
	fake.productFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileStub = nil
	if fake.productFileReturnsOnCall == nil {
		fake.productFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.productFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileForRelease(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	fake.productFileForReleaseMutex.Lock()
	ret, specificReturn := fake.productFileForReleaseReturnsOnCall[len(fake.productFileForReleaseArgsForCall)]
	fake.productFileForReleaseArgsForCall = append(fake.productFileForReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("ProductFileForRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.productFileForReleaseMutex.Unlock()
	if fake.ProductFileForReleaseStub != nil {
		return fake.ProductFileForReleaseStub(productSlug, releaseID, productFileID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFileForReleaseReturns.result1, fake.productFileForReleaseReturns.result2
}

func (fake *FakePivnetClient) ProductFileForReleaseCallCount() int {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return len(fake.productFileForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFileForReleaseArgsForCall(i int) (string, int, int) {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return fake.productFileForReleaseArgsForCall[i].productSlug, fake.productFileForReleaseArgsForCall[i].releaseID, fake.productFileForReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) ProductFileForReleaseReturns(result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileForReleaseStub = nil
	fake.productFileForReleaseReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileForReleaseReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileForReleaseStub = nil
	if fake.productFileForReleaseReturnsOnCall == nil {
		fake.productFileForReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.productFileForReleaseReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateProductFile(config pivnet.CreateProductFileConfig) (pivnet.ProductFile, error) {
	fake.createProductFileMutex.Lock()
	ret, specificReturn := fake.createProductFileReturnsOnCall[len(fake.createProductFileArgsForCall)]
	fake.createProductFileArgsForCall = append(fake.createProductFileArgsForCall, struct {
		config pivnet.CreateProductFileConfig
	}{config})
	fake.recordInvocation("CreateProductFile", []interface{}{config})
	fake.createProductFileMutex.Unlock()
	if fake.CreateProductFileStub != nil {
		return fake.CreateProductFileStub(config)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createProductFileReturns.result1, fake.createProductFileReturns.result2
}

func (fake *FakePivnetClient) CreateProductFileCallCount() int {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return len(fake.createProductFileArgsForCall)
}

func (fake *FakePivnetClient) CreateProductFileArgsForCall(i int) pivnet.CreateProductFileConfig {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return fake.createProductFileArgsForCall[i].config
}

func (fake *FakePivnetClient) CreateProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.CreateProductFileStub = nil
	fake.createProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) CreateProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.CreateProductFileStub = nil
	if fake.createProductFileReturnsOnCall == nil {
		fake.createProductFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.createProductFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateProductFile(productSlug string, productFile pivnet.ProductFile) (pivnet.ProductFile, error) {
	fake.updateProductFileMutex.Lock()
	ret, specificReturn := fake.updateProductFileReturnsOnCall[len(fake.updateProductFileArgsForCall)]
	fake.updateProductFileArgsForCall = append(fake.updateProductFileArgsForCall, struct {
		productSlug string
		productFile pivnet.ProductFile
	}{productSlug, productFile})
	fake.recordInvocation("UpdateProductFile", []interface{}{productSlug, productFile})
	fake.updateProductFileMutex.Unlock()
	if fake.UpdateProductFileStub != nil {
		return fake.UpdateProductFileStub(productSlug, productFile)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateProductFileReturns.result1, fake.updateProductFileReturns.result2
}

func (fake *FakePivnetClient) UpdateProductFileCallCount() int {
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	return len(fake.updateProductFileArgsForCall)
}

func (fake *FakePivnetClient) UpdateProductFileArgsForCall(i int) (string, pivnet.ProductFile) {
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	return fake.updateProductFileArgsForCall[i].productSlug, fake.updateProductFileArgsForCall[i].productFile
}

func (fake *FakePivnetClient) UpdateProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.UpdateProductFileStub = nil
	fake.updateProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpdateProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.UpdateProductFileStub = nil
	if fake.updateProductFileReturnsOnCall == nil {
		fake.updateProductFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.updateProductFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddProductFileToRelease(productSlug string, releaseID int, productFileID int) error {
	fake.addProductFileToReleaseMutex.Lock()
	ret, specificReturn := fake.addProductFileToReleaseReturnsOnCall[len(fake.addProductFileToReleaseArgsForCall)]
	fake.addProductFileToReleaseArgsForCall = append(fake.addProductFileToReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("AddProductFileToRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.addProductFileToReleaseMutex.Unlock()
	if fake.AddProductFileToReleaseStub != nil {
		return fake.AddProductFileToReleaseStub(productSlug, releaseID, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addProductFileToReleaseReturns.result1
}

func (fake *FakePivnetClient) AddProductFileToReleaseCallCount() int {
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	return len(fake.addProductFileToReleaseArgsForCall)
}

func (fake *FakePivnetClient) AddProductFileToReleaseArgsForCall(i int) (string, int, int) {
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	return fake.addProductFileToReleaseArgsForCall[i].productSlug, fake.addProductFileToReleaseArgsForCall[i].releaseID, fake.addProductFileToReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) AddProductFileToReleaseReturns(result1 error) {
	fake.AddProductFileToReleaseStub = nil
	fake.addProductFileToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddProductFileToReleaseReturnsOnCall(i int, result1 error) {
	fake.AddProductFileToReleaseStub = nil
	if fake.addProductFileToReleaseReturnsOnCall == nil {
		fake.addProductFileToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addProductFileToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromRelease(productSlug string, releaseID int, productFileID int) error {
	fake.removeProductFileFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeProductFileFromReleaseReturnsOnCall[len(fake.removeProductFileFromReleaseArgsForCall)]
	fake.removeProductFileFromReleaseArgsForCall = append(fake.removeProductFileFromReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("RemoveProductFileFromRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.removeProductFileFromReleaseMutex.Unlock()
	if fake.RemoveProductFileFromReleaseStub != nil {
		return fake.RemoveProductFileFromReleaseStub(productSlug, releaseID, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeProductFileFromReleaseReturns.result1
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseCallCount() int {
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	return len(fake.removeProductFileFromReleaseArgsForCall)
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseArgsForCall(i int) (string, int, int) {
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	return fake.removeProductFileFromReleaseArgsForCall[i].productSlug, fake.removeProductFileFromReleaseArgsForCall[i].releaseID, fake.removeProductFileFromReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseReturns(result1 error) {
	fake.RemoveProductFileFromReleaseStub = nil
	fake.removeProductFileFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromReleaseReturnsOnCall(i int, result1 error) {
	fake.RemoveProductFileFromReleaseStub = nil
	if fake.removeProductFileFromReleaseReturnsOnCall == nil {
		fake.removeProductFileFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeProductFileFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddProductFileToFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.addProductFileToFileGroupMutex.Lock()
	ret, specificReturn := fake.addProductFileToFileGroupReturnsOnCall[len(fake.addProductFileToFileGroupArgsForCall)]
	fake.addProductFileToFileGroupArgsForCall = append(fake.addProductFileToFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("AddProductFileToFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.addProductFileToFileGroupMutex.Unlock()
	if fake.AddProductFileToFileGroupStub != nil {
		return fake.AddProductFileToFileGroupStub(productSlug, fileGroupID, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addProductFileToFileGroupReturns.result1
}

func (fake *FakePivnetClient) AddProductFileToFileGroupCallCount() int {
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	return len(fake.addProductFileToFileGroupArgsForCall)
}

func (fake *FakePivnetClient) AddProductFileToFileGroupArgsForCall(i int) (string, int, int) {
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	return fake.addProductFileToFileGroupArgsForCall[i].productSlug, fake.addProductFileToFileGroupArgsForCall[i].fileGroupID, fake.addProductFileToFileGroupArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) AddProductFileToFileGroupReturns(result1 error) {
	fake.AddProductFileToFileGroupStub = nil
	fake.addProductFileToFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AddProductFileToFileGroupReturnsOnCall(i int, result1 error) {
	fake.AddProductFileToFileGroupStub = nil
	if fake.addProductFileToFileGroupReturnsOnCall == nil {
		fake.addProductFileToFileGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addProductFileToFileGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroup(productSlug string, fileGroupID int, productFileID int) error {
	fake.removeProductFileFromFileGroupMutex.Lock()
	ret, specificReturn := fake.removeProductFileFromFileGroupReturnsOnCall[len(fake.removeProductFileFromFileGroupArgsForCall)]
	fake.removeProductFileFromFileGroupArgsForCall = append(fake.removeProductFileFromFileGroupArgsForCall, struct {
		productSlug   string
		fileGroupID   int
		productFileID int
	}{productSlug, fileGroupID, productFileID})
	fake.recordInvocation("RemoveProductFileFromFileGroup", []interface{}{productSlug, fileGroupID, productFileID})
	fake.removeProductFileFromFileGroupMutex.Unlock()
	if fake.RemoveProductFileFromFileGroupStub != nil {
		return fake.RemoveProductFileFromFileGroupStub(productSlug, fileGroupID, productFileID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeProductFileFromFileGroupReturns.result1
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupCallCount() int {
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	return len(fake.removeProductFileFromFileGroupArgsForCall)
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupArgsForCall(i int) (string, int, int) {
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	return fake.removeProductFileFromFileGroupArgsForCall[i].productSlug, fake.removeProductFileFromFileGroupArgsForCall[i].fileGroupID, fake.removeProductFileFromFileGroupArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupReturns(result1 error) {
	fake.RemoveProductFileFromFileGroupStub = nil
	fake.removeProductFileFromFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveProductFileFromFileGroupReturnsOnCall(i int, result1 error) {
	fake.RemoveProductFileFromFileGroupStub = nil
	if fake.removeProductFileFromFileGroupReturnsOnCall == nil {
		fake.removeProductFileFromFileGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeProductFileFromFileGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DeleteProductFile(productSlug string, releaseID int) (pivnet.ProductFile, error) {
	fake.deleteProductFileMutex.Lock()
	ret, specificReturn := fake.deleteProductFileReturnsOnCall[len(fake.deleteProductFileArgsForCall)]
	fake.deleteProductFileArgsForCall = append(fake.deleteProductFileArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("DeleteProductFile", []interface{}{productSlug, releaseID})
	fake.deleteProductFileMutex.Unlock()
	if fake.DeleteProductFileStub != nil {
		return fake.DeleteProductFileStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteProductFileReturns.result1, fake.deleteProductFileReturns.result2
}

func (fake *FakePivnetClient) DeleteProductFileCallCount() int {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return len(fake.deleteProductFileArgsForCall)
}

func (fake *FakePivnetClient) DeleteProductFileArgsForCall(i int) (string, int) {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return fake.deleteProductFileArgsForCall[i].productSlug, fake.deleteProductFileArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) DeleteProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.DeleteProductFileStub = nil
	fake.deleteProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DeleteProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.DeleteProductFileStub = nil
	if fake.deleteProductFileReturnsOnCall == nil {
		fake.deleteProductFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.deleteProductFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	ret, specificReturn := fake.acceptEULAReturnsOnCall[len(fake.acceptEULAArgsForCall)]
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.acceptEULAReturns.result1
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AcceptEULAReturnsOnCall(i int, result1 error) {
	fake.AcceptEULAStub = nil
	if fake.acceptEULAReturnsOnCall == nil {
		fake.acceptEULAReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acceptEULAReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DownloadProductFile(location *os.File, productSlug string, releaseID int, productFileID int, progressWriter io.Writer) error {
	fake.downloadProductFileMutex.Lock()
	ret, specificReturn := fake.downloadProductFileReturnsOnCall[len(fake.downloadProductFileArgsForCall)]
	fake.downloadProductFileArgsForCall = append(fake.downloadProductFileArgsForCall, struct {
		location       *os.File
		productSlug    string
		releaseID      int
		productFileID  int
		progressWriter io.Writer
	}{location, productSlug, releaseID, productFileID, progressWriter})
	fake.recordInvocation("DownloadProductFile", []interface{}{location, productSlug, releaseID, productFileID, progressWriter})
	fake.downloadProductFileMutex.Unlock()
	if fake.DownloadProductFileStub != nil {
		return fake.DownloadProductFileStub(location, productSlug, releaseID, productFileID, progressWriter)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downloadProductFileReturns.result1
}

func (fake *FakePivnetClient) DownloadProductFileCallCount() int {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return len(fake.downloadProductFileArgsForCall)
}

func (fake *FakePivnetClient) DownloadProductFileArgsForCall(i int) (*os.File, string, int, int, io.Writer) {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return fake.downloadProductFileArgsForCall[i].location, fake.downloadProductFileArgsForCall[i].productSlug, fake.downloadProductFileArgsForCall[i].releaseID, fake.downloadProductFileArgsForCall[i].productFileID, fake.downloadProductFileArgsForCall[i].progressWriter
}

func (fake *FakePivnetClient) DownloadProductFileReturns(result1 error) {
	fake.DownloadProductFileStub = nil
	fake.downloadProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) DownloadProductFileReturnsOnCall(i int, result1 error) {
	fake.DownloadProductFileStub = nil
	if fake.downloadProductFileReturnsOnCall == nil {
		fake.downloadProductFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadProductFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	fake.updateProductFileMutex.RLock()
	defer fake.updateProductFileMutex.RUnlock()
	fake.addProductFileToReleaseMutex.RLock()
	defer fake.addProductFileToReleaseMutex.RUnlock()
	fake.removeProductFileFromReleaseMutex.RLock()
	defer fake.removeProductFileFromReleaseMutex.RUnlock()
	fake.addProductFileToFileGroupMutex.RLock()
	defer fake.addProductFileToFileGroupMutex.RUnlock()
	fake.removeProductFileFromFileGroupMutex.RLock()
	defer fake.removeProductFileFromFileGroupMutex.RUnlock()
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ productfile.PivnetClient = new(FakePivnetClient)
