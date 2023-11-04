package api

var (
	DRO_URL                 = "https://raw.githubusercontent.com/DATAHOARDERS/dynamic-rules/main/onlyfans.json"
	SUBSCRIPTIONS_URL       = "https://onlyfans.com/api2/v2/subscriptions/subscribes?offset=%s&type=active&limit=10&format=infinite"
	PROFILE_URL             = "https://onlyfans.com/api2/v2/users/%s"
	PINNED_POSTS_URL        = "https://onlyfans.com/api2/v2/users/%d/posts?limit=10&order=publish_date_desc&skip_users=all&skip_users_dups=1&pinned=1&format=infinite"
	TIMELINE_POSTS_URL      = "https://onlyfans.com/api2/v2/users/%d/posts?limit=100&order=publish_date_desc&skip_users=all&skip_users_dups=1&pinned=0&format=infinite"
	TIMELINE_NEXT_POSTS_URL = "https://onlyfans.com/api2/v2/users/%d/posts?limit=100&order=publish_date_desc&skip_users=all&skip_users_dups=1&beforePublishTime=%s&pinned=0&format=infinite"
	ARCHIVED_POSTS_URL      = "https://onlyfans.com/api2/v2/users/%d/posts/archived?limit=100&order=publish_date_desc&skip_users=all&skip_users_dups=1&pinned=0&format=infinite"
	ARCHIVED_NEXT_POSTS_URL = "https://onlyfans.com/api2/v2/users/%d/posts/archived?limit=100&order=publish_date_desc&skip_users=all&skip_users_dups=1&beforePublishTime=%s&pinned=0&format=infinite"
	HIGHLIGHTS_URL          = "https://onlyfans.com/api2/v2/users/%d/stories/highlights?limit=5&offset=0"
	STORIES_URL             = "https://onlyfans.com/api2/v2/users/%d/stories"
	HIGHLIGHT_URL           = "https://onlyfans.com/api2/v2/stories/highlights/%d"
	MESSAGES_URL            = "https://onlyfans.com/api2/v2/chats/%d/messages?limit=10&offset=0&order=desc&skip_users=all&skip_users_dups=1"
	MESSAGES_URL_NEXT       = "https://onlyfans.com/api2/v2/chats/%d/messages?limit=10&offset=0&id=%d&order=desc&skip_users=all&skip_users_dups=1"
)
