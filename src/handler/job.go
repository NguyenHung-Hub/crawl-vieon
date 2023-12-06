package handler

import (
	"context"
	"crawl/src/db"
	"crawl/src/models"
	"crawl/src/util"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

func MainJob(s *Server) {

	exclude_ribbon_id := []string{
		"WATCHING_RIBBON",
		"data-rm1st-becauseyouwatched-DTT",
		"data-recommendforyou-DTT",
		"data-cache-VieONselection",
		"data-suggestforyou-DTT",
		"data-rm1st-becauseyouwatched-DTT",
		"data-recommendforyou-DTT",
		"data-cache-VieONselection",
		"data-suggestforyou-DTT",
	}

	//step 1
	urls := step1_get_url()
	err := util.WriteSliceStrToCSV("data/", "url", urls, true)
	if err != nil {
		log.Printf("Write file error: %s", err)
	}

	//step 2
	ribbonIds := step2_get_id_ribbon(urls)
	ribbonId_set := util.RemoveDuplicates(ribbonIds)
	ribbonId_filter := util.FilterStr(ribbonId_set, exclude_ribbon_id)

	err = util.WriteSliceStrToCSV("data/", "ribbon_id", ribbonId_filter, true)
	if err != nil {
		log.Printf("Write file error: %s", err)
	}

	rc := step3_get_id_content(ribbonId_filter, s.store)
	var contentIds []string
	for _, v := range rc {
		contentIds = append(contentIds, v.ContentID...)
	}
	contentIds_set := util.RemoveDuplicates(contentIds)
	err = util.WriteSliceStrToCSV("data/", "content_id", contentIds_set, true)
	if err != nil {
		log.Printf("Write file error: %s", err)
	}

	rc_unique := filterUniqueContentID(rc)
	contentIdData, err := s.store.GetAllContentID(context.TODO())
	if err != nil {
		log.Println(err)
		return
	}

	allContentID := []string{}
	for _, v := range *contentIdData {
		allContentID = append(allContentID, v.RibbonID)
	}
	contentIdDataUnique := util.RemoveDuplicates(allContentID)

	contentIdNew := removeContentIDDuplicateDatabase(rc_unique, contentIdDataUnique)

	log.Printf(">> %d new content id: ", len(contentIdNew))
	log.Println(">> fetching content ...")
	step4_fetch_content_and_save(s.store, contentIdNew)
	log.Println("--- Finish Job ---")
}

func step1_get_url() []string {

	var id_list []string
	c := colly.NewCollector(colly.AllowedDomains("vieon.vn"))
	c.OnHTML("script#__NEXT_DATA__", func(h *colly.HTMLElement) {

		var data models.NextDataRes
		err := json.Unmarshal([]byte(h.Text), &data)
		if err != nil {
			log.Println(err)
		}

		for _, m := range data.Props.InitialState.Menu.MenuList {
			id_list = append(id_list, m.Seo.ShareURL)
			for _, sm := range m.SubMenu {
				id_list = append(id_list, sm.Seo.ShareURL)
			}

		}

	})

	c.OnRequest(func(r *colly.Request) {
		log.Println(r.ID, " Visiting: ", r.URL)

	})

	c.Visit("https://vieon.vn/")

	return id_list
}

func step2_get_id_ribbon(urls []string) []string {

	log.Println(len(urls))
	var ribbonIds []string
	index := 0

	c := colly.NewCollector(colly.AllowedDomains("vieon.vn"))

	c.OnHTML("section#SECTION_HOME div.rocopa", func(h *colly.HTMLElement) {

		log.Printf("%d %s", index, h.Attr("id"))
		ribbonIds = append(ribbonIds, h.Attr("id"))
		if index < len(urls)-1 {
			index += 1
			c.Visit(urls[index])
		}

	})

	c.Visit(urls[0])
	c.Wait()

	return ribbonIds

}

type RibbonIDContentID struct {
	RibbonID  string
	ContentID []string
}

func step3_get_id_content(ribbonIds []string, db db.Store) []RibbonIDContentID {
	// wp := workerpool.New(3)

	var rc []RibbonIDContentID

	index := 0

	for _, id := range ribbonIds {
		// id := id
		// wp.Submit(func() {
		res, totalPage := fetchAndFilterContentID(id, 0, db)

		rc = append(rc, RibbonIDContentID{RibbonID: id, ContentID: res})

		index += 1
		page := (totalPage / 30) + 1
		log.Printf("%d  %s?page=%d&limit=30", index, id, 0)
		for i := 1; i < int(page); i++ {
			res2, _ := fetchAndFilterContentID(id, i, db)
			// contentIds = append(contentIds, res2...)
			rc = append(rc, RibbonIDContentID{RibbonID: id, ContentID: res2})
			log.Printf("%d  %s?page=%d&limit=30", index, id, i)
			index += 1
		}

		// })

	}

	// wp.StopWait()

	return rc
}

func filterUniqueContentID(data []RibbonIDContentID) []RibbonIDContentID {

	map1 := make(map[string]string)

	for _, v1 := range data {
		for _, v2 := range v1.ContentID {
			map1[v2] = v1.RibbonID
		}
	}

	contentIDMap := make(map[string]map[string]bool)

	for contentID, ribbonID := range map1 {
		if _, ok := contentIDMap[ribbonID]; !ok {
			contentIDMap[ribbonID] = make(map[string]bool)
		}
		contentIDMap[ribbonID][contentID] = true
	}

	var rc2 []RibbonIDContentID
	for ribbonID, contentIDSet := range contentIDMap {
		var uniqueContentIDs []string
		for contentID := range contentIDSet {
			uniqueContentIDs = append(uniqueContentIDs, contentID)
		}
		rc2 = append(rc2, RibbonIDContentID{RibbonID: ribbonID, ContentID: uniqueContentIDs})
	}

	return rc2
}

func removeContentIDDuplicateDatabase(data []RibbonIDContentID, contentIds []string) []RibbonIDContentID {

	var res []RibbonIDContentID

	for _, v := range data {

		ids := util.FindItemNotInSlice(contentIds, v.ContentID)
		if len(ids) > 0 {
			res = append(res, RibbonIDContentID{RibbonID: v.RibbonID, ContentID: ids})
		}

	}

	return res
}

func step4_fetch_content_and_save(db db.Store, data []RibbonIDContentID) {

	// log.Printf("r: %d", len(data))
	// t := 0
	// for _, v := range data {
	// 	t += len(v.ContentID)
	// }
	// log.Printf("c: %d", t)

	for index, item := range data {
		for _, contentID := range item.ContentID {
			url := fmt.Sprintf("https://api.vieon.vn/backend/cm/v5/content/%s", contentID)

			var res models.ContentRes

			err := util.FetchData(url, &res)
			if err != nil {
				log.Printf("%s :%s", err, contentID)
			}

			cateName := ""
			cateID := ""
			for _, v := range res.Tags {
				if v.Name == "category" {
					cateName = v.Name
					cateID = v.ID
				}
			}

			content := &models.Content{
				ContentID:             res.ID,
				CategoryName:          cateName,
				CategoryID:            cateID,
				RibbonID:              item.RibbonID,
				GroupID:               res.GroupID,
				KnownAs:               res.KnownAs,
				Type:                  res.Type,
				Title:                 res.Title,
				CustomAds:             res.CustomAds,
				EnableAds:             res.EnableAds,
				ShortDescription:      res.ShortDescription,
				LongDescription:       res.LongDescription,
				Resolution:            res.Resolution,
				Runtime:               res.Runtime,
				IsPremium:             res.IsPremium,
				IsNew:                 res.IsNew,
				IsTrailer:             res.IsTrailer,
				ReleaseYear:           res.ReleaseYear,
				ThumbnailHotV4:        res.Images.ThumbnailHotV4,
				ThumbnailBigV4:        res.Images.ThumbnailBigV4,
				CarouselWebV4:         res.Images.CarouselWebV4,
				ThumbnailV4:           res.Images.ThumbnailV4,
				PosterV4:              res.Images.PosterV4,
				TitleCardLight:        res.Images.TitleCardLight,
				TitleCardDark:         res.Images.TitleCardDark,
				ThumbnailHotV4Ntc:     res.Images.ThumbnailHotV4Ntc,
				ThumbnailBigV4Ntc:     res.Images.ThumbnailBigV4Ntc,
				CarouselWebV4Ntc:      res.Images.CarouselWebV4Ntc,
				ThumbnailV4Ntc:        res.Images.ThumbnailV4Ntc,
				AvgRate:               res.AvgRate,
				TotalRate:             res.TotalRate,
				IsWatchlater:          res.IsWatchlater,
				UserRating:            res.UserRating,
				Episode:               res.Episode,
				CurrentEpisode:        res.CurrentEpisode,
				ContentProviderID:     res.ContentProviderID,
				Slug:                  res.Slug,
				SeasonSlugSeo:         res.SeasonSlugSeo,
				IsEnd:                 res.IsEnd,
				IsDownloadable:        res.IsDownloadable,
				NoSeeker:              res.NoSeeker,
				GeoCheck:              res.GeoCheck,
				RangePageIndex:        res.RangePageIndex,
				DrmServiceName:        res.DrmServiceName,
				HaveTrailer:           res.HaveTrailer,
				IsVip:                 res.IsVip,
				IntroStart:            res.IntroStart,
				IntroEnd:              res.IntroEnd,
				OuttroStart:           res.OuttroStart,
				OuttroEnd:             res.OuttroEnd,
				Category:              res.Category,
				SeasonName:            res.SeasonName,
				ShowName:              res.ShowName,
				CreatedAt:             res.CreatedAt,
				RatingValue:           res.RatingValue,
				AuthorName:            res.AuthorName,
				ReviewBody:            res.ReviewBody,
				IsComingSoon:          res.IsComingSoon,
				AgeRange:              res.AgeRange,
				LabelSubtitleAudio:    res.LabelSubtitleAudio,
				LabelPublicDay:        res.LabelPublicDay,
				Ranking:               res.Ranking,
				AwardDescription:      res.AwardDescription,
				IsPremiumDisplay:      res.IsPremiumDisplay,
				Status:                res.Status,
				TrialDuration:         res.TrialDuration,
				ExpireDate:            res.ExpireDate,
				ForceLogin:            res.ForceLogin,
				WindowingMessage:      res.WindowingMessage,
				IsPremiere:            res.IsPremiere,
				ButtonLive:            res.ButtonLive,
				PositionLiveCcu:       res.PositionLiveCcu,
				IsSimulcast:           res.IsSimulcast,
				IsTriggerToApp:        res.IsTriggerToApp,
				NumberTrialEpisode:    res.NumberTrialEpisode,
				IsTriggerRegister:     res.IsTriggerRegister,
				ShowAotp:              res.ShowAotp,
				TimeOnTriggerProgress: res.TimeOnTriggerProgress,
				AgeRestricted:         res.AgeRestricted,
				AllowsKid:             res.AllowsKid,
				HasObjectDetection:    res.HasObjectDetection,
				WarningMessage:        res.WarningMessage,
				WarningScreen:         res.WarningScreen,
				WarningTag:            res.WarningTag,
				CanLimit:              res.CanLimit,
			}
			err = db.CreateContent(context.TODO(), content)
			if err != nil {
				log.Println(err)
			}

			time.Sleep(100 * time.Millisecond)

		}

		if index%100 == 0 {
			time.Sleep(10000 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
		}

	}
}

func fetchAndFilterContentID(id string, page int, db db.Store) ([]string, int) {
	url := fmt.Sprintf("https://api.vieon.vn/backend/cm/v5/ribbon/%s?page=%d&limit=30", id, page)
	var ribbonRes models.RibbonRes
	err := util.FetchData(url, &ribbonRes)
	if err != nil {
		log.Printf("fetch failed: %s", id)
	}

	Ids := []string{}

	if id != "" && !db.IsExistsRibbon(context.TODO(), id) {
		ribbon := &models.Ribbon{
			RibbonID:          id,
			Name:              ribbonRes.Name,
			Type:              ribbonRes.Type,
			DisplayImageType:  ribbonRes.DisplayImageType,
			IsPremium:         ribbonRes.IsPremium,
			ThumbnailHotV4:    ribbonRes.Images.ThumbnailHotV4,
			ThumbnailBigV4:    ribbonRes.Images.ThumbnailBigV4,
			ThumbnailV4:       ribbonRes.Images.ThumbnailV4,
			ThumbnailHotV4Ntc: ribbonRes.Images.ThumbnailHotV4Ntc,
			ThumbnailBigV4Ntc: ribbonRes.Images.ThumbnailBigV4Ntc,
			ThumbnailV4Ntc:    ribbonRes.Images.ThumbnailV4Ntc,
		}
		log.Printf("Insert ribbon id:%s", ribbon.RibbonID)
		err = db.CreateRibbon(context.TODO(), ribbon)
		if err != nil {
			log.Println(err)
		}
	}

	for _, v := range ribbonRes.Items {
		// log.Println(v.ID)
		Ids = append(Ids, v.ID)
	}

	return Ids, int(ribbonRes.Metadata.Total)
}
