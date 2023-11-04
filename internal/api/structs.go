package api

import "time"

type Profile struct {
	View                      string    `json:"view"`
	Avatar                    any       `json:"avatar"`
	AvatarThumbs              any       `json:"avatarThumbs"`
	Header                    any       `json:"header"`
	HeaderSize                any       `json:"headerSize"`
	HeaderThumbs              any       `json:"headerThumbs"`
	ID                        int       `json:"id"`
	Name                      string    `json:"name"`
	Username                  string    `json:"username"`
	CanLookStory              bool      `json:"canLookStory"`
	CanCommentStory           bool      `json:"canCommentStory"`
	HasNotViewedStory         bool      `json:"hasNotViewedStory"`
	IsVerified                bool      `json:"isVerified"`
	CanPayInternal            bool      `json:"canPayInternal"`
	CanSendChatToAll          bool      `json:"canSendChatToAll"`
	CreditsMin                int       `json:"creditsMin"`
	CreditsMax                int       `json:"creditsMax"`
	IsPaywallPassed           bool      `json:"isPaywallPassed"`
	CreditsMinAlternatives    int       `json:"creditsMinAlternatives"`
	Unprofitable              bool      `json:"unprofitable"`
	ListsSort                 string    `json:"listsSort"`
	ListsSortOrder            string    `json:"listsSortOrder"`
	CanCreateLists            bool      `json:"canCreateLists"`
	JoinDate                  time.Time `json:"joinDate"`
	IsReferrerAllowed         bool      `json:"isReferrerAllowed"`
	About                     string    `json:"about"`
	RawAbout                  string    `json:"rawAbout"`
	Website                   any       `json:"website"`
	Wishlist                  any       `json:"wishlist"`
	Location                  any       `json:"location"`
	PostsCount                int       `json:"postsCount"`
	ArchivedPostsCount        int       `json:"archivedPostsCount"`
	PrivateArchivedPostsCount int       `json:"privateArchivedPostsCount"`
	PhotosCount               int       `json:"photosCount"`
	VideosCount               int       `json:"videosCount"`
	AudiosCount               int       `json:"audiosCount"`
	MediasCount               int       `json:"mediasCount"`
	LastSeen                  time.Time `json:"lastSeen"`
	FavoritesCount            int       `json:"favoritesCount"`
	FavoritedCount            int       `json:"favoritedCount"`
	ShowPostsInFeed           bool      `json:"showPostsInFeed"`
	CanReceiveChatMessage     bool      `json:"canReceiveChatMessage"`
	IsPerformer               bool      `json:"isPerformer"`
	IsRealPerformer           bool      `json:"isRealPerformer"`
	IsSpotifyConnected        bool      `json:"isSpotifyConnected"`
	SubscribersCount          int       `json:"subscribersCount"`
	HasPinnedPosts            bool      `json:"hasPinnedPosts"`
	HasLabels                 bool      `json:"hasLabels"`
	CanChat                   bool      `json:"canChat"`
	NewTagsCount              int       `json:"newTagsCount"`
	HasTags                   bool      `json:"hasTags"`
	CanAddPhone               bool      `json:"canAddPhone"`
	PhoneLast4                any       `json:"phoneLast4"`
	PhoneMask                 any       `json:"phoneMask"`
	HasNewTicketReplies       struct {
		Open   bool `json:"open"`
		Solved bool `json:"solved"`
		Closed bool `json:"closed"`
	} `json:"hasNewTicketReplies"`
	HasInternalPayments             bool     `json:"hasInternalPayments"`
	IsCreditsEnabled                bool     `json:"isCreditsEnabled"`
	CreditBalance                   int      `json:"creditBalance"`
	IsMakePayment                   bool     `json:"isMakePayment"`
	IsAgeVerified                   bool     `json:"isAgeVerified"`
	AgeVerificationRequired         bool     `json:"ageVerificationRequired"`
	IsOtpEnabled                    bool     `json:"isOtpEnabled"`
	Email                           string   `json:"email"`
	IsEmailChecked                  bool     `json:"isEmailChecked"`
	IsLegalApprovedAllowed          bool     `json:"isLegalApprovedAllowed"`
	IsTwitterConnected              bool     `json:"isTwitterConnected"`
	TwitterUsername                 any      `json:"twitterUsername"`
	IsAllowTweets                   bool     `json:"isAllowTweets"`
	IsPaymentCardConnected          bool     `json:"isPaymentCardConnected"`
	ReferalURL                      string   `json:"referalUrl"`
	IsVisibleOnline                 bool     `json:"isVisibleOnline"`
	SubscribesCount                 int      `json:"subscribesCount"`
	CanPinPost                      bool     `json:"canPinPost"`
	HasNewAlerts                    bool     `json:"hasNewAlerts"`
	HasNewHints                     bool     `json:"hasNewHints"`
	HasNewChangedPriceSubscriptions bool     `json:"hasNewChangedPriceSubscriptions"`
	NotificationsCount              int      `json:"notificationsCount"`
	HasSystemNotifications          bool     `json:"hasSystemNotifications"`
	ChatMessagesCount               int      `json:"chatMessagesCount"`
	CountPinnedChat                 int      `json:"countPinnedChat"`
	IsWantComments                  bool     `json:"isWantComments"`
	WatermarkText                   string   `json:"watermarkText"`
	CustomWatermarkText             any      `json:"customWatermarkText"`
	HasWatermarkPhoto               bool     `json:"hasWatermarkPhoto"`
	HasWatermarkVideo               bool     `json:"hasWatermarkVideo"`
	IsTelegramConnected             bool     `json:"isTelegramConnected"`
	AdvBlock                        []string `json:"advBlock"`
	HasPurchasedPosts               bool     `json:"hasPurchasedPosts"`
	IsEmailRequired                 bool     `json:"isEmailRequired"`
	PayoutLegalApproveState         string   `json:"payoutLegalApproveState"`
	PayoutLegalApproveRejectReason  any      `json:"payoutLegalApproveRejectReason"`
	EnabledImageEditorForChat       bool     `json:"enabledImageEditorForChat"`
	ShouldReceiveLessNotifications  bool     `json:"shouldReceiveLessNotifications"`
	CanCalling                      bool     `json:"canCalling"`
	PaidFeed                        bool     `json:"paidFeed"`
	CanSendSms                      bool     `json:"canSendSms"`
	CanAddFriends                   bool     `json:"canAddFriends"`
	IsRealCardConnected             bool     `json:"isRealCardConnected"`
	CountPriorityChat               int      `json:"countPriorityChat"`
	HasScenario                     bool     `json:"hasScenario"`
	WsAuthToken                     string   `json:"wsAuthToken"`
	CanAddCard                      bool     `json:"canAddCard"`
	IsWalletAutorecharge            bool     `json:"isWalletAutorecharge"`
	WalletAutorechargeAmount        int      `json:"walletAutorechargeAmount"`
	WalletAutorechargeMin           int      `json:"walletAutorechargeMin"`
	WalletFirstRebills              bool     `json:"walletFirstRebills"`
	CanAlternativeWalletTopUp       bool     `json:"canAlternativeWalletTopUp"`
	NeedIVApprove                   bool     `json:"needIVApprove"`
	IvStatus                        any      `json:"ivStatus"`
	IvFailReason                    any      `json:"ivFailReason"`
	FaceIDAvailable                 bool     `json:"faceIdAvailable"`
	IvCountry                       any      `json:"ivCountry"`
	IvForcedVerified                bool     `json:"ivForcedVerified"`
	IvFlow                          string   `json:"ivFlow"`
	ForceFaceOtp                    bool     `json:"forceFaceOtp"`
	FaceIDRegular                   struct {
		NeedToShow  bool `json:"needToShow"`
		CanPostpone bool `json:"canPostpone"`
	} `json:"faceIdRegular"`
	ConnectedOfAccounts []any `json:"connectedOfAccounts"`
	HasPassword         bool  `json:"hasPassword"`
	CanConnectOfAccount bool  `json:"canConnectOfAccount"`
	PinnedPostsCount    int   `json:"pinnedPostsCount"`
	MaxPinnedPostsCount int   `json:"maxPinnedPostsCount"`
	IsDeleteInitiated   bool  `json:"isDeleteInitiated"`
}

