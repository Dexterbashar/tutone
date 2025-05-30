package schema

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/violentbact/tutone/internal/config"
)

// Field is an attribute of a schema Type object.
type Field struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Kind        Kind   `json:"kind,omitempty"`

	Type         TypeRef     `json:"type"`
	Args         []Field     `json:"args,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
}

// GetDescription formats the description into a GoDoc comment.
func (f *Field) GetDescription() string {
	if strings.TrimSpace(f.Description) == "" {
		return ""
	}

	return formatDescription("", f.Description)
}

// GetTypeNameWithOverride returns the typeName, taking into consideration any FieldTypeOverride specified in the PackageConfig.
func (f *Field) GetTypeNameWithOverride(pkgConfig *config.PackageConfig) (string, error) {
	var typeName string
	var overrideType string
	var err error

	// Discover any FieldTypeOverride override for the current field.
	nameToMatch := f.Type.GetTypeName()

	for _, p := range pkgConfig.Types {
		if p.Name == nameToMatch {
			if p.FieldTypeOverride != "" {
				log.WithFields(log.Fields{
					"name":                nameToMatch,
					"field_type_override": p.FieldTypeOverride,
				}).Trace("overriding typeref")
				overrideType = p.FieldTypeOverride
			}
		}
	}

	// Set the typeName to the override or use what is specified in the schema.
	if overrideType != "" {
		typeName = overrideType
	} else {
		typeName, _, err = f.Type.GetType()
		if err != nil {
			return "", err
		}
	}

	return typeName, nil
}

// GetName returns a recusive lookup of the type name
func (f *Field) GetName() string {
	return formatGoName(f.Name)
}

func (f *Field) GetTagsWithOverrides(parentType Type, pkgConfig *config.PackageConfig) string {
	if f == nil {
		return ""
	}

	// Get the parent type config to apply any field struct tag overrides
	parentTypeConfig := pkgConfig.GetTypeConfigByName(parentType.Name)

	var tags string
	if parentTypeConfig != nil && parentTypeConfig.StructTags != nil {
		tags = f.buildStructTags(f.Name, *parentTypeConfig.StructTags)
	}

	if tags == "" {
		return f.GetTags()
	}

	return tags
}

func (f *Field) buildStructTags(fieldName string, structTags config.StructTags) string {
	tagsString := "`"

	tagsCount := len(structTags.Tags)
	if tagsCount == 0 && structTags.OmitEmpty == nil {
		return f.GetTags()
	}

	if tagsCount == 0 && structTags.OmitEmpty != nil {
		structTags.Tags = []string{"json"} // default is to include json struct tags
	}

	canIncludeOmitEmpty := true // default is to add `omitempty` to struct tags
	if structTags.OmitEmpty != nil {
		canIncludeOmitEmpty = *structTags.OmitEmpty
	}

	for i, tagType := range structTags.Tags {
		tagEnd := "\" "
		if i == tagsCount-1 {
			tagEnd = "\"" // no trailing space if last tag
		}

		tagsString = tagsString + tagType + ":\"" + f.Name

		if canIncludeOmitEmpty && (f.Type.IsInputObject() || !f.Type.IsNonNull()) {
			tagsString = tagsString + ",omitempty"
		}

		tagsString = tagsString + tagEnd
	}

	// Add closing back tick
	tagsString += "`"

	return tagsString
}

// GetTags is used to return the Go struct tags for a field.
func (f *Field) GetTags() string {
	if f == nil {
		return ""
	}

	jsonTag := "`json:\"" + f.Name

	if f.Type.IsInputObject() || !f.Type.IsNonNull() {
		jsonTag += ",omitempty"
	}

	tags := jsonTag + "\"`"

	return tags
}

func (f *Field) IsPrimitiveType() bool {
	goTypes := []string{
		"int",
		"string",
		"bool",
		"boolean",
	}

	name := strings.ToLower(f.Type.GetTypeName())

	for _, x := range goTypes {
		if x == name {
			return true
		}
	}

	return false
}

// Convenience method that proxies to TypeRef method
func (f *Field) IsScalarID() bool {
	return f.Type.IsScalarID()
}

// Convenience method that proxies to TypeRef method
func (f *Field) IsRequired() bool {
	return f.Type.IsNonNull()
}

func (f *Field) IsEnum() bool {
	if f.Type.Kind == KindENUM {
		return true
	}

	return f.Type.OfType != nil && f.Type.OfType.Kind == KindENUM
}

func (f *Field) HasRequiredArg() bool {
	for _, a := range f.Args {
		if a.IsRequired() {
			return true
		}
	}

	return false
}
