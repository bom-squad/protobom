// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Document is the predicate function for document builders.
type Document func(*sql.Selector)

// DocumentType is the predicate function for documenttype builders.
type DocumentType func(*sql.Selector)

// Edge is the predicate function for edge builders.
type Edge func(*sql.Selector)

// ExternalReference is the predicate function for externalreference builders.
type ExternalReference func(*sql.Selector)

// HashesEntry is the predicate function for hashesentry builders.
type HashesEntry func(*sql.Selector)

// IdentifiersEntry is the predicate function for identifiersentry builders.
type IdentifiersEntry func(*sql.Selector)

// Metadata is the predicate function for metadata builders.
type Metadata func(*sql.Selector)

// Node is the predicate function for node builders.
type Node func(*sql.Selector)

// NodeList is the predicate function for nodelist builders.
type NodeList func(*sql.Selector)

// Person is the predicate function for person builders.
type Person func(*sql.Selector)

// Timestamp is the predicate function for timestamp builders.
type Timestamp func(*sql.Selector)

// Tool is the predicate function for tool builders.
type Tool func(*sql.Selector)
