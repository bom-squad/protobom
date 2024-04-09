// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileName: ent/schema/metadata.go
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
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/bom-squad/protobom/pkg/sbom"
)

type Metadata struct {
	ent.Schema
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("version"),
		field.String("name"),
		field.Time("date"),
		field.String("comment"),
		field.JSON("tools", []*sbom.Tool{}),
		field.JSON("authors", []*sbom.Person{}),
		field.JSON("document_types", []*sbom.DocumentType{}),
	}
}

func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata_tools", Tool.Type),
		edge.To("metadata_authors", Person.Type),
		edge.To("metadata_document_types", DocumentType.Type),
		edge.To("document", Document.Type).Unique(),
	}
}

func (Metadata) Annotations() []schema.Annotation { return nil }
