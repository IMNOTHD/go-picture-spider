package model

type PixivAjaxResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		IllustDetails struct {
			URL          string   `json:"url"`
			Tags         []string `json:"tags"`
			IllustImages []struct {
				IllustImageWidth  string `json:"illust_image_width"`
				IllustImageHeight string `json:"illust_image_height"`
			} `json:"illust_images"`
			ReuploadTimestamp bool `json:"reupload_timestamp"`
			MangaA            []struct {
				Page     int    `json:"page"`
				URL      string `json:"url"`
				URLSmall string `json:"url_small"`
				URLBig   string `json:"url_big"`
			} `json:"manga_a"`
			DisplayTags []struct {
				Tag                     string `json:"tag"`
				IsPixpediaArticleExists bool   `json:"is_pixpedia_article_exists"`
				SetByAuthor             bool   `json:"set_by_author"`
				IsLocked                bool   `json:"is_locked"`
				IsDeletable             bool   `json:"is_deletable"`
				Translation             string `json:"translation,omitempty"`
			} `json:"display_tags"`
			TagsEditable      bool        `json:"tags_editable"`
			BookmarkUserTotal int         `json:"bookmark_user_total"`
			URLS              string      `json:"url_s"`
			URLSs             string      `json:"url_ss"`
			URLBig            string      `json:"url_big"`
			URLPlaceholder    string      `json:"url_placeholder"`
			UgoiraMeta        interface{} `json:"ugoira_meta"`
			ShareText         string      `json:"share_text"`
			Meta              struct {
				TwitterCard struct {
					Card              string `json:"card"`
					Site              string `json:"site"`
					URL               string `json:"url"`
					Title             string `json:"title"`
					Description       string `json:"description"`
					Image             string `json:"image"`
					AppNameIphone     string `json:"app:name:iphone"`
					AppIDIphone       string `json:"app:id:iphone"`
					AppURLIphone      string `json:"app:url:iphone"`
					AppNameIpad       string `json:"app:name:ipad"`
					AppIDIpad         string `json:"app:id:ipad"`
					AppURLIpad        string `json:"app:url:ipad"`
					AppNameGoogleplay string `json:"app:name:googleplay"`
					AppIDGoogleplay   string `json:"app:id:googleplay"`
					AppURLGoogleplay  string `json:"app:url:googleplay"`
				} `json:"twitter_card"`
				Ogp struct {
					Title       string `json:"title"`
					Type        string `json:"type"`
					Image       string `json:"image"`
					Description string `json:"description"`
				} `json:"ogp"`
				Title             string `json:"title"`
				Description       string `json:"description"`
				DescriptionHeader string `json:"description_header"`
				Canonical         string `json:"canonical"`
			} `json:"meta"`
			IsRated                 bool          `json:"is_rated"`
			ResponseGet             []interface{} `json:"response_get"`
			ResponseSend            []interface{} `json:"response_send"`
			TitleCaptionTranslation struct {
				WorkTitle   interface{} `json:"work_title"`
				WorkCaption interface{} `json:"work_caption"`
			} `json:"title_caption_translation"`
			IsMypixiv       bool     `json:"is_mypixiv"`
			IsPrivate       bool     `json:"is_private"`
			IsHowto         bool     `json:"is_howto"`
			IsOriginal      bool     `json:"is_original"`
			Alt             string   `json:"alt"`
			StorableTags    []string `json:"storable_tags"`
			UploadTimestamp int      `json:"upload_timestamp"`
			ID              string   `json:"id"`
			UserID          string   `json:"user_id"`
			Title           string   `json:"title"`
			Width           string   `json:"width"`
			Height          string   `json:"height"`
			Restrict        string   `json:"restrict"`
			XRestrict       string   `json:"x_restrict"`
			Type            string   `json:"type"`
			Sl              int      `json:"sl"`
			PageCount       string   `json:"page_count"`
			Comment         string   `json:"comment"`
			RatingCount     string   `json:"rating_count"`
			RatingView      string   `json:"rating_view"`
			CommentHTML     string   `json:"comment_html"`
			AuthorDetails   struct {
				UserID      string `json:"user_id"`
				UserName    string `json:"user_name"`
				UserAccount string `json:"user_account"`
			} `json:"author_details"`
		} `json:"illust_details"`
		AuthorDetails struct {
			UserID      string `json:"user_id"`
			UserStatus  string `json:"user_status"`
			UserAccount string `json:"user_account"`
			UserName    string `json:"user_name"`
			UserPremium string `json:"user_premium"`
			ProfileImg  struct {
				Main string `json:"main"`
			} `json:"profile_img"`
			ExternalSiteWorksStatus struct {
				Booth    bool `json:"booth"`
				Sketch   bool `json:"sketch"`
				VroidHub bool `json:"vroidHub"`
			} `json:"external_site_works_status"`
			FanboxDetails interface{} `json:"fanbox_details"`
			IsFollowed    bool        `json:"is_followed"`
			AcceptRequest bool        `json:"accept_request"`
		} `json:"author_details"`
	} `json:"body"`
}

