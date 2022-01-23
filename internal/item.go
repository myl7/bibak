package internal

import (
	"encoding/json"
	"fmt"
	"github.com/mmcdole/gofeed"
	"github.com/myl7/bibak/internal/db"
	"github.com/myl7/bibak/internal/model"
	"os"
	"os/exec"
	"strings"
)

func handleItem(item *gofeed.Item) {
	f := model.FavItem{
		Title:       item.Title,
		Description: item.Description,
		AddFavTime:  *item.PublishedParsed,
		Link:        item.Link,
		Author:      item.Authors[0].Name,
		BID:         getBID(item.Link),
	}

	_, err := db.SelectFavItem(f.BID)
	if err == nil {
		return
	} else if err.Error() != "no result" {
		panic(err)
	}

	err = bbdown(&f)
	if err != nil {
		panic(err)
	}

	err = db.InsertFavItem(&f)
	if err != nil {
		panic(err)
	}
}

func getBID(link string) string {
	i := strings.LastIndex(link, "/")
	if i == -1 {
		panic(fmt.Errorf("invalid link: %s", link))
	}

	return link[i+1:]
}

func bbdown(f *model.FavItem) error {
	err := os.Mkdir(f.BID, 755)
	if err != nil {
		return err
	}

	c := exec.Command("BBDown", "-o", fmt.Sprintf("%s/%s.mp4", f.BID, f.BID), f.Link)
	err = c.Run()
	if err != nil {
		return err
	}

	b, err := json.Marshal(f)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("%s/meta.json", f.BID), b, 0644)
	if err != nil {
		return err
	}

	return nil
}
