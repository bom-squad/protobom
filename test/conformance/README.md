# Protobom Conformance Testing

This directory contains the protobom conformance test framework. It is still
under construction but it already runs along the other project tests.

## Conformance Testing Overview

The main goal of the protobom conformance testing framework is to ensure that
each protobom serialization or unserialization result in equivalent documents
every time.

SBOM diffing is a complex problem. SBOMs cannot be compared as simple strings or
comparing data structs. Verifying that documents are equivalent needs more precise,
semantic comparison.

The approach the protrobom conformance tests take is to ask the same group of
questions to a generated and to its equivalent golden sample and contrast the
answers. The more questions we ask, the stronger asurance the project can give
of consistent parsing and serialization.

Some sample questions that can be asked of documents:

- Are the number of nodes the same?
- Are the NodeLists of both documents `Equal()`?
- Do the hashes of the files match?
- Do both documents have the same structure?
- Does the generated document contain the same identifiers?

Adding more tests that ask questions like these and compare the answers can be
done at any time and it can be a good opportunity for new contributors. Please
check the open issues, there may be some ideas already there.

## Running the Conformance Tests

Conformance tests will run along all go test runs in the project presubmits.
To manually run the conformance tests just run:

```
make conformance-test
```

## Test Data Structure

The test fixtures for the conformance suite are stored in directories broken
into the following pattern: format → version → encoding. Here is how the spdx
2.3 json samples are stored:

```
test/conformance/
└── testdata
    └── spdx
        └── 2.3
            └── json
                ├── curl.spdx.json
                └── curl.spdx.json.proto


```

The conformance suite will descend into all known format directories and use the
sample SBOMs and autogenerated protobuf blobs to run the conformance tests. There
can be more than one sbom/proto sample per directory, in fact it is encouraged
to add more data.

### Generating the protrobuf Blobs

When there is a structural change in protobom that results in expected changes to
the parsed documents or serialized SBOMs, the protobuf blobs need to be regenerated.
The conformance directory contains a generator utility in test/conformance/generator.

Run the utility by running `make conformance`, it will cycle through all the
supported formats and regenerate all protobuf blobs from sample SBOMs it finds.

```
make conformance

INFO[0000] Wrote text/spdx+json;version=2.3 sample to test/conformance/testdata/spdx/2.3/json/curl.spdx.json.proto
```

Please note that rebuilds of the blobs are not expected to be reproducible so
all executions of the generator will result in a sizeable diff.
