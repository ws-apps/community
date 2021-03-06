// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

package doc

import (
	"strings"
	"time"

	"github.com/documize/community/model"
	"github.com/documize/community/model/workflow"
)

// Document represents the purpose of Documize.
type Document struct {
	model.BaseEntity
	OrgID        string              `json:"orgId"`
	LabelID      string              `json:"folderId"`
	UserID       string              `json:"userId"`
	Job          string              `json:"job"`
	Location     string              `json:"location"`
	Title        string              `json:"name"`
	Excerpt      string              `json:"excerpt"`
	Slug         string              `json:"-"`
	Tags         string              `json:"tags"`
	Template     bool                `json:"template"`
	Protection   workflow.Protection `json:"protection"`
	Approval     workflow.Approval   `json:"approval"`
	Lifecycle    workflow.Lifecycle  `json:"lifecycle"`
	Versioned    bool                `json:"versioned"`
	VersionID    string              `json:"versionId"`
	VersionOrder int                 `json:"versionOrder"`
	GroupID      string              `json:"groupId"`
}

// SetDefaults ensures on blanks and cleans.
func (d *Document) SetDefaults() {
	d.Title = strings.TrimSpace(d.Title)

	if len(d.Title) == 0 {
		d.Title = "Document"
	}
}

// ByTitle sorts a collection of documents by document title.
type ByTitle []Document

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return strings.ToLower(a[i].Title) < strings.ToLower(a[j].Title) }

// DocumentMeta details who viewed the document.
type DocumentMeta struct {
	Viewers []DocumentMetaViewer `json:"viewers"`
	Editors []DocumentMetaEditor `json:"editors"`
}

// DocumentMetaViewer contains the "view" metatdata content.
type DocumentMetaViewer struct {
	UserID    string    `json:"userId"`
	Created   time.Time `json:"created"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

// DocumentMetaEditor contains the "edit" metatdata content.
type DocumentMetaEditor struct {
	PageID    string    `json:"pageId"`
	UserID    string    `json:"userId"`
	Action    string    `json:"action"`
	Created   time.Time `json:"created"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
}

// UploadModel details the job ID of an uploaded document.
type UploadModel struct {
	JobID string `json:"jobId"`
}

// SitemapDocument details a document that can be exposed via Sitemap.
type SitemapDocument struct {
	DocumentID string
	Document   string
	FolderID   string
	Folder     string
	Revised    time.Time
}

// Version points to a version of a document.
type Version struct {
	VersionID  string `json:"versionId"`
	DocumentID string `json:"documentId"`
}
