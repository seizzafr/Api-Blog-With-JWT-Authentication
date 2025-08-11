package utils

import (
    "fmt"
    "gorm.io/gorm"
    "strings"
    "unicode"
)

// Slugify mengubah string biasa jadi slug URL-friendly (huruf kecil, strip)
func Slugify(s string) string {
    var slug strings.Builder
    for _, r := range s {
        switch {
        case unicode.IsLetter(r) || unicode.IsDigit(r):
            slug.WriteRune(unicode.ToLower(r))
        case unicode.IsSpace(r) || r == '-' || r == '_':
            slug.WriteRune('-')
        }
    }

    return strings.Trim(slug.String(), "-")
}

func GenerateUniqueSlug(db *gorm.DB, base string, model interface{}, field string) string {
    slug := Slugify(base)
    uniqueSlug := slug
    counter := 1

    for {
        var count int64
        query := fmt.Sprintf("%s = ?", field)
        db.Model(model).Where(query, uniqueSlug).Count(&count)

        if count == 0 {
            break
        }

        uniqueSlug = fmt.Sprintf("%s-%d", slug, counter)
        counter++
    }

    return uniqueSlug
}
