package queries

import (
	"log"
)

// ReadPost is function to read single post
func (c *Client) ReadPost(id int64) (post Post, err error) {
	// run the query with given id
	row := c.DB.QueryRow(`
			SELECT id, content, create_time
			FROM post
			WHERE id = ?`,
		id)

	// assign the query result to 'post' object
	err = row.Scan(&post.ID,
		&post.Content,
		&post.CreateTime)

	if err != nil {
		log.Println(err)
		return
	}

	return
}
