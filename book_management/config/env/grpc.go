package env

import "strconv"

type GRPCConfig struct {
	BookServicePort     string
	AuthorServicePort   string
	CategoryServicePort string
}

func (c *GRPCConfig) Load() {
	c.BookServicePort = ":" + strconv.Itoa(getIntEnv("BOOK_PORT"))
	c.AuthorServicePort = ":" + strconv.Itoa(getIntEnv("AUTHOR_PORT"))
	c.CategoryServicePort = ":" + strconv.Itoa(getIntEnv("CATEGORY_PORT"))
}
