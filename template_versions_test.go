package mailgun_test

import (
	"context"
	"testing"

	"github.com/mailgun/mailgun-go/v5"
	"github.com/mailgun/mailgun-go/v5/mtypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplateVersionsCRUD(t *testing.T) {
	mg := mailgun.NewMailgun(testKey)
	err := mg.SetAPIBase(server.URL())
	require.NoError(t, err)

	ctx := context.Background()

	findVersion := func(templateName, tag string) bool {
		it := mg.ListTemplateVersions(testDomain, templateName, nil)

		var page []mtypes.TemplateVersion
		for it.Next(ctx, &page) {
			for _, v := range page {
				if v.Tag == tag {
					return true
				}
			}
		}
		require.NoError(t, it.Err())
		return false
	}

	const (
		Comment        = "Mailgun-Go TestTemplateVersionsCRUD"
		UpdatedComment = "Mailgun-Go Test Version Updated"
		Template       = "{{.Name}}"
		Tag            = "v1"
	)

	tmpl := mtypes.Template{
		Name: randomString(10, "Mailgun-go-TestTemplateVersionsCRUD-"),
	}

	// Create a template
	require.NoError(t, mg.CreateTemplate(ctx, testDomain, &tmpl))

	version := mtypes.TemplateVersion{
		Tag:      Tag,
		Comment:  Comment,
		Template: Template,
		Active:   true,
		Engine:   mtypes.TemplateEngineGo,
	}

	// Add a version version
	require.NoError(t, mg.AddTemplateVersion(ctx, testDomain, tmpl.Name, &version))
	assert.Equal(t, Tag, version.Tag)
	assert.Equal(t, Comment, version.Comment)
	assert.Equal(t, mtypes.TemplateEngineGo, version.Engine)

	// Ensure the version is in the list
	require.True(t, findVersion(tmpl.Name, version.Tag))

	// Update the Comment
	version.Comment = UpdatedComment
	version.Template = Template + "updated"
	require.NoError(t, mg.UpdateTemplateVersion(ctx, testDomain, tmpl.Name, &version))

	// Ensure update took
	updated, err := mg.GetTemplateVersion(ctx, testDomain, tmpl.Name, version.Tag)

	require.NoError(t, err)
	assert.Equal(t, UpdatedComment, updated.Comment)
	assert.Equal(t, Template+"updated", updated.Template)

	// Add a new active Version
	version2 := mtypes.TemplateVersion{
		Tag:      "v2",
		Comment:  Comment,
		Template: Template,
		Active:   true,
		Engine:   mtypes.TemplateEngineGo,
	}
	require.NoError(t, mg.AddTemplateVersion(ctx, testDomain, tmpl.Name, &version2))

	// Ensure the version is in the list
	require.True(t, findVersion(tmpl.Name, version2.Tag))

	// Delete the first version
	require.NoError(t, mg.DeleteTemplateVersion(ctx, testDomain, tmpl.Name, version.Tag))

	// Ensure version was deleted
	require.False(t, findVersion(tmpl.Name, version.Tag))

	// Delete the template
	require.NoError(t, mg.DeleteTemplate(ctx, testDomain, tmpl.Name))
}