type Subscriptions struct {
	List []struct {
		View         string `json:"view"`
		Avatar       string `json:"avatar"`
		AvatarThumbs struct {
			C50  string `json:"c50"`
			C144 string `json:"c144"`
		} `json:"avatarThumbs"`
		Header     string `json:"header"`
		HeaderSize struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"headerSize"`
		HeaderThumbs struct {
			W480 string `json:"w480"`
			W760 string `json:"w760"`
		} `json:"headerThumbs"`
		ID                 int     `json:"id"`
		Name               string  `json:"name"`
		Username           string  `json:"username"`
		CanLookStory       bool    `json:"canLookStory"`
		CanCommentStory    bool    `json:"canCommentStory"`
		HasNotViewedStory  bool    `json:"hasNotViewedStory"`
		IsVerified         bool    `json:"isVerified"`
		CanPayInternal     bool    `json:"canPayInternal"`
		HasScheduledStream bool    `json:"hasScheduledStream"`
		HasStream          bool    `json:"hasStream"`
		HasStories         bool    `json:"hasStories"`
		TipsEnabled        bool    `json:"tipsEnabled"`
		TipsTextEnabled    bool    `json:"tipsTextEnabled"`
		TipsMin            float64 `json:"tipsMin"`
		TipsMinInternal    float64 `json:"tipsMinInternal"`
		TipsMax            float64 `json:"tipsMax"`
		CanEarn            bool    `json:"canEarn"`
		CanAddSubscriber   bool    `json:"canAddSubscriber"`
		SubscribePrice     float64 `json:"subscribePrice"`
		IsPaywallRequired  bool    `json:"isPaywallRequired"`
		Unprofitable       bool    `json:"unprofitable,omitempty"`
		ListsStates        []struct {
			ID         string `json:"id"`
			Type       string `json:"type"`
			Name       string `json:"name"`
			HasUser    bool   `json:"hasUser"`
			CanAddUser bool   `json:"canAddUser"`
		} `json:"listsStates"`
		IsRestricted            bool      `json:"isRestricted"`
		CanRestrict             bool      `json:"canRestrict"`
		SubscribedBy            bool      `json:"subscribedBy"`
		SubscribedByExpire      bool      `json:"subscribedByExpire"`
		SubscribedByExpireDate  time.Time `json:"subscribedByExpireDate"`
		SubscribedByAutoprolong bool      `json:"subscribedByAutoprolong"`
		SubscribedIsExpiredNow  bool      `json:"subscribedIsExpiredNow"`
		CurrentSubscribePrice   int       `json:"currentSubscribePrice"`
		SubscribedOn            bool      `json:"subscribedOn"`
		SubscribedOnExpiredNow  any       `json:"subscribedOnExpiredNow"`
		SubscribedOnDuration    any       `json:"subscribedOnDuration"`
		CanReport               bool      `json:"canReport"`
		CanReceiveChatMessage   bool      `json:"canReceiveChatMessage"`
		HideChat                bool      `json:"hideChat"`
		LastSeen                time.Time `json:"lastSeen"`
		IsPerformer             bool      `json:"isPerformer"`
		IsRealPerformer         bool      `json:"isRealPerformer"`
		SubscribedByData        struct {
			Price              float64   `json:"price"`
			NewPrice           float64   `json:"newPrice"`
			RegularPrice       float64   `json:"regularPrice"`
			SubscribePrice     float64   `json:"subscribePrice"`
			DiscountPercent    int       `json:"discountPercent"`
			DiscountPeriod     int       `json:"discountPeriod"`
			SubscribeAt        time.Time `json:"subscribeAt"`
			ExpiredAt          time.Time `json:"expiredAt"`
			RenewedAt          time.Time `json:"renewedAt"`
			DiscountFinishedAt any       `json:"discountFinishedAt"`
			DiscountStartedAt  any       `json:"discountStartedAt"`
			Status             any       `json:"status"`
			IsMuted            bool      `json:"isMuted"`
			UnsubscribeReason  string    `json:"unsubscribeReason"`
			Duration           string    `json:"duration"`
			ShowPostsInFeed    bool      `json:"showPostsInFeed"`
			Subscribes         []struct {
				ID           int64     `json:"id"`
				UserID       int       `json:"userId"`
				SubscriberID int       `json:"subscriberId"`
				Date         time.Time `json:"date"`
				Duration     int       `json:"duration"`
				StartDate    time.Time `json:"startDate"`
				ExpireDate   time.Time `json:"expireDate"`
				CancelDate   any       `json:"cancelDate"`
				Price        float64   `json:"price"`
				RegularPrice float64   `json:"regularPrice"`
				Discount     int       `json:"discount"`
				EarningID    int       `json:"earningId"`
				Action       string    `json:"action"`
				Type         string    `json:"type"`
				OfferStart   any       `json:"offerStart"`
				OfferEnd     any       `json:"offerEnd"`
				IsCurrent    bool      `json:"isCurrent"`
			} `json:"subscribes"`
			HasActivePaidSubscriptions bool `json:"hasActivePaidSubscriptions"`
		} `json:"subscribedByData,omitempty"`
		SubscribedOnData    any  `json:"subscribedOnData"`
		CanTrialSend        bool `json:"canTrialSend"`
		IsBlocked           bool `json:"isBlocked"`
		SubscriptionBundles []struct {
			ID       int     `json:"id"`
			Discount int     `json:"discount"`
			Duration int     `json:"duration"`
			Price    float64 `json:"price"`
			CanBuy   bool    `json:"canBuy"`
		} `json:"subscriptionBundles,omitempty"`
	} `json:"list"`
	HasMore bool `json:"hasMore"`
}

type Messsages struct {
	List []struct {
		ResponseType string `json:"responseType"`
		Text         string `json:"text"`
		GiphyID      any    `json:"giphyId"`
		LockedText   bool   `json:"lockedText"`
		IsFree       bool   `json:"isFree"`
		Price        any    `json:"price"`
		IsMediaReady bool   `json:"isMediaReady"`
		MediaCount   int    `json:"mediaCount"`
		Media        []struct {
			ID               int64  `json:"id"`
			CanView          bool   `json:"canView"`
			Type             string `json:"type"`
			Src              string `json:"src"`
			Preview          string `json:"preview"`
			Thumb            string `json:"thumb"`
			Locked           any    `json:"locked"`
			Duration         int    `json:"duration"`
			HasError         bool   `json:"hasError"`
			SquarePreview    string `json:"squarePreview,omitempty"`
			HasCustomPreview bool   `json:"hasCustomPreview,omitempty"`
			VideoSources     struct {
				Num240 any `json:"240"`
				Num720 any `json:"720"`
			} `json:"videoSources"`
			Source struct {
				Source string `json:"source"`
			} `json:"source"`
			Info struct {
				Source struct {
					Width  int `json:"width"`
					Height int `json:"height"`
					Size   int `json:"size"`
				} `json:"source"`
				Preview struct {
					Width  int `json:"width"`
					Height int `json:"height"`
					Size   int `json:"size"`
				} `json:"preview"`
			} `json:"info"`
		} `json:"media"`
		Previews            []int64 `json:"previews"`
		IsTip               bool    `json:"isTip"`
		IsReportedByMe      bool    `json:"isReportedByMe"`
		IsCouplePeopleMedia bool    `json:"isCouplePeopleMedia"`
		QueueID             int64   `json:"queueId"`
		FromUser            struct {
			ID   int    `json:"id"`
			View string `json:"_view"`
		} `json:"fromUser"`
		IsFromQueue        bool      `json:"isFromQueue"`
		CanUnsendQueue     bool      `json:"canUnsendQueue"`
		UnsendSecondsQueue int       `json:"unsendSecondsQueue"`
		ID                 int64     `json:"id"`
		IsOpened           bool      `json:"isOpened"`
		IsNew              bool      `json:"isNew"`
		CreatedAt          time.Time `json:"createdAt"`
		ChangedAt          time.Time `json:"changedAt"`
		CancelSeconds      int       `json:"cancelSeconds"`
		IsLiked            bool      `json:"isLiked"`
		CanPurchase        bool      `json:"canPurchase"`
		CanPurchaseReason  string    `json:"canPurchaseReason"`
		CanReport          bool      `json:"canReport"`
		CanBePinned        bool      `json:"canBePinned"`
		IsPinned           bool      `json:"isPinned"`
	} `json:"list"`
	HasMore bool `json:"hasMore"`
}

type Posts struct {
	List []struct {
		ResponseType    string    `json:"responseType"`
		ID              int       `json:"id"`
		PostedAt        time.Time `json:"postedAt"`
		PostedAtPrecise string    `json:"postedAtPrecise"`
		ExpiredAt       any       `json:"expiredAt"`
		Author          struct {
			ID   int    `json:"id"`
			View string `json:"_view"`
		} `json:"author"`
		Text                string `json:"text"`
		RawText             string `json:"rawText"`
		LockedText          bool   `json:"lockedText"`
		IsFavorite          bool   `json:"isFavorite"`
		CanReport           bool   `json:"canReport"`
		CanDelete           bool   `json:"canDelete"`
		CanComment          bool   `json:"canComment"`
		CanEdit             bool   `json:"canEdit"`
		IsPinned            bool   `json:"isPinned"`
		FavoritesCount      int    `json:"favoritesCount"`
		MediaCount          int    `json:"mediaCount"`
		IsMediaReady        bool   `json:"isMediaReady"`
		Voting              []any  `json:"voting"`
		IsOpened            bool   `json:"isOpened"`
		CanToggleFavorite   bool   `json:"canToggleFavorite"`
		StreamID            any    `json:"streamId"`
		Price               any    `json:"price"`
		HasVoting           bool   `json:"hasVoting"`
		IsAddedToBookmarks  bool   `json:"isAddedToBookmarks"`
		IsArchived          bool   `json:"isArchived"`
		IsPrivateArchived   bool   `json:"isPrivateArchived"`
		IsDeleted           bool   `json:"isDeleted"`
		HasURL              bool   `json:"hasUrl"`
		IsCouplePeopleMedia bool   `json:"isCouplePeopleMedia"`
		CommentsCount       int    `json:"commentsCount"`
		MentionedUsers      []any  `json:"mentionedUsers"`
		LinkedUsers         []any  `json:"linkedUsers"`
		LinkedPosts         []any  `json:"linkedPosts"`
		TipsAmount          string `json:"tipsAmount"`
		TipsAmountRaw       int    `json:"tipsAmountRaw"`
		Media               []struct {
			ID               int64     `json:"id"`
			Type             string    `json:"type"`
			ConvertedToVideo bool      `json:"convertedToVideo"`
			CanView          bool      `json:"canView"`
			HasError         bool      `json:"hasError"`
			CreatedAt        time.Time `json:"createdAt"`
			Info             struct {
				Source struct {
					Source   string `json:"source"`
					Width    int    `json:"width"`
					Height   int    `json:"height"`
					Size     int    `json:"size"`
					Duration int    `json:"duration"`
				} `json:"source"`
				Preview struct {
					Width  int `json:"width"`
					Height int `json:"height"`
					Size   int `json:"size"`
				} `json:"preview"`
			} `json:"info"`
			Source struct {
				Source   string `json:"source"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Size     int    `json:"size"`
				Duration int    `json:"duration"`
			} `json:"source"`
			SquarePreview    string `json:"squarePreview"`
			Full             string `json:"full"`
			Preview          string `json:"preview"`
			Thumb            string `json:"thumb"`
			HasCustomPreview bool   `json:"hasCustomPreview"`
			Files            struct {
				Preview struct {
					URL string `json:"url"`
				} `json:"preview"`
			} `json:"files"`
			VideoSources struct {
				Num240 string `json:"240"`
				Num720 string `json:"720"`
			} `json:"videoSources"`
		} `json:"media"`
		CanViewMedia bool    `json:"canViewMedia"`
		Preview      []int64 `json:"preview"`
	} `json:"list"`
	HasMore    bool   `json:"hasMore"`
	HeadMarker string `json:"headMarker"`
	TailMarker string `json:"tailMarker"`
	Counters   struct {
		AudiosCount        int `json:"audiosCount"`
		PhotosCount        int `json:"photosCount"`
		VideosCount        int `json:"videosCount"`
		MediasCount        int `json:"mediasCount"`
		PostsCount         int `json:"postsCount"`
		StreamsCount       int `json:"streamsCount"`
		ArchivedPostsCount int `json:"archivedPostsCount"`
	} `json:"counters"`
}