type MedibangAjaxResponse struct {
	Result            string      `json:"result"`
	Messages          interface{} `json:"messages"`
	PictureDetailBean struct {
		ContentsID      string `json:"contentsId"`
		HandleName      string `json:"handleName"`
		AuthorBadge     string `json:"authorBadge"`
		Title           string `json:"title"`
		DescriptiveText string `json:"descriptiveText"`
		PublishDate     int64  `json:"publishDate"`
		TagList         []struct {
			ContentsID           string      `json:"contentsId"`
			ContentsTagID        int         `json:"contentsTagId"`
			Tag                  string      `json:"tag"`
			Path                 interface{} `json:"path"`
			ResizedFileName      interface{} `json:"resizedFileName"`
			ThumbnailURL         interface{} `json:"thumbnailUrl"`
			TranslationTagID     interface{} `json:"translationTagId"`
			TranslationTag       interface{} `json:"translationTag"`
			TagLockFlag          string      `json:"tagLockFlag"`
			OfficialTagFlag      string      `json:"officialTagFlag"`
			PopularContentsID    interface{} `json:"popularContentsId"`
			ContestTagFlag       bool        `json:"contestTagFlag"`
			HasLocaleContestTag  interface{} `json:"hasLocaleContestTag"`
			ContestMasterCode    string      `json:"contestMasterCode"`
			ContestStatus        string      `json:"contestStatus"`
			ContestSectionID     interface{} `json:"contestSectionId"`
			IndexPageURL         interface{} `json:"indexPageUrl"`
			IllustEstimatedCount int         `json:"illustEstimatedCount"`
			FavoriteTagFlag      interface{} `json:"favoriteTagFlag"`
			TagCount             interface{} `json:"tagCount"`
			Height               interface{} `json:"height"`
			Width                interface{} `json:"width"`
			FileName             interface{} `json:"fileName"`
			Mimetype             interface{} `json:"mimetype"`
			SearchedAt           interface{} `json:"searchedAt"`
			ProductsNewCount     interface{} `json:"productsNewCount"`
			ProductsAdultsType   interface{} `json:"productsAdultsType"`
			ProductsAdsType      interface{} `json:"productsAdsType"`
		} `json:"tagList"`
		AwardName               interface{} `json:"awardName"`
		AozoraContentsID        interface{} `json:"aozoraContentsId"`
		UserID                  int         `json:"userId"`
		AuthorRole              string      `json:"authorRole"`
		Author                  interface{} `json:"author"`
		FreshArrivalDisplayFlag string      `json:"freshArrivalDisplayFlag"`
		AozoraUserID            interface{} `json:"aozoraUserId"`
		ContributionDate        int64       `json:"contributionDate"`
		ImageURLOriginal        string      `json:"imageUrlOriginal"`
		ImageURLMedium          string      `json:"imageUrlMedium"`
		MyselfURL               string      `json:"myselfUrl"`
		LikeCount               int         `json:"likeCount"`
		AvailableLikeCount      interface{} `json:"availableLikeCount"`
		LikeCountMyself         bool        `json:"likeCountMyself"`
		UserProfileBean         struct {
			UserID                int         `json:"userId"`
			HeaderFileType        string      `json:"headerFileType"`
			HeaderFileID          int         `json:"headerFileId"`
			UserComment           string      `json:"userComment"`
			ProfileFileID         int         `json:"profileFileId"`
			BackgroundColor       string      `json:"backgroundColor"`
			BackgroundType        interface{} `json:"backgroundType"`
			BackgroundFileID      interface{} `json:"backgroundFileId"`
			BackgroundRepeatType  string      `json:"backgroundRepeatType"`
			BackgroundPosition    string      `json:"backgroundPosition"`
			CrownWord             string      `json:"crownWord"`
			JobType               interface{} `json:"jobType"`
			WorkPlace             int         `json:"workPlace"`
			WorkPlaceName         string      `json:"workPlaceName"`
			WorkPlacePubLevel     string      `json:"workPlacePubLevel"`
			Sex                   string      `json:"sex"`
			GenderText            interface{} `json:"genderText"`
			SexPubLevel           string      `json:"sexPubLevel"`
			WorkEnvFileID         interface{} `json:"workEnvFileId"`
			WorkEnvComment        string      `json:"workEnvComment"`
			FreeTime              interface{} `json:"freeTime"`
			FreeTimeUnit          string      `json:"freeTimeUnit"`
			FreeTimeCycle         string      `json:"freeTimeCycle"`
			JobStatus             string      `json:"jobStatus"`
			WorkableText          string      `json:"workableText"`
			HandleName            string      `json:"handleName"`
			AuthorBadge           string      `json:"authorBadge"`
			UserAlias             string      `json:"userAlias"`
			UserInfo              string      `json:"userInfo"`
			FriendCount           int         `json:"friendCount"`
			FollowCount           int         `json:"followCount"`
			FollowerCount         int         `json:"followerCount"`
			UserType              string      `json:"userType"`
			AuthorBadgePrizeDate  interface{} `json:"authorBadgePrizeDate"`
			OrizinalHeaderFileID  int         `json:"orizinalHeaderFileId"`
			OrizinalProfileFileID int         `json:"orizinalProfileFileId"`
			IsBlocking            bool        `json:"isBlocking"`
			IsBlockProcessing     bool        `json:"isBlockProcessing"`
			IsBlocked             bool        `json:"isBlocked"`
			IsMuting              bool        `json:"isMuting"`
			DisplayJobType        interface{} `json:"displayJobType"`
			JobTypeEnum           interface{} `json:"jobTypeEnum"`
			JobTypeName           string      `json:"jobTypeName"`
		} `json:"userProfileBean"`
		FileSize       float64     `json:"fileSize"`
		ProfileFileURL string      `json:"profileFileUrl"`
		IsMe           bool        `json:"isMe"`
		FriendStatus   interface{} `json:"friendStatus"`
		Follow         bool        `json:"follow"`
		ForAdultsFlag  string      `json:"forAdultsFlag"`
		Locale         string      `json:"locale"`
		ViewCount      int         `json:"viewCount"`
	} `json:"pictureDetailBean"`
	ShowAgeAuthenticate bool        `json:"showAgeAuthenticate"`
	LoginUser           bool        `json:"loginUser"`
	CoauthorList        interface{} `json:"coauthorList"`
	TeamInfoBean        interface{} `json:"teamInfoBean"`
	Favorite            bool        `json:"favorite"`
	CanReadR18          bool        `json:"canReadR18"`
}
