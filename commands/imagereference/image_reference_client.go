package imagereference

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/go-pivnet/v2/logger"
	"github.com/pivotal-cf/pivnet-cli/errorhandler"
	"github.com/pivotal-cf/pivnet-cli/printer"
)

//go:generate counterfeiter . PivnetClient
type PivnetClient interface {
	ReleaseForVersion(productSlug string, releaseVersion string) (pivnet.Release, error)
	ImageReferences(productSlug string) ([]pivnet.ImageReference, error)
	ImageReferencesForRelease(productSlug string, releaseID int) ([]pivnet.ImageReference, error)
	CreateImageReference(config pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error)
}

type ImageReferenceClient struct {
	pivnetClient     PivnetClient
	eh               errorhandler.ErrorHandler
	format           string
	outputWriter     io.Writer
	logWriter        io.Writer
	printer          printer.Printer
	l                logger.Logger
}

func NewImageReferenceClient(
	pivnetClient PivnetClient,
	eh errorhandler.ErrorHandler,
	format string,
	outputWriter io.Writer,
	logWriter io.Writer,
	printer printer.Printer,
	l logger.Logger,
) *ImageReferenceClient {
	return &ImageReferenceClient{
		pivnetClient:     pivnetClient,
		eh:               eh,
		format:           format,
		outputWriter:     outputWriter,
		logWriter:        logWriter,
		printer:          printer,
		l:                l,
	}
}

func (c *ImageReferenceClient) List(productSlug string, releaseVersion string) error {
	if releaseVersion == "" {
		imageReferences, err := c.pivnetClient.ImageReferences(productSlug)
		if err != nil {
			return c.eh.HandleError(err)
		}

		return c.printImageReferences(imageReferences)
	}

	release, err := c.pivnetClient.ReleaseForVersion(productSlug, releaseVersion)
	if err != nil {
		return c.eh.HandleError(err)
	}

	imageReferences, err := c.pivnetClient.ImageReferencesForRelease(
		productSlug,
		release.ID,
	)
	if err != nil {
		return c.eh.HandleError(err)
	}

	return c.printImageReferences(imageReferences)
}

func (c *ImageReferenceClient) printImageReferences(imageReferences []pivnet.ImageReference) error {
	switch c.format {

	case printer.PrintAsTable:
		table := tablewriter.NewWriter(c.outputWriter)
		table.SetHeader([]string{
			"ID",
			"Name",
			"Image Path",
			"Digest",
		})

		for _, imageReference := range imageReferences {
			imageReferenceAsString := []string{
				strconv.Itoa(imageReference.ID),
				imageReference.Name,
				imageReference.ImagePath,
				imageReference.Digest,
			}
			table.Append(imageReferenceAsString)
		}
		table.Render()
		return nil
	case printer.PrintAsJSON:
		return c.printer.PrintJSON(imageReferences)
	case printer.PrintAsYAML:
		return c.printer.PrintYAML(imageReferences)
	}

	return nil
}

func (c *ImageReferenceClient) printImageReference(imageReference pivnet.ImageReference) error {
	switch c.format {
	case printer.PrintAsTable:
		table := tablewriter.NewWriter(c.outputWriter)
		table.SetHeader([]string{
			"ID",
			"Name",
			"Image Path",
			"Digest",
		})

		imageReferenceAsString := []string{
			strconv.Itoa(imageReference.ID),
			imageReference.Name,
			imageReference.ImagePath,
			imageReference.Digest,
		}
		table.Append(imageReferenceAsString)
		table.Render()
		return nil
	case printer.PrintAsJSON:
		return c.printer.PrintJSON(imageReference)
	case printer.PrintAsYAML:
		return c.printer.PrintYAML(imageReference)
	}

	return nil
}

func (c *ImageReferenceClient) Create(config pivnet.CreateImageReferenceConfig) error {
	imageReference, err := c.pivnetClient.CreateImageReference(config)
	if err != nil {
		return c.eh.HandleError(err)
	}

	return c.printImageReference(imageReference)
}
