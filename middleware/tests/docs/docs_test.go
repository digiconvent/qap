package docs_test

import (
	"flag"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
)

const root = "../../../docs"

const referencedIn = "\n---\n\n**Referenced in:**\n"

var update = flag.Bool("update", false, "rewrite the codes, link texts and Referenced in sections of every page")

var (
	linkPattern    = regexp.MustCompile(`\]\(([^)\s]+)\)`)
	headingPattern = regexp.MustCompile(`(?m)^#+ (.+)$`)
	slugDrop       = regexp.MustCompile(`[^a-z0-9 -]`)
)

type page struct {
	path  string
	title string
	body  string
}

type link struct {
	target string
	anchor string
}

func loadPages(t *testing.T) map[string]*page {
	t.Helper()
	pages := map[string]*page{}
	err := filepath.WalkDir(root, func(file string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative := filepath.ToSlash(strings.TrimPrefix(file, root+string(filepath.Separator)))
		if entry.IsDir() && relative == "development" {
			return filepath.SkipDir
		}
		if entry.IsDir() || !strings.HasSuffix(relative, ".md") {
			return nil
		}
		raw, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		text := string(raw)
		title, _, _ := strings.Cut(strings.TrimPrefix(text, "# "), "\n")
		body, _, _ := strings.Cut(text, referencedIn)
		pages[relative] = &page{path: relative, title: title, body: body}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return pages
}

func isListing(p *page) bool {
	return slices.Contains([]string{"index.md", "summary.md", "flow.md"}, path.Base(p.path))
}

func linksOf(p *page) []link {
	var links []link
	for _, match := range linkPattern.FindAllStringSubmatch(p.body, -1) {
		target, anchor, _ := strings.Cut(match[1], "#")
		if strings.Contains(target, "://") {
			continue
		}
		if target == "" {
			target = p.path
		} else {
			target = path.Clean(path.Join(path.Dir(p.path), target))
		}
		links = append(links, link{target: target, anchor: anchor})
	}
	return links
}

func slug(heading string) string {
	return strings.ReplaceAll(slugDrop.ReplaceAllString(strings.ToLower(heading), ""), " ", "-")
}

func anchorsOf(p *page) []string {
	var anchors []string
	for _, match := range headingPattern.FindAllStringSubmatch(p.body, -1) {
		anchors = append(anchors, slug(match[1]))
	}
	return anchors
}

func relativeLink(from string, to string) string {
	fromParts := strings.Split(path.Dir(from), "/")
	toParts := strings.Split(to, "/")
	if path.Dir(from) == "." {
		fromParts = nil
	}
	common := 0
	for common < len(fromParts) && common < len(toParts)-1 && fromParts[common] == toParts[common] {
		common++
	}
	return strings.Repeat("../", len(fromParts)-common) + strings.Join(toParts[common:], "/")
}

func sortedPaths(pages map[string]*page) []string {
	paths := make([]string, 0, len(pages))
	for p := range pages {
		paths = append(paths, p)
	}
	slices.Sort(paths)
	return paths
}

func referrersOf(pages map[string]*page) map[string][]*page {
	referrers := map[string][]*page{}
	for _, p := range sortedPaths(pages) {
		from := pages[p]
		if isListing(from) {
			continue
		}
		for _, l := range linksOf(from) {
			if l.target == from.path || slices.Contains(referrers[l.target], from) {
				continue
			}
			referrers[l.target] = append(referrers[l.target], from)
		}
	}
	for target := range referrers {
		slices.SortFunc(referrers[target], byCode)
	}
	return referrers
}

func referencedInSection(p *page, referrers []*page) string {
	if len(referrers) == 0 {
		return ""
	}
	var section strings.Builder
	section.WriteString(referencedIn + "\n")
	for _, r := range referrers {
		section.WriteString("- [" + r.title + "](" + relativeLink(p.path, r.path) + ")\n")
	}
	return section.String()
}

// every link points to an existing page, and every anchor to a heading on it
func TestDocs_LinksResolve(t *testing.T) {
	pages := loadPages(t)
	for _, p := range sortedPaths(pages) {
		for _, l := range linksOf(pages[p]) {
			target, ok := pages[l.target]
			if !ok {
				if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(l.target))); err != nil {
					t.Errorf("%s links to %s, which does not exist", p, l.target)
				}
				continue
			}
			if l.anchor != "" && !slices.Contains(anchorsOf(target), l.anchor) {
				t.Errorf("%s links to %s#%s, which has no such heading", p, l.target, l.anchor)
			}
		}
	}
}

// every page starts with its title and a description, and a page that is not a listing then gives its kind below a line
func TestDocs_PageShape(t *testing.T) {
	pages := loadPages(t)
	for _, p := range sortedPaths(pages) {
		text := pages[p].body
		heading, rest, _ := strings.Cut(text, "\n\n")
		if !strings.HasPrefix(heading, "# ") || strings.Contains(heading, "\n") {
			t.Errorf("%s does not start with a title line and a blank line", p)
			continue
		}
		if strings.HasPrefix(rest, "#") || strings.HasPrefix(rest, "---") || strings.HasPrefix(rest, "-") || strings.HasPrefix(rest, "|") || strings.HasPrefix(rest, "`") {
			t.Errorf("%s does not describe its name in a paragraph below the title", p)
		}
		if !isListing(pages[p]) && !strings.Contains(text, "\n---\n\n**Kind:** ") {
			t.Errorf("%s has no **Kind:** below a horizontal line", p)
		}
	}
}

