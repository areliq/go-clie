package cmd

import (
	"context"
	"log"

	c "clie/pkg/crawler"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(crawlCmd)
}

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Get stories",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.TODO()
		ids, err := c.GetTopStories(ctx)

		if err != nil {
			log.Fatalln(err)
		}

		for _, id := range ids[:3] {
			story, err := c.GetItem(ctx, id)

			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("%#v", story)
		}

		// log.Printf("ids: %#v", ids[:10])
	},
}
