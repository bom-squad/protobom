package writer

import (
	"fmt"
	"io"
	"os"

	"github.com/bom-squad/protobom/pkg/formats"
	"github.com/bom-squad/protobom/pkg/sbom"
	"github.com/bom-squad/protobom/pkg/writer/options"
	"github.com/sirupsen/logrus"
)

type writerImplementation interface {
	GetFormatSerializer(formats.Format) (Serializer, error)
	SerializeSBOM(options.Options, Serializer, *sbom.Document, io.WriteCloser) error
	OpenFile(string) (*os.File, error)
}

type defaultWriterImplementation struct{}

func (di *defaultWriterImplementation) GetFormatSerializer(formatOpt formats.Format) (Serializer, error) {
	switch formatOpt.Type() {
	case formats.CDXFORMAT:
		logrus.Infof("Serializing to %s", formatOpt)
		return &SerializerCDX14{}, nil
	case formats.SPDXFORMAT:
		logrus.Infof("Serializing to %s", formats.SPDX23JSON)
		if formatOpt.Encoding() == formats.JSON {
			return &SerializerSPDX23{}, nil
		}
	}

	return nil, fmt.Errorf("no serializer supports rendering to %s", formatOpt)
}

// SerializeSBOM takes an SBOM in protobuf and a serializer and uses it to render
// the document into the serializer format.
func (di *defaultWriterImplementation) SerializeSBOM(opts options.Options, serializer Serializer, bom *sbom.Document, wr io.WriteCloser) error {
	nativeDoc, err := serializer.Serialize(opts, bom)
	if err != nil {
		return fmt.Errorf("serializing SBOM to native format: %w", err)
	}
	if err := serializer.Render(opts, nativeDoc, wr); err != nil {
		return fmt.Errorf("writing rendered document to string: %w", err)
	}
	return nil
}

// OpenFile opens the file at path and returns it
func (di *defaultWriterImplementation) OpenFile(path string) (*os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	return f, nil
}