// every summary lists exactly the pages of its folder
func TestDocs_SummariesListTheirFolder(t *testing.T) {
	pages := loadPages(t)
	for _, p := range sortedPaths(pages) {
		if path.Base(p) != "summary.md" {
			continue
		}
		var listed, folder []string
		for _, l := range linksOf(pages[p]) {
			listed = append(listed, l.target)
		}
		for _, other := range sortedPaths(pages) {
			if path.Dir(other) == path.Dir(p) && !isListing(pages[other]) {
				folder = append(folder, other)
			}
		}
		slices.Sort(listed)
		if !slices.Equal(listed, folder) {
			t.Errorf("%s lists %v, but its folder holds %v", p, listed, folder)
		}
	}
}

// the index of each tree links to every page in that tree
func TestDocs_IndexListsEveryPage(t *testing.T) {
	pages := loadPages(t)
	for _, index := range sortedPaths(pages) {
		if path.Base(index) != "index.md" || path.Dir(index) == "." {
			continue
		}
		var listed []string
		for _, l := range linksOf(pages[index]) {
			listed = append(listed, l.target)
		}
		for _, p := range sortedPaths(pages) {
			if p != index && strings.HasPrefix(p, path.Dir(index)+"/") && !slices.Contains(listed, p) {
				t.Errorf("%s does not link to %s", index, p)
			}
		}
	}
}

var layerLetters = map[string]string{
	"middleware/layer-0": "U", "middleware/layer-1": "E", "middleware/layer-2": "I", "middleware/layer-3": "B",
	"middleware/layer-4": "A", "middleware/concepts": "C", "middleware/host": "H",
	"client/layer-0": "cU", "client/layer-1": "cE", "client/layer-2": "cM", "client/layer-3": "cB",
	"client/layer-4": "cA", "client/states": "cS", "client/concepts": "cC", "client/app": "cH",
}

var codePattern = regexp.MustCompile(`^c?[A-Z]\d+ `)

var codeParts = regexp.MustCompile(`^(c?[A-Z])(\d+) `)

func byCode(a, b *page) int {
	partsA, partsB := codeParts.FindStringSubmatch(a.title), codeParts.FindStringSubmatch(b.title)
	if partsA == nil || partsB == nil || partsA[1] != partsB[1] {
		return strings.Compare(strings.ToLower(a.title), strings.ToLower(b.title))
	}
	numberA, _ := strconv.Atoi(partsA[2])
	numberB, _ := strconv.Atoi(partsB[2])
	return numberA - numberB
}

func codesOf(t *testing.T, pages map[string]*page) map[string]string {
	t.Helper()
	codes := map[string]string{}
	for folder, letter := range layerLetters {
		summary, ok := pages[folder+"/summary.md"]
		if !ok {
			t.Fatalf("%s has no summary", folder)
		}
		for i, l := range linksOf(summary) {
			codes[l.target] = letter + strconv.Itoa(i+1)
		}
	}
	return codes
}

func bareTitle(title string) string {
	return codePattern.ReplaceAllString(title, "")
}

func withLinkTexts(text string, textFor func(target string) (string, bool), from string) string {
	return linkTextPattern.ReplaceAllStringFunc(text, func(match string) string {
		parts := linkTextPattern.FindStringSubmatch(match)
		target, _, _ := strings.Cut(parts[2], "#")
		if want, ok := textFor(path.Clean(path.Join(path.Dir(from), target))); ok {
			return "[" + want + "](" + parts[2] + ")"
		}
		return match
	})
}

var linkTextPattern = regexp.MustCompile(`\[([^\]]*)\]\(([^)\s]+)\)`)

// every page carries its code in its title, and indexes, summaries and the flow name pages by their full title; run with -update to rewrite them
func TestDocs_Codes(t *testing.T) {
	pages := loadPages(t)
	codes := codesOf(t, pages)
	titles := map[string]string{}
	for p, code := range codes {
		titles[p] = code + " " + bareTitle(pages[p].title)
	}
	for _, p := range sortedPaths(pages) {
		file := filepath.Join(root, filepath.FromSlash(p))
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		want := string(raw)
		if title, coded := titles[p]; coded {
			_, rest, _ := strings.Cut(want, "\n")
			want = "# " + title + "\n" + rest
		} else if !isListing(pages[p]) {
			t.Errorf("%s is not listed in the summary of its folder, so it has no code", p)
			continue
		}
		if isListing(pages[p]) {
			want = withLinkTexts(want, func(target string) (string, bool) { title, ok := titles[target]; return title, ok }, p)
		}
		if want == string(raw) {
			continue
		}
		if *update {
			if err := os.WriteFile(file, []byte(want), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		t.Errorf("%s has an outdated code or link text; run go test ./middleware/tests/docs -update", p)
	}
}

// every page ends with the pages that link to it, sorted by title; run with -update to rewrite them
func TestDocs_ReferencedIn(t *testing.T) {
	pages := loadPages(t)
	referrers := referrersOf(pages)
	for _, p := range sortedPaths(pages) {
		file := filepath.Join(root, filepath.FromSlash(p))
		raw, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		want := strings.TrimRight(pages[p].body, "\n") + "\n" + referencedInSection(pages[p], referrers[p])
		if string(raw) == want {
			continue
		}
		if *update {
			if err := os.WriteFile(file, []byte(want), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		t.Errorf("%s has an outdated Referenced in section; run go test ./middleware/tests/docs -update", p)
	}
}