type Highlights struct {
	List []struct {
		ID           int       `json:"id"`
		UserID       int       `json:"userId"`
		Title        string    `json:"title"`
		CoverStoryID int       `json:"coverStoryId"`
		Cover        string    `json:"cover"`
		StoriesCount int       `json:"storiesCount"`
		CreatedAt    time.Time `json:"createdAt"`
	} `json:"list"`
	HasMore bool `json:"hasMore"`
}

type Story struct {
	ID        int  `json:"id"`
	UserID    int  `json:"userId"`
	IsWatched bool `json:"isWatched"`
	IsReady   bool `json:"isReady"`
	Media     []struct {
		ID               int64     `json:"id"`
		Type             string    `json:"type"`
		ConvertedToVideo bool      `json:"convertedToVideo"`
		CanView          bool      `json:"canView"`
		HasError         bool      `json:"hasError"`
		CreatedAt        time.Time `json:"createdAt"`
		Files            struct {
			Source struct {
				URL      string `json:"url"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Duration int    `json:"duration"`
				Size     int    `json:"size"`
				Sources  struct {
					Num240 any `json:"240"`
					Num720 any `json:"720"`
				} `json:"sources"`
			} `json:"source"`
			Thumb struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				Size   int    `json:"size"`
			} `json:"thumb"`
			Preview struct {
				URL     string `json:"url"`
				Width   int    `json:"width"`
				Height  int    `json:"height"`
				Size    int    `json:"size"`
				Sources struct {
					W150 string `json:"w150"`
				} `json:"sources"`
			} `json:"preview"`
			SquarePreview struct {
				URL     string `json:"url"`
				Width   int    `json:"width"`
				Height  int    `json:"height"`
				Size    int    `json:"size"`
				Sources struct {
					W150 string `json:"w150"`
					W480 string `json:"w480"`
				} `json:"sources"`
			} `json:"squarePreview"`
		} `json:"files"`
		HasCustomPreview bool `json:"hasCustomPreview"`
	} `json:"media"`
	CreatedAt time.Time `json:"createdAt"`
	Question  any       `json:"question"`
	CanLike   bool      `json:"canLike"`
	IsLiked   bool      `json:"isLiked"`
}

type Highlight struct {
	ID           int       `json:"id"`
	UserID       int       `json:"userId"`
	Title        string    `json:"title"`
	CoverStoryID int       `json:"coverStoryId"`
	Cover        string    `json:"cover"`
	StoriesCount int       `json:"storiesCount"`
	CreatedAt    time.Time `json:"createdAt"`
	Stories      []Story   `json:"stories"`
}

type Media struct {
	URL     string
	MediaID int
	PostID  int
	Type    string
}
